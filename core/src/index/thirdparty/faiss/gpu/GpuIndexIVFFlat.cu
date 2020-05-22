/**
 * Copyright (c) Facebook, Inc. and its affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */


#include <faiss/gpu/GpuIndexIVFFlat.h>
#include <faiss/IndexFlat.h>
#include <faiss/IndexIVFFlat.h>
#include <faiss/gpu/GpuIndexFlat.h>
#include <faiss/gpu/GpuResources.h>
#include <faiss/gpu/impl/IVFFlat.cuh>
#include <faiss/gpu/utils/CopyUtils.cuh>
#include <faiss/gpu/utils/DeviceUtils.h>
#include <faiss/gpu/utils/Float16.cuh>

#include <limits>

namespace faiss { namespace gpu {

GpuIndexIVFFlat::GpuIndexIVFFlat(GpuResources* resources,
                                 const faiss::IndexIVFFlat* index,
                                 GpuIndexIVFFlatConfig config) :
    GpuIndexIVF(resources,
                index->d,
                index->metric_type,
                index->metric_arg,
                index->nlist,
                config),
    ivfFlatConfig_(config),
    reserveMemoryVecs_(0),
    index_(nullptr) {
  copyFrom(index);
}

GpuIndexIVFFlat::GpuIndexIVFFlat(GpuResources* resources,
                                 int dims,
                                 int nlist,
                                 faiss::MetricType metric,
                                 GpuIndexIVFFlatConfig config) :
    GpuIndexIVF(resources, dims, metric, 0, nlist, config),
    ivfFlatConfig_(config),
    reserveMemoryVecs_(0),
    index_(nullptr) {

  // faiss::Index params
  this->is_trained = false;

  // We haven't trained ourselves, so don't construct the IVFFlat
  // index yet
}

GpuIndexIVFFlat::~GpuIndexIVFFlat() {
  delete index_;
}

void
GpuIndexIVFFlat::reserveMemory(size_t numVecs) {
  reserveMemoryVecs_ = numVecs;
  if (index_) {
    DeviceScope scope(device_);
    index_->reserveMemory(numVecs);
  }
}

void
GpuIndexIVFFlat::copyFrom(const faiss::IndexIVFFlat* index) {
  DeviceScope scope(device_);

  GpuIndexIVF::copyFrom(index);

  // Clear out our old data
  delete index_;
  index_ = nullptr;

  // The other index might not be trained
  if (!index->is_trained) {
    FAISS_ASSERT(!is_trained);
    return;
  }

  // Otherwise, we can populate ourselves from the other index
  FAISS_ASSERT(is_trained);

  // Copy our lists as well
  index_ = new IVFFlat(resources_,
                       quantizer->getGpuData(),
                       index->metric_type,
                       index->metric_arg,
                       false, // no residual
                       nullptr, // no scalar quantizer
                       ivfFlatConfig_.indicesOptions,
                       memorySpace_);
  InvertedLists *ivf = index->invlists;

  if (ReadOnlyArrayInvertedLists* rol = dynamic_cast<ReadOnlyArrayInvertedLists*>(ivf)) {
    index_->copyCodeVectorsFromCpu((const float* )(rol->pin_readonly_codes->data),
                                   (const long *)(rol->pin_readonly_ids->data), rol->readonly_length);
    /* double t0 = getmillisecs(); */
    /* std::cout << "Readonly Takes " << getmillisecs() - t0 << " ms" << std::endl; */
  } else {
    for (size_t i = 0; i < ivf->nlist; ++i) {
      auto numVecs = ivf->list_size(i);

      // GPU index can only support max int entries per list
      FAISS_THROW_IF_NOT_FMT(numVecs <=
                             (size_t) std::numeric_limits<int>::max(),
                             "GPU inverted list can only support "
                             "%zu entries; %zu found",
                             (size_t) std::numeric_limits<int>::max(),
                             numVecs);

      index_->addCodeVectorsFromCpu(i,
                                    (const unsigned char*)(ivf->get_codes(i)),
                                    ivf->get_ids(i),
                                    numVecs);
    }
  }
}

void
GpuIndexIVFFlat::copyTo(faiss::IndexIVFFlat* index) const {
  DeviceScope scope(device_);

  // We must have the indices in order to copy to ourselves
  FAISS_THROW_IF_NOT_MSG(ivfFlatConfig_.indicesOptions != INDICES_IVF,
                         "Cannot copy to CPU as GPU index doesn't retain "
                         "indices (INDICES_IVF)");

  GpuIndexIVF::copyTo(index);
  index->code_size = this->d * sizeof(float);

  InvertedLists *ivf = new ArrayInvertedLists(nlist, index->code_size);
  index->replace_invlists(ivf, true);

  // Copy the inverted lists
  if (index_) {
    for (int i = 0; i < nlist; ++i) {
      auto listIndices = index_->getListIndices(i);
      auto listData = index_->getListVectors(i);

      ivf->add_entries(i,
                       listIndices.size(),
                       listIndices.data(),
                       (const uint8_t*) listData.data());
    }
  }
}

size_t
GpuIndexIVFFlat::reclaimMemory() {
  if (index_) {
    DeviceScope scope(device_);

    return index_->reclaimMemory();
  }

  return 0;
}

void
GpuIndexIVFFlat::reset() {
  if (index_) {
    DeviceScope scope(device_);

    index_->reset();
    this->ntotal = 0;
  } else {
    FAISS_ASSERT(this->ntotal == 0);
  }
}

void
GpuIndexIVFFlat::train(Index::idx_t n, const float* x) {
  DeviceScope scope(device_);

  if (this->is_trained) {
    FAISS_ASSERT(quantizer->is_trained);
    FAISS_ASSERT(quantizer->ntotal == nlist);
    FAISS_ASSERT(index_);
    return;
  }

  FAISS_ASSERT(!index_);

  trainQuantizer_(n, x);

  // The quantizer is now trained; construct the IVF index
  index_ = new IVFFlat(resources_,
                       quantizer->getGpuData(),
                       this->metric_type,
                       this->metric_arg,
                       false, // no residual
                       nullptr, // no scalar quantizer
                       ivfFlatConfig_.indicesOptions,
                       memorySpace_);

  if (reserveMemoryVecs_) {
    index_->reserveMemory(reserveMemoryVecs_);
  }

  this->is_trained = true;
}

void
GpuIndexIVFFlat::addImpl_(int n,
                          const float* x,
                          const Index::idx_t* xids) {
  // Device is already set in GpuIndex::add
  FAISS_ASSERT(index_);
  FAISS_ASSERT(n > 0);

  auto stream = resources_->getDefaultStream(device_);

  // Data is already resident on the GPU
  Tensor<float, 2, true> data(const_cast<float*>(x), {n, (int) this->d});

  static_assert(sizeof(long) == sizeof(Index::idx_t), "size mismatch");
  Tensor<long, 1, true> labels(const_cast<long*>(xids), {n});

  // Not all vectors may be able to be added (some may contain NaNs etc)
  index_->classifyAndAddVectors(data, labels);

  // but keep the ntotal based on the total number of vectors that we attempted
  // to add
  ntotal += n;
}

void
GpuIndexIVFFlat::searchImpl_(int n,
                             const float* x,
                             int k,
                             float* distances,
                             Index::idx_t* labels,
                             ConcurrentBitsetPtr bitset) const {
  // Device is already set in GpuIndex::search
  FAISS_ASSERT(index_);
  FAISS_ASSERT(n > 0);

  auto stream = resources_->getDefaultStream(device_);

  // Data is already resident on the GPU
  Tensor<float, 2, true> queries(const_cast<float*>(x), {n, (int) this->d});
  Tensor<float, 2, true> outDistances(distances, {n, k});

  static_assert(sizeof(long) == sizeof(Index::idx_t), "size mismatch");
  Tensor<long, 2, true> outLabels(const_cast<long*>(labels), {n, k});

  if (!bitset) {
    auto bitsetDevice = toDevice<uint8_t, 1>(resources_, device_, nullptr, stream, {0});
    index_->query(queries, bitsetDevice, nprobe, k, outDistances, outLabels);
  } else {
    auto bitsetDevice = toDevice<uint8_t, 1>(resources_, device_,
                                             const_cast<uint8_t*>(bitset->data()), stream,
                                             {(int) bitset->size()});
    index_->query(queries, bitsetDevice, nprobe, k, outDistances, outLabels);
  }
}


} } // namespace
