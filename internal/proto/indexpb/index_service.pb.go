// Code generated by protoc-gen-go. DO NOT EDIT.
// source: index_service.proto

package indexpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	commonpb "github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	internalpb2 "github.com/zilliztech/milvus-distributed/internal/proto/internalpb2"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RegisterNodeRequest struct {
	Base                 *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Address              *commonpb.Address `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RegisterNodeRequest) Reset()         { *m = RegisterNodeRequest{} }
func (m *RegisterNodeRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterNodeRequest) ProtoMessage()    {}
func (*RegisterNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{0}
}

func (m *RegisterNodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterNodeRequest.Unmarshal(m, b)
}
func (m *RegisterNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterNodeRequest.Marshal(b, m, deterministic)
}
func (m *RegisterNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterNodeRequest.Merge(m, src)
}
func (m *RegisterNodeRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterNodeRequest.Size(m)
}
func (m *RegisterNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterNodeRequest proto.InternalMessageInfo

func (m *RegisterNodeRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *RegisterNodeRequest) GetAddress() *commonpb.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

type RegisterNodeResponse struct {
	Status               *commonpb.Status        `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	InitParams           *internalpb2.InitParams `protobuf:"bytes,2,opt,name=init_params,json=initParams,proto3" json:"init_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RegisterNodeResponse) Reset()         { *m = RegisterNodeResponse{} }
func (m *RegisterNodeResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterNodeResponse) ProtoMessage()    {}
func (*RegisterNodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{1}
}

func (m *RegisterNodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterNodeResponse.Unmarshal(m, b)
}
func (m *RegisterNodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterNodeResponse.Marshal(b, m, deterministic)
}
func (m *RegisterNodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterNodeResponse.Merge(m, src)
}
func (m *RegisterNodeResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterNodeResponse.Size(m)
}
func (m *RegisterNodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterNodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterNodeResponse proto.InternalMessageInfo

func (m *RegisterNodeResponse) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *RegisterNodeResponse) GetInitParams() *internalpb2.InitParams {
	if m != nil {
		return m.InitParams
	}
	return nil
}

type IndexStatesRequest struct {
	IndexID              []int64  `protobuf:"varint,1,rep,packed,name=indexID,proto3" json:"indexID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexStatesRequest) Reset()         { *m = IndexStatesRequest{} }
func (m *IndexStatesRequest) String() string { return proto.CompactTextString(m) }
func (*IndexStatesRequest) ProtoMessage()    {}
func (*IndexStatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{2}
}

