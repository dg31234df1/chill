#include <iostream>
#include <sstream>
#include <thread>

#include <easylogging++.h>

#include "MemManager.h"
#include "Meta.h"
#include "FaissSerializer.h"


namespace zilliz {
namespace vecwise {
namespace engine {

MemVectors::MemVectors(const std::shared_ptr<meta::Meta>& meta_ptr,
        const meta::GroupFileSchema& schema, const Options& options)
  : pMeta_(meta_ptr),
    options_(options),
    schema_(schema),
    _pIdGenerator(new SimpleIDGenerator()),
    pSerializer_(new FaissSerializer(schema_.dimension, schema_.location)) {
}

void MemVectors::add(size_t n_, const float* vectors_, IDNumbers& vector_ids_) {
    _pIdGenerator->getNextIDNumbers(n_, vector_ids_);
    pSerializer_->AddWithIds(n_, vectors_, vector_ids_.data());
    for(auto i=0 ; i<n_; i++) {
        vector_ids_.push_back(i);
    }
}

size_t MemVectors::total() const {
    return pSerializer_->Count();
}

size_t MemVectors::approximate_size() const {
    return pSerializer_->Size();
}

Status MemVectors::serialize(std::string& group_id) {
    group_id = schema_.group_id;
    auto rows = approximate_size();
    pSerializer_->Serialize();
    schema_.rows = rows;
    schema_.file_type = (rows >= options_.index_trigger_size) ?
        meta::GroupFileSchema::TO_INDEX : meta::GroupFileSchema::RAW;

    auto status = pMeta_->update_group_file(schema_);

    pSerializer_->Cache();

    return status;
}

MemVectors::~MemVectors() {
    if (_pIdGenerator != nullptr) {
        delete _pIdGenerator;
        _pIdGenerator = nullptr;
    }
}

/*
 * MemManager
 */

VectorsPtr MemManager::get_mem_by_group(const std::string& group_id) {
    auto memIt = _memMap.find(group_id);
    if (memIt != _memMap.end()) {
        return memIt->second;
    }

    meta::GroupFileSchema group_file;
    group_file.group_id = group_id;
    auto status = _pMeta->add_group_file(group_file);
    if (!status.ok()) {
        return nullptr;
    }

    _memMap[group_id] = std::shared_ptr<MemVectors>(new MemVectors(_pMeta, group_file, options_));
    return _memMap[group_id];
}

Status MemManager::add_vectors(const std::string& group_id_,
        size_t n_,
        const float* vectors_,
        IDNumbers& vector_ids_) {
    std::unique_lock<std::mutex> lock(_mutex);
    return add_vectors_no_lock(group_id_, n_, vectors_, vector_ids_);
}

Status MemManager::add_vectors_no_lock(const std::string& group_id,
        size_t n,
        const float* vectors,
        IDNumbers& vector_ids) {
    std::shared_ptr<MemVectors> mem = get_mem_by_group(group_id);
    if (mem == nullptr) {
        return Status::NotFound("Group " + group_id + " not found!");
    }
    mem->add(n, vectors, vector_ids);

    return Status::OK();
}

Status MemManager::mark_memory_as_immutable() {
    std::unique_lock<std::mutex> lock(_mutex);
    for (auto& kv: _memMap) {
        _immMems.push_back(kv.second);
    }

    _memMap.clear();
    return Status::OK();
}

/* bool MemManager::need_serialize(double interval) { */
/*     if (_immMems.size() > 0) { */
/*         return false; */
/*     } */

/*     auto diff = std::difftime(std::time(nullptr), _last_compact_time); */
/*     if (diff >= interval) { */
/*         return true; */
/*     } */

/*     return false; */
/* } */

Status MemManager::serialize(std::vector<std::string>& group_ids) {
    mark_memory_as_immutable();
    std::unique_lock<std::mutex> lock(serialization_mtx_);
    std::string group_id;
    group_ids.clear();
    for (auto& mem : _immMems) {
        mem->serialize(group_id);
        group_ids.push_back(group_id);
    }
    _immMems.clear();
    return Status::OK();
}


} // namespace engine
} // namespace vecwise
} // namespace zilliz
