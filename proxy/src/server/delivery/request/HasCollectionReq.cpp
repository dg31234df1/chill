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

#include "server/delivery/request/HasCollectionReq.h"
#include "server/ValidationUtil.h"
#include "utils/Log.h"
#include "utils/TimeRecorder.h"


namespace milvus {
namespace server {

HasCollectionReq::HasCollectionReq(const ContextPtr& context, const std::string& collection_name, bool& exist)
    : BaseReq(context, ReqType::kHasCollection), collection_name_(collection_name), exist_(exist) {
}

BaseReqPtr
HasCollectionReq::Create(const ContextPtr& context, const std::string& collection_name, bool& exist) {
    return std::shared_ptr<BaseReq>(new HasCollectionReq(context, collection_name, exist));
}

Status
HasCollectionReq::OnExecute() {

    return Status::OK();
}

}  // namespace server
}  // namespace milvus
