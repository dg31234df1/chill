// This autogenerated skeleton file illustrates how to build a server.
// You should copy it to another filename to avoid overwriting it.

#include "VecService.h"
#include <thrift/protocol/TBinaryProtocol.h>
#include <thrift/server/TSimpleServer.h>
#include <thrift/transport/TServerSocket.h>
#include <thrift/transport/TBufferTransports.h>

using namespace ::apache::thrift;
using namespace ::apache::thrift::protocol;
using namespace ::apache::thrift::transport;
using namespace ::apache::thrift::server;

using namespace  ::zilliz;

class VecServiceHandler : virtual public VecServiceIf {
 public:
  VecServiceHandler() {
    // Your initialization goes here
  }

  /**
   * group interfaces
   * 
   * @param group
   */
  void add_group(const VecGroup& group) {
    // Your implementation goes here
    printf("add_group\n");
  }

  void get_group(VecGroup& _return, const std::string& group_id) {
    // Your implementation goes here
    printf("get_group\n");
  }

  void del_group(const std::string& group_id) {
    // Your implementation goes here
    printf("del_group\n");
  }

  /**
   * insert vector interfaces
   * 
   * 
   * @param group_id
   * @param tensor
   */
  void add_vector(std::string& _return, const std::string& group_id, const VecTensor& tensor) {
    // Your implementation goes here
    printf("add_vector\n");
  }

  void add_vector_batch(std::vector<std::string> & _return, const std::string& group_id, const VecTensorList& tensor_list) {
    // Your implementation goes here
    printf("add_vector_batch\n");
  }

  void add_binary_vector(std::string& _return, const std::string& group_id, const VecBinaryTensor& tensor) {
    // Your implementation goes here
    printf("add_binary_vector\n");
  }

  void add_binary_vector_batch(std::vector<std::string> & _return, const std::string& group_id, const VecBinaryTensorList& tensor_list) {
    // Your implementation goes here
    printf("add_binary_vector_batch\n");
  }

  /**
   * search interfaces
   * you can use filter to reduce search result
   * filter.attrib_filter can specify which attribute you need, for example:
   * set attrib_filter = {"color":""} means you want to get "color" attribute for result vector
   * set attrib_filter = {"color":"red"} means you want to get vectors which has attribute "color" equals "red"
   * if filter.time_range is empty, engine will search without time limit
   * 
   * @param group_id
   * @param top_k
   * @param tensor
   * @param filter
   */
  void search_vector(VecSearchResult& _return, const std::string& group_id, const int64_t top_k, const VecTensor& tensor, const VecSearchFilter& filter) {
    // Your implementation goes here
    printf("search_vector\n");
  }

  void search_vector_batch(VecSearchResultList& _return, const std::string& group_id, const int64_t top_k, const VecTensorList& tensor_list, const VecSearchFilter& filter) {
    // Your implementation goes here
    printf("search_vector_batch\n");
  }

  void search_binary_vector(VecSearchResult& _return, const std::string& group_id, const int64_t top_k, const VecBinaryTensor& tensor, const VecSearchFilter& filter) {
    // Your implementation goes here
    printf("search_binary_vector\n");
  }

  void search_binary_vector_batch(VecSearchResultList& _return, const std::string& group_id, const int64_t top_k, const VecBinaryTensorList& tensor_list, const VecSearchFilter& filter) {
    // Your implementation goes here
    printf("search_binary_vector_batch\n");
  }

};

int main(int argc, char **argv) {
  int port = 9090;
  ::apache::thrift::stdcxx::shared_ptr<VecServiceHandler> handler(new VecServiceHandler());
  ::apache::thrift::stdcxx::shared_ptr<TProcessor> processor(new VecServiceProcessor(handler));
  ::apache::thrift::stdcxx::shared_ptr<TServerTransport> serverTransport(new TServerSocket(port));
  ::apache::thrift::stdcxx::shared_ptr<TTransportFactory> transportFactory(new TBufferedTransportFactory());
  ::apache::thrift::stdcxx::shared_ptr<TProtocolFactory> protocolFactory(new TBinaryProtocolFactory());

  TSimpleServer server(processor, serverTransport, transportFactory, protocolFactory);
  server.serve();
  return 0;
}

