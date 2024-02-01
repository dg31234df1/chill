// Copyright (C) 2019-2020 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License

#include <boost/filesystem/operations.hpp>
#include <chrono>
#include <arrow/array/builder_binary.h>
#include <arrow/array/builder_primitive.h>
#include <arrow/record_batch.h>
#include <arrow/type.h>
#include <arrow/type_fwd.h>
#include <gtest/gtest.h>
#include <cstdint>
#include <memory>
#include <memory_resource>
#include <string>
#include <fstream>
#include <vector>
#include <unistd.h>

#include "common/EasyAssert.h"
#include "common/FieldDataInterface.h"
#include "common/Slice.h"
#include "common/Common.h"
#include "common/Types.h"
#include "storage/ChunkManager.h"
#include "storage/DataCodec.h"
#include "storage/InsertData.h"
#include "storage/ThreadPool.h"
#include "storage/Types.h"
#include "storage/options.h"
#include "storage/schema.h"
#include "storage/space.h"
#include "storage/Util.h"
#include "storage/DiskFileManagerImpl.h"
#include "storage/LocalChunkManagerSingleton.h"

#include "test_utils/storage_test_utils.h"

using namespace std;
using namespace milvus;
using namespace milvus::storage;
using namespace knowhere;

class DiskAnnFileManagerTest : public testing::Test {
 public:
    DiskAnnFileManagerTest() {
    }
    ~DiskAnnFileManagerTest() {
    }

    virtual void
    SetUp() {
        cm_ = storage::CreateChunkManager(get_default_local_storage_config());
    }

 protected:
    ChunkManagerPtr cm_;
};

TEST_F(DiskAnnFileManagerTest, AddFilePositiveParallel) {
    auto lcm = LocalChunkManagerSingleton::GetInstance().GetChunkManager();
    std::string indexFilePath = "/tmp/diskann/index_files/1000/index";
    auto exist = lcm->Exist(indexFilePath);
    EXPECT_EQ(exist, false);
    uint64_t index_size = 50 << 20;
    lcm->CreateFile(indexFilePath);
    std::vector<uint8_t> data(index_size);
    lcm->Write(indexFilePath, data.data(), index_size);

    // collection_id: 1, partition_id: 2, segment_id: 3
    // field_id: 100, index_build_id: 1000, index_version: 1
    FieldDataMeta filed_data_meta = {1, 2, 3, 100};
    IndexMeta index_meta = {3, 100, 1000, 1, "index"};

    int64_t slice_size = milvus::FILE_SLICE_SIZE;
    auto diskAnnFileManager = std::make_shared<DiskFileManagerImpl>(
        storage::FileManagerContext(filed_data_meta, index_meta, cm_));
    auto ok = diskAnnFileManager->AddFile(indexFilePath);
    EXPECT_EQ(ok, true);

    auto remote_files_to_size = diskAnnFileManager->GetRemotePathsToFileSize();
    auto num_slice = index_size / slice_size;
    EXPECT_EQ(remote_files_to_size.size(),
              index_size % slice_size == 0 ? num_slice : num_slice + 1);

    std::vector<std::string> remote_files;
    for (auto& file2size : remote_files_to_size) {
        std::cout << file2size.first << std::endl;
        remote_files.emplace_back(file2size.first);
    }
    diskAnnFileManager->CacheIndexToDisk(remote_files);
    auto local_files = diskAnnFileManager->GetLocalFilePaths();
    for (auto& file : local_files) {
        auto file_size = lcm->Size(file);
        auto buf = std::unique_ptr<uint8_t[]>(new uint8_t[file_size]);
        lcm->Read(file, buf.get(), file_size);

        auto index = milvus::storage::CreateFieldData(storage::DataType::INT8);
        index->FillFieldData(buf.get(), file_size);
        auto rows = index->get_num_rows();
        auto rawData = (uint8_t*)(index->Data());

        EXPECT_EQ(rows, index_size);
        EXPECT_EQ(rawData[0], data[0]);
        EXPECT_EQ(rawData[4], data[4]);
    }

    for (auto file : local_files) {
        cm_->Remove(file);
    }
}

int
test_worker(string s) {
    std::cout << s << std::endl;
    std::this_thread::sleep_for(std::chrono::seconds(4));
    std::cout << s << std::endl;
    return 1;
}

int
compute(int a) {
    return a + 10;
}

