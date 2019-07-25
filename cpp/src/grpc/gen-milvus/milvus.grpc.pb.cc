// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: milvus.proto

#include "milvus.pb.h"
#include "milvus.grpc.pb.h"

#include <functional>
#include <grpcpp/impl/codegen/async_stream.h>
#include <grpcpp/impl/codegen/async_unary_call.h>
#include <grpcpp/impl/codegen/channel_interface.h>
#include <grpcpp/impl/codegen/client_unary_call.h>
#include <grpcpp/impl/codegen/client_callback.h>
#include <grpcpp/impl/codegen/method_handler_impl.h>
#include <grpcpp/impl/codegen/rpc_service_method.h>
#include <grpcpp/impl/codegen/server_callback.h>
#include <grpcpp/impl/codegen/service_type.h>
#include <grpcpp/impl/codegen/sync_stream.h>
namespace milvus {
namespace grpc {

static const char* MilvusService_method_names[] = {
  "/milvus.grpc.MilvusService/CreateTable",
  "/milvus.grpc.MilvusService/HasTable",
  "/milvus.grpc.MilvusService/DropTable",
  "/milvus.grpc.MilvusService/BuildIndex",
  "/milvus.grpc.MilvusService/InsertVector",
  "/milvus.grpc.MilvusService/SearchVector",
  "/milvus.grpc.MilvusService/SearchVectorInFiles",
  "/milvus.grpc.MilvusService/DescribeTable",
  "/milvus.grpc.MilvusService/GetTableRowCount",
  "/milvus.grpc.MilvusService/ShowTables",
  "/milvus.grpc.MilvusService/Ping",
};

std::unique_ptr< MilvusService::Stub> MilvusService::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< MilvusService::Stub> stub(new MilvusService::Stub(channel));
  return stub;
}

MilvusService::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel)
  : channel_(channel), rpcmethod_CreateTable_(MilvusService_method_names[0], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_HasTable_(MilvusService_method_names[1], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_DropTable_(MilvusService_method_names[2], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_BuildIndex_(MilvusService_method_names[3], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_InsertVector_(MilvusService_method_names[4], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_SearchVector_(MilvusService_method_names[5], ::grpc::internal::RpcMethod::SERVER_STREAMING, channel)
  , rpcmethod_SearchVectorInFiles_(MilvusService_method_names[6], ::grpc::internal::RpcMethod::SERVER_STREAMING, channel)
  , rpcmethod_DescribeTable_(MilvusService_method_names[7], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_GetTableRowCount_(MilvusService_method_names[8], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_ShowTables_(MilvusService_method_names[9], ::grpc::internal::RpcMethod::SERVER_STREAMING, channel)
  , rpcmethod_Ping_(MilvusService_method_names[10], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status MilvusService::Stub::CreateTable(::grpc::ClientContext* context, const ::milvus::grpc::TableSchema& request, ::milvus::grpc::Status* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_CreateTable_, context, request, response);
}

void MilvusService::Stub::experimental_async::CreateTable(::grpc::ClientContext* context, const ::milvus::grpc::TableSchema* request, ::milvus::grpc::Status* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_CreateTable_, context, request, response, std::move(f));
}

void MilvusService::Stub::experimental_async::CreateTable(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::milvus::grpc::Status* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_CreateTable_, context, request, response, std::move(f));
}

void MilvusService::Stub::experimental_async::CreateTable(::grpc::ClientContext* context, const ::milvus::grpc::TableSchema* request, ::milvus::grpc::Status* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_CreateTable_, context, request, response, reactor);
}

void MilvusService::Stub::experimental_async::CreateTable(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::milvus::grpc::Status* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_CreateTable_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::milvus::grpc::Status>* MilvusService::Stub::AsyncCreateTableRaw(::grpc::ClientContext* context, const ::milvus::grpc::TableSchema& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::milvus::grpc::Status>::Create(channel_.get(), cq, rpcmethod_CreateTable_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::milvus::grpc::Status>* MilvusService::Stub::PrepareAsyncCreateTableRaw(::grpc::ClientContext* context, const ::milvus::grpc::TableSchema& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::milvus::grpc::Status>::Create(channel_.get(), cq, rpcmethod_CreateTable_, context, request, false);
}

::grpc::Status MilvusService::Stub::HasTable(::grpc::ClientContext* context, const ::milvus::grpc::TableName& request, ::milvus::grpc::BoolReply* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_HasTable_, context, request, response);
}

void MilvusService::Stub::experimental_async::HasTable(::grpc::ClientContext* context, const ::milvus::grpc::TableName* request, ::milvus::grpc::BoolReply* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_HasTable_, context, request, response, std::move(f));
}

void MilvusService::Stub::experimental_async::HasTable(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::milvus::grpc::BoolReply* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_HasTable_, context, request, response, std::move(f));
}

void MilvusService::Stub::experimental_async::HasTable(::grpc::ClientContext* context, const ::milvus::grpc::TableName* request, ::milvus::grpc::BoolReply* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_HasTable_, context, request, response, reactor);
}

void MilvusService::Stub::experimental_async::HasTable(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::milvus::grpc::BoolReply* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_HasTable_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::milvus::grpc::BoolReply>* MilvusService::Stub::AsyncHasTableRaw(::grpc::ClientContext* context, const ::milvus::grpc::TableName& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::milvus::grpc::BoolReply>::Create(channel_.get(), cq, rpcmethod_HasTable_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::milvus::grpc::BoolReply>* MilvusService::Stub::PrepareAsyncHasTableRaw(::grpc::ClientContext* context, const ::milvus::grpc::TableName& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::milvus::grpc::BoolReply>::Create(channel_.get(), cq, rpcmethod_HasTable_, context, request, false);
}

::grpc::Status MilvusService::Stub::DropTable(::grpc::ClientContext* context, const ::milvus::grpc::TableName& request, ::milvus::grpc::Status* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_DropTable_, context, request, response);
}

void MilvusService::Stub::experimental_async::DropTable(::grpc::ClientContext* context, const ::milvus::grpc::TableName* request, ::milvus::grpc::Status* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_DropTable_, context, request, response, std::move(f));
}

void MilvusService::Stub::experimental_async::DropTable(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::milvus::grpc::Status* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_DropTable_, context, request, response, std::move(f));
}

void MilvusService::Stub::experimental_async::DropTable(::grpc::ClientContext* context, const ::milvus::grpc::TableName* request, ::milvus::grpc::Status* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_DropTable_, context, request, response, reactor);
}

void MilvusService::Stub::experimental_async::DropTable(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::milvus::grpc::Status* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_DropTable_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::milvus::grpc::Status>* MilvusService::Stub::AsyncDropTableRaw(::grpc::ClientContext* context, const ::milvus::grpc::TableName& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::milvus::grpc::Status>::Create(channel_.get(), cq, rpcmethod_DropTable_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::milvus::grpc::Status>* MilvusService::Stub::PrepareAsyncDropTableRaw(::grpc::ClientContext* context, const ::milvus::grpc::TableName& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::milvus::grpc::Status>::Create(channel_.get(), cq, rpcmethod_DropTable_, context, request, false);
}

::grpc::Status MilvusService::Stub::BuildIndex(::grpc::ClientContext* context, const ::milvus::grpc::TableName& request, ::milvus::grpc::Status* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_BuildIndex_, context, request, response);
}

void MilvusService::Stub::experimental_async::BuildIndex(::grpc::ClientContext* context, const ::milvus::grpc::TableName* request, ::milvus::grpc::Status* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_BuildIndex_, context, request, response, std::move(f));
}

void MilvusService::Stub::experimental_async::BuildIndex(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::milvus::grpc::Status* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_BuildIndex_, context, request, response, std::move(f));
}

void MilvusService::Stub::experimental_async::BuildIndex(::grpc::ClientContext* context, const ::milvus::grpc::TableName* request, ::milvus::grpc::Status* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_BuildIndex_, context, request, response, reactor);
}

void MilvusService::Stub::experimental_async::BuildIndex(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::milvus::grpc::Status* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_BuildIndex_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::milvus::grpc::Status>* MilvusService::Stub::AsyncBuildIndexRaw(::grpc::ClientContext* context, const ::milvus::grpc::TableName& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::milvus::grpc::Status>::Create(channel_.get(), cq, rpcmethod_BuildIndex_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::milvus::grpc::Status>* MilvusService::Stub::PrepareAsyncBuildIndexRaw(::grpc::ClientContext* context, const ::milvus::grpc::TableName& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::milvus::grpc::Status>::Create(channel_.get(), cq, rpcmethod_BuildIndex_, context, request, false);
}

::grpc::Status MilvusService::Stub::InsertVector(::grpc::ClientContext* context, const ::milvus::grpc::InsertInfos& request, ::milvus::grpc::VectorIds* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_InsertVector_, context, request, response);
}

void MilvusService::Stub::experimental_async::InsertVector(::grpc::ClientContext* context, const ::milvus::grpc::InsertInfos* request, ::milvus::grpc::VectorIds* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_InsertVector_, context, request, response, std::move(f));
}

void MilvusService::Stub::experimental_async::InsertVector(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::milvus::grpc::VectorIds* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_InsertVector_, context, request, response, std::move(f));
}

void MilvusService::Stub::experimental_async::InsertVector(::grpc::ClientContext* context, const ::milvus::grpc::InsertInfos* request, ::milvus::grpc::VectorIds* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_InsertVector_, context, request, response, reactor);
}

