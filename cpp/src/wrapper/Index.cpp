////////////////////////////////////////////////////////////////////////////////
// Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited.
// Proprietary and confidential.
////////////////////////////////////////////////////////////////////////////////

#ifdef CUDA_VERSION
#include "faiss/gpu/GpuAutoTune.h"
#include "faiss/gpu/StandardGpuResources.h"
#include "faiss/gpu/utils/DeviceUtils.h"
#endif

#include "Index.h"

namespace zilliz {
namespace vecwise {
namespace engine {

using std::string;
using std::unordered_map;
using std::vector;

Index::Index(const std::shared_ptr<faiss::Index> &raw_index) {
    index_ = raw_index;
    dim = index_->d;
    ntotal = index_->ntotal;
    store_on_gpu = false;
}

bool Index::reset() {
    try {
        index_->reset();
        ntotal = index_->ntotal;
    }
    catch (std::exception &e) {
//        LOG(ERROR) << e.what();
        return false;
    }
    return true;
}

bool Index::add_with_ids(idx_t n, const float *xdata, const long *xids) {
    try {
        index_->add_with_ids(n, xdata, xids);
        ntotal += n;
    }
    catch (std::exception &e) {
//        LOG(ERROR) << e.what();
        return false;
    }
    return true;
}

bool Index::search(idx_t n, const float *data, idx_t k, float *distances, long *labels) const {
    try {
        index_->search(n, data, k, distances, labels);
    }
    catch (std::exception &e) {
//        LOG(ERROR) << e.what();
        return false;
    }
    return true;
}

}
}
}
