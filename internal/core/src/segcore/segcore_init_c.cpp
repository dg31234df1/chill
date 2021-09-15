// Copyright (C) 2019-2020 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License

#include <iostream>

#include "exceptions/EasyAssert.h"
#include "knowhere/archive/KnowhereConfig.h"
#include "segcore/segcore_init_c.h"
#include "segcore/SegcoreConfig.h"
#include "utils/Log.h"

namespace milvus::segcore {
static void
SegcoreInitImpl() {
    namespace eg = milvus::engine;
    eg::KnowhereConfig::SetSimdType(eg::KnowhereConfig::SimdType::AUTO);
    eg::KnowhereConfig::SetBlasThreshold(16384);
    eg::KnowhereConfig::SetEarlyStopThreshold(0);
    eg::KnowhereConfig::SetLogHandler();
    eg::KnowhereConfig::SetStatisticsLevel(0);
    el::Configurations el_conf;
    el_conf.setGlobally(el::ConfigurationType::Enabled, std::to_string(false));
}
}  // namespace milvus::segcore

extern "C" void
SegcoreInit() {
    milvus::segcore::SegcoreInitImpl();
}

extern "C" void
SegcoreSetChunkSize(const int64_t value) {
    milvus::segcore::SegcoreConfig& config = milvus::segcore::SegcoreConfig::default_config();
    config.set_size_per_chunk(value);
    std::cout << "set config chunk_size: " << config.get_size_per_chunk() << std::endl;
}
