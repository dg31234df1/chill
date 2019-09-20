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


#include <faiss/gpu/GpuIndexIVFPQ.h>
#include <faiss/gpu/GpuAutoTune.h>
#include <faiss/IndexIVFPQ.h>

#include "IndexGPUIVFPQ.h"
#include "knowhere/common/Exception.h"
#include "knowhere/adapter/VectorAdapter.h"


namespace zilliz {
namespace knowhere {

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

} // knowhere
} // zilliz
