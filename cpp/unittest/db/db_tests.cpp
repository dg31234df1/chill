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

#include "db/utils.h"
#include "db/DB.h"
#include "db/DBImpl.h"
#include "db/meta/MetaConsts.h"
#include "db/DBFactory.h"
#include "cache/CpuCacheMgr.h"
#include "utils/CommonUtil.h"

#include <gtest/gtest.h>
#include <boost/filesystem.hpp>
#include <thread>
#include <random>

namespace {

namespace ms = zilliz::milvus;

static const char *TABLE_NAME = "test_group";
static constexpr int64_t TABLE_DIM = 256;
static constexpr int64_t VECTOR_COUNT = 25000;
static constexpr int64_t INSERT_LOOP = 1000;
static constexpr int64_t SECONDS_EACH_HOUR = 3600;
static constexpr int64_t DAY_SECONDS = 24 * 60 * 60;

ms::engine::meta::TableSchema
BuildTableSchema() {
    ms::engine::meta::TableSchema table_info;
    table_info.dimension_ = TABLE_DIM;
    table_info.table_id_ = TABLE_NAME;
    return table_info;
}

void
BuildVectors(int64_t n, std::vector<float> &vectors) {
    vectors.clear();
    vectors.resize(n * TABLE_DIM);
    float *data = vectors.data();
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < TABLE_DIM; j++) data[TABLE_DIM * i + j] = drand48();
        data[TABLE_DIM * i] += i / 2000.;
    }
}

std::string
CurrentTmDate(int64_t offset_day = 0) {
    time_t tt;
    time(&tt);
    tt = tt + 8 * SECONDS_EACH_HOUR;
    tt = tt + 24 * SECONDS_EACH_HOUR * offset_day;
    tm t;
    gmtime_r(&tt, &t);

    std::string str = std::to_string(t.tm_year + 1900) + "-" + std::to_string(t.tm_mon + 1)
        + "-" + std::to_string(t.tm_mday);

    return str;
}

void
ConvertTimeRangeToDBDates(const std::string &start_value,
                          const std::string &end_value,
                          std::vector<ms::engine::meta::DateT> &dates) {
    dates.clear();

    time_t tt_start, tt_end;
    tm tm_start, tm_end;
    if (!zilliz::milvus::server::CommonUtil::TimeStrToTime(start_value, tt_start, tm_start)) {
        return;
    }

    if (!zilliz::milvus::server::CommonUtil::TimeStrToTime(end_value, tt_end, tm_end)) {
        return;
    }

    int64_t days = (tt_end > tt_start) ? (tt_end - tt_start) / DAY_SECONDS : (tt_start - tt_end) /
        DAY_SECONDS;
    if (days == 0) {
        return;
    }

    for (int64_t i = 0; i < days; i++) {
        time_t tt_day = tt_start + DAY_SECONDS * i;
        tm tm_day;
        zilliz::milvus::server::CommonUtil::ConvertTime(tt_day, tm_day);

        int64_t date = tm_day.tm_year * 10000 + tm_day.tm_mon * 100 +
            tm_day.tm_mday;//according to db logic
        dates.push_back(date);
    }
}

} // namespace

TEST_F(DBTest, CONFIG_TEST) {
    {
        ASSERT_ANY_THROW(ms::engine::ArchiveConf conf("wrong"));
        /* EXPECT_DEATH(engine::ArchiveConf conf("wrong"), ""); */
    }
    {
        ms::engine::ArchiveConf conf("delete");
        ASSERT_EQ(conf.GetType(), "delete");
        auto criterias = conf.GetCriterias();
        ASSERT_EQ(criterias.size(), 0);
    }
    {
        ms::engine::ArchiveConf conf("swap");
        ASSERT_EQ(conf.GetType(), "swap");
        auto criterias = conf.GetCriterias();
        ASSERT_EQ(criterias.size(), 0);
    }
    {
        ASSERT_ANY_THROW(ms::engine::ArchiveConf conf1("swap", "disk:"));
        ASSERT_ANY_THROW(ms::engine::ArchiveConf conf2("swap", "disk:a"));
        ms::engine::ArchiveConf conf("swap", "disk:1024");
        auto criterias = conf.GetCriterias();
        ASSERT_EQ(criterias.size(), 1);
        ASSERT_EQ(criterias["disk"], 1024);
    }
    {
        ASSERT_ANY_THROW(ms::engine::ArchiveConf conf1("swap", "days:"));
        ASSERT_ANY_THROW(ms::engine::ArchiveConf conf2("swap", "days:a"));
        ms::engine::ArchiveConf conf("swap", "days:100");
        auto criterias = conf.GetCriterias();
        ASSERT_EQ(criterias.size(), 1);
        ASSERT_EQ(criterias["days"], 100);
    }
    {
        ASSERT_ANY_THROW(ms::engine::ArchiveConf conf1("swap", "days:"));
        ASSERT_ANY_THROW(ms::engine::ArchiveConf conf2("swap", "days:a"));
        ms::engine::ArchiveConf conf("swap", "days:100;disk:200");
        auto criterias = conf.GetCriterias();
        ASSERT_EQ(criterias.size(), 2);
        ASSERT_EQ(criterias["days"], 100);
        ASSERT_EQ(criterias["disk"], 200);
    }
}