TEST_F(DiskAnnFileManagerTest, TestThreadPoolBase) {
    auto thread_pool = std::make_shared<milvus::ThreadPool>(10, "test1");
    std::cout << "current thread num" << thread_pool->GetThreadNum()
              << std::endl;
    auto thread_num_1 = thread_pool->GetThreadNum();
    EXPECT_GT(thread_num_1, 0);

    auto fut = thread_pool->Submit(compute, 10);
    auto res = fut.get();
    EXPECT_EQ(res, 20);

    std::vector<std::future<int>> futs;
    for (int i = 0; i < 10; ++i) {
        futs.push_back(thread_pool->Submit(compute, i));
    }
    std::this_thread::sleep_for(std::chrono::milliseconds(300));
    std::cout << "current thread num" << thread_pool->GetThreadNum()
              << std::endl;

    for (int i = 0; i < 10; ++i) {
        std::cout << futs[i].get() << std::endl;
    }

    std::this_thread::sleep_for(std::chrono::milliseconds(5000));
    std::cout << "current thread num" << thread_pool->GetThreadNum()
              << std::endl;
    auto thread_num_2 = thread_pool->GetThreadNum();
    EXPECT_EQ(thread_num_2, thread_num_1);
}

TEST_F(DiskAnnFileManagerTest, TestThreadPool) {
    auto thread_pool = std::make_shared<milvus::ThreadPool>(50, "test");
    std::vector<std::future<int>> futures;
    auto start = chrono::system_clock::now();
    for (int i = 0; i < 100; i++) {
        futures.push_back(
            thread_pool->Submit(test_worker, "test_id" + std::to_string(i)));
    }
    for (auto& future : futures) {
        EXPECT_EQ(future.get(), 1);
    }
    auto end = chrono::system_clock::now();
    auto duration = chrono::duration_cast<chrono::microseconds>(end - start);
    auto second = double(duration.count()) * chrono::microseconds::period::num /
                  chrono::microseconds::period::den;
    std::cout << "cost time:" << second << std::endl;
    EXPECT_LT(second, 4 * 100);
}

int
test_exception(string s) {
    if (s == "test_id60") {
        throw SegcoreError(ErrorCode::UnexpectedError, "run time error");
    }
    return 1;
}

TEST_F(DiskAnnFileManagerTest, TestThreadPoolException) {
    try {
        auto thread_pool = std::make_shared<milvus::ThreadPool>(50, "test");
        std::vector<std::future<int>> futures;
        for (int i = 0; i < 100; i++) {
            futures.push_back(thread_pool->Submit(
                test_exception, "test_id" + std::to_string(i)));
        }
        for (auto& future : futures) {
            future.get();
        }
    } catch (std::exception& e) {
        EXPECT_EQ(std::string(e.what()), "run time error");
    }
}

const int64_t kOptFieldId = 123456;
const std::string kOptFieldName = "opt_field_name";
const DataType kOptFiledType = DataType::INT64;
const int64_t kOptFieldDataRange = 10000;
const std::string kOptFieldPath = "/tmp/diskann/opt_field/123123";
// const std::string kOptFieldPath = "/tmp/diskann/index_files/1000/index";
const size_t kEntityCnt = 1000 * 1000;
const DataType kOptFieldDataType = DataType::INT64;
const FieldDataMeta kOptVecFieldDataMeta = {1, 2, 3, 100};
using OffsetT = uint32_t;

auto
CreateFileManager(const ChunkManagerPtr& cm)
    -> std::shared_ptr<DiskFileManagerImpl> {
    // collection_id: 1, partition_id: 2, segment_id: 3
    // field_id: 100, index_build_id: 1000, index_version: 1
    IndexMeta index_meta = {
        3, 100, 1000, 1, "opt_fields", "field_name", DataType::VECTOR_FLOAT, 1};
    int64_t slice_size = milvus::FILE_SLICE_SIZE;
    return std::make_shared<DiskFileManagerImpl>(
        storage::FileManagerContext(kOptVecFieldDataMeta, index_meta, cm));
}

auto
PrepareRawFieldData() -> std::vector<int64_t> {
    std::vector<int64_t> data(kEntityCnt);
    int64_t field_val = 0;
    for (size_t i = 0; i < kEntityCnt; ++i) {
        data[i] = field_val++;
        if (field_val >= kOptFieldDataRange) {
            field_val = 0;
        }
    }
    return data;
}