void MilvusService::Stub::experimental_async::InsertVector(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::milvus::grpc::VectorIds* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_InsertVector_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::milvus::grpc::VectorIds>* MilvusService::Stub::AsyncInsertVectorRaw(::grpc::ClientContext* context, const ::milvus::grpc::InsertInfos& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::milvus::grpc::VectorIds>::Create(channel_.get(), cq, rpcmethod_InsertVector_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::milvus::grpc::VectorIds>* MilvusService::Stub::PrepareAsyncInsertVectorRaw(::grpc::ClientContext* context, const ::milvus::grpc::InsertInfos& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::milvus::grpc::VectorIds>::Create(channel_.get(), cq, rpcmethod_InsertVector_, context, request, false);
}

::grpc::ClientReader< ::milvus::grpc::TopKQueryResult>* MilvusService::Stub::SearchVectorRaw(::grpc::ClientContext* context, const ::milvus::grpc::SearchVectorInfos& request) {
  return ::grpc::internal::ClientReaderFactory< ::milvus::grpc::TopKQueryResult>::Create(channel_.get(), rpcmethod_SearchVector_, context, request);
}

void MilvusService::Stub::experimental_async::SearchVector(::grpc::ClientContext* context, ::milvus::grpc::SearchVectorInfos* request, ::grpc::experimental::ClientReadReactor< ::milvus::grpc::TopKQueryResult>* reactor) {
  ::grpc::internal::ClientCallbackReaderFactory< ::milvus::grpc::TopKQueryResult>::Create(stub_->channel_.get(), stub_->rpcmethod_SearchVector_, context, request, reactor);
}

::grpc::ClientAsyncReader< ::milvus::grpc::TopKQueryResult>* MilvusService::Stub::AsyncSearchVectorRaw(::grpc::ClientContext* context, const ::milvus::grpc::SearchVectorInfos& request, ::grpc::CompletionQueue* cq, void* tag) {
  return ::grpc::internal::ClientAsyncReaderFactory< ::milvus::grpc::TopKQueryResult>::Create(channel_.get(), cq, rpcmethod_SearchVector_, context, request, true, tag);
}

::grpc::ClientAsyncReader< ::milvus::grpc::TopKQueryResult>* MilvusService::Stub::PrepareAsyncSearchVectorRaw(::grpc::ClientContext* context, const ::milvus::grpc::SearchVectorInfos& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncReaderFactory< ::milvus::grpc::TopKQueryResult>::Create(channel_.get(), cq, rpcmethod_SearchVector_, context, request, false, nullptr);
}

::grpc::ClientReader< ::milvus::grpc::TopKQueryResult>* MilvusService::Stub::SearchVectorInFilesRaw(::grpc::ClientContext* context, const ::milvus::grpc::SearchVectorInFilesInfos& request) {
  return ::grpc::internal::ClientReaderFactory< ::milvus::grpc::TopKQueryResult>::Create(channel_.get(), rpcmethod_SearchVectorInFiles_, context, request);
}

void MilvusService::Stub::experimental_async::SearchVectorInFiles(::grpc::ClientContext* context, ::milvus::grpc::SearchVectorInFilesInfos* request, ::grpc::experimental::ClientReadReactor< ::milvus::grpc::TopKQueryResult>* reactor) {
  ::grpc::internal::ClientCallbackReaderFactory< ::milvus::grpc::TopKQueryResult>::Create(stub_->channel_.get(), stub_->rpcmethod_SearchVectorInFiles_, context, request, reactor);
}

::grpc::ClientAsyncReader< ::milvus::grpc::TopKQueryResult>* MilvusService::Stub::AsyncSearchVectorInFilesRaw(::grpc::ClientContext* context, const ::milvus::grpc::SearchVectorInFilesInfos& request, ::grpc::CompletionQueue* cq, void* tag) {
  return ::grpc::internal::ClientAsyncReaderFactory< ::milvus::grpc::TopKQueryResult>::Create(channel_.get(), cq, rpcmethod_SearchVectorInFiles_, context, request, true, tag);
}

::grpc::ClientAsyncReader< ::milvus::grpc::TopKQueryResult>* MilvusService::Stub::PrepareAsyncSearchVectorInFilesRaw(::grpc::ClientContext* context, const ::milvus::grpc::SearchVectorInFilesInfos& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncReaderFactory< ::milvus::grpc::TopKQueryResult>::Create(channel_.get(), cq, rpcmethod_SearchVectorInFiles_, context, request, false, nullptr);
}

::grpc::Status MilvusService::Stub::DescribeTable(::grpc::ClientContext* context, const ::milvus::grpc::TableName& request, ::milvus::grpc::TableSchema* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_DescribeTable_, context, request, response);
}

void MilvusService::Stub::experimental_async::DescribeTable(::grpc::ClientContext* context, const ::milvus::grpc::TableName* request, ::milvus::grpc::TableSchema* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_DescribeTable_, context, request, response, std::move(f));
}

void MilvusService::Stub::experimental_async::DescribeTable(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::milvus::grpc::TableSchema* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_DescribeTable_, context, request, response, std::move(f));
}

