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

#pragma once

#include <memory>
#include <string>
#include "storage/IOWriter.h"

namespace milvus {
namespace storage {

class S3IOWriter : public IOWriter {
 public:
    S3IOWriter() = default;
    ~S3IOWriter() = default;

    // No copy and move
    S3IOWriter(const S3IOWriter&) = delete;
    S3IOWriter(S3IOWriter&&) = delete;

    S3IOWriter&
    operator=(const S3IOWriter&) = delete;
    S3IOWriter&
    operator=(S3IOWriter&&) = delete;

    bool
    open(const std::string& name) override;

    void
    write(void* ptr, int64_t size) override;

    int64_t
    length() override;

    void
    close() override;

 public:
    std::string name_;
    int64_t len_;
    std::string buffer_;
};

using S3IOWriterPtr = std::shared_ptr<S3IOWriter>;

}  // namespace storage
}  // namespace milvus
