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

#include "wrapper/ConfAdapter.h"
#include "WrapperException.h"
#include "knowhere/index/vector_index/helpers/IndexParameter.h"
#include "utils/Log.h"

#include <cmath>
#include <memory>
#include <vector>

// TODO(lxj): add conf checker

namespace milvus {
namespace engine {

#if CUDA_VERSION > 9000
#define GPU_MAX_NRPOBE 2048
#else
#define GPU_MAX_NRPOBE 1024
#endif

void
ConfAdapter::MatchBase(knowhere::Config conf) {
    if (conf->metric_type == knowhere::DEFAULT_TYPE)
        conf->metric_type = knowhere::METRICTYPE::L2;
    if (conf->gpu_id == knowhere::INVALID_VALUE)
        conf->gpu_id = 0;
}

knowhere::Config
ConfAdapter::Match(const TempMetaConf& metaconf) {
    auto conf = std::make_shared<knowhere::Cfg>();
    conf->d = metaconf.dim;
    conf->metric_type = metaconf.metric_type;
    conf->gpu_id = metaconf.gpu_id;
    conf->k = metaconf.k;
    MatchBase(conf);
    return conf;
}

knowhere::Config
ConfAdapter::MatchSearch(const TempMetaConf& metaconf, const IndexType& type) {
    auto conf = std::make_shared<knowhere::Cfg>();
    conf->k = metaconf.k;
    return conf;
}

knowhere::Config
IVFConfAdapter::Match(const TempMetaConf& metaconf) {
    auto conf = std::make_shared<knowhere::IVFCfg>();
    conf->nlist = MatchNlist(metaconf.size, metaconf.nlist);
    conf->d = metaconf.dim;
    conf->metric_type = metaconf.metric_type;
    conf->gpu_id = metaconf.gpu_id;
    MatchBase(conf);
    return conf;
}

static constexpr float TYPICAL_COUNT = 1000000.0;

int64_t
IVFConfAdapter::MatchNlist(const int64_t& size, const int64_t& nlist) {
    if (size <= TYPICAL_COUNT / 16384 + 1) {
        // handle less row count, avoid nlist set to 0
        return 1;
    } else if (int(size / TYPICAL_COUNT) * nlist <= 0) {
        // calculate a proper nlist if nlist not specified or size less than TYPICAL_COUNT
        return int(size / TYPICAL_COUNT * 16384);
    }
    return nlist;
}

knowhere::Config
IVFConfAdapter::MatchSearch(const TempMetaConf& metaconf, const IndexType& type) {
    auto conf = std::make_shared<knowhere::IVFCfg>();
    conf->k = metaconf.k;

    if (metaconf.nprobe <= 0)
        conf->nprobe = 16;  // hardcode here
    else
        conf->nprobe = metaconf.nprobe;

    switch (type) {
        case IndexType::FAISS_IVFFLAT_GPU:
        case IndexType::FAISS_IVFSQ8_GPU:
        case IndexType::FAISS_IVFPQ_GPU:
            if (conf->nprobe > GPU_MAX_NRPOBE) {
                WRAPPER_LOG_WARNING << "When search with GPU, nprobe shoud be no more than " << GPU_MAX_NRPOBE
                                    << ", but you passed " << conf->nprobe << ". Search with " << GPU_MAX_NRPOBE
                                    << " instead";
                conf->nprobe = GPU_MAX_NRPOBE;
            }
    }
    return conf;
}

knowhere::Config
IVFSQConfAdapter::Match(const TempMetaConf& metaconf) {
    auto conf = std::make_shared<knowhere::IVFSQCfg>();
    conf->nlist = MatchNlist(metaconf.size, metaconf.nlist);
    conf->d = metaconf.dim;
    conf->metric_type = metaconf.metric_type;
    conf->gpu_id = metaconf.gpu_id;
    conf->nbits = 8;
    MatchBase(conf);
    return conf;
}

knowhere::Config
IVFPQConfAdapter::Match(const TempMetaConf& metaconf) {
    auto conf = std::make_shared<knowhere::IVFPQCfg>();
    conf->nlist = MatchNlist(metaconf.size, metaconf.nlist);
    conf->d = metaconf.dim;
    conf->metric_type = metaconf.metric_type;
    conf->gpu_id = metaconf.gpu_id;
    conf->nbits = 8;
    MatchBase(conf);

    /*
     * Faiss 1.6
     * Only 1, 2, 3, 4, 6, 8, 10, 12, 16, 20, 24, 28, 32 dims per sub-quantizer are currently supporte with
     * no precomputed codes. Precomputed codes supports any number of dimensions, but will involve memory overheads.
     */
    static std::vector<int64_t> support_dim_per_subquantizer{32, 28, 24, 20, 16, 12, 10, 8, 6, 4, 3, 2, 1};
    static std::vector<int64_t> support_subquantizer{96, 64, 56, 48, 40, 32, 28, 24, 20, 16, 12, 8, 4, 3, 2, 1};
    std::vector<int64_t> resset;
    for (const auto& dimperquantizer : support_dim_per_subquantizer) {
        if (!(conf->d % dimperquantizer)) {
            auto subquantzier_num = conf->d / dimperquantizer;
            auto finder = std::find(support_subquantizer.begin(), support_subquantizer.end(), subquantzier_num);
            if (finder != support_subquantizer.end()) {
                resset.push_back(subquantzier_num);
            }
        }
    }

    if (resset.empty()) {
        // todo(linxj): throw exception here.
        return nullptr;
    }

    static int64_t compression_level = 1; // 1:low, 2:high
    if (compression_level == 1) {
        conf->m = resset[int(resset.size()/2)];
        WRAPPER_LOG_DEBUG << "PQ m = " << conf->m << ", compression radio = " <<  conf->d / conf->m * 4;
    }
    return conf;
}

knowhere::Config
IVFPQConfAdapter::MatchSearch(const TempMetaConf& metaconf, const IndexType& type) {
    auto conf = std::make_shared<knowhere::IVFPQCfg>();
    conf->k = metaconf.k;

    if (metaconf.nprobe <= 0) {
        WRAPPER_LOG_ERROR << "The nprobe of PQ is wrong!";
        throw WrapperException("The nprobe of PQ is wrong!");
    } else {
        conf->nprobe = metaconf.nprobe;
    }

    return conf;
}

int64_t
IVFPQConfAdapter::MatchNlist(const int64_t& size, const int64_t& nlist) {
    if (size <= TYPICAL_COUNT / 16384 + 1) {
        // handle less row count, avoid nlist set to 0
        return 1;
    } else if (int(size / TYPICAL_COUNT) * nlist <= 0) {
        // calculate a proper nlist if nlist not specified or size less than TYPICAL_COUNT
        return int(size / TYPICAL_COUNT * 16384);
    }
    return nlist;
}

knowhere::Config
NSGConfAdapter::Match(const TempMetaConf& metaconf) {
    auto conf = std::make_shared<knowhere::NSGCfg>();
    conf->nlist = MatchNlist(metaconf.size, metaconf.nlist);
    conf->d = metaconf.dim;
    conf->metric_type = metaconf.metric_type;
    conf->gpu_id = metaconf.gpu_id;
    conf->k = metaconf.k;

    auto scale_factor = round(metaconf.dim / 128.0);
    scale_factor = scale_factor >= 4 ? 4 : scale_factor;
    conf->nprobe = int64_t(conf->nlist * 0.01);
    conf->knng = 40 + 10 * scale_factor;  // the size of knng
    conf->search_length = 40 + 5 * scale_factor;
    conf->out_degree = 50 + 5 * scale_factor;
    conf->candidate_pool_size = 200 + 100 * scale_factor;
    MatchBase(conf);
    return conf;
}

knowhere::Config
NSGConfAdapter::MatchSearch(const TempMetaConf& metaconf, const IndexType& type) {
    auto conf = std::make_shared<knowhere::NSGCfg>();
    conf->k = metaconf.k;
    conf->search_length = metaconf.search_length;
    if (metaconf.search_length == TEMPMETA_DEFAULT_VALUE) {
        conf->search_length = 30;  // TODO(linxj): hardcode here.
    }
    return conf;
}

knowhere::Config
SPTAGKDTConfAdapter::Match(const TempMetaConf& metaconf) {
    auto conf = std::make_shared<knowhere::KDTCfg>();
    conf->d = metaconf.dim;
    conf->metric_type = metaconf.metric_type;
    return conf;
}

knowhere::Config
SPTAGKDTConfAdapter::MatchSearch(const TempMetaConf& metaconf, const IndexType& type) {
    auto conf = std::make_shared<knowhere::KDTCfg>();
    conf->k = metaconf.k;
    return conf;
}

knowhere::Config
SPTAGBKTConfAdapter::Match(const TempMetaConf& metaconf) {
    auto conf = std::make_shared<knowhere::BKTCfg>();
    conf->d = metaconf.dim;
    conf->metric_type = metaconf.metric_type;
    return conf;
}

knowhere::Config
SPTAGBKTConfAdapter::MatchSearch(const TempMetaConf& metaconf, const IndexType& type) {
    auto conf = std::make_shared<knowhere::BKTCfg>();
    conf->k = metaconf.k;
    return conf;
}

}  // namespace engine
}  // namespace milvus
