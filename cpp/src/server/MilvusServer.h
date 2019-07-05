/*******************************************************************************
 * Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 ******************************************************************************/
#pragma once

#include <cstdint>
#include <string>

namespace zilliz {
namespace milvus {
namespace server {

class MilvusServer {
public:
    static void StartService();
    static void StopService();
};

}
}
}
