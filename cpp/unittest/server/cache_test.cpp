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

#include <gtest/gtest.h>
#include "cache/CpuCacheMgr.h"
#include "cache/GpuCacheMgr.h"
#include "utils/Error.h"
#include "wrapper/VecIndex.h"

namespace {

namespace ms = zilliz::milvus;

class InvalidCacheMgr : public ms::cache::CacheMgr<ms::cache::DataObjPtr> {
 public:
    InvalidCacheMgr() {
    }
};

class LessItemCacheMgr : public ms::cache::CacheMgr<ms::cache::DataObjPtr> {
 public:
    LessItemCacheMgr() {
        cache_ = std::make_shared<ms::cache::Cache<ms::cache::DataObjPtr>>(1UL << 12, 10);
    }
};

class MockVecIndex : public ms::engine::VecIndex {
 public:
    MockVecIndex(int64_t dim, int64_t total)
        : dimension_(dim),
          ntotal_(total) {
    }

    virtual ms::Status BuildAll(const int64_t &nb,
                                const float *xb,
                                const int64_t *ids,
                                const ms::engine::Config &cfg,
                                const int64_t &nt = 0,
                                const float *xt = nullptr) {
        return ms::Status();
    }

    ms::engine::VecIndexPtr Clone() override {
        return zilliz::milvus::engine::VecIndexPtr();
    }

    int64_t GetDeviceId() override {
        return 0;
    }

    ms::engine::IndexType GetType() override {
        return ms::engine::IndexType::INVALID;
    }

    virtual ms::Status Add(const int64_t &nb,
                           const float *xb,
                           const int64_t *ids,
                           const ms::engine::Config &cfg = ms::engine::Config()) {
        return ms::Status();
    }

    virtual ms::Status Search(const int64_t &nq,
                              const float *xq,
                              float *dist,
                              int64_t *ids,
                              const ms::engine::Config &cfg = ms::engine::Config()) {
        return ms::Status();
    }

    ms::engine::VecIndexPtr CopyToGpu(const int64_t &device_id,
                                      const ms::engine::Config &cfg) override {
        return nullptr;
    }

    ms::engine::VecIndexPtr CopyToCpu(const ms::engine::Config &cfg) override {
        return nullptr;
    }

    virtual int64_t Dimension() {
        return dimension_;
    }

    virtual int64_t Count() {
        return ntotal_;
    }

    virtual zilliz::knowhere::BinarySet Serialize() {
        zilliz::knowhere::BinarySet binset;
        return binset;
    }

    virtual ms::Status Load(const zilliz::knowhere::BinarySet &index_binary) {
        return ms::Status();
    }

 public:
    int64_t dimension_ = 256;
    int64_t ntotal_ = 0;
};

} // namespace

TEST(CacheTest, DUMMY_TEST) {
    ms::engine::Config cfg;
    MockVecIndex mock_index(256, 1000);
    mock_index.Dimension();
    mock_index.Count();
    mock_index.Add(1, nullptr, nullptr);
    mock_index.BuildAll(1, nullptr, nullptr, cfg);
    mock_index.Search(1, nullptr, nullptr, nullptr, cfg);
    mock_index.Clone();
    mock_index.CopyToCpu(cfg);
    mock_index.CopyToGpu(1, cfg);
    mock_index.GetDeviceId();
    mock_index.GetType();
    zilliz::knowhere::BinarySet index_binary;
    mock_index.Load(index_binary);
    mock_index.Serialize();
}