void MilvusService::Stub::experimental_async::DescribeTable(::grpc::ClientContext* context, const ::milvus::grpc::TableName* request, ::milvus::grpc::TableSchema* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_DescribeTable_, context, request, response, reactor);
}

void MilvusService::Stub::experimental_async::DescribeTable(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::milvus::grpc::TableSchema* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_DescribeTable_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::milvus::grpc::TableSchema>* MilvusService::Stub::AsyncDescribeTableRaw(::grpc::ClientContext* context, const ::milvus::grpc::TableName& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::milvus::grpc::TableSchema>::Create(channel_.get(), cq, rpcmethod_DescribeTable_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::milvus::grpc::TableSchema>* MilvusService::Stub::PrepareAsyncDescribeTableRaw(::grpc::ClientContext* context, const ::milvus::grpc::TableName& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::milvus::grpc::TableSchema>::Create(channel_.get(), cq, rpcmethod_DescribeTable_, context, request, false);
}

::grpc::Status MilvusService::Stub::GetTableRowCount(::grpc::ClientContext* context, const ::milvus::grpc::TableName& request, ::milvus::grpc::TableRowCount* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_GetTableRowCount_, context, request, response);
}

void MilvusService::Stub::experimental_async::GetTableRowCount(::grpc::ClientContext* context, const ::milvus::grpc::TableName* request, ::milvus::grpc::TableRowCount* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_GetTableRowCount_, context, request, response, std::move(f));
}

void MilvusService::Stub::experimental_async::GetTableRowCount(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::milvus::grpc::TableRowCount* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_GetTableRowCount_, context, request, response, std::move(f));
}