func (m *IndexStatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexStatesRequest.Unmarshal(m, b)
}
func (m *IndexStatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexStatesRequest.Marshal(b, m, deterministic)
}
func (m *IndexStatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexStatesRequest.Merge(m, src)
}
func (m *IndexStatesRequest) XXX_Size() int {
	return xxx_messageInfo_IndexStatesRequest.Size(m)
}
func (m *IndexStatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexStatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IndexStatesRequest proto.InternalMessageInfo

func (m *IndexStatesRequest) GetIndexID() []int64 {
	if m != nil {
		return m.IndexID
	}
	return nil
}

type IndexInfo struct {
	State                commonpb.IndexState `protobuf:"varint,1,opt,name=state,proto3,enum=milvus.proto.common.IndexState" json:"state,omitempty"`
	IndexID              int64               `protobuf:"varint,2,opt,name=indexID,proto3" json:"indexID,omitempty"`
	Reason               string              `protobuf:"bytes,3,opt,name=Reason,proto3" json:"Reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IndexInfo) Reset()         { *m = IndexInfo{} }
func (m *IndexInfo) String() string { return proto.CompactTextString(m) }
func (*IndexInfo) ProtoMessage()    {}
func (*IndexInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{3}
}

func (m *IndexInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexInfo.Unmarshal(m, b)
}
func (m *IndexInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexInfo.Marshal(b, m, deterministic)
}
func (m *IndexInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexInfo.Merge(m, src)
}
func (m *IndexInfo) XXX_Size() int {
	return xxx_messageInfo_IndexInfo.Size(m)
}
func (m *IndexInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IndexInfo proto.InternalMessageInfo

func (m *IndexInfo) GetState() commonpb.IndexState {
	if m != nil {
		return m.State
	}
	return commonpb.IndexState_NONE
}

func (m *IndexInfo) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

func (m *IndexInfo) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type IndexStatesResponse struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	States               []*IndexInfo     `protobuf:"bytes,2,rep,name=states,proto3" json:"states,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *IndexStatesResponse) Reset()         { *m = IndexStatesResponse{} }
func (m *IndexStatesResponse) String() string { return proto.CompactTextString(m) }
func (*IndexStatesResponse) ProtoMessage()    {}
func (*IndexStatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{4}
}

func (m *IndexStatesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexStatesResponse.Unmarshal(m, b)
}
func (m *IndexStatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexStatesResponse.Marshal(b, m, deterministic)
}
func (m *IndexStatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexStatesResponse.Merge(m, src)
}
func (m *IndexStatesResponse) XXX_Size() int {
	return xxx_messageInfo_IndexStatesResponse.Size(m)
}
func (m *IndexStatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexStatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IndexStatesResponse proto.InternalMessageInfo

func (m *IndexStatesResponse) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *IndexStatesResponse) GetStates() []*IndexInfo {
	if m != nil {
		return m.States
	}
	return nil
}

type BuildIndexRequest struct {
	DataPaths            []string                 `protobuf:"bytes,2,rep,name=data_paths,json=dataPaths,proto3" json:"data_paths,omitempty"`
	TypeParams           []*commonpb.KeyValuePair `protobuf:"bytes,3,rep,name=type_params,json=typeParams,proto3" json:"type_params,omitempty"`
	IndexParams          []*commonpb.KeyValuePair `protobuf:"bytes,4,rep,name=index_params,json=indexParams,proto3" json:"index_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *BuildIndexRequest) Reset()         { *m = BuildIndexRequest{} }
func (m *BuildIndexRequest) String() string { return proto.CompactTextString(m) }
func (*BuildIndexRequest) ProtoMessage()    {}
func (*BuildIndexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{5}
}

func (m *BuildIndexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildIndexRequest.Unmarshal(m, b)
}
func (m *BuildIndexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildIndexRequest.Marshal(b, m, deterministic)
}
func (m *BuildIndexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildIndexRequest.Merge(m, src)
}
func (m *BuildIndexRequest) XXX_Size() int {
	return xxx_messageInfo_BuildIndexRequest.Size(m)
}
func (m *BuildIndexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildIndexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuildIndexRequest proto.InternalMessageInfo

func (m *BuildIndexRequest) GetDataPaths() []string {
	if m != nil {
		return m.DataPaths
	}
	return nil
}

func (m *BuildIndexRequest) GetTypeParams() []*commonpb.KeyValuePair {
	if m != nil {
		return m.TypeParams
	}
	return nil
}

func (m *BuildIndexRequest) GetIndexParams() []*commonpb.KeyValuePair {
	if m != nil {
		return m.IndexParams
	}
	return nil
}

type BuildIndexResponse struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IndexID              int64            `protobuf:"varint,2,opt,name=indexID,proto3" json:"indexID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BuildIndexResponse) Reset()         { *m = BuildIndexResponse{} }
func (m *BuildIndexResponse) String() string { return proto.CompactTextString(m) }
func (*BuildIndexResponse) ProtoMessage()    {}
func (*BuildIndexResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{6}
}

func (m *BuildIndexResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildIndexResponse.Unmarshal(m, b)
}
func (m *BuildIndexResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildIndexResponse.Marshal(b, m, deterministic)
}
func (m *BuildIndexResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildIndexResponse.Merge(m, src)
}
func (m *BuildIndexResponse) XXX_Size() int {
	return xxx_messageInfo_BuildIndexResponse.Size(m)
}
func (m *BuildIndexResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildIndexResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BuildIndexResponse proto.InternalMessageInfo

func (m *BuildIndexResponse) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *BuildIndexResponse) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

type BuildIndexCmd struct {
	IndexID              int64              `protobuf:"varint,1,opt,name=indexID,proto3" json:"indexID,omitempty"`
	Req                  *BuildIndexRequest `protobuf:"bytes,2,opt,name=req,proto3" json:"req,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BuildIndexCmd) Reset()         { *m = BuildIndexCmd{} }
func (m *BuildIndexCmd) String() string { return proto.CompactTextString(m) }
func (*BuildIndexCmd) ProtoMessage()    {}
func (*BuildIndexCmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{7}
}

func (m *BuildIndexCmd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildIndexCmd.Unmarshal(m, b)
}
func (m *BuildIndexCmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildIndexCmd.Marshal(b, m, deterministic)
}
func (m *BuildIndexCmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildIndexCmd.Merge(m, src)
}
func (m *BuildIndexCmd) XXX_Size() int {
	return xxx_messageInfo_BuildIndexCmd.Size(m)
}
func (m *BuildIndexCmd) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildIndexCmd.DiscardUnknown(m)
}

var xxx_messageInfo_BuildIndexCmd proto.InternalMessageInfo

func (m *BuildIndexCmd) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

func (m *BuildIndexCmd) GetReq() *BuildIndexRequest {
	if m != nil {
		return m.Req
	}
	return nil
}

type BuildIndexNotification struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IndexID              int64            `protobuf:"varint,2,opt,name=indexID,proto3" json:"indexID,omitempty"`
	IndexFilePaths       []string         `protobuf:"bytes,3,rep,name=index_file_paths,json=indexFilePaths,proto3" json:"index_file_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BuildIndexNotification) Reset()         { *m = BuildIndexNotification{} }
func (m *BuildIndexNotification) String() string { return proto.CompactTextString(m) }
func (*BuildIndexNotification) ProtoMessage()    {}
func (*BuildIndexNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{8}
}

func (m *BuildIndexNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildIndexNotification.Unmarshal(m, b)
}
func (m *BuildIndexNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildIndexNotification.Marshal(b, m, deterministic)
}
func (m *BuildIndexNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildIndexNotification.Merge(m, src)
}
func (m *BuildIndexNotification) XXX_Size() int {
	return xxx_messageInfo_BuildIndexNotification.Size(m)
}
func (m *BuildIndexNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildIndexNotification.DiscardUnknown(m)
}

var xxx_messageInfo_BuildIndexNotification proto.InternalMessageInfo

func (m *BuildIndexNotification) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *BuildIndexNotification) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

func (m *BuildIndexNotification) GetIndexFilePaths() []string {
	if m != nil {
		return m.IndexFilePaths
	}
	return nil
}

type IndexFilePathRequest struct {
	IndexID              int64    `protobuf:"varint,1,opt,name=indexID,proto3" json:"indexID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexFilePathRequest) Reset()         { *m = IndexFilePathRequest{} }
func (m *IndexFilePathRequest) String() string { return proto.CompactTextString(m) }
func (*IndexFilePathRequest) ProtoMessage()    {}
func (*IndexFilePathRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{9}
}

func (m *IndexFilePathRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexFilePathRequest.Unmarshal(m, b)
}
func (m *IndexFilePathRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexFilePathRequest.Marshal(b, m, deterministic)
}
func (m *IndexFilePathRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexFilePathRequest.Merge(m, src)
}
func (m *IndexFilePathRequest) XXX_Size() int {
	return xxx_messageInfo_IndexFilePathRequest.Size(m)
}
func (m *IndexFilePathRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexFilePathRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IndexFilePathRequest proto.InternalMessageInfo

func (m *IndexFilePathRequest) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

type IndexFilePathsResponse struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IndexID              int64            `protobuf:"varint,2,opt,name=indexID,proto3" json:"indexID,omitempty"`
	IndexFilePaths       []string         `protobuf:"bytes,3,rep,name=index_file_paths,json=indexFilePaths,proto3" json:"index_file_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *IndexFilePathsResponse) Reset()         { *m = IndexFilePathsResponse{} }
func (m *IndexFilePathsResponse) String() string { return proto.CompactTextString(m) }
func (*IndexFilePathsResponse) ProtoMessage()    {}
func (*IndexFilePathsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{10}
}

func (m *IndexFilePathsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexFilePathsResponse.Unmarshal(m, b)
}
func (m *IndexFilePathsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexFilePathsResponse.Marshal(b, m, deterministic)
}
func (m *IndexFilePathsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexFilePathsResponse.Merge(m, src)
}
func (m *IndexFilePathsResponse) XXX_Size() int {
	return xxx_messageInfo_IndexFilePathsResponse.Size(m)
}
func (m *IndexFilePathsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexFilePathsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IndexFilePathsResponse proto.InternalMessageInfo

func (m *IndexFilePathsResponse) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *IndexFilePathsResponse) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

func (m *IndexFilePathsResponse) GetIndexFilePaths() []string {
	if m != nil {
		return m.IndexFilePaths
	}
	return nil
}

type IndexMeta struct {
	State                commonpb.IndexState `protobuf:"varint,1,opt,name=state,proto3,enum=milvus.proto.common.IndexState" json:"state,omitempty"`
	IndexID              int64               `protobuf:"varint,2,opt,name=indexID,proto3" json:"indexID,omitempty"`
	EnqueTime            int64               `protobuf:"varint,3,opt,name=enque_time,json=enqueTime,proto3" json:"enque_time,omitempty"`
	ScheduleTime         int64               `protobuf:"varint,4,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	BuildCompleteTime    int64               `protobuf:"varint,5,opt,name=build_complete_time,json=buildCompleteTime,proto3" json:"build_complete_time,omitempty"`
	Req                  *BuildIndexRequest  `protobuf:"bytes,6,opt,name=req,proto3" json:"req,omitempty"`
	IndexFilePaths       []string            `protobuf:"bytes,7,rep,name=index_file_paths,json=indexFilePaths,proto3" json:"index_file_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IndexMeta) Reset()         { *m = IndexMeta{} }
func (m *IndexMeta) String() string { return proto.CompactTextString(m) }
func (*IndexMeta) ProtoMessage()    {}
func (*IndexMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d2036b4df73e0a, []int{11}
}

func (m *IndexMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexMeta.Unmarshal(m, b)
}
func (m *IndexMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexMeta.Marshal(b, m, deterministic)
}
func (m *IndexMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexMeta.Merge(m, src)
}
func (m *IndexMeta) XXX_Size() int {
	return xxx_messageInfo_IndexMeta.Size(m)
}
func (m *IndexMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexMeta.DiscardUnknown(m)
}

var xxx_messageInfo_IndexMeta proto.InternalMessageInfo

func (m *IndexMeta) GetState() commonpb.IndexState {
	if m != nil {
		return m.State
	}
	return commonpb.IndexState_NONE
}

func (m *IndexMeta) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

func (m *IndexMeta) GetEnqueTime() int64 {
	if m != nil {
		return m.EnqueTime
	}
	return 0
}

func (m *IndexMeta) GetScheduleTime() int64 {
	if m != nil {
		return m.ScheduleTime
	}
	return 0
}

func (m *IndexMeta) GetBuildCompleteTime() int64 {
	if m != nil {
		return m.BuildCompleteTime
	}
	return 0
}

func (m *IndexMeta) GetReq() *BuildIndexRequest {
	if m != nil {
		return m.Req
	}
	return nil
}

func (m *IndexMeta) GetIndexFilePaths() []string {
	if m != nil {
		return m.IndexFilePaths
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterNodeRequest)(nil), "milvus.proto.index.RegisterNodeRequest")
	proto.RegisterType((*RegisterNodeResponse)(nil), "milvus.proto.index.RegisterNodeResponse")
	proto.RegisterType((*IndexStatesRequest)(nil), "milvus.proto.index.IndexStatesRequest")
	proto.RegisterType((*IndexInfo)(nil), "milvus.proto.index.IndexInfo")
	proto.RegisterType((*IndexStatesResponse)(nil), "milvus.proto.index.IndexStatesResponse")
	proto.RegisterType((*BuildIndexRequest)(nil), "milvus.proto.index.BuildIndexRequest")
	proto.RegisterType((*BuildIndexResponse)(nil), "milvus.proto.index.BuildIndexResponse")
	proto.RegisterType((*BuildIndexCmd)(nil), "milvus.proto.index.BuildIndexCmd")
	proto.RegisterType((*BuildIndexNotification)(nil), "milvus.proto.index.BuildIndexNotification")
	proto.RegisterType((*IndexFilePathRequest)(nil), "milvus.proto.index.IndexFilePathRequest")
	proto.RegisterType((*IndexFilePathsResponse)(nil), "milvus.proto.index.IndexFilePathsResponse")
	proto.RegisterType((*IndexMeta)(nil), "milvus.proto.index.IndexMeta")
}

func init() { proto.RegisterFile("index_service.proto", fileDescriptor_a5d2036b4df73e0a) }

var fileDescriptor_a5d2036b4df73e0a = []byte{
	// 757 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x5d, 0x4f, 0xdb, 0x4a,
	0x10, 0xc5, 0x38, 0x04, 0x65, 0x12, 0x22, 0xd8, 0x20, 0x14, 0xe5, 0x5e, 0x74, 0xc1, 0x57, 0x17,
	0x22, 0xa4, 0xeb, 0xa0, 0x20, 0xda, 0xc7, 0x8a, 0x80, 0x5a, 0x45, 0x15, 0x08, 0xb9, 0x55, 0x1f,
	0x5a, 0x55, 0x91, 0x63, 0x0f, 0x64, 0x55, 0x7f, 0x04, 0xef, 0x1a, 0x15, 0x5e, 0xaa, 0xaa, 0x3f,
	0xa0, 0xea, 0x6f, 0xe9, 0x6b, 0x7f, 0x5c, 0xe5, 0xdd, 0x75, 0x12, 0x83, 0x49, 0x40, 0xd0, 0x37,
	0xef, 0xee, 0x99, 0x33, 0xb3, 0xe7, 0xcc, 0xac, 0xa1, 0x46, 0x03, 0x17, 0x3f, 0xf7, 0x18, 0x46,
	0x97, 0xd4, 0x41, 0x73, 0x18, 0x85, 0x3c, 0x24, 0xc4, 0xa7, 0xde, 0x65, 0xcc, 0xe4, 0xca, 0x14,
	0x88, 0x46, 0xc5, 0x09, 0x7d, 0x3f, 0x0c, 0xe4, 0x5e, 0xa3, 0x4a, 0x03, 0x8e, 0x51, 0x60, 0x7b,
	0x72, 0x6d, 0x7c, 0x81, 0x9a, 0x85, 0xe7, 0x94, 0x71, 0x8c, 0x4e, 0x42, 0x17, 0x2d, 0xbc, 0x88,
	0x91, 0x71, 0xb2, 0x0b, 0x85, 0xbe, 0xcd, 0xb0, 0xae, 0x6d, 0x68, 0xcd, 0x72, 0xfb, 0x6f, 0x33,
	0xc3, 0xab, 0x08, 0x8f, 0xd9, 0x79, 0xc7, 0x66, 0x68, 0x09, 0x24, 0x79, 0x06, 0x8b, 0xb6, 0xeb,
	0x46, 0xc8, 0x58, 0x7d, 0x7e, 0x4a, 0xd0, 0x81, 0xc4, 0x58, 0x29, 0xd8, 0xf8, 0xae, 0xc1, 0x6a,
	0xb6, 0x02, 0x36, 0x0c, 0x03, 0x86, 0x64, 0x0f, 0x8a, 0x8c, 0xdb, 0x3c, 0x66, 0xaa, 0x88, 0xbf,
	0x72, 0xf9, 0xde, 0x08, 0x88, 0xa5, 0xa0, 0xa4, 0x03, 0x65, 0x1a, 0x50, 0xde, 0x1b, 0xda, 0x91,
	0xed, 0xa7, 0x95, 0x6c, 0x9a, 0x37, 0x64, 0x51, 0x0a, 0x74, 0x03, 0xca, 0x4f, 0x05, 0xd0, 0x02,
	0x3a, 0xfa, 0x36, 0x4c, 0x20, 0xdd, 0x44, 0xb9, 0x84, 0x1a, 0x59, 0xaa, 0x48, 0x1d, 0x16, 0x85,
	0x9e, 0xdd, 0xa3, 0xba, 0xb6, 0xa1, 0x37, 0x75, 0x2b, 0x5d, 0x1a, 0x1c, 0x4a, 0x02, 0xdf, 0x0d,
	0xce, 0x42, 0xb2, 0x0f, 0x0b, 0x49, 0x29, 0x52, 0xb9, 0x6a, 0xfb, 0x9f, 0xdc, 0xa2, 0xc7, 0xf4,
	0x96, 0x44, 0x4f, 0xb2, 0x27, 0x35, 0x8f, 0xd9, 0xc9, 0x1a, 0x14, 0x2d, 0xb4, 0x59, 0x18, 0xd4,
	0xf5, 0x0d, 0xad, 0x59, 0xb2, 0xd4, 0xca, 0xf8, 0xaa, 0x41, 0x2d, 0x53, 0xe6, 0x63, 0x64, 0xdb,
	0x97, 0x41, 0x98, 0x28, 0xa6, 0x37, 0xcb, 0xed, 0x75, 0xf3, 0x76, 0x23, 0x99, 0xa3, 0x4b, 0x5a,
	0x0a, 0x6c, 0xfc, 0xd2, 0x60, 0xa5, 0x13, 0x53, 0xcf, 0x15, 0x47, 0xa9, 0x52, 0xeb, 0x00, 0xae,
	0xcd, 0xed, 0xde, 0xd0, 0xe6, 0x03, 0x49, 0x58, 0xb2, 0x4a, 0xc9, 0xce, 0x69, 0xb2, 0x91, 0x58,
	0xc4, 0xaf, 0x86, 0x98, 0x5a, 0xa4, 0x8b, 0x84, 0x9b, 0xb9, 0x55, 0xbe, 0xc6, 0xab, 0x77, 0xb6,
	0x17, 0xe3, 0xa9, 0x4d, 0x23, 0x0b, 0x92, 0x28, 0x69, 0x11, 0x39, 0x82, 0x8a, 0x6c, 0x7f, 0x45,
	0x52, 0xb8, 0x2f, 0x49, 0x59, 0x84, 0x29, 0xa3, 0x1d, 0x20, 0x93, 0xd5, 0x3f, 0x46, 0xc0, 0x3b,
	0xfd, 0x33, 0xfa, 0xb0, 0x34, 0x4e, 0x72, 0xe8, 0xbb, 0xd9, 0x46, 0xca, 0x58, 0xfd, 0x1c, 0xf4,
	0x08, 0x2f, 0x54, 0xd3, 0xfe, 0x97, 0x67, 0xc1, 0x2d, 0xb1, 0xad, 0x24, 0xc2, 0xf8, 0xa1, 0xc1,
	0xda, 0xf8, 0xe8, 0x24, 0xe4, 0xf4, 0x8c, 0x3a, 0x36, 0xa7, 0x61, 0xf0, 0xc4, 0xb7, 0x21, 0x4d,
	0x58, 0x96, 0xc2, 0x9f, 0x51, 0x0f, 0x95, 0xc3, 0xba, 0x70, 0xb8, 0x2a, 0xf6, 0x5f, 0x52, 0x0f,
	0x85, 0xcd, 0xc6, 0x2e, 0xac, 0x76, 0x27, 0x77, 0x72, 0xe7, 0x28, 0xa3, 0x54, 0x72, 0x8b, 0x4c,
	0x08, 0xfb, 0x43, 0x9e, 0x3c, 0xe0, 0x16, 0x3f, 0xe7, 0xd5, 0x70, 0x1f, 0x23, 0xb7, 0x9f, 0x7e,
	0xb8, 0xd7, 0x01, 0x30, 0xb8, 0x88, 0xb1, 0xc7, 0xa9, 0x8f, 0x62, 0xc0, 0x75, 0xab, 0x24, 0x76,
	0xde, 0x52, 0x1f, 0xc9, 0xbf, 0xb0, 0xc4, 0x9c, 0x01, 0xba, 0xb1, 0xa7, 0x10, 0x05, 0x81, 0xa8,
	0xa4, 0x9b, 0x02, 0x64, 0x42, 0xad, 0x9f, 0x78, 0xdf, 0x73, 0x42, 0x7f, 0xe8, 0x21, 0x57, 0xd0,
	0x05, 0x01, 0x5d, 0x11, 0x47, 0x87, 0xea, 0x44, 0xe0, 0x55, 0x97, 0x15, 0x1f, 0xda, 0x65, 0xb9,
	0xaa, 0x2d, 0xe6, 0xa9, 0xd6, 0xfe, 0x56, 0x80, 0x8a, 0x94, 0x41, 0xfe, 0x9d, 0x88, 0x03, 0x95,
	0xc9, 0x37, 0x9e, 0x6c, 0xe7, 0xa5, 0xcd, 0xf9, 0x0f, 0x35, 0x9a, 0xb3, 0x81, 0xb2, 0x45, 0x8c,
	0x39, 0xf2, 0x11, 0x60, 0x5c, 0x39, 0xb9, 0xdf, 0xcd, 0x1a, 0x5b, 0xb3, 0x60, 0x23, 0x7a, 0x07,
	0xaa, 0xaf, 0x90, 0x4f, 0x3c, 0xb9, 0x64, 0xeb, 0xce, 0x57, 0x32, 0xf3, 0xeb, 0x68, 0x6c, 0xcf,
	0xc4, 0x8d, 0x92, 0x7c, 0x82, 0x95, 0x34, 0xc9, 0x48, 0x4e, 0xd2, 0xbc, 0x33, 0xfe, 0xc6, 0x70,
	0x35, 0x76, 0x66, 0x22, 0x59, 0x46, 0xb0, 0x65, 0xf1, 0x56, 0x5c, 0x4d, 0xc8, 0xb6, 0x33, 0x5d,
	0x8f, 0xc9, 0xb7, 0xa5, 0x31, 0x6d, 0x0a, 0x8d, 0xb9, 0xf6, 0x07, 0x35, 0x3a, 0xc2, 0xf1, 0x93,
	0x8c, 0x39, 0x9b, 0xd3, 0xb3, 0x1c, 0xfa, 0xee, 0x0c, 0xf2, 0xce, 0xc1, 0xfb, 0x17, 0xe7, 0x94,
	0x0f, 0xe2, 0x7e, 0x72, 0xd2, 0xba, 0xa6, 0x9e, 0x47, 0xaf, 0x39, 0x3a, 0x83, 0x96, 0x8c, 0xfa,
	0xdf, 0xa5, 0x8c, 0x47, 0xb4, 0x1f, 0x73, 0x74, 0x5b, 0xe9, 0x0f, 0xbf, 0x25, 0xa8, 0x5a, 0x22,
	0xdb, 0xb0, 0xdf, 0x2f, 0x8a, 0xe5, 0xde, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x32, 0xc8,
	0x07, 0x4a, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IndexServiceClient is the client API for IndexService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IndexServiceClient interface {
	//*
	// @brief This method is used to create collection
	//
	// @param CollectionSchema, use to provide collection information to be created.
	//
	// @return Status
	RegisterNode(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*RegisterNodeResponse, error)
	BuildIndex(ctx context.Context, in *BuildIndexRequest, opts ...grpc.CallOption) (*BuildIndexResponse, error)
	GetIndexStates(ctx context.Context, in *IndexStatesRequest, opts ...grpc.CallOption) (*IndexStatesResponse, error)
	GetIndexFilePaths(ctx context.Context, in *IndexFilePathRequest, opts ...grpc.CallOption) (*IndexFilePathsResponse, error)
	NotifyBuildIndex(ctx context.Context, in *BuildIndexNotification, opts ...grpc.CallOption) (*commonpb.Status, error)
}

type indexServiceClient struct {
	cc *grpc.ClientConn
}

func NewIndexServiceClient(cc *grpc.ClientConn) IndexServiceClient {
	return &indexServiceClient{cc}
}

func (c *indexServiceClient) RegisterNode(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*RegisterNodeResponse, error) {
	out := new(RegisterNodeResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.index.IndexService/RegisterNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) BuildIndex(ctx context.Context, in *BuildIndexRequest, opts ...grpc.CallOption) (*BuildIndexResponse, error) {
	out := new(BuildIndexResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.index.IndexService/BuildIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) GetIndexStates(ctx context.Context, in *IndexStatesRequest, opts ...grpc.CallOption) (*IndexStatesResponse, error) {
	out := new(IndexStatesResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.index.IndexService/GetIndexStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) GetIndexFilePaths(ctx context.Context, in *IndexFilePathRequest, opts ...grpc.CallOption) (*IndexFilePathsResponse, error) {
	out := new(IndexFilePathsResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.index.IndexService/GetIndexFilePaths", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) NotifyBuildIndex(ctx context.Context, in *BuildIndexNotification, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.index.IndexService/NotifyBuildIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexServiceServer is the server API for IndexService service.
type IndexServiceServer interface {
	//*
	// @brief This method is used to create collection
	//
	// @param CollectionSchema, use to provide collection information to be created.
	//
	// @return Status
	RegisterNode(context.Context, *RegisterNodeRequest) (*RegisterNodeResponse, error)
	BuildIndex(context.Context, *BuildIndexRequest) (*BuildIndexResponse, error)
	GetIndexStates(context.Context, *IndexStatesRequest) (*IndexStatesResponse, error)
	GetIndexFilePaths(context.Context, *IndexFilePathRequest) (*IndexFilePathsResponse, error)
	NotifyBuildIndex(context.Context, *BuildIndexNotification) (*commonpb.Status, error)
}

// UnimplementedIndexServiceServer can be embedded to have forward compatible implementations.
type UnimplementedIndexServiceServer struct {
}

func (*UnimplementedIndexServiceServer) RegisterNode(ctx context.Context, req *RegisterNodeRequest) (*RegisterNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNode not implemented")
}
func (*UnimplementedIndexServiceServer) BuildIndex(ctx context.Context, req *BuildIndexRequest) (*BuildIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildIndex not implemented")
}
func (*UnimplementedIndexServiceServer) GetIndexStates(ctx context.Context, req *IndexStatesRequest) (*IndexStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndexStates not implemented")
}
func (*UnimplementedIndexServiceServer) GetIndexFilePaths(ctx context.Context, req *IndexFilePathRequest) (*IndexFilePathsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndexFilePaths not implemented")
}
func (*UnimplementedIndexServiceServer) NotifyBuildIndex(ctx context.Context, req *BuildIndexNotification) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyBuildIndex not implemented")
}

func RegisterIndexServiceServer(s *grpc.Server, srv IndexServiceServer) {
	s.RegisterService(&_IndexService_serviceDesc, srv)
}

func _IndexService_RegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).RegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.index.IndexService/RegisterNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).RegisterNode(ctx, req.(*RegisterNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_BuildIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).BuildIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.index.IndexService/BuildIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).BuildIndex(ctx, req.(*BuildIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_GetIndexStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).GetIndexStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.index.IndexService/GetIndexStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).GetIndexStates(ctx, req.(*IndexStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_GetIndexFilePaths_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexFilePathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).GetIndexFilePaths(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.index.IndexService/GetIndexFilePaths",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).GetIndexFilePaths(ctx, req.(*IndexFilePathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_NotifyBuildIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildIndexNotification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).NotifyBuildIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.index.IndexService/NotifyBuildIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).NotifyBuildIndex(ctx, req.(*BuildIndexNotification))
	}
	return interceptor(ctx, in, info, handler)
}

var _IndexService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "milvus.proto.index.IndexService",
	HandlerType: (*IndexServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNode",
			Handler:    _IndexService_RegisterNode_Handler,
		},
		{
			MethodName: "BuildIndex",
			Handler:    _IndexService_BuildIndex_Handler,
		},
		{
			MethodName: "GetIndexStates",
			Handler:    _IndexService_GetIndexStates_Handler,
		},
		{
			MethodName: "GetIndexFilePaths",
			Handler:    _IndexService_GetIndexFilePaths_Handler,
		},
		{
			MethodName: "NotifyBuildIndex",
			Handler:    _IndexService_NotifyBuildIndex_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "index_service.proto",
}

// IndexNodeClient is the client API for IndexNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IndexNodeClient interface {
	//*
	// @brief This method is used to create collection
	//
	// @param CollectionSchema, use to provide collection information to be created.
	//
	// @return Status
	BuildIndex(ctx context.Context, in *BuildIndexCmd, opts ...grpc.CallOption) (*commonpb.Status, error)
}

type indexNodeClient struct {
	cc *grpc.ClientConn
}

func NewIndexNodeClient(cc *grpc.ClientConn) IndexNodeClient {
	return &indexNodeClient{cc}
}

func (c *indexNodeClient) BuildIndex(ctx context.Context, in *BuildIndexCmd, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.index.IndexNode/BuildIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexNodeServer is the server API for IndexNode service.
type IndexNodeServer interface {
	//*
	// @brief This method is used to create collection
	//
	// @param CollectionSchema, use to provide collection information to be created.
	//
	// @return Status
	BuildIndex(context.Context, *BuildIndexCmd) (*commonpb.Status, error)
}

// UnimplementedIndexNodeServer can be embedded to have forward compatible implementations.
type UnimplementedIndexNodeServer struct {
}

func (*UnimplementedIndexNodeServer) BuildIndex(ctx context.Context, req *BuildIndexCmd) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildIndex not implemented")
}

func RegisterIndexNodeServer(s *grpc.Server, srv IndexNodeServer) {
	s.RegisterService(&_IndexNode_serviceDesc, srv)
}

func _IndexNode_BuildIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildIndexCmd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexNodeServer).BuildIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.index.IndexNode/BuildIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexNodeServer).BuildIndex(ctx, req.(*BuildIndexCmd))
	}
	return interceptor(ctx, in, info, handler)
}

var _IndexNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "milvus.proto.index.IndexNode",
	HandlerType: (*IndexNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuildIndex",
			Handler:    _IndexNode_BuildIndex_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "index_service.proto",
}
