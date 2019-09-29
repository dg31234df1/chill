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

#pragma once

#include <yaml-cpp/yaml.h>
#include <mutex>
#include <string>
#include <unordered_map>
#include <vector>

#include "config/ConfigNode.h"
#include "utils/Status.h"

namespace milvus {
namespace server {

/* server config */
static const char* CONFIG_SERVER = "server_config";
static const char* CONFIG_SERVER_ADDRESS = "address";
static const char* CONFIG_SERVER_ADDRESS_DEFAULT = "127.0.0.1";
static const char* CONFIG_SERVER_PORT = "port";
static const char* CONFIG_SERVER_PORT_DEFAULT = "19530";
static const char* CONFIG_SERVER_DEPLOY_MODE = "deploy_mode";
static const char* CONFIG_SERVER_DEPLOY_MODE_DEFAULT = "single";
static const char* CONFIG_SERVER_TIME_ZONE = "time_zone";
static const char* CONFIG_SERVER_TIME_ZONE_DEFAULT = "UTC+8";

/* db config */
static const char* CONFIG_DB = "db_config";
static const char* CONFIG_DB_PRIMARY_PATH = "primary_path";
static const char* CONFIG_DB_PRIMARY_PATH_DEFAULT = "/tmp/milvus";
static const char* CONFIG_DB_SECONDARY_PATH = "secondary_path";
static const char* CONFIG_DB_SECONDARY_PATH_DEFAULT = "";
static const char* CONFIG_DB_BACKEND_URL = "backend_url";
static const char* CONFIG_DB_BACKEND_URL_DEFAULT = "sqlite://:@:/";
static const char* CONFIG_DB_ARCHIVE_DISK_THRESHOLD = "archive_disk_threshold";
static const char* CONFIG_DB_ARCHIVE_DISK_THRESHOLD_DEFAULT = "0";
static const char* CONFIG_DB_ARCHIVE_DAYS_THRESHOLD = "archive_days_threshold";
static const char* CONFIG_DB_ARCHIVE_DAYS_THRESHOLD_DEFAULT = "0";
static const char* CONFIG_DB_INSERT_BUFFER_SIZE = "insert_buffer_size";
static const char* CONFIG_DB_INSERT_BUFFER_SIZE_DEFAULT = "4";
static const char* CONFIG_DB_BUILD_INDEX_GPU = "build_index_gpu";
static const char* CONFIG_DB_BUILD_INDEX_GPU_DEFAULT = "0";

/* cache config */
static const char* CONFIG_CACHE = "cache_config";
static const char* CONFIG_CACHE_CPU_CACHE_CAPACITY = "cpu_cache_capacity";
static const char* CONFIG_CACHE_CPU_CACHE_CAPACITY_DEFAULT = "16";
static const char* CONFIG_CACHE_GPU_CACHE_CAPACITY = "gpu_cache_capacity";
static const char* CONFIG_CACHE_GPU_CACHE_CAPACITY_DEFAULT = "0";
static const char* CONFIG_CACHE_CPU_CACHE_THRESHOLD = "cpu_mem_threshold";
static const char* CONFIG_CACHE_CPU_CACHE_THRESHOLD_DEFAULT = "0.85";
static const char* CONFIG_CACHE_GPU_CACHE_THRESHOLD = "gpu_mem_threshold";
static const char* CONFIG_CACHE_GPU_CACHE_THRESHOLD_DEFAULT = "0.85";
static const char* CONFIG_CACHE_CACHE_INSERT_DATA = "cache_insert_data";
static const char* CONFIG_CACHE_CACHE_INSERT_DATA_DEFAULT = "false";

/* metric config */
static const char* CONFIG_METRIC = "metric_config";
static const char* CONFIG_METRIC_ENABLE_MONITOR = "enable_monitor";
static const char* CONFIG_METRIC_ENABLE_MONITOR_DEFAULT = "false";
static const char* CONFIG_METRIC_COLLECTOR = "collector";
static const char* CONFIG_METRIC_COLLECTOR_DEFAULT = "prometheus";
static const char* CONFIG_METRIC_PROMETHEUS = "prometheus_config";
static const char* CONFIG_METRIC_PROMETHEUS_PORT = "port";
static const char* CONFIG_METRIC_PROMETHEUS_PORT_DEFAULT = "8080";

/* engine config */
static const char* CONFIG_ENGINE = "engine_config";
static const char* CONFIG_ENGINE_USE_BLAS_THRESHOLD = "use_blas_threshold";
static const char* CONFIG_ENGINE_USE_BLAS_THRESHOLD_DEFAULT = "20";
static const char* CONFIG_ENGINE_OMP_THREAD_NUM = "omp_thread_num";
static const char* CONFIG_ENGINE_OMP_THREAD_NUM_DEFAULT = "0";

/* resource config */
static const char* CONFIG_RESOURCE = "resource_config";
static const char* CONFIG_RESOURCE_MODE = "mode";
static const char* CONFIG_RESOURCE_MODE_DEFAULT = "simple";
static const char* CONFIG_RESOURCE_POOL = "resource_pool";

class Config {
 public:
    static Config&
    GetInstance();
    Status
    LoadConfigFile(const std::string& filename);
    Status
    ValidateConfig();
    Status
    ResetDefaultConfig();
    void
    PrintAll();

