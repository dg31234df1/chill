// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#pragma once

#include <memory>
#include <boost/dynamic_bitset.hpp>
#include "knowhere/dataset.h"
#include "common/Types.h"

namespace milvus::index {

class IndexBase {
 public:
    virtual ~IndexBase() = default;

    virtual BinarySet
    Serialize(const Config& config) = 0;

    virtual void
    Load(const BinarySet& binary_set, const Config& config = {}) = 0;

    virtual void
    BuildWithRawData(size_t n, const void* values, const Config& config = {}) = 0;

    virtual void
    BuildWithDataset(const DatasetPtr& dataset, const Config& config = {}) = 0;

    virtual int64_t
    Count() = 0;

 protected:
    IndexType index_type_ = "";
    IndexMode index_mode_ = IndexMode::MODE_CPU;
};

using IndexBasePtr = std::unique_ptr<IndexBase>;
}  // namespace milvus::index
