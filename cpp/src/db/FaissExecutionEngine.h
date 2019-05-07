/*******************************************************************************
 * Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 ******************************************************************************/
#pragma once

#include <memory>
#include <string>

#include "ExecutionEngine.h"

namespace faiss {
    class Index;
}

namespace zilliz {
namespace vecwise {
namespace engine {


template<class IndexTrait>
class FaissExecutionEngine : public ExecutionEngine<FaissExecutionEngine<IndexTrait>> {
public:
    typedef std::shared_ptr<FaissExecutionEngine<IndexTrait>> Ptr;

    FaissExecutionEngine(uint16_t dimension, const std::string& location);
    FaissExecutionEngine(std::shared_ptr<faiss::Index> index, const std::string& location);

    Status AddWithIds(const std::vector<float>& vectors,
                              const std::vector<long>& vector_ids);

    Status AddWithIds(long n, const float *xdata, const long *xids);

    size_t Count() const;

    size_t Size() const;

    size_t PhysicalSize() const;

    Status Serialize();

    Status Load();

    Status Merge(const std::string& location);

    Status Search(long n,
                  const float *data,
                  long k,
                  float *distances,
                  long *labels) const;

    Ptr BuildIndex(const std::string&);

    Status Cache();
protected:
    std::shared_ptr<faiss::Index> pIndex_;
    std::string location_;
};


} // namespace engine
} // namespace vecwise
} // namespace zilliz

#include "FaissExecutionEngine.cpp"
