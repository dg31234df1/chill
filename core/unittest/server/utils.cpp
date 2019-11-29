// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

#include "server/utils.h"
#include "utils/CommonUtil.h"

#include <fstream>
#include <iostream>
#include <string>
#include <thread>

namespace {

static const char* VALID_CONFIG_STR =
    "# Default values are used when you make no changes to the following parameters.\n"
    "\n"
    "version: 0.1"
    "\n"
    "server_config:\n"
    "  address: 0.0.0.0                  # milvus server ip address (IPv4)\n"
    "  port: 19530                       # port range: 1025 ~ 65534\n"
    "  deploy_mode: single               \n"
    "  time_zone: UTC+8\n"
    "\n"
    "db_config:\n"
    "  primary_path: /tmp/milvus    # path used to store data and meta\n"
    "  secondary_path:                   # path used to store data only, split by semicolon\n"
    "\n"
    "  backend_url: sqlite://:@:/        \n"
    "\n"
    "  insert_buffer_size: 4             # GB, maximum insert buffer size allowed\n"
    "  preload_table:                    \n"
    "\n"
    "metric_config:\n"
    "  enable_monitor: false             # enable monitoring or not\n"
    "  collector: prometheus             # prometheus\n"
    "  prometheus_config:\n"
    "    port: 8080                      # port prometheus uses to fetch metrics\n"
    "\n"
    "cache_config:\n"
    "  cpu_cache_capacity: 16            # GB, CPU memory used for cache\n"
    "  cpu_cache_threshold: 0.85         \n"
    "  gpu_cache_capacity: 4             # GB, GPU memory used for cache\n"
    "  gpu_cache_threshold: 0.85         \n"
    "  cache_insert_data: false          # whether to load inserted data into cache\n"
    "\n"
    "engine_config:\n"
    "  use_blas_threshold: 20            \n"
    "\n"
    "resource_config:\n"
#ifdef MILVUS_CPU_VERSION
    "  search_resources:\n"
    "    - cpu\n"
    "  index_build_device: cpu           # CPU used for building index";
#else
    "  search_resources:\n"
    "    - cpu\n"
    "    - gpu0\n"
    "  index_build_device: gpu0          # GPU used for building index";
#endif

static const char* INVALID_CONFIG_STR = "*INVALID*";

void
WriteToFile(const std::string& file_path, const char* content) {
    std::fstream fs(file_path.c_str(), std::ios_base::out);

    // write data to file
    fs << content;
    fs.close();
}

}  // namespace

void
ConfigTest::SetUp() {
    std::string config_path(CONFIG_PATH);
    milvus::server::CommonUtil::CreateDirectory(config_path);
    WriteToFile(config_path + VALID_CONFIG_FILE, VALID_CONFIG_STR);
    WriteToFile(config_path + INVALID_CONFIG_FILE, INVALID_CONFIG_STR);
}

void
ConfigTest::TearDown() {
    std::string config_path(CONFIG_PATH);
    milvus::server::CommonUtil::DeleteDirectory(config_path);
}
