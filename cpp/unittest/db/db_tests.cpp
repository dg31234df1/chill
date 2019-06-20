////////////////////////////////////////////////////////////////////////////////
// Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited.
// Proprietary and confidential.
////////////////////////////////////////////////////////////////////////////////
#include <gtest/gtest.h>
#include <thread>
#include <easylogging++.h>

#include "utils.h"
#include "db/DB.h"
#include "db/DBImpl.h"
#include "db/MetaConsts.h"

using namespace zilliz::milvus;

TEST_F(DBTest, CONFIG_TEST) {
    {
        ASSERT_ANY_THROW(engine::ArchiveConf conf("wrong"));
        /* EXPECT_DEATH(engine::ArchiveConf conf("wrong"), ""); */
    }
    {
        engine::ArchiveConf conf("delete");
        ASSERT_EQ(conf.GetType(), "delete");
        auto criterias = conf.GetCriterias();
        ASSERT_TRUE(criterias.size() == 1);
        ASSERT_TRUE(criterias["disk"] == 512);
    }
    {
        engine::ArchiveConf conf("swap");
        ASSERT_EQ(conf.GetType(), "swap");
        auto criterias = conf.GetCriterias();
        ASSERT_TRUE(criterias.size() == 1);
        ASSERT_TRUE(criterias["disk"] == 512);
    }
    {
        ASSERT_ANY_THROW(engine::ArchiveConf conf1("swap", "disk:"));
        ASSERT_ANY_THROW(engine::ArchiveConf conf2("swap", "disk:a"));
        engine::ArchiveConf conf("swap", "disk:1024");
        auto criterias = conf.GetCriterias();
        ASSERT_TRUE(criterias.size() == 1);
        ASSERT_TRUE(criterias["disk"] == 1024);
    }
    {
        ASSERT_ANY_THROW(engine::ArchiveConf conf1("swap", "days:"));
        ASSERT_ANY_THROW(engine::ArchiveConf conf2("swap", "days:a"));
        engine::ArchiveConf conf("swap", "days:100");
        auto criterias = conf.GetCriterias();
        ASSERT_TRUE(criterias.size() == 1);
        ASSERT_TRUE(criterias["days"] == 100);
    }
    {
        ASSERT_ANY_THROW(engine::ArchiveConf conf1("swap", "days:"));
        ASSERT_ANY_THROW(engine::ArchiveConf conf2("swap", "days:a"));
        engine::ArchiveConf conf("swap", "days:100;disk:200");
        auto criterias = conf.GetCriterias();
        ASSERT_TRUE(criterias.size() == 2);
        ASSERT_TRUE(criterias["days"] == 100);
        ASSERT_TRUE(criterias["disk"] == 200);
    }
}

TEST_F(DBTest2, ARHIVE_DISK_CHECK) {

    static const std::string group_name = "test_group";
    static const int group_dim = 256;
    uint64_t size;

    engine::meta::TableSchema group_info;
    group_info.dimension_ = group_dim;
    group_info.table_id_ = group_name;
    group_info.engine_type_ = (int)engine::EngineType::FAISS_IVFFLAT;
    engine::Status stat = db_->CreateTable(group_info);

    engine::meta::TableSchema group_info_get;
    group_info_get.table_id_ = group_name;
    stat = db_->DescribeTable(group_info_get);
    ASSERT_STATS(stat);
    ASSERT_EQ(group_info_get.dimension_, group_dim);

    engine::IDNumbers vector_ids;
    engine::IDNumbers target_ids;

    db_->Size(size);
    int d = 256;
    int nb = 20;
    float *xb = new float[d * nb];
    for(int i = 0; i < nb; i++) {
        for(int j = 0; j < d; j++) xb[d * i + j] = drand48();
        xb[d * i] += i / 2000.;
    }

    int loop = 100000;

    for (auto i=0; i<loop; ++i) {
        db_->InsertVectors(group_name, nb, xb, vector_ids);
        std::this_thread::sleep_for(std::chrono::microseconds(1));
    }

    std::this_thread::sleep_for(std::chrono::seconds(1));

    db_->Size(size);
    LOG(DEBUG) << "size=" << size;
    ASSERT_TRUE(size < 1 * engine::meta::G);

    delete [] xb;
};


