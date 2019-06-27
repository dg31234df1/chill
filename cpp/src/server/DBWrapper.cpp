/*******************************************************************************
 * Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 ******************************************************************************/

#include "DBWrapper.h"
#include "ServerConfig.h"
#include "utils/CommonUtil.h"
#include "utils/Log.h"

namespace zilliz {
namespace milvus {
namespace server {

DBWrapper::DBWrapper() {
    zilliz::milvus::engine::Options opt;
    ConfigNode& config = ServerConfig::GetInstance().GetConfig(CONFIG_DB);
    opt.meta.backend_uri = config.GetValue(CONFIG_DB_URL);
    std::string db_path = config.GetValue(CONFIG_DB_PATH);
    opt.meta.path = db_path + "/db";
    int64_t index_size = config.GetInt64Value(CONFIG_DB_INDEX_TRIGGER_SIZE);
    if(index_size > 0) {//ensure larger than zero, unit is MB
        opt.index_trigger_size = (size_t)index_size * engine::ONE_MB;
    }

    //set archive config
    engine::ArchiveConf::CriteriaT criterial;
    int64_t disk = config.GetInt64Value(CONFIG_DB_ARCHIVE_DISK, 0);
    int64_t days = config.GetInt64Value(CONFIG_DB_ARCHIVE_DAYS, 0);
    if(disk > 0) {
        criterial[engine::ARCHIVE_CONF_DISK] = disk;
    }
    if(days > 0) {
        criterial[engine::ARCHIVE_CONF_DAYS] = days;
    }
    opt.meta.archive_conf.SetCriterias(criterial);

    //create db root folder
    ServerError err = CommonUtil::CreateDirectory(opt.meta.path);
    if(err != SERVER_SUCCESS) {
        std::cout << "ERROR! Failed to create database root path: " << opt.meta.path << std::endl;
        kill(0, SIGUSR1);
    }

    zilliz::milvus::engine::DB::Open(opt, &db_);
    if(db_ == nullptr) {
        std::cout << "ERROR! Failed to open database" << std::endl;
        kill(0, SIGUSR1);
    }
}

DBWrapper::~DBWrapper() {
    delete db_;
}

}
}
}