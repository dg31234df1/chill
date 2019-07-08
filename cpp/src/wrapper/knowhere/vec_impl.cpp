////////////////////////////////////////////////////////////////////////////////
// Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited.
// Proprietary and confidential.
////////////////////////////////////////////////////////////////////////////////

#include "knowhere/index/vector_index/idmap.h"

#include "vec_impl.h"
#include "data_transfer.h"


namespace zilliz {
namespace milvus {
namespace engine {

using namespace zilliz::knowhere;

void VecIndexImpl::BuildAll(const long &nb,
                            const float *xb,
                            const long *ids,
                            const Config &cfg,
                            const long &nt,
                            const float *xt) {
    dim = cfg["dim"].as<int>();
    auto dataset = GenDatasetWithIds(nb, dim, xb, ids);

    auto preprocessor = index_->BuildPreprocessor(dataset, cfg);
    index_->set_preprocessor(preprocessor);
    auto model = index_->Train(dataset, cfg);
    index_->set_index_model(model);
    index_->Add(dataset, cfg);
}

void VecIndexImpl::Add(const long &nb, const float *xb, const long *ids, const Config &cfg) {
    // TODO(linxj): Assert index is trained;

    auto d = cfg.get_with_default("dim", dim);
    auto dataset = GenDatasetWithIds(nb, d, xb, ids);

    index_->Add(dataset, cfg);
}

void VecIndexImpl::Search(const long &nq, const float *xq, float *dist, long *ids, const Config &cfg) {
    // TODO: Assert index is trained;

    auto k = cfg["k"].as<int>();
    auto d = cfg.get_with_default("dim", dim);
    auto dataset = GenDataset(nq, d, xq);

    Config search_cfg;
    auto res = index_->Search(dataset, cfg);
    auto ids_array = res->array()[0];
    auto dis_array = res->array()[1];

    //{
    //    auto& ids = ids_array;
    //    auto& dists = dis_array;
    //    std::stringstream ss_id;
    //    std::stringstream ss_dist;
    //    for (auto i = 0; i < 10; i++) {
    //        for (auto j = 0; j < k; ++j) {
    //            ss_id << *(ids->data()->GetValues<int64_t>(1, i * k + j)) << " ";
    //            ss_dist << *(dists->data()->GetValues<float>(1, i * k + j)) << " ";
    //        }
    //        ss_id << std::endl;
    //        ss_dist << std::endl;
    //    }
    //    std::cout << "id\n" << ss_id.str() << std::endl;
    //    std::cout << "dist\n" << ss_dist.str() << std::endl;
    //}

    auto p_ids = ids_array->data()->GetValues<int64_t>(1, 0);
    auto p_dist = ids_array->data()->GetValues<float>(1, 0);

    // TODO(linxj): avoid copy here.
    memcpy(ids, p_ids, sizeof(int64_t) * nq * k);
    memcpy(dist, p_dist, sizeof(float) * nq * k);
}

zilliz::knowhere::BinarySet VecIndexImpl::Serialize() {
    return index_->Serialize();
}

void VecIndexImpl::Load(const zilliz::knowhere::BinarySet &index_binary) {
    index_->Load(index_binary);
}

int64_t VecIndexImpl::Dimension() {
    return index_->Dimension();
}

int64_t VecIndexImpl::Count() {
    return index_->Count();
}

float *BFIndex::GetRawVectors() {
    return std::static_pointer_cast<IDMAP>(index_)->GetRawVectors();
}

int64_t *BFIndex::GetRawIds() {
    return std::static_pointer_cast<IDMAP>(index_)->GetRawIds();
}

void BFIndex::Build(const int64_t &d) {
    dim = d;
    std::static_pointer_cast<IDMAP>(index_)->Train(dim);
}

}
}
}
