/*******************************************************************************
 * Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 ******************************************************************************/
#pragma once

#include <string>

#include "Resource.h"


namespace zilliz {
namespace milvus {
namespace engine {

class CpuResource : public Resource {
public:
    explicit
    CpuResource(std::string name)
        : Resource(std::move(name), ResourceType::CPU) {}

protected:
    void
    LoadFile(TaskPtr task) override {
//        if (src.type == DISK) {
//            fd = open(filename);
//            content = fd.read();
//            close(fd);
//        } else if (src.type == CPU) {
//            memcpy(src, dest, len);
//        } else if (src.type == GPU) {
//            cudaMemcpyD2H(src, dest);
//        } else {
//            // unknown type, exception
//        }
    }

    void
    Process(TaskPtr task) override {
        task->Execute();
    }
};

}
}
}
