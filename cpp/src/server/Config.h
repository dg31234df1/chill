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

#include <mutex>
#include <unordered_map>
#include "yaml-cpp/yaml.h"
#include "utils/Status.h"
#include "config/ConfigNode.h"


namespace zilliz {
namespace milvus {
namespace server {

/* server config */
static const char* CONFIG_SERVER = "server_config";
static const char* CONFIG_SERVER_ADDRESS = "address";
static const char* CONFIG_SERVER_ADDRESS_DEFAULT = "127.0.0.1";
static const char* CONFIG_SERVER_PORT = "port";
static const char* CONFIG_SERVER_PORT_DEFAULT = "19530";
static const char* CONFIG_SERVER_MODE = "mode";
static const char* CONFIG_SERVER_MODE_DEFAULT = "single";
static const char* CONFIG_SERVER_TIME_ZONE = "time_zone";
static const char* CONFIG_SERVER_TIME_ZONE_DEFAULT = "UTC+8";

/* db config */
static const char* CONFIG_DB = "db_config";
static const char* CONFIG_DB_PATH = "path";
static const char* CONFIG_DB_PATH_DEFAULT = "/tmp/milvus";
static const char* CONFIG_DB_SLAVE_PATH = "slave_path";
static const char* CONFIG_DB_SLAVE_PATH_DEFAULT = "";
static const char* CONFIG_DB_BACKEND_URL = "backend_url";
static const char* CONFIG_DB_BACKEND_URL_DEFAULT = "sqlite://:@:/";
static const char* CONFIG_DB_ARCHIVE_DISK_THRESHOLD = "archive_disk_threshold";
static const char* CONFIG_DB_ARCHIVE_DISK_THRESHOLD_DEFAULT = "0";
static const char* CONFIG_DB_ARCHIVE_DAYS_THRESHOLD = "archive_days_threshold";
static const char* CONFIG_DB_ARCHIVE_DAYS_THRESHOLD_DEFAULT = "0";
static const char* CONFIG_DB_BUFFER_SIZE = "buffer_size";
static const char* CONFIG_DB_BUFFER_SIZE_DEFAULT = "4";
static const char* CONFIG_DB_BUILD_INDEX_GPU = "build_index_gpu";
static const char* CONFIG_DB_BUILD_INDEX_GPU_DEFAULT = "0";

/* cache config */
static const char* CONFIG_CACHE = "cache_config";
static const char* CONFIG_CACHE_CPU_MEM_CAPACITY = "cpu_mem_capacity";
static const char* CONFIG_CACHE_CPU_MEM_CAPACITY_DEFAULT = "16";
static const char* CONFIG_CACHE_GPU_MEM_CAPACITY = "gpu_mem_capacity";
static const char* CONFIG_CACHE_GPU_MEM_CAPACITY_DEFAULT = "0";
static const char* CONFIG_CACHE_CPU_MEM_THRESHOLD = "cpu_mem_threshold";
static const char* CONFIG_CACHE_CPU_MEM_THRESHOLD_DEFAULT = "0.85";
static const char* CONFIG_CACHE_GPU_MEM_THRESHOLD = "gpu_mem_threshold";
static const char* CONFIG_CACHE_GPU_MEM_THRESHOLD_DEFAULT = "0.85";
static const char* CONFIG_CACHE_CACHE_INSERT_DATA = "cache_insert_data";
static const char* CONFIG_CACHE_CACHE_INSERT_DATA_DEFAULT = "false";

/* metric config */
static const char* CONFIG_METRIC = "metric_config";
static const char* CONFIG_METRIC_AUTO_BOOTUP = "auto_bootup";
static const char* CONFIG_METRIC_AUTO_BOOTUP_DEFAULT = "false";
static const char* CONFIG_METRIC_COLLECTOR = "collector";
static const char* CONFIG_METRIC_COLLECTOR_DEFAULT = "prometheus";
static const char* CONFIG_METRIC_PROMETHEUS = "prometheus_config";
static const char* CONFIG_METRIC_PROMETHEUS_PORT = "port";
static const char* CONFIG_METRIC_PROMETHEUS_PORT_DEFAULT = "8080";

/* engine config */
static const char* CONFIG_ENGINE = "engine_config";
static const char* CONFIG_ENGINE_BLAS_THRESHOLD = "blas_threshold";
static const char* CONFIG_ENGINE_BLAS_THRESHOLD_DEFAULT = "20";
static const char* CONFIG_ENGINE_OMP_THREAD_NUM = "omp_thread_num";
static const char* CONFIG_ENGINE_OMP_THREAD_NUM_DEFAULT = "0";

/* resource config */
static const char* CONFIG_RESOURCE = "resource_config";
static const char* CONFIG_RESOURCE_MODE = "mode";
static const char* CONFIG_RESOURCE_MODE_DEFAULT = "simple";
static const char* CONFIG_RESOURCE_POOL = "pool";


class Config {
 public:
    static Config& GetInstance();
    Status LoadConfigFile(const std::string& filename);
    void PrintAll();

