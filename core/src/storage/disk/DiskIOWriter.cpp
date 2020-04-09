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

#include "storage/disk/DiskIOWriter.h"

namespace milvus {
namespace storage {

bool
DiskIOWriter::open(const std::string& name) {
    name_ = name;
    len_ = 0;
    fs_ = std::fstream(name_, std::ios::out | std::ios::binary);
    return fs_.good();
}

void
DiskIOWriter::write(void* ptr, int64_t size) {
    fs_.write(reinterpret_cast<char*>(ptr), size);
    len_ += size;
}

int64_t
DiskIOWriter::length() {
    return len_;
}

void
DiskIOWriter::close() {
    fs_.close();
}

}  // namespace storage
}  // namespace milvus
