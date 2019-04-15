#ifndef STORAGE_VECENGINE_MEMVECTORS_H_
#define STORAGE_VECENGINE_MEMVECTORS_H_

#include <map>
#include <string>
#include <ctime>
#include <memory>
#include <mutex>
#include "id_generators.h"
#include "status.h"

namespace faiss {
    class IndexIDMap;
    class Index;
}


namespace zilliz {
namespace vecwise {
namespace engine {

class MemVectors {
public:
    explicit MemVectors(size_t dimension_, const std::string& file_location_);

    void add(size_t n_, const float* vectors_, IDNumbers& vector_ids_);

    size_t total() const;

    size_t approximate_size() const;

    void serialize();

    ~MemVectors();

    const std::string& location() const { return _file_location; }

private:
    MemVectors() = delete;
    MemVectors(const MemVectors&) = delete;
    MemVectors& operator=(const MemVectors&) = delete;

    const std::string _file_location;
    IDGenerator* _pIdGenerator;
    size_t _dimension;
    faiss::Index* _pInnerIndex;
    faiss::IndexIDMap* _pIdMapIndex;

}; // MemVectors


class Meta;
typedef std::shared_ptr<MemVectors> VectorsPtr;

class MemManager {
public:
    MemManager(const std::shared_ptr<Meta>& meta_)
        : _pMeta(meta_) /*_last_compact_time(std::time(nullptr))*/ {}

    VectorsPtr get_mem_by_group(const std::string& group_id_);

    Status add_vectors(const std::string& group_id_,
            size_t n_, const float* vectors_, IDNumbers& vector_ids_);

    Status serialize();

private:
    Status add_vectors_no_lock(const std::string& group_id_,
            size_t n_, const float* vectors_, IDNumbers& vector_ids_);
    Status mark_memory_as_immutable();

    typedef std::map<std::string, VectorsPtr> MemMap;
    typedef std::vector<VectorsPtr> ImmMemPool;
    MemMap _memMap;
    ImmMemPool _immMems;
    std::shared_ptr<Meta> _pMeta;
    /* std::time_t _last_compact_time; */
    std::mutex _mutex;
}; // MemManager


} // namespace engine
} // namespace vecwise
} // namespace zilliz

#endif