auto
PrepareInsertData() -> std::string {
    size_t sz = sizeof(int64_t) * kEntityCnt;
    std::vector<int64_t> data = PrepareRawFieldData();
    auto field_data =
        storage::CreateFieldData(kOptFieldDataType, 1, kEntityCnt);
    field_data->FillFieldData(data.data(), kEntityCnt);
    storage::InsertData insert_data(field_data);
    insert_data.SetFieldDataMeta(kOptVecFieldDataMeta);
    insert_data.SetTimestamps(0, 100);
    auto serialized_data = insert_data.Serialize(storage::StorageType::Remote);

    auto chunk_manager =
        storage::CreateChunkManager(get_default_local_storage_config());

    std::string path = kOptFieldPath + "0";
    boost::filesystem::remove_all(path);
    chunk_manager->Write(path, serialized_data.data(), serialized_data.size());
    return path;
}

auto
PrepareInsertDataSpace()
    -> std::pair<std::string, std::shared_ptr<milvus_storage::Space>> {
    std::string path = kOptFieldPath + "1";
    arrow::FieldVector arrow_fields{
        arrow::field("pk", arrow::int64()),
        arrow::field("ts", arrow::int64()),
        arrow::field(kOptFieldName, arrow::int64()),
        arrow::field("vec", arrow::fixed_size_binary(1))};
    auto arrow_schema = std::make_shared<arrow::Schema>(arrow_fields);
    auto schema_options = std::make_shared<milvus_storage::SchemaOptions>();
    schema_options->primary_column = "pk";
    schema_options->version_column = "ts";
    schema_options->vector_column = "vec";
    auto schema =
        std::make_shared<milvus_storage::Schema>(arrow_schema, schema_options);
    boost::filesystem::remove_all(path);
    boost::filesystem::create_directories(path);
    EXPECT_TRUE(schema->Validate().ok());
    auto opt_space = milvus_storage::Space::Open(
        "file://" + boost::filesystem::canonical(path).string(),
        milvus_storage::Options{schema});
    EXPECT_TRUE(opt_space.has_value());
    auto space = std::move(opt_space.value());
    const auto data = PrepareRawFieldData();
    arrow::Int64Builder pk_builder;
    arrow::Int64Builder ts_builder;
    arrow::NumericBuilder<arrow::Int64Type> scalar_builder;
    arrow::FixedSizeBinaryBuilder vec_builder(arrow::fixed_size_binary(1));
    const uint8_t kByteZero = 0;
    for (size_t i = 0; i < kEntityCnt; ++i) {
        EXPECT_TRUE(pk_builder.Append(i).ok());
        EXPECT_TRUE(ts_builder.Append(i).ok());
        EXPECT_TRUE(vec_builder.Append(&kByteZero).ok());
    }
    for (size_t i = 0; i < kEntityCnt; ++i) {
        EXPECT_TRUE(scalar_builder.Append(data[i]).ok());
    }
    std::shared_ptr<arrow::Array> pk_array;
    EXPECT_TRUE(pk_builder.Finish(&pk_array).ok());
    std::shared_ptr<arrow::Array> ts_array;
    EXPECT_TRUE(ts_builder.Finish(&ts_array).ok());
    std::shared_ptr<arrow::Array> scalar_array;
    EXPECT_TRUE(scalar_builder.Finish(&scalar_array).ok());
    std::shared_ptr<arrow::Array> vec_array;
    EXPECT_TRUE(vec_builder.Finish(&vec_array).ok());
    auto batch =
        arrow::RecordBatch::Make(arrow_schema,
                                 kEntityCnt,
                                 {pk_array, ts_array, scalar_array, vec_array});
    auto write_opt = milvus_storage::WriteOption{kEntityCnt};
    space->Write(arrow::RecordBatchReader::Make({batch}, arrow_schema)
                     .ValueOrDie()
                     .get(),
                 &write_opt);
    return {path, std::move(space)};
}

auto
PrepareOptionalField(const std::shared_ptr<DiskFileManagerImpl>& file_manager,
                     const std::string& insert_file_path) -> OptFieldT {
    OptFieldT opt_field;
    std::vector<std::string> insert_files;
    insert_files.emplace_back(insert_file_path);
    opt_field[kOptFieldId] = {kOptFieldName, kOptFiledType, insert_files};
    return opt_field;
}

