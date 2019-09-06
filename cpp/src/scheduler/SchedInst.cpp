/*******************************************************************************
 * Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 ******************************************************************************/

#include "SchedInst.h"
#include "server/ServerConfig.h"
#include "ResourceFactory.h"
#include "knowhere/index/vector_index/gpu_ivf.h"

namespace zilliz {
namespace milvus {
namespace engine {

ResourceMgrPtr ResMgrInst::instance = nullptr;
std::mutex ResMgrInst::mutex_;

SchedulerPtr SchedInst::instance = nullptr;
std::mutex SchedInst::mutex_;

void
StartSchedulerService() {
    try {
        server::ConfigNode &config = server::ServerConfig::GetInstance().GetConfig(server::CONFIG_RESOURCE);

        //TODO: change const char * to standard
        if (config.GetChildren().empty()) throw "resource_config null exception";

        auto resources = config.GetChild(server::CONFIG_RESOURCES).GetChildren();

        if (resources.empty()) throw "Children of resource_config null exception";

        for (auto &resource : resources) {
            auto &resname = resource.first;
            auto &resconf = resource.second;
            auto type = resconf.GetValue(server::CONFIG_RESOURCE_TYPE);
//        auto memory = resconf.GetInt64Value(server::CONFIG_RESOURCE_MEMORY);
            auto device_id = resconf.GetInt64Value(server::CONFIG_RESOURCE_DEVICE_ID);
            auto enable_loader = resconf.GetBoolValue(server::CONFIG_RESOURCE_ENABLE_LOADER);
            auto enable_executor = resconf.GetBoolValue(server::CONFIG_RESOURCE_ENABLE_EXECUTOR);
            auto pinned_memory = resconf.GetInt64Value(server::CONFIG_RESOURCE_PIN_MEMORY);
            auto temp_memory = resconf.GetInt64Value(server::CONFIG_RESOURCE_TEMP_MEMORY);
            auto resource_num = resconf.GetInt64Value(server::CONFIG_RESOURCE_NUM);

            auto res = ResMgrInst::GetInstance()->Add(ResourceFactory::Create(resname,
                                                                              type,
                                                                              device_id,
                                                                              enable_loader,
                                                                              enable_executor));

            if (res.lock()->Type() == ResourceType::GPU) {
                auto pinned_memory = resconf.GetInt64Value(server::CONFIG_RESOURCE_PIN_MEMORY, 300);
                auto temp_memory = resconf.GetInt64Value(server::CONFIG_RESOURCE_TEMP_MEMORY, 300);
                auto resource_num = resconf.GetInt64Value(server::CONFIG_RESOURCE_NUM, 2);
                pinned_memory = 1024 * 1024 * pinned_memory;
                temp_memory = 1024 * 1024 * temp_memory;
                knowhere::FaissGpuResourceMgr::GetInstance().InitDevice(device_id,
                                                                        pinned_memory,
                                                                        temp_memory,
                                                                        resource_num);
            }
        }

        knowhere::FaissGpuResourceMgr::GetInstance().InitResource();

        auto connections = config.GetChild(server::CONFIG_RESOURCE_CONNECTIONS).GetChildren();
        if(connections.empty()) throw "connections config null exception";
        for (auto &conn : connections) {
            auto &connect_name = conn.first;
            auto &connect_conf = conn.second;
            auto connect_speed = connect_conf.GetInt64Value(server::CONFIG_SPEED_CONNECTIONS);
            auto connect_endpoint = connect_conf.GetValue(server::CONFIG_ENDPOINT_CONNECTIONS);

            std::string delimiter = "===";
            std::string left = connect_endpoint.substr(0, connect_endpoint.find(delimiter));
            std::string right = connect_endpoint.substr(connect_endpoint.find(delimiter) + 3,
                                                        connect_endpoint.length());

            auto connection = Connection(connect_name, connect_speed);
            ResMgrInst::GetInstance()->Connect(left, right, connection);
        }
    } catch (const char* msg) {
        SERVER_LOG_ERROR << msg;
        exit(-1);
    }

    ResMgrInst::GetInstance()->Start();
    SchedInst::GetInstance()->Start();
}

void
StopSchedulerService() {
    ResMgrInst::GetInstance()->Stop();
    SchedInst::GetInstance()->Stop();
}
}
}
}
