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

#include <map>
#include <memory>
#include <string>
#include <utility>
#include <vector>

#include "Id.h"

namespace knowhere {

struct Binary {
    ID id;
    std::shared_ptr<uint8_t> data;
    int64_t size = 0;
};
using BinaryPtr = std::shared_ptr<Binary>;

class BinarySet {
 public:
    BinaryPtr
    GetByName(const std::string& name) const {
        return binary_map_.at(name);
    }

    void
    Append(const std::string& name, BinaryPtr binary) {
        binary_map_[name] = std::move(binary);
    }

    void
    Append(const std::string& name, std::shared_ptr<uint8_t> data, int64_t size) {
        auto binary = std::make_shared<Binary>();
        binary->data = data;
        binary->size = size;
        binary_map_[name] = std::move(binary);
    }

    // void
    // Append(const std::string &name, void *data, int64_t size, ID id) {
    //    Binary binary;
    //    binary.data = data;
    //    binary.size = size;
    //    binary.id = id;
    //    binary_map_[name] = binary;
    //}

    void
    clear() {
        binary_map_.clear();
    }

 public:
    std::map<std::string, BinaryPtr> binary_map_;
};

}  // namespace knowhere
