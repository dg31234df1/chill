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

#pragma once

#include <memory>

namespace zilliz {
namespace milvus {
namespace engine {

enum class ScheduleContextType {
    kUnknown = 0,
    kSearch,
    kDelete,
};

class IScheduleContext {
public:
    IScheduleContext(ScheduleContextType type)
    : type_(type) {
    }

    virtual ~IScheduleContext() = default;

    ScheduleContextType type() const { return type_; }

protected:
    ScheduleContextType type_;
};

using ScheduleContextPtr = std::shared_ptr<IScheduleContext>;

}
}
}
