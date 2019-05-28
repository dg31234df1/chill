/*******************************************************************************
 * Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 ******************************************************************************/
#include <chrono>
#include <chrono>
#include <map>
#include <memory>
#include <string>
#include <thread>
#include <gtest/gtest.h>
//#include "prometheus/registry.h"
//#include "prometheus/exposer.h"
#include <cache/CpuCacheMgr.h>

#include "metrics/Metrics.h"
#include "../db/utils.h"
#include "db/DB.h"
#include "db/DBMetaImpl.h"
#include "db/Factories.h"


using namespace zilliz::vecwise;

TEST_F(DBTest, Metric_Test) {

    using namespace zilliz::vecwise;
//    server::Metrics::GetInstance().Init();
    server::Metrics::GetInstance().Init();
//    server::PrometheusMetrics::GetInstance().exposer_ptr()->RegisterCollectable(server::PrometheusMetrics::GetInstance().registry_ptr());
//    server::Metrics::GetInstance().exposer_ptr()->RegisterCollectable(server::Metrics::GetInstance().registry_ptr());

    static const std::string group_name = "test_group";
    static const int group_dim = 256;

    engine::meta::GroupSchema group_info;
    group_info.dimension = group_dim;
    group_info.group_id = group_name;
    engine::Status stat = db_->add_group(group_info);

    engine::meta::GroupSchema group_info_get;
    group_info_get.group_id = group_name;

//    int iter = 600000;
//    for (int i = 0; i < iter; ++i) {
//        db_->get_group(group_info);
//        bool b = i % 2;
//        db_->has_group(std::to_string(i),b);
//        db_->add_group(group_info);
//
//        std::this_thread::sleep_for(std::chrono::microseconds(1));
//    }
    stat = db_->get_group(group_info_get);
    ASSERT_STATS(stat);
    ASSERT_EQ(group_info_get.dimension, group_dim);

    engine::IDNumbers vector_ids;
    engine::IDNumbers target_ids;

    int d = 256;
    int nb = 50;
    float *xb = new float[d * nb];
    for(int i = 0; i < nb; i++) {
        for(int j = 0; j < d; j++) xb[d * i + j] = drand48();
        xb[d * i] += i / 2000.;
    }

    int qb = 5;
    float *qxb = new float[d * qb];
    for(int i = 0; i < qb; i++) {
        for(int j = 0; j < d; j++) qxb[d * i + j] = drand48();
        qxb[d * i] += i / 2000.;
    }

    int loop =100000;

    for (auto i=0; i<loop; ++i) {
        if (i==40) {
            db_->add_vectors(group_name, qb, qxb, target_ids);
            ASSERT_EQ(target_ids.size(), qb);
        } else {
            db_->add_vectors(group_name, nb, xb, vector_ids);
        }
        std::this_thread::sleep_for(std::chrono::microseconds(5));
    }

//    search.join();

    delete [] xb;
    delete [] qxb;
};


TEST_F(DBTest, Metric_Tes) {


//    server::Metrics::GetInstance().Init();
//    server::Metrics::GetInstance().exposer_ptr()->RegisterCollectable(server::Metrics::GetInstance().registry_ptr());
    server::Metrics::GetInstance().Init();
//    server::PrometheusMetrics::GetInstance().exposer_ptr()->RegisterCollectable(server::PrometheusMetrics::GetInstance().registry_ptr());
    zilliz::vecwise::cache::CpuCacheMgr::GetInstance()->SetCapacity(1*1024*1024*1024);
    std::cout<<zilliz::vecwise::cache::CpuCacheMgr::GetInstance()->CacheCapacity()<<std::endl;
    static const std::string group_name = "test_group";
    static const int group_dim = 256;

    engine::meta::GroupSchema group_info;
    group_info.dimension = group_dim;
    group_info.group_id = group_name;
    engine::Status stat = db_->add_group(group_info);

    engine::meta::GroupSchema group_info_get;
    group_info_get.group_id = group_name;
    stat = db_->get_group(group_info_get);
    ASSERT_STATS(stat);
    ASSERT_EQ(group_info_get.dimension, group_dim);

    engine::IDNumbers vector_ids;
    engine::IDNumbers target_ids;

    int d = 256;
    int nb = 50;
    float *xb = new float[d * nb];
    for(int i = 0; i < nb; i++) {
        for(int j = 0; j < d; j++) xb[d * i + j] = drand48();
        xb[d * i] += i / 2000.;
    }

    int qb = 5;
    float *qxb = new float[d * qb];
    for(int i = 0; i < qb; i++) {
        for(int j = 0; j < d; j++) qxb[d * i + j] = drand48();
        qxb[d * i] += i / 2000.;
    }

    std::thread search([&]() {
        engine::QueryResults results;
        int k = 10;
        std::this_thread::sleep_for(std::chrono::seconds(2));

        INIT_TIMER;
        std::stringstream ss;
        long count = 0;
        long prev_count = -1;

        for (auto j=0; j<10; ++j) {
            ss.str("");
            db_->count(group_name, count);
            prev_count = count;

            START_TIMER;
            stat = db_->search(group_name, k, qb, qxb, results);
            ss << "Search " << j << " With Size " << (float)(count*group_dim*sizeof(float))/(1024*1024) << " M";
//            STOP_TIMER(ss.str());

            ASSERT_STATS(stat);
            for (auto k=0; k<qb; ++k) {
                ASSERT_EQ(results[k][0], target_ids[k]);
                ss.str("");
                ss << "Result [" << k << "]:";
                for (auto result : results[k]) {
                    ss << result << " ";
                }
                /* LOG(DEBUG) << ss.str(); */
            }
            ASSERT_TRUE(count >= prev_count);
            std::this_thread::sleep_for(std::chrono::seconds(1));
        }
    });

    int loop = 100000;

    for (auto i=0; i<loop; ++i) {
        if (i==40) {
            db_->add_vectors(group_name, qb, qxb, target_ids);
            ASSERT_EQ(target_ids.size(), qb);
        } else {
            db_->add_vectors(group_name, nb, xb, vector_ids);
        }
        std::this_thread::sleep_for(std::chrono::microseconds(1));
    }

    search.join();

    delete [] xb;
    delete [] qxb;
};


