/*******************************************************************************
 * Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 ******************************************************************************/
#include "ClientTest.h"
#include "MilvusApi.h"

#include <iostream>
#include <time.h>
#include <unistd.h>

using namespace ::milvus;

namespace {
    std::string GetTableName();

    static const std::string TABLE_NAME = GetTableName();
    static constexpr int64_t TABLE_DIMENSION = 512;
    static constexpr int64_t BATCH_ROW_COUNT = 100000;
    static constexpr int64_t NQ = 10;
    static constexpr int64_t TOP_K = 10;
    static constexpr int64_t SEARCH_TARGET = 5000; //change this value, result is different
    static constexpr int64_t ADD_VECTOR_LOOP = 5;

#define BLOCK_SPLITER std::cout << "===========================================" << std::endl;

    void PrintTableSchema(const TableSchema& tb_schema) {
        BLOCK_SPLITER
        std::cout << "Table name: " << tb_schema.table_name << std::endl;
        std::cout << "Table index type: " << (int)tb_schema.index_type << std::endl;
        std::cout << "Table dimension: " << tb_schema.dimension << std::endl;
        std::cout << "Table store raw data: " << (tb_schema.store_raw_vector ? "true" : "false") << std::endl;
        BLOCK_SPLITER
    }

    void PrintRecordIdArray(const std::vector<int64_t>& record_ids) {
        BLOCK_SPLITER
        std::cout << "Returned id array count: " << record_ids.size() << std::endl;
#if 0
        for(auto id : record_ids) {
            std::cout << std::to_string(id) << std::endl;
        }
#endif
        BLOCK_SPLITER
    }

    void PrintSearchResult(const std::vector<TopKQueryResult>& topk_query_result_array) {
        BLOCK_SPLITER
        std::cout << "Returned result count: " << topk_query_result_array.size() << std::endl;

        int32_t index = 0;
        for(auto& result : topk_query_result_array) {
            index++;
            std::cout << "No." << std::to_string(index) << " vector top "
                << std::to_string(result.query_result_arrays.size())
                << " search result:" << std::endl;
            for(auto& item : result.query_result_arrays) {
                std::cout << "\t" << std::to_string(item.id) << "\tscore:" << std::to_string(item.score);
                std::cout << std::endl;
            }
        }

        BLOCK_SPLITER
    }

    std::string CurrentTime() {
        time_t tt;
        time( &tt );
        tt = tt + 8*3600;
        tm* t= gmtime( &tt );

        std::string str = std::to_string(t->tm_year + 1900) + "_" + std::to_string(t->tm_mon + 1)
                          + "_" + std::to_string(t->tm_mday) + "_" + std::to_string(t->tm_hour)
                          + "_" + std::to_string(t->tm_min) + "_" + std::to_string(t->tm_sec);

        return str;
    }

    std::string CurrentTmDate() {
        time_t tt;
        time( &tt );
        tt = tt + 8*3600;
        tm* t= gmtime( &tt );

        std::string str = std::to_string(t->tm_year + 1900) + "-" + std::to_string(t->tm_mon + 1)
                          + "-" + std::to_string(t->tm_mday);

        return str;
    }

    std::string GetTableName() {
        static std::string s_id(CurrentTime());
        return s_id;
    }

    TableSchema BuildTableSchema() {
        TableSchema tb_schema;
        tb_schema.table_name = TABLE_NAME;
        tb_schema.index_type = IndexType::cpu_idmap;
        tb_schema.dimension = TABLE_DIMENSION;
        tb_schema.store_raw_vector = true;

        return tb_schema;
    }

    void BuildVectors(int64_t from, int64_t to,
                      std::vector<RowRecord>& vector_record_array) {
        if(to <= from){
            return;
        }

        vector_record_array.clear();
        for (int64_t k = from; k < to; k++) {
            RowRecord record;
            record.data.resize(TABLE_DIMENSION);
            for(int64_t i = 0; i < TABLE_DIMENSION; i++) {
                record.data[i] = (float)(k%(i+1));
            }

            vector_record_array.emplace_back(record);
        }
    }

    void Sleep(int seconds) {
        std::cout << "Waiting " << seconds << " seconds ..." << std::endl;
        sleep(seconds);
    }
}

void
ClientTest::Test(const std::string& address, const std::string& port) {
    std::shared_ptr<Connection> conn = Connection::Create();

    {//connect server
        ConnectParam param = {address, port};
        Status stat = conn->Connect(param);
        std::cout << "Connect function call status: " << stat.ToString() << std::endl;
    }

    {//server version
        std::string version = conn->ServerVersion();
        std::cout << "Server version: " << version << std::endl;
    }

    {//sdk version
        std::string version = conn->ClientVersion();
        std::cout << "SDK version: " << version << std::endl;
    }

    {
        std::vector<std::string> tables;
        Status stat = conn->ShowTables(tables);
        std::cout << "ShowTables function call status: " << stat.ToString() << std::endl;
        std::cout << "All tables: " << std::endl;
        for(auto& table : tables) {
            int64_t row_count = 0;
            stat = conn->GetTableRowCount(table, row_count);
            std::cout << "\t" << table << "(" << row_count << " rows)" << std::endl;
        }
    }

    {//create table
        TableSchema tb_schema = BuildTableSchema();
        Status stat = conn->CreateTable(tb_schema);
        std::cout << "CreateTable function call status: " << stat.ToString() << std::endl;
        PrintTableSchema(tb_schema);

        bool has_table = conn->HasTable(tb_schema.table_name);
        if(has_table) {
            std::cout << "Table is created" << std::endl;
        }
    }

    {//describe table
        TableSchema tb_schema;
        Status stat = conn->DescribeTable(TABLE_NAME, tb_schema);
        std::cout << "DescribeTable function call status: " << stat.ToString() << std::endl;
        PrintTableSchema(tb_schema);
    }

    for(int i = 0; i < ADD_VECTOR_LOOP; i++){//add vectors
        std::vector<RowRecord> record_array;
        BuildVectors(i*BATCH_ROW_COUNT, (i+1)*BATCH_ROW_COUNT, record_array);
        std::vector<int64_t> record_ids;
        Status stat = conn->AddVector(TABLE_NAME, record_array, record_ids);
        std::cout << "AddVector function call status: " << stat.ToString() << std::endl;
        PrintRecordIdArray(record_ids);
    }

    {//search vectors
        Sleep(2);

        std::vector<RowRecord> record_array;
        BuildVectors(SEARCH_TARGET, SEARCH_TARGET + NQ, record_array);

        std::vector<Range> query_range_array;
        Range rg;
        rg.start_value = CurrentTmDate();
        rg.end_value = CurrentTmDate();
        query_range_array.emplace_back(rg);
        std::vector<TopKQueryResult> topk_query_result_array;
        Status stat = conn->SearchVector(TABLE_NAME, record_array, query_range_array, TOP_K, topk_query_result_array);
        std::cout << "SearchVector function call status: " << stat.ToString() << std::endl;
        PrintSearchResult(topk_query_result_array);
    }

    {//delete table
        Status stat = conn->DeleteTable(TABLE_NAME);
        std::cout << "DeleteTable function call status: " << stat.ToString() << std::endl;
    }

    {//server status
        std::string status = conn->ServerStatus();
        std::cout << "Server status before disconnect: " << status << std::endl;
    }
    Connection::Destroy(conn);
    {//server status
        std::string status = conn->ServerStatus();
        std::cout << "Server status after disconnect: " << status << std::endl;
    }
}