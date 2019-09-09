/*******************************************************************************
 * Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 ******************************************************************************/

#include "scheduler/resource/CpuResource.h"
#include "scheduler/resource/GpuResource.h"
#include "scheduler/resource/DiskResource.h"
#include "scheduler/resource/TestResource.h"
#include "scheduler/task/TestTask.h"
#include "scheduler/ResourceMgr.h"
#include <gtest/gtest.h>


namespace zilliz {
namespace milvus {
namespace engine {


/************ ResourceMgrBaseTest ************/
class ResourceMgrBaseTest : public testing::Test {
protected:
    void
    SetUp() override {
        empty_mgr_ = std::make_shared<ResourceMgr>();
        mgr1_ = std::make_shared<ResourceMgr>();
        disk_res = std::make_shared<DiskResource>("disk", 0, true, false);
        cpu_res = std::make_shared<CpuResource>("cpu", 1, true, false);
        gpu_res = std::make_shared<GpuResource>("gpu", 2, true, true);
        mgr1_->Add(ResourcePtr(disk_res));
        mgr1_->Add(ResourcePtr(cpu_res));
        mgr1_->Add(ResourcePtr(gpu_res));
    }

    void
    TearDown() override {
    }

    ResourceMgrPtr empty_mgr_;
    ResourceMgrPtr mgr1_;
    ResourcePtr disk_res;
    ResourcePtr cpu_res;
    ResourcePtr gpu_res;
};

TEST_F(ResourceMgrBaseTest, add) {
    auto resource = std::make_shared<TestResource>("test", 0, true, true);
    auto ret = empty_mgr_->Add(ResourcePtr(resource));
    ASSERT_EQ(ret.lock(), resource);
}

TEST_F(ResourceMgrBaseTest, add_disk) {
    auto resource = std::make_shared<DiskResource>("disk", 0, true, true);
    auto ret = empty_mgr_->Add(ResourcePtr(resource));
    ASSERT_EQ(ret.lock(), resource);
}

TEST_F(ResourceMgrBaseTest, connect) {
    auto resource1 = std::make_shared<TestResource>("resource1", 0, true, true);
    auto resource2 = std::make_shared<TestResource>("resource2", 2, true, true);
    empty_mgr_->Add(resource1);
    empty_mgr_->Add(resource2);
    Connection io("io", 500.0);
    ASSERT_TRUE(empty_mgr_->Connect("resource1", "resource2", io));
}


TEST_F(ResourceMgrBaseTest, invalid_connect) {
    auto resource1 = std::make_shared<TestResource>("resource1", 0, true, true);
    auto resource2 = std::make_shared<TestResource>("resource2", 2, true, true);
    empty_mgr_->Add(resource1);
    empty_mgr_->Add(resource2);
    Connection io("io", 500.0);
    ASSERT_FALSE(empty_mgr_->Connect("xx", "yy", io));
}


TEST_F(ResourceMgrBaseTest, clear) {
    ASSERT_EQ(mgr1_->GetNumOfResource(), 3);
    mgr1_->Clear();
    ASSERT_EQ(mgr1_->GetNumOfResource(), 0);
}

TEST_F(ResourceMgrBaseTest, get_disk_resources) {
    auto disks = mgr1_->GetDiskResources();
    ASSERT_EQ(disks.size(), 1);
    ASSERT_EQ(disks[0].lock(), disk_res);
}

TEST_F(ResourceMgrBaseTest, get_all_resources) {
    bool disk = false, cpu = false, gpu = false;
    auto resources = mgr1_->GetAllResources();
    ASSERT_EQ(resources.size(), 3);
    for (auto &res : resources) {
        if (res->type() == ResourceType::DISK) disk = true;
        if (res->type() == ResourceType::CPU) cpu = true;
        if (res->type() == ResourceType::GPU) gpu = true;
    }

    ASSERT_TRUE(disk);
    ASSERT_TRUE(cpu);
    ASSERT_TRUE(gpu);
}

TEST_F(ResourceMgrBaseTest, get_compute_resources) {
    auto compute_resources = mgr1_->GetComputeResources();
    ASSERT_EQ(compute_resources.size(), 1);
    ASSERT_EQ(compute_resources[0], gpu_res);
}

TEST_F(ResourceMgrBaseTest, get_resource_by_type_and_deviceid) {
    auto cpu = mgr1_->GetResource(ResourceType::CPU, 1);
    ASSERT_EQ(cpu, cpu_res);

    auto invalid = mgr1_->GetResource(ResourceType::GPU, 1);
    ASSERT_EQ(invalid, nullptr);
}

TEST_F(ResourceMgrBaseTest, get_resource_by_name) {
    auto disk = mgr1_->GetResource("disk");
    ASSERT_EQ(disk, disk_res);

    auto invalid = mgr1_->GetResource("invalid");
    ASSERT_EQ(invalid, nullptr);
}

TEST_F(ResourceMgrBaseTest, get_num_of_resource) {
    ASSERT_EQ(empty_mgr_->GetNumOfResource(), 0);
    ASSERT_EQ(mgr1_->GetNumOfResource(), 3);
}

TEST_F(ResourceMgrBaseTest, get_num_of_compute_resource) {
    ASSERT_EQ(empty_mgr_->GetNumOfComputeResource(), 0);
    ASSERT_EQ(mgr1_->GetNumOfComputeResource(), 1);
}

TEST_F(ResourceMgrBaseTest, get_num_of_gpu_resource) {
    ASSERT_EQ(empty_mgr_->GetNumGpuResource(), 0);
    ASSERT_EQ(mgr1_->GetNumGpuResource(), 1);
}

TEST_F(ResourceMgrBaseTest, dump) {
    ASSERT_FALSE(mgr1_->Dump().empty());
}

TEST_F(ResourceMgrBaseTest, dump_tasktables) {
    ASSERT_FALSE(mgr1_->DumpTaskTables().empty());
}

/************ ResourceMgrAdvanceTest ************/

class ResourceMgrAdvanceTest : public testing::Test {
    protected:
    void
    SetUp() override {
        mgr1_ = std::make_shared<ResourceMgr>();
        disk_res = std::make_shared<DiskResource>("disk", 0, true, false);
        mgr1_->Add(ResourcePtr(disk_res));
        mgr1_->Start();
    }

    void
    TearDown() override {
        mgr1_->Stop();
    }

    ResourceMgrPtr mgr1_;
    ResourcePtr disk_res;
};

TEST_F(ResourceMgrAdvanceTest, register_subscriber) {
    bool flag = false;
    auto callback = [&](EventPtr event) {
        flag = true;
    };
    mgr1_->RegisterSubscriber(callback);
    TableFileSchemaPtr dummy = nullptr;
    disk_res->task_table().Put(std::make_shared<TestTask>(dummy));
    sleep(1);
    ASSERT_TRUE(flag);
}


}
}
}