TEST_F(DBTest, DB_TEST) {
    ms::engine::meta::TableSchema table_info = BuildTableSchema();
    auto stat = db_->CreateTable(table_info);

    ms::engine::meta::TableSchema table_info_get;
    table_info_get.table_id_ = TABLE_NAME;
    stat = db_->DescribeTable(table_info_get);
    ASSERT_TRUE(stat.ok());
    ASSERT_EQ(table_info_get.dimension_, TABLE_DIM);

    ms::engine::IDNumbers vector_ids;
    ms::engine::IDNumbers target_ids;

    int64_t nb = 50;
    std::vector<float> xb;
    BuildVectors(nb, xb);

    int64_t qb = 5;
    std::vector<float> qxb;
    BuildVectors(qb, qxb);

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
            stat = db_->Query(TABLE_NAME, k, qb, 10, qxb.data(), results);
            ss << "Search " << j << " With Size " << count / ms::engine::meta::M << " M";
            STOP_TIMER(ss.str());

            ASSERT_TRUE(stat.ok());
            for (auto k = 0; k < qb; ++k) {
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

    int loop = INSERT_LOOP;

    for (auto i = 0; i < loop; ++i) {
        if (i == 40) {
            db_->InsertVectors(TABLE_NAME, qb, qxb.data(), target_ids);
            ASSERT_EQ(target_ids.size(), qb);
        } else {
            db_->InsertVectors(TABLE_NAME, nb, xb.data(), vector_ids);
        }
        std::this_thread::sleep_for(std::chrono::microseconds(1));
    }

    search.join();

    uint64_t count;
    stat = db_->GetTableRowCount(TABLE_NAME, count);
    ASSERT_TRUE(stat.ok());
    ASSERT_GT(count, 0);
}

TEST_F(DBTest, SEARCH_TEST) {
    ms::engine::meta::TableSchema table_info = BuildTableSchema();
    auto stat = db_->CreateTable(table_info);

    ms::engine::meta::TableSchema table_info_get;
    table_info_get.table_id_ = TABLE_NAME;
    stat = db_->DescribeTable(table_info_get);
    ASSERT_TRUE(stat.ok());
    ASSERT_EQ(table_info_get.dimension_, TABLE_DIM);

    // prepare raw data
    size_t nb = VECTOR_COUNT;
    size_t nq = 10;
    size_t k = 5;
    std::vector<float> xb(nb * TABLE_DIM);
    std::vector<float> xq(nq * TABLE_DIM);
    std::vector<int64_t> ids(nb);

    std::random_device rd;
    std::mt19937 gen(rd());
    std::uniform_real_distribution<> dis_xt(-1.0, 1.0);
    for (size_t i = 0; i < nb * TABLE_DIM; i++) {
        xb[i] = dis_xt(gen);
        if (i < nb) {
            ids[i] = i;
        }
    }
    for (size_t i = 0; i < nq * TABLE_DIM; i++) {
        xq[i] = dis_xt(gen);
    }

    // result data
    //std::vector<long> nns_gt(k*nq);
    std::vector<int64_t> nns(k * nq);  // nns = nearst neg search
    //std::vector<float> dis_gt(k*nq);
    std::vector<float> dis(k * nq);

    // insert data
    const int batch_size = 100;
    for (int j = 0; j < nb / batch_size; ++j) {
        stat = db_->InsertVectors(TABLE_NAME, batch_size, xb.data() + batch_size * j * TABLE_DIM, ids);
        if (j == 200) { sleep(1); }
        ASSERT_TRUE(stat.ok());
    }

    ms::engine::TableIndex index;
    index.engine_type_ = (int) ms::engine::EngineType::FAISS_IDMAP;
    db_->CreateIndex(TABLE_NAME, index); // wait until build index finish

    {
        ms::engine::QueryResults results;
        stat = db_->Query(TABLE_NAME, k, nq, 10, xq.data(), results);
        ASSERT_TRUE(stat.ok());
    }

    {//search by specify index file
        ms::engine::meta::DatesT dates;
        std::vector<std::string> file_ids = {"1", "2", "3", "4", "5", "6"};
        ms::engine::QueryResults results;
        stat = db_->Query(TABLE_NAME, file_ids, k, nq, 10, xq.data(), dates, results);
        ASSERT_TRUE(stat.ok());
    }

    // TODO(lxj): add groundTruth assert
}

TEST_F(DBTest, PRELOADTABLE_TEST) {
    ms::engine::meta::TableSchema table_info = BuildTableSchema();
    auto stat = db_->CreateTable(table_info);

    ms::engine::meta::TableSchema table_info_get;
    table_info_get.table_id_ = TABLE_NAME;
    stat = db_->DescribeTable(table_info_get);
    ASSERT_TRUE(stat.ok());
    ASSERT_EQ(table_info_get.dimension_, TABLE_DIM);

    int64_t nb = VECTOR_COUNT;
    std::vector<float> xb;
    BuildVectors(nb, xb);

    int loop = 5;
    for (auto i = 0; i < loop; ++i) {
        ms::engine::IDNumbers vector_ids;
        db_->InsertVectors(TABLE_NAME, nb, xb.data(), vector_ids);
        ASSERT_EQ(vector_ids.size(), nb);
    }

    ms::engine::TableIndex index;
    index.engine_type_ = (int) ms::engine::EngineType::FAISS_IDMAP;
    db_->CreateIndex(TABLE_NAME, index); // wait until build index finish

    int64_t prev_cache_usage = ms::cache::CpuCacheMgr::GetInstance()->CacheUsage();
    stat = db_->PreloadTable(TABLE_NAME);
    ASSERT_TRUE(stat.ok());
    int64_t cur_cache_usage = ms::cache::CpuCacheMgr::GetInstance()->CacheUsage();
    ASSERT_TRUE(prev_cache_usage < cur_cache_usage);
}

TEST_F(DBTest, SHUTDOWN_TEST) {
    db_->Stop();

    ms::engine::meta::TableSchema table_info = BuildTableSchema();
    auto stat = db_->CreateTable(table_info);
    ASSERT_FALSE(stat.ok());

    stat = db_->DescribeTable(table_info);
    ASSERT_FALSE(stat.ok());

    bool has_table = false;
    stat = db_->HasTable(table_info.table_id_, has_table);
    ASSERT_FALSE(stat.ok());

    ms::engine::IDNumbers ids;
    stat = db_->InsertVectors(table_info.table_id_, 0, nullptr, ids);
    ASSERT_FALSE(stat.ok());

    stat = db_->PreloadTable(table_info.table_id_);
    ASSERT_FALSE(stat.ok());

    uint64_t row_count = 0;
    stat = db_->GetTableRowCount(table_info.table_id_, row_count);
    ASSERT_FALSE(stat.ok());

    ms::engine::TableIndex index;
    stat = db_->CreateIndex(table_info.table_id_, index);
    ASSERT_FALSE(stat.ok());

    stat = db_->DescribeIndex(table_info.table_id_, index);
    ASSERT_FALSE(stat.ok());

    ms::engine::meta::DatesT dates;
    ms::engine::QueryResults results;
    stat = db_->Query(table_info.table_id_, 1, 1, 1, nullptr, dates, results);
    ASSERT_FALSE(stat.ok());
    std::vector<std::string> file_ids;
    stat = db_->Query(table_info.table_id_, file_ids, 1, 1, 1, nullptr, dates, results);
    ASSERT_FALSE(stat.ok());

    stat = db_->DeleteTable(table_info.table_id_, dates);
    ASSERT_FALSE(stat.ok());
}

TEST_F(DBTest, INDEX_TEST) {
    ms::engine::meta::TableSchema table_info = BuildTableSchema();
    auto stat = db_->CreateTable(table_info);

    int64_t nb = VECTOR_COUNT;
    std::vector<float> xb;
    BuildVectors(nb, xb);

    ms::engine::IDNumbers vector_ids;
    db_->InsertVectors(TABLE_NAME, nb, xb.data(), vector_ids);
    ASSERT_EQ(vector_ids.size(), nb);

    ms::engine::TableIndex index;
    index.engine_type_ = (int) ms::engine::EngineType::FAISS_IVFSQ8;
    index.metric_type_ = (int) ms::engine::MetricType::IP;
    stat = db_->CreateIndex(table_info.table_id_, index);
    ASSERT_TRUE(stat.ok());

    ms::engine::TableIndex index_out;
    stat = db_->DescribeIndex(table_info.table_id_, index_out);
    ASSERT_TRUE(stat.ok());
    ASSERT_EQ(index.engine_type_, index_out.engine_type_);
    ASSERT_EQ(index.nlist_, index_out.nlist_);
    ASSERT_EQ(table_info.metric_type_, index_out.metric_type_);

    stat = db_->DropIndex(table_info.table_id_);
    ASSERT_TRUE(stat.ok());
}

TEST_F(DBTest2, ARHIVE_DISK_CHECK) {
    ms::engine::meta::TableSchema table_info = BuildTableSchema();
    auto stat = db_->CreateTable(table_info);

    std::vector<ms::engine::meta::TableSchema> table_schema_array;
    stat = db_->AllTables(table_schema_array);
    ASSERT_TRUE(stat.ok());
    bool bfound = false;
    for (auto &schema : table_schema_array) {
        if (schema.table_id_ == TABLE_NAME) {
            bfound = true;
            break;
        }
    }
    ASSERT_TRUE(bfound);

    ms::engine::meta::TableSchema table_info_get;
    table_info_get.table_id_ = TABLE_NAME;
    stat = db_->DescribeTable(table_info_get);
    ASSERT_TRUE(stat.ok());
    ASSERT_EQ(table_info_get.dimension_, TABLE_DIM);

    uint64_t size;
    db_->Size(size);

    int64_t nb = 10;
    std::vector<float> xb;
    BuildVectors(nb, xb);

    int loop = INSERT_LOOP;
    for (auto i = 0; i < loop; ++i) {
        ms::engine::IDNumbers vector_ids;
        db_->InsertVectors(TABLE_NAME, nb, xb.data(), vector_ids);
        std::this_thread::sleep_for(std::chrono::microseconds(1));
    }

    std::this_thread::sleep_for(std::chrono::seconds(1));

    db_->Size(size);
    LOG(DEBUG) << "size=" << size;
    ASSERT_LE(size, 1 * ms::engine::meta::G);
}

TEST_F(DBTest2, DELETE_TEST) {
    ms::engine::meta::TableSchema table_info = BuildTableSchema();
    auto stat = db_->CreateTable(table_info);

    ms::engine::meta::TableSchema table_info_get;
    table_info_get.table_id_ = TABLE_NAME;
    stat = db_->DescribeTable(table_info_get);
    ASSERT_TRUE(stat.ok());

    bool has_table = false;
    db_->HasTable(TABLE_NAME, has_table);
    ASSERT_TRUE(has_table);

    uint64_t size;
    db_->Size(size);

    int64_t nb = VECTOR_COUNT;
    std::vector<float> xb;
    BuildVectors(nb, xb);

    ms::engine::IDNumbers vector_ids;
    stat = db_->InsertVectors(TABLE_NAME, nb, xb.data(), vector_ids);
    ms::engine::TableIndex index;
    stat = db_->CreateIndex(TABLE_NAME, index);

    std::vector<ms::engine::meta::DateT> dates;
    stat = db_->DeleteTable(TABLE_NAME, dates);
    std::this_thread::sleep_for(std::chrono::seconds(2));
    ASSERT_TRUE(stat.ok());

    db_->HasTable(TABLE_NAME, has_table);
    ASSERT_FALSE(has_table);
}

TEST_F(DBTest2, DELETE_BY_RANGE_TEST) {
    ms::engine::meta::TableSchema table_info = BuildTableSchema();
    auto stat = db_->CreateTable(table_info);

    ms::engine::meta::TableSchema table_info_get;
    table_info_get.table_id_ = TABLE_NAME;
    stat = db_->DescribeTable(table_info_get);
    ASSERT_TRUE(stat.ok());

    bool has_table = false;
    db_->HasTable(TABLE_NAME, has_table);
    ASSERT_TRUE(has_table);

    uint64_t size;
    db_->Size(size);
    ASSERT_EQ(size, 0UL);

    int64_t nb = VECTOR_COUNT;
    std::vector<float> xb;
    BuildVectors(nb, xb);

    ms::engine::IDNumbers vector_ids;
    stat = db_->InsertVectors(TABLE_NAME, nb, xb.data(), vector_ids);
    ms::engine::TableIndex index;
    stat = db_->CreateIndex(TABLE_NAME, index);

    db_->Size(size);
    ASSERT_NE(size, 0UL);

    std::vector<ms::engine::meta::DateT> dates;
    std::string start_value = CurrentTmDate();
    std::string end_value = CurrentTmDate(1);
    ConvertTimeRangeToDBDates(start_value, end_value, dates);

    stat = db_->DeleteTable(TABLE_NAME, dates);
    ASSERT_TRUE(stat.ok());

    uint64_t row_count = 0;
    db_->GetTableRowCount(TABLE_NAME, row_count);
    ASSERT_EQ(row_count, 0UL);
}
