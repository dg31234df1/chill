////////////////////////////////////////////////////////////////////////////////
// Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited.
// Proprietary and confidential.
////////////////////////////////////////////////////////////////////////////////

#include "knowhere/index/vector_index/nsg_index.h"
#include "knowhere/index/vector_index/nsg/nsg.h"
#include "knowhere/index/vector_index/nsg/nsg_io.h"
#include "knowhere/index/vector_index/idmap.h"
#include "knowhere/index/vector_index/ivf.h"
#include "knowhere/index/vector_index/gpu_ivf.h"
#include "knowhere/adapter/faiss_adopt.h"
#include "knowhere/common/exception.h"
#include "knowhere/common/timer.h"


namespace zilliz {
namespace knowhere {

BinarySet NSG::Serialize() {
    if (!index_ || !index_->is_trained) {
        KNOWHERE_THROW_MSG("index not initialize or trained");
    }

    try {
        algo::NsgIndex *index = index_.get();

        MemoryIOWriter writer;
        algo::write_index(index, writer);
        auto data = std::make_shared<uint8_t>();
        data.reset(writer.data_);

        BinarySet res_set;
        res_set.Append("NSG", data, writer.total);
        return res_set;
    } catch (std::exception &e) {
        KNOWHERE_THROW_MSG(e.what());
    }
}

void NSG::Load(const BinarySet &index_binary) {
    try {
        auto binary = index_binary.GetByName("NSG");

        MemoryIOReader reader;
        reader.total = binary->size;
        reader.data_ = binary->data.get();

        auto index = algo::read_index(reader);
        index_.reset(index);
    } catch (std::exception &e) {
        KNOWHERE_THROW_MSG(e.what());
    }
}

DatasetPtr NSG::Search(const DatasetPtr &dataset, const Config &config) {
    if (!index_ || !index_->is_trained) {
        KNOWHERE_THROW_MSG("index not initialize or trained");
    }

    // Required
    // if not found throw exception here.
    auto k = config["k"].as<size_t>();
    auto search_length = config.get_with_default("search_length", 30);

    GETTENSOR(dataset)

    auto elems = rows * k;
    auto res_ids = (int64_t *) malloc(sizeof(int64_t) * elems);
    auto res_dis = (float *) malloc(sizeof(float) * elems);

    // TODO(linxj): get from config
    algo::SearchParams s_params;
    s_params.search_length = search_length;
    index_->Search((float *) p_data, rows, dim, k, res_dis, res_ids, s_params);

    auto id_buf = MakeMutableBufferSmart((uint8_t *) res_ids, sizeof(int64_t) * elems);
    auto dist_buf = MakeMutableBufferSmart((uint8_t *) res_dis, sizeof(float) * elems);

    // TODO: magic
    std::vector<BufferPtr> id_bufs{nullptr, id_buf};
    std::vector<BufferPtr> dist_bufs{nullptr, dist_buf};

    auto int64_type = std::make_shared<arrow::Int64Type>();
    auto float_type = std::make_shared<arrow::FloatType>();

    auto id_array_data = arrow::ArrayData::Make(int64_type, elems, id_bufs);
    auto dist_array_data = arrow::ArrayData::Make(float_type, elems, dist_bufs);

    auto ids = std::make_shared<NumericArray<arrow::Int64Type>>(id_array_data);
    auto dists = std::make_shared<NumericArray<arrow::FloatType>>(dist_array_data);
    std::vector<ArrayPtr> array{ids, dists};

    return std::make_shared<Dataset>(array, nullptr);
}

IndexModelPtr NSG::Train(const DatasetPtr &dataset, const Config &config) {
    TimeRecorder rc("Interface");

    auto metric_type = config["metric_type"].as_string();
    if (metric_type != "L2") { KNOWHERE_THROW_MSG("NSG not support this kind of metric type");}

    // TODO(linxj): dev IndexFactory, support more IndexType
    auto preprocess_index = std::make_shared<GPUIVF>(0);
    //auto preprocess_index = std::make_shared<IVF>();
    auto model = preprocess_index->Train(dataset, config);
    preprocess_index->set_index_model(model);
    preprocess_index->AddWithoutIds(dataset, config);
    rc.RecordSection("build ivf");

    auto k = config["knng"].as<int64_t>();
    Graph knng;
    preprocess_index->GenGraph(k, knng, dataset, config);
    rc.RecordSection("build knng");

    GETTENSOR(dataset)
    auto array = dataset->array()[0];
    auto p_ids = array->data()->GetValues<long>(1, 0);

    algo::BuildParams b_params;
    b_params.candidate_pool_size = config["candidate_pool_size"].as<size_t>();
    b_params.out_degree = config["out_degree"].as<size_t>();
    b_params.search_length = config["search_length"].as<size_t>();

    index_ = std::make_shared<algo::NsgIndex>(dim, rows);
    index_->SetKnnGraph(knng);
    index_->Build_with_ids(rows, (float *) p_data, (long *) p_ids, b_params);
    rc.RecordSection("build nsg");
    rc.ElapseFromBegin("total cost");
    return nullptr; // TODO(linxj): support serialize
}

void NSG::Add(const DatasetPtr &dataset, const Config &config) {
    // TODO(linxj): support incremental index.

    //KNOWHERE_THROW_MSG("Not support yet");
}

int64_t NSG::Count() {
    return index_->ntotal;
}

int64_t NSG::Dimension() {
    return index_->dimension;
}
VectorIndexPtr NSG::Clone() {
    KNOWHERE_THROW_MSG("not support");
}

void NSG::Seal() {
    // do nothing
}

}
}

