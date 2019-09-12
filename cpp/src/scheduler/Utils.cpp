/*******************************************************************************
 * Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 ******************************************************************************/

#include "Utils.h"

#include <chrono>
#include <cuda_runtime.h>


namespace zilliz {
namespace milvus {
namespace engine {

uint64_t
get_current_timestamp() {
    std::chrono::time_point<std::chrono::system_clock> now = std::chrono::system_clock::now();
    auto duration = now.time_since_epoch();
    auto millis = std::chrono::duration_cast<std::chrono::milliseconds>(duration).count();
    return millis;
}

uint64_t
get_num_gpu() {
    int n_devices = 0;
    cudaGetDeviceCount(&n_devices);
    return n_devices;
}

}
}
}