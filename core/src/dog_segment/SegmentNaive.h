#pragma once
#include <tbb/concurrent_priority_queue.h>
#include <tbb/concurrent_unordered_map.h>
#include <tbb/concurrent_vector.h>

#include <shared_mutex>

#include "AckResponder.h"
#include "ConcurrentVector.h"
#include "dog_segment/SegmentBase.h"
// #include "knowhere/index/structured_index/StructuredIndex.h"
#include "query/GeneralQuery.h"
#include "utils/Status.h"
using idx_t = int64_t;

namespace milvus::dog_segment {
struct ColumnBasedDataChunk {
    std::vector<std::vector<float>> entity_vecs;
    static ColumnBasedDataChunk
    from(const DogDataChunk& source, const Schema& schema) {
        ColumnBasedDataChunk dest;
        auto count = source.count;
        auto raw_data = reinterpret_cast<const char*>(source.raw_data);
        auto align = source.sizeof_per_row;
        for (auto& field : schema) {
            auto len = field.get_sizeof();
            assert(len % sizeof(float) == 0);
            std::vector<float> new_col(len * count / sizeof(float));
            for (int64_t i = 0; i < count; ++i) {
                memcpy(new_col.data() + i * len / sizeof(float), raw_data + i * align, len);
            }
            dest.entity_vecs.push_back(std::move(new_col));
            // offset the raw_data
            raw_data += len / sizeof(float);
        }
        return dest;
    }
};

class SegmentNaive : public SegmentBase {
 public:
    virtual ~SegmentNaive() = default;

    // SegmentBase(std::shared_ptr<FieldsInfo> collection);

    int64_t PreInsert(int64_t size) override;

    // TODO: originally, id should be put into data_chunk
    // TODO: Is it ok to put them the other side?
    Status
    Insert(int64_t reserverd_offset, int64_t size, const int64_t* primary_keys, const Timestamp* timestamps, const DogDataChunk& values) override;

    int64_t PreDelete(int64_t size) override;

    // TODO: add id into delete log, possibly bitmap
    Status
    Delete(int64_t reserverd_offset, int64_t size, const int64_t* primary_keys, const Timestamp* timestamps) override;

    // query contains metadata of
    Status
    Query(const query::QueryPtr& query, Timestamp timestamp, QueryResult& results) override;

    // stop receive insert requests
    // will move data to immutable vector or something
    Status
    Close() override;

    // using IndexType = knowhere::IndexType;
    // using IndexMode = knowhere::IndexMode;
    // using IndexConfig = knowhere::Config;
    // BuildIndex With Paramaters, must with Frozen State
    // NOTE: index_params contains serveral policies for several index
    // TODO: currently, index has to be set at startup, and can't be modified
    // AddIndex and DropIndex will be added later
    Status
    BuildIndex() override;

    Status
    DropRawData(std::string_view field_name) override {
        // TODO: NO-OP
        return Status::OK();
    }

    Status
    LoadRawData(std::string_view field_name, const char* blob, int64_t blob_size) override {
        // TODO: NO-OP
        return Status::OK();
    }

 private:
    struct MutableRecord {
        ConcurrentVector<uint64_t> uids_;
        tbb::concurrent_vector<Timestamp> timestamps_;
        std::vector<tbb::concurrent_vector<float>> entity_vecs_;
        MutableRecord(int entity_size) : entity_vecs_(entity_size) {
        }
    };

    struct Record {
        std::atomic<int64_t> reserved = 0;
        AckResponder ack_responder_;
        ConcurrentVector<Timestamp, true> timestamps_;
        ConcurrentVector<idx_t, true> uids_;
        std::vector<std::shared_ptr<VectorBase>> entity_vec_;
        Record(const Schema& schema);
    };

    Status
    QueryImpl(const query::QueryPtr& query, Timestamp timestamp, QueryResult& results);

 public:
    ssize_t
    get_row_count() const override {
        return record_.ack_responder_.GetAck();
    }
    SegmentState
    get_state() const override {
        return state_.load(std::memory_order_relaxed);
    }
    ssize_t
    get_deleted_count() const override {
        return 0;
    }

 public:
    friend std::unique_ptr<SegmentBase>
    CreateSegment(SchemaPtr schema, IndexMetaPtr index_meta);
    explicit SegmentNaive(SchemaPtr schema, IndexMetaPtr index_meta)
        : schema_(schema), index_meta_(index_meta), record_(*schema) {
    }

 private:
    SchemaPtr schema_;
    IndexMetaPtr index_meta_;
    std::atomic<SegmentState> state_ = SegmentState::Open;
    Record record_;

    //  tbb::concurrent_unordered_map<uint64_t, int> internal_indexes_;
    //  std::shared_ptr<MutableRecord> record_mutable_;
    //  // to determined that if immutable data if available
    //  std::shared_ptr<ImmutableRecord> record_immutable_ = nullptr;
    //  std::unordered_map<int, knowhere::VecIndexPtr> vec_indexings_;
    //  // TODO: scalar indexing
    //  // std::unordered_map<int, knowhere::IndexPtr> scalar_indexings_;
    //  tbb::concurrent_unordered_multimap<int, Timestamp> delete_logs_;
};
}  // namespace milvus::dog_segment
