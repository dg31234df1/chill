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
#include <vector>

#include "index/VectorIndex.h"
#include "storage/DiskFileManagerImpl.h"
#include "storage/space.h"

namespace milvus::index {

template <typename T>
class VectorDiskAnnIndex : public VectorIndex {
 public:
    explicit VectorDiskAnnIndex(
        const IndexType& index_type,
        const MetricType& metric_type,
        const IndexVersion& version,
        const storage::FileManagerContext& file_manager_context =
            storage::FileManagerContext());

    explicit VectorDiskAnnIndex(
        const IndexType& index_type,
        const MetricType& metric_type,
        const IndexVersion& version,
        std::shared_ptr<milvus_storage::Space> space,
        const storage::FileManagerContext& file_manager_context =
            storage::FileManagerContext());

    BinarySet
    Serialize(const Config& config) override {  // deprecated
        BinarySet binary_set;
        index_.Serialize(binary_set);
        auto remote_paths_to_size = file_manager_->GetRemotePathsToFileSize();
        for (auto& file : remote_paths_to_size) {
            binary_set.Append(file.first, nullptr, file.second);
        }

        return binary_set;
    }

    BinarySet
    Upload(const Config& config = {}) override;

    BinarySet
    UploadV2(const Config& config = {}) override;

    int64_t
    Count() override {
        return index_.Count();
    }

    void
    Load(const BinarySet& binary_set /* not used */,
         const Config& config = {}) override;

    void
    Load(milvus::tracer::TraceContext ctx, const Config& config = {}) override;

    void
    LoadV2(const Config& config = {}) override;

    void
    BuildWithDataset(const DatasetPtr& dataset,
                     const Config& config = {}) override;

    void
    Build(const Config& config = {}) override;

    void
    BuildV2(const Config& config = {}) override;

    std::unique_ptr<SearchResult>
    Query(const DatasetPtr dataset,
          const SearchInfo& search_info,
          const BitsetView& bitset) override;

    const bool
    HasRawData() const override;

    std::vector<uint8_t>
    GetVector(const DatasetPtr dataset) const override;

    void
    CleanLocalData() override;

 private:
    knowhere::Json
    update_load_json(const Config& config);

 private:
    knowhere::Index<knowhere::IndexNode> index_;
    std::shared_ptr<storage::DiskFileManagerImpl> file_manager_;
    uint32_t search_beamwidth_ = 8;
    std::shared_ptr<milvus_storage::Space> space_;
};

template <typename T>
using VectorDiskAnnIndexPtr = std::unique_ptr<VectorDiskAnnIndex<T>>;

}  // namespace milvus::index
