// This autogenerated skeleton file illustrates how to build a server.
// You should copy it to another filename to avoid overwriting it.

#include "MegasearchService.h"
#include <thrift/protocol/TBinaryProtocol.h>
#include <thrift/server/TSimpleServer.h>
#include <thrift/transport/TServerSocket.h>
#include <thrift/transport/TBufferTransports.h>

using namespace ::apache::thrift;
using namespace ::apache::thrift::protocol;
using namespace ::apache::thrift::transport;
using namespace ::apache::thrift::server;

using namespace  ::megasearch::thrift;

class MegasearchServiceHandler : virtual public MegasearchServiceIf {
 public:
  MegasearchServiceHandler() {
    // Your initialization goes here
  }

  /**
   * @brief Create table method
   * 
   * This method is used to create table
   * 
   * @param param, use to provide table information to be created.
   * 
   * 
   * @param param
   */
  void CreateTable(const TableSchema& param) {
    // Your implementation goes here
    printf("CreateTable\n");
  }

  /**
   * @brief Delete table method
   * 
   * This method is used to delete table.
   * 
   * @param table_name, table name is going to be deleted.
   * 
   * 
   * @param table_name
   */
  void DeleteTable(const std::string& table_name) {
    // Your implementation goes here
    printf("DeleteTable\n");
  }

  /**
   * @brief Create table partition
   * 
   * This method is used to create table partition.
   * 
   * @param param, use to provide partition information to be created.
   * 
   * 
   * @param param
   */
  void CreateTablePartition(const CreateTablePartitionParam& param) {
    // Your implementation goes here
    printf("CreateTablePartition\n");
  }

  /**
   * @brief Delete table partition
   * 
   * This method is used to delete table partition.
   * 
   * @param param, use to provide partition information to be deleted.
   * 
   * 
   * @param param
   */
  void DeleteTablePartition(const DeleteTablePartitionParam& param) {
    // Your implementation goes here
    printf("DeleteTablePartition\n");
  }

  /**
   * @brief Add vector array to table
   * 
   * This method is used to add vector array to table.
   * 
   * @param table_name, table_name is inserted.
   * @param record_array, vector array is inserted.
   * 
   * @return vector id array
   * 
   * @param table_name
   * @param record_array
   */
  void AddVector(std::vector<int64_t> & _return, const std::string& table_name, const std::vector<RowRecord> & record_array) {
    // Your implementation goes here
    printf("AddVector\n");
  }

  /**
   * @brief Query vector
   * 
   * This method is used to query vector in table.
   * 
   * @param table_name, table_name is queried.
   * @param query_record_array, all vector are going to be queried.
   * @param topk, how many similarity vectors will be searched.
   * 
   * @return query result array.
   * 
   * @param table_name
   * @param query_record_array
   * @param topk
   */
  void SearchVector(std::vector<TopKQueryResult> & _return, const std::string& table_name, const std::vector<QueryRecord> & query_record_array, const int64_t topk) {
    // Your implementation goes here
    printf("SearchVector\n");
  }

  /**
   * @brief Show table information
   * 
   * This method is used to show table information.
   * 
   * @param table_name, which table is show.
   * 
   * @return table schema
   * 
   * @param table_name
   */
  void DescribeTable(TableSchema& _return, const std::string& table_name) {
    // Your implementation goes here
    printf("DescribeTable\n");
  }

  /**
   * @brief List all tables in database
   * 
   * This method is used to list all tables.
   * 
   * 
   * @return table names.
   */
  void ShowTables(std::vector<std::string> & _return) {
    // Your implementation goes here
    printf("ShowTables\n");
  }

  /**
   * @brief Give the server status
   * 
   * This method is used to give the server status.
   * 
   * @return Server status.
   * 
   * @param cmd
   */
  void Ping(std::string& _return, const std::string& cmd) {
    // Your implementation goes here
    printf("Ping\n");
  }

};

int main(int argc, char **argv) {
  int port = 9090;
  ::apache::thrift::stdcxx::shared_ptr<MegasearchServiceHandler> handler(new MegasearchServiceHandler());
  ::apache::thrift::stdcxx::shared_ptr<TProcessor> processor(new MegasearchServiceProcessor(handler));
  ::apache::thrift::stdcxx::shared_ptr<TServerTransport> serverTransport(new TServerSocket(port));
  ::apache::thrift::stdcxx::shared_ptr<TTransportFactory> transportFactory(new TBufferedTransportFactory());
  ::apache::thrift::stdcxx::shared_ptr<TProtocolFactory> protocolFactory(new TBinaryProtocolFactory());

  TSimpleServer server(processor, serverTransport, transportFactory, protocolFactory);
  server.serve();
  return 0;
}