 private:
    ConfigNode& GetConfigNode(const std::string& name);

    Status GetConfigValueInMem(const std::string& parent_key,
                               const std::string& child_key,
                               std::string& value);

    void   SetConfigValueInMem(const std::string& parent_key,
                               const std::string& child_key,
                               const std::string& value);

    void   PrintConfigSection(const std::string& config_node_name);

    ///////////////////////////////////////////////////////////////////////////
    /* server config */
    Status CheckServerConfigAddress(const std::string& value);
    Status CheckServerConfigPort(const std::string& value);
    Status CheckServerConfigMode(const std::string& value);
    Status CheckServerConfigTimeZone(const std::string& value);

    /* db config */
    Status CheckDBConfigPath(const std::string& value);
    Status CheckDBConfigSlavePath(const std::string& value);
    Status CheckDBConfigBackendUrl(const std::string& value);
    Status CheckDBConfigArchiveDiskThreshold(const std::string& value);
    Status CheckDBConfigArchiveDaysThreshold(const std::string& value);
    Status CheckDBConfigBufferSize(const std::string& value);
    Status CheckDBConfigBuildIndexGPU(const std::string& value);

    /* metric config */
    Status CheckMetricConfigAutoBootup(const std::string& value);
    Status CheckMetricConfigCollector(const std::string& value);
    Status CheckMetricConfigPrometheusPort(const std::string& value);

    /* cache config */
    Status CheckCacheConfigCpuMemCapacity(const std::string& value);
    Status CheckCacheConfigCpuMemThreshold(const std::string& value);
    Status CheckCacheConfigGpuMemCapacity(const std::string& value);
    Status CheckCacheConfigGpuMemThreshold(const std::string& value);
    Status CheckCacheConfigCacheInsertData(const std::string& value);

    /* engine config */
    Status CheckEngineConfigBlasThreshold(const std::string& value);
    Status CheckEngineConfigOmpThreadNum(const std::string& value);

    /* resource config */
    Status CheckResourceConfigMode(const std::string& value);
    Status CheckResourceConfigPool(const std::vector<std::string>& value);

    ///////////////////////////////////////////////////////////////////////////
    /* server config */
    std::string GetServerConfigStrAddress();
    std::string GetServerConfigStrPort();
    std::string GetServerConfigStrMode();
    std::string GetServerConfigStrTimeZone();

    /* db config */
    std::string GetDBConfigStrPath();
    std::string GetDBConfigStrSlavePath();
    std::string GetDBConfigStrBackendUrl();
    std::string GetDBConfigStrArchiveDiskThreshold();
    std::string GetDBConfigStrArchiveDaysThreshold();
    std::string GetDBConfigStrBufferSize();
    std::string GetDBConfigStrBuildIndexGPU();