void
CheckOptFieldCorrectness(const std::string& local_file_path) {
    std::ifstream ifs(local_file_path);
    if (!ifs.is_open()) {
        FAIL() << "open file failed: " << local_file_path << std::endl;
        return;
    }
    uint8_t meta_version;
    uint32_t meta_num_of_fields, num_of_unique_field_data;
    int64_t field_id;
    ifs.read(reinterpret_cast<char*>(&meta_version), sizeof(meta_version));
    EXPECT_EQ(meta_version, 0);
    ifs.read(reinterpret_cast<char*>(&meta_num_of_fields),
             sizeof(meta_num_of_fields));
    EXPECT_EQ(meta_num_of_fields, 1);
    ifs.read(reinterpret_cast<char*>(&field_id), sizeof(field_id));
    EXPECT_EQ(field_id, kOptFieldId);
    ifs.read(reinterpret_cast<char*>(&num_of_unique_field_data),
             sizeof(num_of_unique_field_data));
    EXPECT_EQ(num_of_unique_field_data, kOptFieldDataRange);

    uint32_t expected_single_category_offset_cnt =
        kEntityCnt / kOptFieldDataRange;
    uint32_t read_single_category_offset_cnt;
    std::vector<OffsetT> single_category_offsets(
        expected_single_category_offset_cnt);
    for (uint32_t i = 0; i < num_of_unique_field_data; ++i) {
        ifs.read(reinterpret_cast<char*>(&read_single_category_offset_cnt),
                 sizeof(read_single_category_offset_cnt));
        ASSERT_EQ(read_single_category_offset_cnt,
                  expected_single_category_offset_cnt);
        ifs.read(reinterpret_cast<char*>(single_category_offsets.data()),
                 read_single_category_offset_cnt * sizeof(OffsetT));

        OffsetT first_offset = 0;
        if (read_single_category_offset_cnt > 0) {
            first_offset = single_category_offsets[0];
        }
        for (size_t j = 1; j < read_single_category_offset_cnt; ++j) {
            ASSERT_EQ(single_category_offsets[j] % kOptFieldDataRange,
                      first_offset % kOptFieldDataRange);
        }
    }
}

TEST_F(DiskAnnFileManagerTest, CacheOptFieldToDiskFieldEmpty) {
    auto file_manager = CreateFileManager(cm_);
    const auto& [insert_file_space_path, space] = PrepareInsertDataSpace();
    OptFieldT opt_fields;
    EXPECT_TRUE(file_manager->CacheOptFieldToDisk(opt_fields).empty());
    EXPECT_TRUE(file_manager->CacheOptFieldToDisk(space, opt_fields).empty());
}

TEST_F(DiskAnnFileManagerTest, CacheOptFieldToDiskSpaceEmpty) {
    auto file_manager = CreateFileManager(cm_);
    auto opt_fileds = PrepareOptionalField(file_manager, "");
    auto res = file_manager->CacheOptFieldToDisk(nullptr, opt_fileds);
    EXPECT_TRUE(res.empty());
}

TEST_F(DiskAnnFileManagerTest, CacheOptFieldToDiskOptFieldMoreThanOne) {
    auto file_manager = CreateFileManager(cm_);
    const auto insert_file_path = PrepareInsertData();
    const auto& [insert_file_space_path, space] = PrepareInsertDataSpace();

    OptFieldT opt_fields = PrepareOptionalField(file_manager, insert_file_path);
    opt_fields[kOptFieldId + 1] = {
        kOptFieldName + "second", kOptFiledType, {insert_file_space_path}};
    EXPECT_THROW(file_manager->CacheOptFieldToDisk(opt_fields), SegcoreError);
    EXPECT_THROW(file_manager->CacheOptFieldToDisk(space, opt_fields),
                 SegcoreError);
}

TEST_F(DiskAnnFileManagerTest, CacheOptFieldToDiskCorrect) {
    auto file_manager = CreateFileManager(cm_);
    const auto insert_file_path = PrepareInsertData();
    auto opt_fileds = PrepareOptionalField(file_manager, insert_file_path);
    auto res = file_manager->CacheOptFieldToDisk(opt_fileds);
    ASSERT_FALSE(res.empty());
    CheckOptFieldCorrectness(res);
}

TEST_F(DiskAnnFileManagerTest, CacheOptFieldToDiskSpaceCorrect) {
    auto file_manager = CreateFileManager(cm_);
    const auto& [insert_file_path, space] = PrepareInsertDataSpace();
    auto opt_fileds = PrepareOptionalField(file_manager, insert_file_path);
    auto res = file_manager->CacheOptFieldToDisk(space, opt_fileds);
    ASSERT_FALSE(res.empty());
    CheckOptFieldCorrectness(res);
}