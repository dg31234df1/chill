#pragma once
#include "AckResponder.h"
#include <tbb/concurrent_vector.h>
#include "common/Schema.h"
#include <optional>
#include "InsertRecord.h"
#include <knowhere/index/vector_index/IndexIVF.h>
#include <knowhere/index/structured_index_simple/StructuredIndexSort.h>

namespace milvus::segcore {

// this should be concurrent
// All concurrent
class IndexingEntry {
 public:
    explicit IndexingEntry(const FieldMeta& field_meta) : field_meta_(field_meta) {
    }
    IndexingEntry(const IndexingEntry&) = delete;
    IndexingEntry&
    operator=(const IndexingEntry&) = delete;

    // Do this in parallel
    virtual void
    BuildIndexRange(int64_t ack_beg, int64_t ack_end, const VectorBase* vec_base) = 0;

    const FieldMeta&
    get_field_meta() {
        return field_meta_;
    }

 protected:
    // additional info
    const FieldMeta& field_meta_;
};
template <typename T>
class ScalarIndexingEntry : public IndexingEntry {
 public:
    using IndexingEntry::IndexingEntry;

    void
    BuildIndexRange(int64_t ack_beg, int64_t ack_end, const VectorBase* vec_base) override;

    // concurrent
    knowhere::scalar::StructuredIndex<T>*
    get_indexing(int64_t chunk_id) const {
        Assert(!field_meta_.is_vector());
        return data_.at(chunk_id).get();
    }

 private:
    tbb::concurrent_vector<std::unique_ptr<knowhere::scalar::StructuredIndex<T>>> data_;
};

class VecIndexingEntry : public IndexingEntry {
 public:
    using IndexingEntry::IndexingEntry;

    void
    BuildIndexRange(int64_t ack_beg, int64_t ack_end, const VectorBase* vec_base) override;

    // concurrent
    knowhere::VecIndex*
    get_vec_indexing(int64_t chunk_id) const {
        Assert(field_meta_.is_vector());
        return data_.at(chunk_id).get();
    }

    knowhere::Config
    get_build_conf() const;
    knowhere::Config
    get_search_conf(int top_k) const;

 private:
    tbb::concurrent_vector<std::unique_ptr<knowhere::VecIndex>> data_;
};

std::unique_ptr<IndexingEntry>
CreateIndex(const FieldMeta& field_meta);

class IndexingRecord {
 public:
    explicit IndexingRecord(const Schema& schema) : schema_(schema) {
        Initialize();
    }

    void
    Initialize() {
        int offset = 0;
        for (auto& field : schema_) {
            entries_.try_emplace(offset, CreateIndex(field));
            ++offset;
        }
        assert(offset == schema_.size());
    }

    // concurrent, reentrant
    void
    UpdateResourceAck(int64_t chunk_ack, const InsertRecord& record);

    // concurrent
    int64_t
    get_finished_ack() const {
        return finished_ack_.GetAck();
    }

    const IndexingEntry&
    get_entry(int field_offset) const {
        assert(entries_.count(field_offset));
        return *entries_.at(field_offset);
    }

    const VecIndexingEntry&
    get_vec_entry(int field_offset) const {
        auto& entry = get_entry(field_offset);
        auto ptr = dynamic_cast<const VecIndexingEntry*>(&entry);
        AssertInfo(ptr, "invalid indexing");
        return *ptr;
    }
    template <typename T>
    auto
    get_scalar_entry(int field_offset) const -> const ScalarIndexingEntry<T>& {
        auto& entry = get_entry(field_offset);
        auto ptr = dynamic_cast<const ScalarIndexingEntry<T>*>(&entry);
        AssertInfo(ptr, "invalid indexing");
        return *ptr;
    }

 private:
    const Schema& schema_;

 private:
    // control info
    std::atomic<int64_t> resource_ack_ = 0;
    //    std::atomic<int64_t> finished_ack_ = 0;
    AckResponder finished_ack_;
    std::mutex mutex_;

 private:
    // field_offset => indexing
    std::map<int, std::unique_ptr<IndexingEntry>> entries_;
};

}  // namespace milvus::segcore