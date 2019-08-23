/*******************************************************************************
 * Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 ******************************************************************************/
#pragma once

#include "utils/Error.h"
#include "config/ConfigNode.h"

#include "yaml-cpp/yaml.h"

namespace zilliz {
namespace milvus {
namespace server {

static const char* CONFIG_SERVER = "server_config";
static const char* CONFIG_SERVER_ADDRESS = "address";
static const char* CONFIG_SERVER_PORT = "port";
static const char* CONFIG_SERVER_PROTOCOL = "transfer_protocol";
static const char* CONFIG_CLUSTER_MODE = "mode";
static const char* CONFIG_GPU_INDEX = "gpu_index";

static const char* CONFIG_DB = "db_config";
static const char* CONFIG_DB_URL = "db_backend_url";
static const char* CONFIG_DB_PATH = "db_path";
static const char* CONFIG_DB_SLAVE_PATH = "db_slave_path";
static const char* CONFIG_DB_INDEX_TRIGGER_SIZE = "index_building_threshold";
static const char* CONFIG_DB_ARCHIVE_DISK = "archive_disk_threshold";
static const char* CONFIG_DB_ARCHIVE_DAYS = "archive_days_threshold";
static const char* CONFIG_DB_INSERT_BUFFER_SIZE = "insert_buffer_size";
static const char* CONFIG_DB_PARALLEL_REDUCE = "parallel_reduce";

static const char* CONFIG_LOG = "log_config";

static const char* CONFIG_CACHE = "cache_config";
static const char* CONFIG_CPU_CACHE_CAPACITY = "cpu_cache_capacity";
static const char* CONFIG_GPU_CACHE_CAPACITY = "gpu_cache_capacity";
static const char* CACHE_FREE_PERCENT = "cpu_cache_free_percent";
static const char* CONFIG_INSERT_CACHE_IMMEDIATELY = "insert_cache_immediately";
static const char* CONFIG_GPU_IDS = "gpu_ids";
static const char *GPU_CACHE_FREE_PERCENT = "gpu_cache_free_percent";

static const char* CONFIG_LICENSE = "license_config";
static const char* CONFIG_LICENSE_PATH = "license_path";

static const char* CONFIG_METRIC = "metric_config";
static const char* CONFIG_METRIC_IS_STARTUP = "is_startup";
static const char* CONFIG_METRIC_COLLECTOR = "collector";
static const char* CONFIG_PROMETHEUS = "prometheus_config";
static const char* CONFIG_METRIC_PROMETHEUS_PORT = "port";

static const std::string CONFIG_ENGINE = "engine_config";
static const std::string CONFIG_NPROBE = "nprobe";
static const std::string CONFIG_NLIST = "nlist";
static const std::string CONFIG_DCBT = "use_blas_threshold";
static const std::string CONFIG_METRICTYPE = "metric_type";
static const std::string CONFIG_OMP_THREAD_NUM = "omp_thread_num";
static const std::string CONFIG_USE_HYBRID_INDEX = "use_hybrid_index";
static const std::string CONFIG_HYBRID_INDEX_GPU = "hybrid_index_gpu";

static const char* CONFIG_RESOURCE = "resource_config";
static const char* CONFIG_RESOURCES = "resources";
static const char* CONFIG_RESOURCE_TYPE = "type";
static const char* CONFIG_RESOURCE_MEMORY = "memory";
static const char* CONFIG_RESOURCE_DEVICE_ID = "device_id";
static const char* CONFIG_RESOURCE_ENABLE_LOADER = "enable_loader";
static const char* CONFIG_RESOURCE_ENABLE_EXECUTOR = "enable_executor";
static const char* CONFIG_RESOURCE_CONNECTIONS = "connections";


class ServerConfig {
 public:
    static ServerConfig &GetInstance();

    ServerError LoadConfigFile(const std::string& config_filename);
    ServerError ValidateConfig() const;
    void PrintAll() const;

    ConfigNode GetConfig(const std::string& name) const;
    ConfigNode& GetConfig(const std::string& name);
};

}
}
}

