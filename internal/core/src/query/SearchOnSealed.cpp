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

//
// Created by mike on 12/26/20.
//

#include "query/SearchOnSealed.h"
#include <knowhere/index/vector_index/VecIndex.h>
#include "knowhere/index/vector_index/helpers/IndexParameter.h"
#include "knowhere/index/vector_index/adapter/VectorAdapter.h"
#include <boost_ext/dynamic_bitset_ext.hpp>

namespace milvus::query {

// negate bitset, and merge them into one
aligned_vector<uint8_t>
AssembleNegBitset(const BitsetSimple& bitset_simple) {
    int64_t N = 0;

    for (auto& bitset : bitset_simple) {
        N += bitset.size();
    }
    aligned_vector<uint8_t> result(upper_align(upper_div(N, 8), 64));

    auto acc_byte_count = 0;
    for (auto& bitset : bitset_simple) {
        auto size = bitset.size();
        Assert(size % 8 == 0);
        auto byte_count = size / 8;
        auto src_ptr = boost_ext::get_data(bitset);
        memcpy(result.data() + acc_byte_count, src_ptr, byte_count);
        acc_byte_count += byte_count;
    }

    // revert the bitset
    for (int64_t i = 0; i < result.size(); ++i) {
        result[i] = ~result[i];
    }
    return result;
}

void
SearchOnSealed(const Schema& schema,
               const segcore::SealedIndexingRecord& record,
               const QueryInfo& query_info,
               const void* query_data,
               int64_t num_queries,
               const faiss::BitsetView& bitset,
               QueryResult& result) {
    auto topK = query_info.topK_;

    auto field_offset = query_info.field_offset_;
    auto& field = schema[field_offset];
    // Assert(field.get_data_type() == DataType::VECTOR_FLOAT);
    auto dim = field.get_dim();

    Assert(record.is_ready(field_offset));
    auto field_indexing = record.get_field_indexing(field_offset);
    Assert(field_indexing->metric_type_ == query_info.metric_type_);

    auto final = [&] {
        auto ds = knowhere::GenDataset(num_queries, dim, query_data);

        auto conf = query_info.search_params_;
        conf[milvus::knowhere::meta::TOPK] = query_info.topK_;
        conf[milvus::knowhere::Metric::TYPE] = MetricTypeToName(field_indexing->metric_type_);
        return field_indexing->indexing_->Query(ds, conf, bitset);
    }();

    auto ids = final->Get<idx_t*>(knowhere::meta::IDS);
    auto distances = final->Get<float*>(knowhere::meta::DISTANCE);

    auto total_num = num_queries * topK;
    result.internal_seg_offsets_.resize(total_num);
    result.result_distances_.resize(total_num);
    result.num_queries_ = num_queries;
    result.topK_ = topK;

    std::copy_n(ids, total_num, result.internal_seg_offsets_.data());
    std::copy_n(distances, total_num, result.result_distances_.data());
}
}  // namespace milvus::query
