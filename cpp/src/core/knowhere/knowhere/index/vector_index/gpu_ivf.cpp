// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.



#include <faiss/gpu/GpuIndexFlat.h>
#include <faiss/gpu/GpuIndexIVF.h>
#include <faiss/gpu/GpuIndexIVFFlat.h>
#include <faiss/gpu/GpuIndexIVFPQ.h>
#include <faiss/gpu/GpuAutoTune.h>
#include <faiss/IndexIVFPQ.h>
#include <faiss/index_io.h>


#include "knowhere/common/exception.h"
#include "cloner.h"
#include "knowhere/adapter/faiss_adopt.h"
#include "gpu_ivf.h"

#include <algorithm>

namespace zilliz {
namespace knowhere {

IndexModelPtr GPUIVF::Train(const DatasetPtr &dataset, const Config &config) {
    auto nlist = config["nlist"].as<size_t>();
    gpu_id_ = config.get_with_default("gpu_id", gpu_id_);
    auto metric_type = config["metric_type"].as_string() == "L2" ?
                       faiss::METRIC_L2 : faiss::METRIC_INNER_PRODUCT;

    GETTENSOR(dataset)

    auto temp_resource = FaissGpuResourceMgr::GetInstance().GetRes(gpu_id_);
    if (temp_resource != nullptr) {
        ResScope rs(temp_resource, gpu_id_, true);
        faiss::gpu::GpuIndexIVFFlatConfig idx_config;
        idx_config.device = gpu_id_;
        faiss::gpu::GpuIndexIVFFlat device_index(temp_resource->faiss_res.get(), dim, nlist, metric_type, idx_config);
        device_index.train(rows, (float *) p_data);

        std::shared_ptr<faiss::Index> host_index = nullptr;
        host_index.reset(faiss::gpu::index_gpu_to_cpu(&device_index));

        return std::make_shared<IVFIndexModel>(host_index);
    } else {
        KNOWHERE_THROW_MSG("Build IVF can't get gpu resource");
    }
}

void GPUIVF::set_index_model(IndexModelPtr model) {
    std::lock_guard<std::mutex> lk(mutex_);

    auto host_index = std::static_pointer_cast<IVFIndexModel>(model);
    if (auto gpures = FaissGpuResourceMgr::GetInstance().GetRes(gpu_id_)) {
        ResScope rs(gpures, gpu_id_, false);
        auto device_index = faiss::gpu::index_cpu_to_gpu(gpures->faiss_res.get(), gpu_id_, host_index->index_.get());
        index_.reset(device_index);
        res_ = gpures;
    } else {
        KNOWHERE_THROW_MSG("load index model error, can't get gpu_resource");
    }
}

BinarySet GPUIVF::SerializeImpl() {
    if (!index_ || !index_->is_trained) {
        KNOWHERE_THROW_MSG("index not initialize or trained");
    }

    try {
        MemoryIOWriter writer;
        {
            faiss::Index *index = index_.get();
            faiss::Index *host_index = faiss::gpu::index_gpu_to_cpu(index);

            SealImpl();

            faiss::write_index(host_index, &writer);
            delete host_index;
        }
        auto data = std::make_shared<uint8_t>();
        data.reset(writer.data_);

        BinarySet res_set;
        res_set.Append("IVF", data, writer.rp);

        return res_set;
    } catch (std::exception &e) {
        KNOWHERE_THROW_MSG(e.what());
    }
}

void GPUIVF::LoadImpl(const BinarySet &index_binary) {
    auto binary = index_binary.GetByName("IVF");
    MemoryIOReader reader;
    {
        reader.total = binary->size;
        reader.data_ = binary->data.get();

        faiss::Index *index = faiss::read_index(&reader);

        if (auto temp_res = FaissGpuResourceMgr::GetInstance().GetRes(gpu_id_)) {
            ResScope rs(temp_res, gpu_id_, false);
            auto device_index = faiss::gpu::index_cpu_to_gpu(temp_res->faiss_res.get(), gpu_id_, index);
            index_.reset(device_index);
            res_ = temp_res;
        } else {
            KNOWHERE_THROW_MSG("Load error, can't get gpu resource");
        }

        delete index;
    }
}

IVFIndexPtr GPUIVF::Copy_index_gpu_to_cpu() {
    std::lock_guard<std::mutex> lk(mutex_);

    faiss::Index *device_index = index_.get();
    faiss::Index *host_index = faiss::gpu::index_gpu_to_cpu(device_index);

    std::shared_ptr<faiss::Index> new_index;
    new_index.reset(host_index);
    return std::make_shared<IVF>(new_index);
}

void GPUIVF::search_impl(int64_t n,
                         const float *data,
                         int64_t k,
                         float *distances,
                         int64_t *labels,
                         const Config &cfg) {
    std::lock_guard<std::mutex> lk(mutex_);

    if (auto device_index = std::static_pointer_cast<faiss::gpu::GpuIndexIVF>(index_)) {
        auto nprobe = cfg.get_with_default("nprobe", size_t(1));
        device_index->setNumProbes(nprobe);

        {
            // TODO(linxj): allocate mem
            ResScope rs(res_, gpu_id_);
            device_index->search(n, (float *) data, k, distances, labels);
        }
    }
}

VectorIndexPtr GPUIVF::CopyGpuToCpu(const Config &config) {
    std::lock_guard<std::mutex> lk(mutex_);

    faiss::Index *device_index = index_.get();
    faiss::Index *host_index = faiss::gpu::index_gpu_to_cpu(device_index);

    std::shared_ptr<faiss::Index> new_index;
    new_index.reset(host_index);
    return std::make_shared<IVF>(new_index);
}

VectorIndexPtr GPUIVF::Clone() {
    auto cpu_idx = CopyGpuToCpu(Config());
    return ::zilliz::knowhere::CopyCpuToGpu(cpu_idx, gpu_id_, Config());
}

VectorIndexPtr GPUIVF::CopyGpuToGpu(const int64_t &device_id, const Config &config) {
    auto host_index = CopyGpuToCpu(config);
    return std::static_pointer_cast<IVF>(host_index)->CopyCpuToGpu(device_id, config);
}

void GPUIVF::Add(const DatasetPtr &dataset, const Config &config) {
    if (auto spt = res_.lock()) {
        ResScope rs(res_, gpu_id_);
        IVF::Add(dataset, config);
    }
    else {
        KNOWHERE_THROW_MSG("Add IVF can't get gpu resource");
    }
}

IndexModelPtr GPUIVFPQ::Train(const DatasetPtr &dataset, const Config &config) {
    auto nlist = config["nlist"].as<size_t>();
    auto M = config["M"].as<size_t>();        // number of subquantizers(subvectors)
    auto nbits = config["nbits"].as<size_t>();// number of bit per subvector index
    auto gpu_num = config.get_with_default("gpu_id", gpu_id_);
    auto metric_type = config["metric_type"].as_string() == "L2" ?
                       faiss::METRIC_L2 : faiss::METRIC_L2; // IP not support.

    GETTENSOR(dataset)

    // TODO(linxj): set device here.
    // TODO(linxj): set gpu resource here.
    faiss::gpu::StandardGpuResources res;
    faiss::gpu::GpuIndexIVFPQ device_index(&res, dim, nlist, M, nbits, metric_type);
    device_index.train(rows, (float *) p_data);

    std::shared_ptr<faiss::Index> host_index = nullptr;
    host_index.reset(faiss::gpu::index_gpu_to_cpu(&device_index));

    return std::make_shared<IVFIndexModel>(host_index);
}

std::shared_ptr<faiss::IVFSearchParameters> GPUIVFPQ::GenParams(const Config &config) {
    auto params = std::make_shared<faiss::IVFPQSearchParameters>();
    params->nprobe = config.get_with_default("nprobe", size_t(1));
    //params->scan_table_threshold = 0;
    //params->polysemous_ht = 0;
    //params->max_codes = 0;

    return params;
}

VectorIndexPtr GPUIVFPQ::CopyGpuToCpu(const Config &config) {
    KNOWHERE_THROW_MSG("not support yet");
}

IndexModelPtr GPUIVFSQ::Train(const DatasetPtr &dataset, const Config &config) {
    auto nlist = config["nlist"].as<size_t>();
    auto nbits = config["nbits"].as<size_t>(); // TODO(linxj):  gpu only support SQ4 SQ8 SQ16
    gpu_id_ = config.get_with_default("gpu_id", gpu_id_);
    auto metric_type = config["metric_type"].as_string() == "L2" ?
                       faiss::METRIC_L2 : faiss::METRIC_INNER_PRODUCT;

    GETTENSOR(dataset)

    std::stringstream index_type;
    index_type << "IVF" << nlist << "," << "SQ" << nbits;
    auto build_index = faiss::index_factory(dim, index_type.str().c_str(), metric_type);

    auto temp_resource = FaissGpuResourceMgr::GetInstance().GetRes(gpu_id_);
    if (temp_resource != nullptr) {
        ResScope rs(temp_resource, gpu_id_, true);
        auto device_index = faiss::gpu::index_cpu_to_gpu(temp_resource->faiss_res.get(), gpu_id_, build_index);
        device_index->train(rows, (float *) p_data);

        std::shared_ptr<faiss::Index> host_index = nullptr;
        host_index.reset(faiss::gpu::index_gpu_to_cpu(device_index));

        delete device_index;
        delete build_index;

        return std::make_shared<IVFIndexModel>(host_index);
    } else {
        KNOWHERE_THROW_MSG("Build IVFSQ can't get gpu resource");
    }
}

VectorIndexPtr GPUIVFSQ::CopyGpuToCpu(const Config &config) {
    std::lock_guard<std::mutex> lk(mutex_);

    faiss::Index *device_index = index_.get();
    faiss::Index *host_index = faiss::gpu::index_gpu_to_cpu(device_index);

    std::shared_ptr<faiss::Index> new_index;
    new_index.reset(host_index);
    return std::make_shared<IVFSQ>(new_index);
}

void GPUIndex::SetGpuDevice(const int &gpu_id) {
    gpu_id_ = gpu_id;
}

const int64_t &GPUIndex::GetGpuDevice() {
    return gpu_id_;
}

}
}
