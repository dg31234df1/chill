#ifndef STORAGE_VECENGINE_MEMVECTORS_H_
#define STORAGE_VECENGINE_MEMVECTORS_H_

#include <map>
#include <string>
#include "id_generators.h"
#include "status.h"

class faiss::IndexIDMap;
class faiss::Index;


namespace zilliz {
namespace vecwise {
namespace engine {

class MemVectors {
public:
    explicit MemVectors(size_t dimension_, const std::string& file_location_);

    IDNumbers&& add(size_t n, const float* vectors);

    size_t total() const;

    size_t approximate_size() const;

    void serialize();

    ~MemVectors();

private:
    std::string _file_location;
    IDGenerator* _pIdGenerator;
    size_t _dimension;
    faiss::Index* _pInnerIndex;
    faiss::IndexIDMap* _pIdMapIndex;

}; // MemVectors


class Meta;

class MemManager {
public:
    MemManager(const std::shared_ptr<Meta>& meta_) : _pMeta(meta_) {}

    MemVectors* get_mem_by_group(const std::string& group_id_);

    Status add_vectors(const std::string& group_id_,
            size_t n_, const float* vectors_, IDNumbers& vector_ids_);

private:
    Status add_vectors_no_lock(const std::string& group_id_,
            size_t n_, const float* vectors_, IDNumbers& vector_ids_);

    typedef std::map<std::string, MemVectors> MemMap;
    MemMap _memMap;
    std::shared_ptr<Meta> _pMeta;
}; // MemManager


} // namespace engine
} // namespace vecwise
} // namespace zilliz

#endif
