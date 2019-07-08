////////////////////////////////////////////////////////////////////////////////
// Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited.
// Proprietary and confidential.
////////////////////////////////////////////////////////////////////////////////

#pragma once

#include "knowhere/index/vector_index/vector_index.h"

#include "vec_index.h"


namespace zilliz {
namespace milvus {
namespace engine {

class VecIndexImpl : public VecIndex {
 public:
    explicit VecIndexImpl(std::shared_ptr<zilliz::knowhere::VectorIndex> index) : index_(std::move(index)) {};
    void BuildAll(const long &nb,
                  const float *xb,
                  const long *ids,
                  const Config &cfg,
                  const long &nt,
                  const float *xt) override;
    int64_t Dimension() override;
    int64_t Count() override;
    void Add(const long &nb, const float *xb, const long *ids, const Config &cfg) override;
    zilliz::knowhere::BinarySet Serialize() override;
    void Load(const zilliz::knowhere::BinarySet &index_binary) override;
    void Search(const long &nq, const float *xq, float *dist, long *ids, const Config &cfg) override;

 protected:
    int64_t dim = 0;
    std::shared_ptr<zilliz::knowhere::VectorIndex> index_ = nullptr;
};

class BFIndex : public VecIndexImpl {
 public:
    explicit BFIndex(std::shared_ptr<zilliz::knowhere::VectorIndex> index) : VecIndexImpl(std::move(index)) {};
    void Build(const int64_t& d);
    float* GetRawVectors();
    void BuildAll(const long &nb,
                  const float *xb,
                  const long *ids,
                  const Config &cfg,
                  const long &nt,
                  const float *xt) override;
    int64_t* GetRawIds();
};

}
}
}