 private:
    ConfigNode&
    GetConfigNode(const std::string& name);

    Status
    GetConfigValueInMem(const std::string& parent_key, const std::string& child_key, std::string& value);

    void
    SetConfigValueInMem(const std::string& parent_key, const std::string& child_key, const std::string& value);

    void
    PrintConfigSection(const std::string& config_node_name);

    ///////////////////////////////////////////////////////////////////////////
    /* server config */
    Status
    CheckServerConfigAddress(const std::string& value);
    Status
    CheckServerConfigPort(const std::string& value);
    Status
    CheckServerConfigDeployMode(const std::string& value);
    Status
    CheckServerConfigTimeZone(const std::string& value);

    /* db config */
    Status
    CheckDBConfigPrimaryPath(const std::string& value);
    Status
    CheckDBConfigSecondaryPath(const std::string& value);
    Status
    CheckDBConfigBackendUrl(const std::string& value);
    Status
    CheckDBConfigArchiveDiskThreshold(const std::string& value);
    Status
    CheckDBConfigArchiveDaysThreshold(const std::string& value);
    Status
    CheckDBConfigInsertBufferSize(const std::string& value);
    Status
    CheckDBConfigBuildIndexGPU(const std::string& value);

    /* metric config */
    Status
    CheckMetricConfigEnableMonitor(const std::string& value);
    Status
    CheckMetricConfigCollector(const std::string& value);
    Status
    CheckMetricConfigPrometheusPort(const std::string& value);

    /* cache config */
    Status
    CheckCacheConfigCpuCacheCapacity(const std::string& value);
    Status
    CheckCacheConfigCpuCacheThreshold(const std::string& value);
    Status
    CheckCacheConfigGpuCacheCapacity(const std::string& value);
    Status
    CheckCacheConfigGpuCacheThreshold(const std::string& value);
    Status
    CheckCacheConfigCacheInsertData(const std::string& value);

    /* engine config */
    Status
    CheckEngineConfigUseBlasThreshold(const std::string& value);
    Status
    CheckEngineConfigOmpThreadNum(const std::string& value);

    /* resource config */
    Status
    CheckResourceConfigMode(const std::string& value);
    Status
    CheckResourceConfigPool(const std::vector<std::string>& value);

    ///////////////////////////////////////////////////////////////////////////
    /* server config */
    std::string
    GetServerConfigStrAddress();
    std::string
    GetServerConfigStrPort();
    std::string
    GetServerConfigStrDeployMode();
    std::string
    GetServerConfigStrTimeZone();

    /* db config */
    std::string
    GetDBConfigStrPrimaryPath();
    std::string
    GetDBConfigStrSecondaryPath();
    std::string
    GetDBConfigStrBackendUrl();
    std::string
    GetDBConfigStrArchiveDiskThreshold();
    std::string
    GetDBConfigStrArchiveDaysThreshold();
    std::string
    GetDBConfigStrInsertBufferSize();
    std::string
    GetDBConfigStrBuildIndexGPU();

    /* metric config */
    std::string
    GetMetricConfigStrEnableMonitor();
    std::string
    GetMetricConfigStrCollector();
    std::string
    GetMetricConfigStrPrometheusPort();

