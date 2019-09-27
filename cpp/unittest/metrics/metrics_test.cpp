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

#include <chrono>
#include <map>
#include <memory>
#include <string>
#include <thread>
#include <gtest/gtest.h>

#include "cache/CpuCacheMgr.h"
#include "server/Config.h"
#include "metrics/Metrics.h"
#include "metrics/utils.h"
#include "db/DB.h"
#include "db/meta/SqliteMetaImpl.h"

namespace {

namespace ms = zilliz::milvus;

} // namespace

TEST_F(MetricTest, METRIC_TEST) {
    ms::server::Config::GetInstance().SetMetricConfigCollector("zabbix");
    ms::server::Metrics::GetInstance();
    ms::server::Config::GetInstance().SetMetricConfigCollector("prometheus");
    ms::server::Metrics::GetInstance();

    ms::server::SystemInfo::GetInstance().Init();
//    server::Metrics::GetInstance().Init();
//    server::Metrics::GetInstance().exposer_ptr()->RegisterCollectable(server::Metrics::GetInstance().registry_ptr());
    ms::server::Metrics::GetInstance().Init();

//    server::PrometheusMetrics::GetInstance().exposer_ptr()->RegisterCollectable(server::PrometheusMetrics::GetInstance().registry_ptr());
    zilliz::milvus::cache::CpuCacheMgr::GetInstance()->SetCapacity(1UL * 1024 * 1024 * 1024);
    std::cout << zilliz::milvus::cache::CpuCacheMgr::GetInstance()->CacheCapacity() << std::endl;

    static const char *group_name = "test_group";
    static const int group_dim = 256;

    ms::engine::meta::TableSchema group_info;
    group_info.dimension_ = group_dim;
    group_info.table_id_ = group_name;
    auto stat = db_->CreateTable(group_info);

    ms::engine::meta::TableSchema group_info_get;
    group_info_get.table_id_ = group_name;
    stat = db_->DescribeTable(group_info_get);

    ms::engine::IDNumbers vector_ids;
    ms::engine::IDNumbers target_ids;

    int d = 256;
    int nb = 50;
    float *xb = new float[d * nb];
    for (int i = 0; i < nb; i++) {
        for (int j = 0; j < d; j++) xb[d * i + j] = drand48();
        xb[d * i] += i / 2000.;
    }

    int qb = 5;
    float *qxb = new float[d * qb];
    for (int i = 0; i < qb; i++) {
        for (int j = 0; j < d; j++) qxb[d * i + j] = drand48();
        qxb[d * i] += i / 2000.;
    }

    std::thread search([&]() {
        ms::engine::QueryResults results;
        int k = 10;
        std::this_thread::sleep_for(std::chrono::seconds(2));

        INIT_TIMER;
        std::stringstream ss;
        uint64_t count = 0;
        uint64_t prev_count = 0;

        for (auto j = 0; j < 10; ++j) {
            ss.str("");
            db_->Size(count);
            prev_count = count;

            START_TIMER;
//            stat = db_->Query(group_name, k, qb, qxb, results);
            ss << "Search " << j << " With Size " << (float) (count * group_dim * sizeof(float)) / (1024 * 1024)
               << " M";

            for (auto k = 0; k < qb; ++k) {
//                ASSERT_EQ(results[k][0].first, target_ids[k]);
                ss.str("");
                ss << "Result [" << k << "]:";
//                for (auto result : results[k]) {
//                    ss << result.first << " ";
//                }
            }
            ASSERT_TRUE(count >= prev_count);
            std::this_thread::sleep_for(std::chrono::seconds(1));
        }
    });

    int loop = 10000;

    for (auto i = 0; i < loop; ++i) {
        if (i == 40) {
            db_->InsertVectors(group_name, qb, qxb, target_ids);
            ASSERT_EQ(target_ids.size(), qb);
        } else {
            db_->InsertVectors(group_name, nb, xb, vector_ids);
        }
        std::this_thread::sleep_for(std::chrono::microseconds(2000));
    }

    search.join();

    delete[] xb;
    delete[] qxb;
}

TEST_F(MetricTest, COLLECTOR_METRICS_TEST) {
    auto status = ms::Status::OK();
    ms::server::CollectInsertMetrics insert_metrics0(0, status);
    status = ms::Status(ms::DB_ERROR, "error");
    ms::server::CollectInsertMetrics insert_metrics1(0, status);

    ms::server::CollectQueryMetrics query_metrics(10);

    ms::server::CollectMergeFilesMetrics merge_metrics();

    ms::server::CollectBuildIndexMetrics build_index_metrics();

    ms::server::CollectExecutionEngineMetrics execution_metrics(10);

    ms::server::CollectSerializeMetrics serialize_metrics(10);

    ms::server::CollectAddMetrics add_metrics(10, 128);

    ms::server::CollectDurationMetrics duration_metrics_raw(ms::engine::meta::TableFileSchema::RAW);
    ms::server::CollectDurationMetrics duration_metrics_index(ms::engine::meta::TableFileSchema::TO_INDEX);
    ms::server::CollectDurationMetrics duration_metrics_delete(ms::engine::meta::TableFileSchema::TO_DELETE);

    ms::server::CollectSearchTaskMetrics search_metrics_raw(ms::engine::meta::TableFileSchema::RAW);
    ms::server::CollectSearchTaskMetrics search_metrics_index(ms::engine::meta::TableFileSchema::TO_INDEX);
    ms::server::CollectSearchTaskMetrics search_metrics_delete(ms::engine::meta::TableFileSchema::TO_DELETE);

    ms::server::MetricCollector metric_collector();
}


