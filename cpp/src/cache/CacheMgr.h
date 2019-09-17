////////////////////////////////////////////////////////////////////////////////
// Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited.
// Proprietary and confidential.
////////////////////////////////////////////////////////////////////////////////

#pragma once

#include "Cache.h"
#include "utils/Log.h"
#include "metrics/Metrics.h"

namespace zilliz {
namespace milvus {
namespace cache {

template<typename ItemObj>
class CacheMgr {
public:
    virtual uint64_t ItemCount() const;

    virtual bool ItemExists(const std::string& key);

    virtual ItemObj GetItem(const std::string& key);

    virtual void InsertItem(const std::string& key, const ItemObj& data);

    virtual void EraseItem(const std::string& key);

    virtual void PrintInfo();

    virtual void ClearCache();

    int64_t CacheUsage() const;
    int64_t CacheCapacity() const;
    void SetCapacity(int64_t capacity);

protected:
    CacheMgr();
    virtual ~CacheMgr();

protected:
    using CachePtr = std::shared_ptr<Cache<ItemObj>>;
    CachePtr cache_;
};


}
}
}

#include "cache/CacheMgr.inl"