    /* cache config */
    std::string
    GetCacheConfigStrCpuCacheCapacity();
    std::string
    GetCacheConfigStrCpuCacheThreshold();
    std::string
    GetCacheConfigStrGpuCacheCapacity();
    std::string
    GetCacheConfigStrGpuCacheThreshold();
    std::string
    GetCacheConfigStrCacheInsertData();

    /* engine config */
    std::string
    GetEngineConfigStrUseBlasThreshold();
    std::string
    GetEngineConfigStrOmpThreadNum();

    /* resource config */
    std::string
    GetResourceConfigStrMode();

 public:
    /* server config */
    Status
    GetServerConfigAddress(std::string& value);
    Status
    GetServerConfigPort(std::string& value);
    Status
    GetServerConfigDeployMode(std::string& value);
    Status
    GetServerConfigTimeZone(std::string& value);

    /* db config */
    Status
    GetDBConfigPrimaryPath(std::string& value);
    Status
    GetDBConfigSecondaryPath(std::string& value);
    Status
    GetDBConfigBackendUrl(std::string& value);
    Status
    GetDBConfigArchiveDiskThreshold(int32_t& value);
    Status
    GetDBConfigArchiveDaysThreshold(int32_t& value);
    Status
    GetDBConfigInsertBufferSize(int32_t& value);
    Status
    GetDBConfigBuildIndexGPU(int32_t& value);

    /* metric config */
    Status
    GetMetricConfigEnableMonitor(bool& value);
    Status
    GetMetricConfigCollector(std::string& value);
    Status
    GetMetricConfigPrometheusPort(std::string& value);

    /* cache config */
    Status
    GetCacheConfigCpuCacheCapacity(int32_t& value);
    Status
    GetCacheConfigCpuCacheThreshold(float& value);
    Status
    GetCacheConfigGpuCacheCapacity(int32_t& value);
    Status
    GetCacheConfigGpuCacheThreshold(float& value);
    Status
    GetCacheConfigCacheInsertData(bool& value);

    /* engine config */
    Status
    GetEngineConfigUseBlasThreshold(int32_t& value);
    Status
    GetEngineConfigOmpThreadNum(int32_t& value);

    /* resource config */
    Status
    GetResourceConfigMode(std::string& value);
    Status
    GetResourceConfigPool(std::vector<std::string>& value);

 public:
    /* server config */
    Status
    SetServerConfigAddress(const std::string& value);
    Status
    SetServerConfigPort(const std::string& value);
    Status
    SetServerConfigDeployMode(const std::string& value);
    Status
    SetServerConfigTimeZone(const std::string& value);

    /* db config */
    Status
    SetDBConfigPrimaryPath(const std::string& value);
    Status
    SetDBConfigSecondaryPath(const std::string& value);
    Status
    SetDBConfigBackendUrl(const std::string& value);
    Status
    SetDBConfigArchiveDiskThreshold(const std::string& value);
    Status
    SetDBConfigArchiveDaysThreshold(const std::string& value);
    Status
    SetDBConfigInsertBufferSize(const std::string& value);
    Status
    SetDBConfigBuildIndexGPU(const std::string& value);

    /* metric config */
    Status
    SetMetricConfigEnableMonitor(const std::string& value);
    Status
    SetMetricConfigCollector(const std::string& value);
    Status
    SetMetricConfigPrometheusPort(const std::string& value);

    /* cache config */
    Status
    SetCacheConfigCpuCacheCapacity(const std::string& value);
    Status
    SetCacheConfigCpuCacheThreshold(const std::string& value);
    Status
    SetCacheConfigGpuCacheCapacity(const std::string& value);
    Status
    SetCacheConfigGpuCacheThreshold(const std::string& value);
    Status
    SetCacheConfigCacheInsertData(const std::string& value);

    /* engine config */
    Status
    SetEngineConfigUseBlasThreshold(const std::string& value);
    Status
    SetEngineConfigOmpThreadNum(const std::string& value);

    /* resource config */
    Status
    SetResourceConfigMode(const std::string& value);

 private:
    std::unordered_map<std::string, std::unordered_map<std::string, std::string>> config_map_;
    std::mutex mutex_;
};

}  // namespace server
}  // namespace milvus