void MilvusService::Stub::experimental_async::GetTableRowCount(::grpc::ClientContext* context, const ::milvus::grpc::TableName* request, ::milvus::grpc::TableRowCount* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_GetTableRowCount_, context, request, response, reactor);
}

void MilvusService::Stub::experimental_async::GetTableRowCount(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::milvus::grpc::TableRowCount* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_GetTableRowCount_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::milvus::grpc::TableRowCount>* MilvusService::Stub::AsyncGetTableRowCountRaw(::grpc::ClientContext* context, const ::milvus::grpc::TableName& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::milvus::grpc::TableRowCount>::Create(channel_.get(), cq, rpcmethod_GetTableRowCount_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::milvus::grpc::TableRowCount>* MilvusService::Stub::PrepareAsyncGetTableRowCountRaw(::grpc::ClientContext* context, const ::milvus::grpc::TableName& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::milvus::grpc::TableRowCount>::Create(channel_.get(), cq, rpcmethod_GetTableRowCount_, context, request, false);
}

::grpc::ClientReader< ::milvus::grpc::TableName>* MilvusService::Stub::ShowTablesRaw(::grpc::ClientContext* context, const ::milvus::grpc::Command& request) {
  return ::grpc::internal::ClientReaderFactory< ::milvus::grpc::TableName>::Create(channel_.get(), rpcmethod_ShowTables_, context, request);
}

void MilvusService::Stub::experimental_async::ShowTables(::grpc::ClientContext* context, ::milvus::grpc::Command* request, ::grpc::experimental::ClientReadReactor< ::milvus::grpc::TableName>* reactor) {
  ::grpc::internal::ClientCallbackReaderFactory< ::milvus::grpc::TableName>::Create(stub_->channel_.get(), stub_->rpcmethod_ShowTables_, context, request, reactor);
}

::grpc::ClientAsyncReader< ::milvus::grpc::TableName>* MilvusService::Stub::AsyncShowTablesRaw(::grpc::ClientContext* context, const ::milvus::grpc::Command& request, ::grpc::CompletionQueue* cq, void* tag) {
  return ::grpc::internal::ClientAsyncReaderFactory< ::milvus::grpc::TableName>::Create(channel_.get(), cq, rpcmethod_ShowTables_, context, request, true, tag);
}

::grpc::ClientAsyncReader< ::milvus::grpc::TableName>* MilvusService::Stub::PrepareAsyncShowTablesRaw(::grpc::ClientContext* context, const ::milvus::grpc::Command& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncReaderFactory< ::milvus::grpc::TableName>::Create(channel_.get(), cq, rpcmethod_ShowTables_, context, request, false, nullptr);
}

::grpc::Status MilvusService::Stub::Ping(::grpc::ClientContext* context, const ::milvus::grpc::Command& request, ::milvus::grpc::ServerStatus* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_Ping_, context, request, response);
}

void MilvusService::Stub::experimental_async::Ping(::grpc::ClientContext* context, const ::milvus::grpc::Command* request, ::milvus::grpc::ServerStatus* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_Ping_, context, request, response, std::move(f));
}

void MilvusService::Stub::experimental_async::Ping(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::milvus::grpc::ServerStatus* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_Ping_, context, request, response, std::move(f));
}

