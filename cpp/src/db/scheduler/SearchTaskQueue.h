/*******************************************************************************
 * Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 ******************************************************************************/
#pragma once

#include "SearchContext.h"
#include "utils/BlockingQueue.h"
#include "db/ExecutionEngine.h"

#include <memory>

namespace zilliz {
namespace vecwise {
namespace engine {

class SearchTask {
public:
    bool DoSearch();

public:
    size_t index_id_ = 0;
    int index_type_ = 0; //for metrics
    ExecutionEnginePtr index_engine_;
    std::vector<SearchContextPtr> search_contexts_;
};

using SearchTaskPtr = std::shared_ptr<SearchTask>;
using SearchTaskQueue = server::BlockingQueue<SearchTaskPtr>;


}
}
}