TEST(CacheTest, CPU_CACHE_TEST) {
    auto cpu_mgr = ms::cache::CpuCacheMgr::GetInstance();

    const int64_t gbyte = 1024 * 1024 * 1024;
    int64_t g_num = 16;
    int64_t cap = g_num * gbyte;
    cpu_mgr->SetCapacity(cap);
    ASSERT_EQ(cpu_mgr->CacheCapacity(), cap);

    uint64_t item_count = 20;
    for (uint64_t i = 0; i < item_count; i++) {
        //each vector is 1k byte, total size less than 1G
        ms::engine::VecIndexPtr mock_index = std::make_shared<MockVecIndex>(256, 1000000);
        ms::cache::DataObjPtr data_obj = std::make_shared<ms::cache::DataObj>(mock_index);
        cpu_mgr->InsertItem("index_" + std::to_string(i), data_obj);
    }
    ASSERT_LT(cpu_mgr->ItemCount(), g_num);

    auto obj = cpu_mgr->GetIndex("index_0");
    ASSERT_TRUE(obj == nullptr);

    obj = cpu_mgr->GetIndex("index_" + std::to_string(item_count - 1));
    ASSERT_TRUE(obj != nullptr);

    {
        std::string item = "index_15";
        ASSERT_TRUE(cpu_mgr->ItemExists(item));
        cpu_mgr->EraseItem(item);
        ASSERT_FALSE(cpu_mgr->ItemExists(item));
    }

    {
        g_num = 5;
        cpu_mgr->SetCapacity(g_num * gbyte);

        //each vector is 1k byte, total size less than 6G
        ms::engine::VecIndexPtr mock_index = std::make_shared<MockVecIndex>(256, 6000000);
        ms::cache::DataObjPtr data_obj = std::make_shared<ms::cache::DataObj>(mock_index);
        cpu_mgr->InsertItem("index_6g", data_obj);
        ASSERT_TRUE(cpu_mgr->ItemExists("index_6g"));
    }

    cpu_mgr->PrintInfo();
}

TEST(CacheTest, GPU_CACHE_TEST) {
    auto gpu_mgr = ms::cache::GpuCacheMgr::GetInstance(0);

    for (int i = 0; i < 20; i++) {
        //each vector is 1k byte
        ms::engine::VecIndexPtr mock_index = std::make_shared<MockVecIndex>(256, 1000);
        ms::cache::DataObjPtr data_obj = std::make_shared<ms::cache::DataObj>(mock_index);
        gpu_mgr->InsertItem("index_" + std::to_string(i), data_obj);
    }

    auto obj = gpu_mgr->GetItem("index_0");

    gpu_mgr->ClearCache();
    ASSERT_EQ(gpu_mgr->ItemCount(), 0);

    for (auto i = 0; i < 3; i++) {
        // TODO(myh): use gpu index to mock
        //each vector is 1k byte, total size less than 2G
        ms::engine::VecIndexPtr mock_index = std::make_shared<MockVecIndex>(256, 2000000);
        ms::cache::DataObjPtr data_obj = std::make_shared<ms::cache::DataObj>(mock_index);
        std::cout << data_obj->size() << std::endl;
        gpu_mgr->InsertItem("index_" + std::to_string(i), data_obj);
    }

    gpu_mgr->ClearCache();
    ASSERT_EQ(gpu_mgr->ItemCount(), 0);
}

TEST(CacheTest, INVALID_TEST) {
    {
        InvalidCacheMgr mgr;
        ASSERT_EQ(mgr.ItemCount(), 0);
        ASSERT_FALSE(mgr.ItemExists("test"));
        ASSERT_EQ(mgr.GetItem("test"), nullptr);

        mgr.InsertItem("test", ms::cache::DataObjPtr());
        mgr.InsertItem("test", nullptr);
        mgr.EraseItem("test");
        mgr.PrintInfo();
        mgr.ClearCache();
        mgr.SetCapacity(100);
        ASSERT_EQ(mgr.CacheCapacity(), 0);
        ASSERT_EQ(mgr.CacheUsage(), 0);
    }

    {
        LessItemCacheMgr mgr;
        for (int i = 0; i < 20; i++) {
            //each vector is 1k byte
            ms::engine::VecIndexPtr mock_index = std::make_shared<MockVecIndex>(256, 2);
            ms::cache::DataObjPtr data_obj = std::make_shared<ms::cache::DataObj>(mock_index);
            mgr.InsertItem("index_" + std::to_string(i), data_obj);
        }
        ASSERT_EQ(mgr.GetItem("index_0"), nullptr);
    }
}