    /* metric config */
    std::string GetMetricConfigStrAutoBootup();
    std::string GetMetricConfigStrCollector();
    std::string GetMetricConfigStrPrometheusPort();

    /* cache config */
    std::string GetCacheConfigStrCpuMemCapacity();
    std::string GetCacheConfigStrCpuMemThreshold();
    std::string GetCacheConfigStrGpuMemCapacity();
    std::string GetCacheConfigStrGpuMemThreshold();
    std::string GetCacheConfigStrCacheInsertData();

    /* engine config */
    std::string GetEngineConfigStrBlasThreshold();
    std::string GetEngineConfigStrOmpThreadNum();

    /* resource config */
    std::string GetResourceConfigStrMode();

 public:
    /* server config */
    Status GetServerConfigAddress(std::string& value);
    Status GetServerConfigPort(std::string& value);
    Status GetServerConfigMode(std::string& value);
    Status GetServerConfigTimeZone(std::string& value);

    /* db config */
    Status GetDBConfigPath(std::string& value);
    Status GetDBConfigSlavePath(std::string& value);
    Status GetDBConfigBackendUrl(std::string& value);
    Status GetDBConfigArchiveDiskThreshold(int32_t& value);
    Status GetDBConfigArchiveDaysThreshold(int32_t& value);
    Status GetDBConfigBufferSize(int32_t& value);
    Status GetDBConfigBuildIndexGPU(int32_t& value);

    /* metric config */
    Status GetMetricConfigAutoBootup(bool& value);
    Status GetMetricConfigCollector(std::string& value);
    Status GetMetricConfigPrometheusPort(std::string& value);

    /* cache config */
    Status GetCacheConfigCpuMemCapacity(int32_t& value);
    Status GetCacheConfigCpuMemThreshold(float& value);
    Status GetCacheConfigGpuMemCapacity(int32_t& value);
    Status GetCacheConfigGpuMemThreshold(float& value);
    Status GetCacheConfigCacheInsertData(bool& value);

    /* engine config */
    Status GetEngineConfigBlasThreshold(int32_t& value);
    Status GetEngineConfigOmpThreadNum(int32_t& value);

    /* resource config */
    Status GetResourceConfigMode(std::string& value);
    Status GetResourceConfigPool(std::vector<std::string>& value);

 public:
    /* server config */
    Status SetServerConfigAddress(const std::string& value);
    Status SetServerConfigPort(const std::string& value);
    Status SetServerConfigMode(const std::string& value);
    Status SetServerConfigTimeZone(const std::string& value);

    /* db config */
    Status SetDBConfigPath(const std::string& value);
    Status SetDBConfigSlavePath(const std::string& value);
    Status SetDBConfigBackendUrl(const std::string& value);
    Status SetDBConfigArchiveDiskThreshold(const std::string& value);
    Status SetDBConfigArchiveDaysThreshold(const std::string& value);
    Status SetDBConfigBufferSize(const std::string& value);
    Status SetDBConfigBuildIndexGPU(const std::string& value);

    /* metric config */
    Status SetMetricConfigAutoBootup(const std::string& value);
    Status SetMetricConfigCollector(const std::string& value);
    Status SetMetricConfigPrometheusPort(const std::string& value);

    /* cache config */
    Status SetCacheConfigCpuMemCapacity(const std::string& value);
    Status SetCacheConfigCpuMemThreshold(const std::string& value);
    Status SetCacheConfigGpuMemCapacity(const std::string& value);
    Status SetCacheConfigGpuMemThreshold(const std::string& value);
    Status SetCacheConfigCacheInsertData(const std::string& value);

    /* engine config */
    Status SetEngineConfigBlasThreshold(const std::string& value);
    Status SetEngineConfigOmpThreadNum(const std::string& value);

    /* resource config */
    Status SetResourceConfigMode(const std::string& value);

 private:
    std::unordered_map<std::string, std::unordered_map<std::string, std::string>> config_map_;
    std::mutex mutex_;
};

}
}
}