TEST_F(DBTest, DB_TEST) {



    static const std::string group_name = "test_group";
    static const int group_dim = 256;

    engine::meta::TableSchema group_info;
    group_info.dimension_ = group_dim;
    group_info.table_id_ = group_name;
    group_info.engine_type_ = (int)engine::EngineType::FAISS_IVFFLAT;
    engine::Status stat = db_->CreateTable(group_info);

    engine::meta::TableSchema group_info_get;
    group_info_get.table_id_ = group_name;
    stat = db_->DescribeTable(group_info_get);
    ASSERT_STATS(stat);
    ASSERT_EQ(group_info_get.dimension_, group_dim);

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
        uint64_t count = 0;
        uint64_t prev_count = 0;

        for (auto j=0; j<10; ++j) {
            ss.str("");
            db_->Size(count);
            prev_count = count;

            START_TIMER;
            stat = db_->Query(group_name, k, qb, qxb, results);
            ss << "Search " << j << " With Size " << count/engine::meta::M << " M";
            STOP_TIMER(ss.str());

            ASSERT_STATS(stat);
            for (auto k=0; k<qb; ++k) {
                ASSERT_EQ(results[k][0].first, target_ids[k]);
                ss.str("");
                ss << "Result [" << k << "]:";
                for (auto result : results[k]) {
                    ss << result.first << " ";
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
            db_->InsertVectors(group_name, qb, qxb, target_ids);
            ASSERT_EQ(target_ids.size(), qb);
        } else {
            db_->InsertVectors(group_name, nb, xb, vector_ids);
        }
        std::this_thread::sleep_for(std::chrono::microseconds(1));
    }

    search.join();

    delete [] xb;
    delete [] qxb;
};

TEST_F(DBTest, SEARCH_TEST) {
    static const std::string group_name = "test_group";
    static const int group_dim = 256;

    engine::meta::TableSchema group_info;
    group_info.dimension_ = group_dim;
    group_info.table_id_ = group_name;
    group_info.engine_type_ = (int)engine::EngineType::FAISS_IVFFLAT;
    engine::Status stat = db_->CreateTable(group_info);

    engine::meta::TableSchema group_info_get;
    group_info_get.table_id_ = group_name;
    stat = db_->DescribeTable(group_info_get);
    ASSERT_STATS(stat);
    ASSERT_EQ(group_info_get.dimension_, group_dim);

    // prepare raw data
    size_t nb = 250000;
    size_t nq = 10;
    size_t k = 5;
    std::vector<float> xb(nb*group_dim);
    std::vector<float> xq(nq*group_dim);
    std::vector<long> ids(nb);

    std::random_device rd;
    std::mt19937 gen(rd());
    std::uniform_real_distribution<> dis_xt(-1.0, 1.0);
    for (size_t i = 0; i < nb*group_dim; i++) {
        xb[i] = dis_xt(gen);
        if (i < nb){
            ids[i] = i;
        }
    }
    for (size_t i = 0; i < nq*group_dim; i++) {
        xq[i] = dis_xt(gen);
    }

    // result data
    //std::vector<long> nns_gt(k*nq);
    std::vector<long> nns(k*nq);  // nns = nearst neg search
    //std::vector<float> dis_gt(k*nq);
    std::vector<float> dis(k*nq);

    // insert data
    const int batch_size = 100;
    for (int j = 0; j < nb / batch_size; ++j) {
        stat = db_->InsertVectors(group_name, batch_size, xb.data()+batch_size*j*group_dim, ids);
        if (j == 200){ sleep(1);}
        ASSERT_STATS(stat);
    }

    sleep(2); // wait until build index finish

    engine::QueryResults results;
    stat = db_->Query(group_name, k, nq, xq.data(), results);
    ASSERT_STATS(stat);

    // TODO(linxj): add groundTruth assert
};

