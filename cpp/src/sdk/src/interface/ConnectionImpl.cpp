/*******************************************************************************
 * Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 ******************************************************************************/
#include "ConnectionImpl.h"

namespace megasearch {

std::shared_ptr<Connection>
Connection::Create() {
    return std::shared_ptr<Connection>(new ConnectionImpl());
}

Status
Connection::Destroy(std::shared_ptr<megasearch::Connection> connection_ptr) {
    if(connection_ptr != nullptr) {
        return connection_ptr->Disconnect();
    }
    return Status::OK();
}

//////////////////////////////////////////////////////////////////////////////////////////////
ConnectionImpl::ConnectionImpl() {
    client_proxy = std::make_shared<ClientProxy>();
}

Status
ConnectionImpl::Connect(const ConnectParam &param) {
    return client_proxy->Connect(param);
}

Status
ConnectionImpl::Connect(const std::string &uri) {
    return client_proxy->Connect(uri);
}

Status
ConnectionImpl::Connected() const {
    return client_proxy->Connected();
}

Status
ConnectionImpl::Disconnect() {
    return client_proxy->Disconnect();
}

std::string
ConnectionImpl::ClientVersion() const {
    return client_proxy->ClientVersion();
}

Status
ConnectionImpl::CreateTable(const TableSchema &param) {
    return client_proxy->CreateTable(param);
}

Status
ConnectionImpl::CreateTablePartition(const CreateTablePartitionParam &param) {
    return client_proxy->CreateTablePartition(param);
}

Status
ConnectionImpl::DeleteTablePartition(const DeleteTablePartitionParam &param) {
    return client_proxy->DeleteTablePartition(param);
}

Status
ConnectionImpl::DeleteTable(const std::string &table_name) {
    return client_proxy->DeleteTable(table_name);
}

Status
ConnectionImpl::AddVector(const std::string &table_name,
                          const std::vector<RowRecord> &record_array,
                          std::vector<int64_t> &id_array) {
    return client_proxy->AddVector(table_name, record_array, id_array);
}

Status
ConnectionImpl::SearchVector(const std::string &table_name,
                             const std::vector<QueryRecord> &query_record_array,
                             std::vector<TopKQueryResult> &topk_query_result_array,
                             int64_t topk) {
    return client_proxy->SearchVector(table_name, query_record_array, topk_query_result_array, topk);
}

Status
ConnectionImpl::DescribeTable(const std::string &table_name, TableSchema &table_schema) {
    return client_proxy->DescribeTable(table_name, table_schema);
}

Status
ConnectionImpl::ShowTables(std::vector<std::string> &table_array) {
    return client_proxy->ShowTables(table_array);
}

std::string
ConnectionImpl::ServerVersion() const {
    return client_proxy->ServerVersion();
}

std::string
ConnectionImpl::ServerStatus() const {
    return client_proxy->ServerStatus();
}

}