void MilvusService::Stub::experimental_async::Ping(::grpc::ClientContext* context, const ::milvus::grpc::Command* request, ::milvus::grpc::ServerStatus* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_Ping_, context, request, response, reactor);
}

void MilvusService::Stub::experimental_async::Ping(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::milvus::grpc::ServerStatus* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_Ping_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::milvus::grpc::ServerStatus>* MilvusService::Stub::AsyncPingRaw(::grpc::ClientContext* context, const ::milvus::grpc::Command& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::milvus::grpc::ServerStatus>::Create(channel_.get(), cq, rpcmethod_Ping_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::milvus::grpc::ServerStatus>* MilvusService::Stub::PrepareAsyncPingRaw(::grpc::ClientContext* context, const ::milvus::grpc::Command& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::milvus::grpc::ServerStatus>::Create(channel_.get(), cq, rpcmethod_Ping_, context, request, false);
}

MilvusService::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      MilvusService_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< MilvusService::Service, ::milvus::grpc::TableSchema, ::milvus::grpc::Status>(
          std::mem_fn(&MilvusService::Service::CreateTable), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      MilvusService_method_names[1],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< MilvusService::Service, ::milvus::grpc::TableName, ::milvus::grpc::BoolReply>(
          std::mem_fn(&MilvusService::Service::HasTable), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      MilvusService_method_names[2],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< MilvusService::Service, ::milvus::grpc::TableName, ::milvus::grpc::Status>(
          std::mem_fn(&MilvusService::Service::DropTable), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      MilvusService_method_names[3],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< MilvusService::Service, ::milvus::grpc::TableName, ::milvus::grpc::Status>(
          std::mem_fn(&MilvusService::Service::BuildIndex), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      MilvusService_method_names[4],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< MilvusService::Service, ::milvus::grpc::InsertInfos, ::milvus::grpc::VectorIds>(
          std::mem_fn(&MilvusService::Service::InsertVector), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      MilvusService_method_names[5],
      ::grpc::internal::RpcMethod::SERVER_STREAMING,
      new ::grpc::internal::ServerStreamingHandler< MilvusService::Service, ::milvus::grpc::SearchVectorInfos, ::milvus::grpc::TopKQueryResult>(
          std::mem_fn(&MilvusService::Service::SearchVector), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      MilvusService_method_names[6],
      ::grpc::internal::RpcMethod::SERVER_STREAMING,
      new ::grpc::internal::ServerStreamingHandler< MilvusService::Service, ::milvus::grpc::SearchVectorInFilesInfos, ::milvus::grpc::TopKQueryResult>(
          std::mem_fn(&MilvusService::Service::SearchVectorInFiles), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      MilvusService_method_names[7],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< MilvusService::Service, ::milvus::grpc::TableName, ::milvus::grpc::TableSchema>(
          std::mem_fn(&MilvusService::Service::DescribeTable), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      MilvusService_method_names[8],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< MilvusService::Service, ::milvus::grpc::TableName, ::milvus::grpc::TableRowCount>(
          std::mem_fn(&MilvusService::Service::GetTableRowCount), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      MilvusService_method_names[9],
      ::grpc::internal::RpcMethod::SERVER_STREAMING,
      new ::grpc::internal::ServerStreamingHandler< MilvusService::Service, ::milvus::grpc::Command, ::milvus::grpc::TableName>(
          std::mem_fn(&MilvusService::Service::ShowTables), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      MilvusService_method_names[10],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< MilvusService::Service, ::milvus::grpc::Command, ::milvus::grpc::ServerStatus>(
          std::mem_fn(&MilvusService::Service::Ping), this)));
}

MilvusService::Service::~Service() {
}

::grpc::Status MilvusService::Service::CreateTable(::grpc::ServerContext* context, const ::milvus::grpc::TableSchema* request, ::milvus::grpc::Status* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status MilvusService::Service::HasTable(::grpc::ServerContext* context, const ::milvus::grpc::TableName* request, ::milvus::grpc::BoolReply* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status MilvusService::Service::DropTable(::grpc::ServerContext* context, const ::milvus::grpc::TableName* request, ::milvus::grpc::Status* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status MilvusService::Service::BuildIndex(::grpc::ServerContext* context, const ::milvus::grpc::TableName* request, ::milvus::grpc::Status* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status MilvusService::Service::InsertVector(::grpc::ServerContext* context, const ::milvus::grpc::InsertInfos* request, ::milvus::grpc::VectorIds* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status MilvusService::Service::SearchVector(::grpc::ServerContext* context, const ::milvus::grpc::SearchVectorInfos* request, ::grpc::ServerWriter< ::milvus::grpc::TopKQueryResult>* writer) {
  (void) context;
  (void) request;
  (void) writer;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status MilvusService::Service::SearchVectorInFiles(::grpc::ServerContext* context, const ::milvus::grpc::SearchVectorInFilesInfos* request, ::grpc::ServerWriter< ::milvus::grpc::TopKQueryResult>* writer) {
  (void) context;
  (void) request;
  (void) writer;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status MilvusService::Service::DescribeTable(::grpc::ServerContext* context, const ::milvus::grpc::TableName* request, ::milvus::grpc::TableSchema* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status MilvusService::Service::GetTableRowCount(::grpc::ServerContext* context, const ::milvus::grpc::TableName* request, ::milvus::grpc::TableRowCount* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status MilvusService::Service::ShowTables(::grpc::ServerContext* context, const ::milvus::grpc::Command* request, ::grpc::ServerWriter< ::milvus::grpc::TableName>* writer) {
  (void) context;
  (void) request;
  (void) writer;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status MilvusService::Service::Ping(::grpc::ServerContext* context, const ::milvus::grpc::Command* request, ::milvus::grpc::ServerStatus* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace milvus
}  // namespace grpc

