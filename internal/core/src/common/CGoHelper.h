// Copyright (C) 2019-2020 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License

#pragma once

#include <string>
#include "common/type_c.h"

namespace milvus {

inline CProtoResult
AllocCProtoResult(const google::protobuf::Message& msg) {
    auto size = msg.ByteSize();
    void* buffer = malloc(size);
    msg.SerializePartialToArray(buffer, size);
    return CProtoResult{CStatus{Success}, CProto{buffer, size}};
}

inline CStatus
SuccessCStatus() {
    return CStatus{Success, ""};
}

inline CStatus
FailureCStatus(ErrorCode error_code, const std::string& str) {
    auto str_dup = strdup(str.c_str());
    return CStatus{error_code, str_dup};
}

}  // namespace milvus
