////////////////////////////////////////////////////////////////////////////////
// Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited.
// Proprietary and confidential.
////////////////////////////////////////////////////////////////////////////////
#pragma once

#include "Types.h"

#include <cstddef>
#include <vector>


namespace zilliz {
namespace milvus {
namespace engine {

class IDGenerator {
 public:
    virtual IDNumber GetNextIDNumber() = 0;
    virtual void GetNextIDNumbers(size_t n, IDNumbers &ids) = 0;
    virtual ~IDGenerator() = 0;
}; // IDGenerator


class SimpleIDGenerator : public IDGenerator {
 public:
    ~SimpleIDGenerator() override = default;

    IDNumber
    GetNextIDNumber() override;

    void
    GetNextIDNumbers(size_t n, IDNumbers &ids) override;

 private:
    void
    NextIDNumbers(size_t n, IDNumbers &ids);

    static constexpr size_t MAX_IDS_PER_MICRO = 1000;

}; // SimpleIDGenerator


} // namespace engine
} // namespace milvus
} // namespace zilliz
