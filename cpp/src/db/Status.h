/*******************************************************************************
 * Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 ******************************************************************************/
#pragma once

#include <string>


namespace zilliz {
namespace milvus {
namespace engine {

class Status {
 public:
    Status() noexcept : state_(nullptr) {}

    ~Status() { delete[] state_; }

    Status(const Status &rhs);

    Status &
    operator=(const Status &rhs);

    Status(Status &&rhs) noexcept : state_(rhs.state_) { rhs.state_ = nullptr; }

    Status &
    operator=(Status &&rhs_) noexcept;

    static Status
    OK() { return Status(); }

    static Status
    NotFound(const std::string &msg, const std::string &msg2 = "") {
        return Status(kNotFound, msg, msg2);
    }
    static Status
    Error(const std::string &msg, const std::string &msg2 = "") {
        return Status(kError, msg, msg2);
    }

    static Status
    InvalidDBPath(const std::string &msg, const std::string &msg2 = "") {
        return Status(kInvalidDBPath, msg, msg2);
    }
    static Status
    GroupError(const std::string &msg, const std::string &msg2 = "") {
        return Status(kGroupError, msg, msg2);
    }
    static Status
    DBTransactionError(const std::string &msg, const std::string &msg2 = "") {
        return Status(kDBTransactionError, msg, msg2);
    }

    static Status
    AlreadyExist(const std::string &msg, const std::string &msg2 = "") {
        return Status(kAlreadyExist, msg, msg2);
    }

    bool ok() const { return state_ == nullptr; }

    bool IsNotFound() const { return code() == kNotFound; }
    bool IsError() const { return code() == kError; }

    bool IsInvalidDBPath() const { return code() == kInvalidDBPath; }
    bool IsGroupError() const { return code() == kGroupError; }
    bool IsDBTransactionError() const { return code() == kDBTransactionError; }
    bool IsAlreadyExist() const { return code() == kAlreadyExist; }

    std::string ToString() const;

 private:
    const char *state_ = nullptr;

    enum Code {
        kOK = 0,
        kNotFound,
        kError,

        kInvalidDBPath,
        kGroupError,
        kDBTransactionError,

        kAlreadyExist,
    };

    Code code() const {
        return (state_ == nullptr) ? kOK : static_cast<Code>(state_[4]);
    }
    Status(Code code, const std::string &msg, const std::string &msg2);
    static const char *CopyState(const char *s);

}; // Status

inline Status::Status(const Status &rhs) {
    state_ = (rhs.state_ == nullptr) ? nullptr : CopyState(rhs.state_);
}

inline Status &Status::operator=(const Status &rhs) {
    if (state_ != rhs.state_) {
        delete[] state_;
        state_ = (rhs.state_ == nullptr) ? nullptr : CopyState(rhs.state_);
    }
    return *this;
}

inline Status &Status::operator=(Status &&rhs) noexcept {
    std::swap(state_, rhs.state_);
    return *this;
}

} // namespace engine
} // namespace milvus
} // namespace zilliz
