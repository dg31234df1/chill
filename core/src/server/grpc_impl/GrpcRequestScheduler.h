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

#include "grpc/gen-status/status.grpc.pb.h"
#include "grpc/gen-status/status.pb.h"
#include "server/grpc_impl/request/GrpcBaseRequest.h"
#include "utils/BlockingQueue.h"
#include "utils/Status.h"

#include <map>
#include <memory>
#include <string>
#include <thread>
#include <vector>

namespace milvus {
namespace server {
namespace grpc {

using RequestQueue = BlockingQueue<BaseRequestPtr>;
using RequestQueuePtr = std::shared_ptr<RequestQueue>;
using ThreadPtr = std::shared_ptr<std::thread>;

class GrpcRequestScheduler {
 public:
    static GrpcRequestScheduler&
    GetInstance() {
        static GrpcRequestScheduler scheduler;
        return scheduler;
    }

    void
    Start();

    void
    Stop();

    Status
    ExecuteRequest(const BaseRequestPtr& request_ptr);

    static void
    ExecRequest(BaseRequestPtr& request_ptr, ::milvus::grpc::Status* grpc_status);

 protected:
    GrpcRequestScheduler();

    virtual ~GrpcRequestScheduler();

    void
    TakeToExecute(RequestQueuePtr request_queue);

    Status
    PutToQueue(const BaseRequestPtr& request_ptr);

 private:
    mutable std::mutex queue_mtx_;

    std::map<std::string, RequestQueuePtr> request_groups_;

    std::vector<ThreadPtr> execute_threads_;

    bool stopped_;
};

}  // namespace grpc
}  // namespace server
}  // namespace milvus
