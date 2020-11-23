#pragma once
#include "common/Schema.h"
#include "ConcurrentVector.h"
#include "AckResponder.h"
#include "segcore/Record.h"

namespace milvus::segcore {
struct InsertRecord {
    std::atomic<int64_t> reserved = 0;
    AckResponder ack_responder_;
    ConcurrentVector<Timestamp, true> timestamps_;
    ConcurrentVector<idx_t, true> uids_;
    std::vector<std::shared_ptr<VectorBase>> entity_vec_;

    InsertRecord(const Schema& schema);
    template <typename Type>
    auto
    get_scalar_entity(int offset) const {
        auto ptr = std::dynamic_pointer_cast<const ConcurrentVector<Type, true>>(entity_vec_[offset]);
        Assert(ptr);
        return ptr;
    }

    template <typename Type>
    auto
    get_vec_entity(int offset) const {
        auto ptr = std::dynamic_pointer_cast<const ConcurrentVector<Type>>(entity_vec_[offset]);
        Assert(ptr);
        return ptr;
    }

    template <typename Type>
    auto
    get_scalar_entity(int offset) {
        auto ptr = std::dynamic_pointer_cast<ConcurrentVector<Type, true>>(entity_vec_[offset]);
        Assert(ptr);
        return ptr;
    }

    template <typename Type>
    auto
    get_vec_entity(int offset) {
        auto ptr = std::dynamic_pointer_cast<ConcurrentVector<Type>>(entity_vec_[offset]);
        Assert(ptr);
        return ptr;
    }
};
}  // namespace milvus::segcore
