#include <gtest/gtest.h>
#include "query/deprecated/Parser.h"
#include "query/Expr.h"
#include "query/PlanNode.h"
#include "query/generated/ExprVisitor.h"
#include "query/generated/PlanNodeVisitor.h"
#include "test_utils/DataGen.h"
#include "query/generated/ShowPlanNodeVisitor.h"
#include "query/Plan.h"
using namespace milvus;

TEST(Expr, Naive) {
    SUCCEED();
    using namespace milvus::wtf;
    std::string dsl_string = R"(
{
    "bool": {
        "must": [
            {
                "term": {
                    "A": [
                        1,
                        2,
                        5
                    ]
                }
            },
            {
                "range": {
                    "B": {
                        "GT": 1,
                        "LT": 100
                    }
                }
            },
            {
                "vector": {
                    "Vec": {
                        "metric_type": "L2",
                        "params": {
                            "nprobe": 10
                        },
                        "query": "$0",
                        "topk": 10
                    }
                }
            }
        ]
    }
})";
}

TEST(Expr, Range) {
    SUCCEED();
    using namespace milvus;
    using namespace milvus::query;
    using namespace milvus::segcore;
    std::string dsl_string = R"(
{
    "bool": {
        "must": [
            {
                "range": {
                    "age": {
                        "GT": 1,
                        "LT": 100
                    }
                }
            },
            {
                "vector": {
                    "fakevec": {
                        "metric_type": "L2",
                        "params": {
                            "nprobe": 10
                        },
                        "query": "$0",
                        "topk": 10
                    }
                }
            }
        ]
    }
})";
    auto schema = std::make_shared<Schema>();
    schema->AddField("fakevec", DataType::VECTOR_FLOAT, 16);
    schema->AddField("age", DataType::INT32);
    auto plan = CreatePlan(*schema, dsl_string);
    ShowPlanNodeVisitor shower;
    Assert(plan->tag2field_.at("$0") == "fakevec");
    auto out = shower.call_child(*plan->plan_node_);
    std::cout << out.dump(4);
}

TEST(Expr, ShowExecutor) {
    using namespace milvus::query;
    using namespace milvus::segcore;
    auto node = std::make_unique<FloatVectorANNS>();
    auto schema = std::make_shared<Schema>();
    schema->AddField("fakevec", DataType::VECTOR_FLOAT, 16);
    int64_t num_queries = 100L;
    auto raw_data = DataGen(schema, num_queries);
    auto& info = node->query_info_;
    info.metric_type_ = "L2";
    info.topK_ = 20;
    info.field_id_ = "fakevec";
    node->predicate_ = std::nullopt;
    ShowPlanNodeVisitor show_visitor;
    PlanNodePtr base(node.release());
    auto res = show_visitor.call_child(*base);
    auto dup = res;
    dup["data"] = "...collased...";
    std::cout << dup.dump(4);
}
