// Copyright (C) 2019-2020 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License.
#ifdef MILVUS_GPU_VERSION
#include "scheduler/optimizer/FaissIVFPQPass.h"
#include "cache/GpuCacheMgr.h"
#include "scheduler/SchedInst.h"
#include "scheduler/Utils.h"
#include "scheduler/task/SearchTask.h"
#include "scheduler/tasklabel/SpecResLabel.h"
#include "server/Config.h"
#include "utils/Log.h"

#include <fiu-local.h>

namespace milvus {
namespace scheduler {

void
FaissIVFPQPass::Init() {
#ifdef MILVUS_GPU_VERSION
    server::Config& config = server::Config::GetInstance();
    Status s = config.GetEngineConfigGpuSearchThreshold(threshold_);
    if (!s.ok()) {
        threshold_ = std::numeric_limits<int32_t>::max();
    }
    s = config.GetGpuResourceConfigSearchResources(search_gpus_);
    if (!s.ok()) {
        throw std::exception();
    }

    SetIdentity("FaissIVFPQPass");
    AddGpuEnableListener();
    AddGpuSearchThresholdListener();
    AddGpuSearchResListener();
#endif
}

bool
FaissIVFPQPass::Run(const TaskPtr& task) {
    if (task->Type() != TaskType::SearchTask) {
        return false;
    }

    auto search_task = std::static_pointer_cast<XSearchTask>(task);
    if (search_task->file_->engine_type_ != (int)engine::EngineType::FAISS_PQ) {
        return false;
    }

    auto search_job = std::static_pointer_cast<SearchJob>(search_task->job_.lock());
    ResourcePtr res_ptr;
    if (!gpu_enable_) {
        SERVER_LOG_DEBUG << "FaissIVFPQPass: gpu disable, specify cpu to search!";
        res_ptr = ResMgrInst::GetInstance()->GetResource("cpu");
    } else if (search_job->nq() < threshold_) {
        SERVER_LOG_DEBUG << "FaissIVFPQPass: nq < gpu_search_threshold, specify cpu to search!";
        res_ptr = ResMgrInst::GetInstance()->GetResource("cpu");
    } else {
        auto best_device_id = count_ % search_gpus_.size();
        SERVER_LOG_DEBUG << "FaissIVFPQPass: nq > gpu_search_threshold, specify gpu" << best_device_id << " to search!";
        ++count_;
        res_ptr = ResMgrInst::GetInstance()->GetResource(ResourceType::GPU, search_gpus_[best_device_id]);
    }
    auto label = std::make_shared<SpecResLabel>(res_ptr);
    task->label() = label;
    return true;
}

}  // namespace scheduler
}  // namespace milvus
#endif
