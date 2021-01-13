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

#include "utils/Json.h"
#include "query/PlanImpl.h"
#include "segcore/SegmentBase.h"
#include <utility>
#include "query/generated/ExecPlanNodeVisitor.h"
#include "segcore/SegmentSmallIndex.h"
#include "query/generated/ExecExprVisitor.h"
#include "query/Search.h"
#include "query/SearchOnSealed.h"

namespace milvus::query {

#if 1
namespace impl {
// THIS CONTAINS EXTRA BODY FOR VISITOR
// WILL BE USED BY GENERATOR UNDER suvlim/core_gen/
class ExecPlanNodeVisitor : PlanNodeVisitor {
 public:
    using RetType = QueryResult;
    ExecPlanNodeVisitor(segcore::SegmentBase& segment, Timestamp timestamp, const PlaceholderGroup& placeholder_group)
        : segment_(segment), timestamp_(timestamp), placeholder_group_(placeholder_group) {
    }
    // using RetType = nlohmann::json;

    RetType
    get_moved_result(PlanNode& node) {
        assert(!ret_.has_value());
        node.accept(*this);
        assert(ret_.has_value());
        auto ret = std::move(ret_).value();
        ret_ = std::nullopt;
        return ret;
    }

 private:
    // std::optional<RetType> ret_;
    segcore::SegmentBase& segment_;
    Timestamp timestamp_;
    const PlaceholderGroup& placeholder_group_;

    std::optional<RetType> ret_;
};
}  // namespace impl
#endif

void
ExecPlanNodeVisitor::visit(FloatVectorANNS& node) {
    // TODO: optimize here, remove the dynamic cast
    assert(!ret_.has_value());
    auto segment = dynamic_cast<segcore::SegmentSmallIndex*>(&segment_);
    AssertInfo(segment, "support SegmentSmallIndex Only");
    RetType ret;
    auto& ph = placeholder_group_.at(0);
    auto src_data = ph.get_blob<float>();
    auto num_queries = ph.num_of_queries_;

    ExecExprVisitor::RetType bitmap_holder;
    std::optional<const ExecExprVisitor::RetType*> bitset_pack;

    if (node.predicate_.has_value()) {
        bitmap_holder = ExecExprVisitor(*segment).call_child(*node.predicate_.value());
        bitset_pack = &bitmap_holder;
    }

    auto& sealed_indexing = segment->get_sealed_indexing_record();

    if (sealed_indexing.is_ready(node.query_info_.field_offset_)) {
        SearchOnSealed(segment->get_schema(), sealed_indexing, node.query_info_, src_data, num_queries, timestamp_,
                       bitset_pack, ret);
    } else {
        FloatSearch(*segment, node.query_info_, src_data, num_queries, timestamp_, bitset_pack, ret);
    }

    ret_ = ret;
}

void
ExecPlanNodeVisitor::visit(BinaryVectorANNS& node) {
    // TODO: optimize here, remove the dynamic cast
    assert(!ret_.has_value());
    auto segment = dynamic_cast<segcore::SegmentSmallIndex*>(&segment_);
    AssertInfo(segment, "support SegmentSmallIndex Only");
    RetType ret;
    auto& ph = placeholder_group_.at(0);
    auto src_data = ph.get_blob<uint8_t>();
    auto num_queries = ph.num_of_queries_;

    ExecExprVisitor::RetType bitmap_holder;
    std::optional<const ExecExprVisitor::RetType*> bitset_pack;

    if (node.predicate_.has_value()) {
        bitmap_holder = ExecExprVisitor(*segment).call_child(*node.predicate_.value());
        bitset_pack = &bitmap_holder;
    }

    auto& sealed_indexing = segment->get_sealed_indexing_record();
    if (sealed_indexing.is_ready(node.query_info_.field_offset_)) {
        SearchOnSealed(segment->get_schema(), sealed_indexing, node.query_info_, src_data, num_queries, timestamp_,
                       bitset_pack, ret);
    } else {
        BinarySearch(*segment, node.query_info_, src_data, num_queries, timestamp_, bitset_pack, ret);
    }
    ret_ = ret;
}

}  // namespace milvus::query
