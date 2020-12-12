// Code generated by protoc-gen-go. DO NOT EDIT.
// source: index_builder.proto

package indexbuilderpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	commonpb "github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
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

type IndexStatus int32

const (
	IndexStatus_NONE       IndexStatus = 0
	IndexStatus_UNISSUED   IndexStatus = 1
	IndexStatus_INPROGRESS IndexStatus = 2
	IndexStatus_FINISHED   IndexStatus = 3
)

var IndexStatus_name = map[int32]string{
	0: "NONE",
	1: "UNISSUED",
	2: "INPROGRESS",
	3: "FINISHED",
}

var IndexStatus_value = map[string]int32{
	"NONE":       0,
	"UNISSUED":   1,
	"INPROGRESS": 2,
	"FINISHED":   3,
}

func (x IndexStatus) String() string {
	return proto.EnumName(IndexStatus_name, int32(x))
}

func (IndexStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c1d6a79d693ba681, []int{0}
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
	return fileDescriptor_c1d6a79d693ba681, []int{0}
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
	return fileDescriptor_c1d6a79d693ba681, []int{1}
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

type GetIndexFilePathsRequest struct {
	IndexID              int64    `protobuf:"varint,1,opt,name=indexID,proto3" json:"indexID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIndexFilePathsRequest) Reset()         { *m = GetIndexFilePathsRequest{} }
func (m *GetIndexFilePathsRequest) String() string { return proto.CompactTextString(m) }
func (*GetIndexFilePathsRequest) ProtoMessage()    {}
func (*GetIndexFilePathsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1d6a79d693ba681, []int{2}
}

func (m *GetIndexFilePathsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIndexFilePathsRequest.Unmarshal(m, b)
}
func (m *GetIndexFilePathsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIndexFilePathsRequest.Marshal(b, m, deterministic)
}
func (m *GetIndexFilePathsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIndexFilePathsRequest.Merge(m, src)
}
func (m *GetIndexFilePathsRequest) XXX_Size() int {
	return xxx_messageInfo_GetIndexFilePathsRequest.Size(m)
}
func (m *GetIndexFilePathsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIndexFilePathsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetIndexFilePathsRequest proto.InternalMessageInfo

func (m *GetIndexFilePathsRequest) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

type GetIndexFilePathsResponse struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IndexID              int64            `protobuf:"varint,2,opt,name=indexID,proto3" json:"indexID,omitempty"`
	IndexFilePaths       []string         `protobuf:"bytes,3,rep,name=index_file_paths,json=indexFilePaths,proto3" json:"index_file_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetIndexFilePathsResponse) Reset()         { *m = GetIndexFilePathsResponse{} }
func (m *GetIndexFilePathsResponse) String() string { return proto.CompactTextString(m) }
func (*GetIndexFilePathsResponse) ProtoMessage()    {}
func (*GetIndexFilePathsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1d6a79d693ba681, []int{3}
}

func (m *GetIndexFilePathsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIndexFilePathsResponse.Unmarshal(m, b)
}
func (m *GetIndexFilePathsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIndexFilePathsResponse.Marshal(b, m, deterministic)
}
func (m *GetIndexFilePathsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIndexFilePathsResponse.Merge(m, src)
}
func (m *GetIndexFilePathsResponse) XXX_Size() int {
	return xxx_messageInfo_GetIndexFilePathsResponse.Size(m)
}
func (m *GetIndexFilePathsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIndexFilePathsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetIndexFilePathsResponse proto.InternalMessageInfo

func (m *GetIndexFilePathsResponse) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetIndexFilePathsResponse) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

func (m *GetIndexFilePathsResponse) GetIndexFilePaths() []string {
	if m != nil {
		return m.IndexFilePaths
	}
	return nil
}

type DescribleIndexRequest struct {
	IndexID              int64    `protobuf:"varint,1,opt,name=indexID,proto3" json:"indexID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribleIndexRequest) Reset()         { *m = DescribleIndexRequest{} }
func (m *DescribleIndexRequest) String() string { return proto.CompactTextString(m) }
func (*DescribleIndexRequest) ProtoMessage()    {}
func (*DescribleIndexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1d6a79d693ba681, []int{4}
}

func (m *DescribleIndexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribleIndexRequest.Unmarshal(m, b)
}
func (m *DescribleIndexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribleIndexRequest.Marshal(b, m, deterministic)
}
func (m *DescribleIndexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribleIndexRequest.Merge(m, src)
}
func (m *DescribleIndexRequest) XXX_Size() int {
	return xxx_messageInfo_DescribleIndexRequest.Size(m)
}
func (m *DescribleIndexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribleIndexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DescribleIndexRequest proto.InternalMessageInfo

func (m *DescribleIndexRequest) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

type DescribleIndexResponse struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IndexStatus          IndexStatus      `protobuf:"varint,2,opt,name=index_status,json=indexStatus,proto3,enum=milvus.proto.service.IndexStatus" json:"index_status,omitempty"`
	IndexID              int64            `protobuf:"varint,3,opt,name=indexID,proto3" json:"indexID,omitempty"`
	EnqueTime            int64            `protobuf:"varint,4,opt,name=enque_time,json=enqueTime,proto3" json:"enque_time,omitempty"`
	ScheduleTime         int64            `protobuf:"varint,5,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	BuildCompleteTime    int64            `protobuf:"varint,6,opt,name=build_complete_time,json=buildCompleteTime,proto3" json:"build_complete_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DescribleIndexResponse) Reset()         { *m = DescribleIndexResponse{} }
func (m *DescribleIndexResponse) String() string { return proto.CompactTextString(m) }
func (*DescribleIndexResponse) ProtoMessage()    {}
func (*DescribleIndexResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1d6a79d693ba681, []int{5}
}

func (m *DescribleIndexResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribleIndexResponse.Unmarshal(m, b)
}
func (m *DescribleIndexResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribleIndexResponse.Marshal(b, m, deterministic)
}
func (m *DescribleIndexResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribleIndexResponse.Merge(m, src)
}
func (m *DescribleIndexResponse) XXX_Size() int {
	return xxx_messageInfo_DescribleIndexResponse.Size(m)
}
func (m *DescribleIndexResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribleIndexResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DescribleIndexResponse proto.InternalMessageInfo

func (m *DescribleIndexResponse) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribleIndexResponse) GetIndexStatus() IndexStatus {
	if m != nil {
		return m.IndexStatus
	}
	return IndexStatus_NONE
}

func (m *DescribleIndexResponse) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

func (m *DescribleIndexResponse) GetEnqueTime() int64 {
	if m != nil {
		return m.EnqueTime
	}
	return 0
}

func (m *DescribleIndexResponse) GetScheduleTime() int64 {
	if m != nil {
		return m.ScheduleTime
	}
	return 0
}

func (m *DescribleIndexResponse) GetBuildCompleteTime() int64 {
	if m != nil {
		return m.BuildCompleteTime
	}
	return 0
}

type IndexMeta struct {
	Status               IndexStatus        `protobuf:"varint,1,opt,name=status,proto3,enum=milvus.proto.service.IndexStatus" json:"status,omitempty"`
	IndexID              int64              `protobuf:"varint,2,opt,name=indexID,proto3" json:"indexID,omitempty"`
	EnqueTime            int64              `protobuf:"varint,3,opt,name=enque_time,json=enqueTime,proto3" json:"enque_time,omitempty"`
	ScheduleTime         int64              `protobuf:"varint,4,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	BuildCompleteTime    int64              `protobuf:"varint,5,opt,name=build_complete_time,json=buildCompleteTime,proto3" json:"build_complete_time,omitempty"`
	Req                  *BuildIndexRequest `protobuf:"bytes,6,opt,name=req,proto3" json:"req,omitempty"`
	IndexFilePaths       []string           `protobuf:"bytes,7,rep,name=index_file_paths,json=indexFilePaths,proto3" json:"index_file_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IndexMeta) Reset()         { *m = IndexMeta{} }
func (m *IndexMeta) String() string { return proto.CompactTextString(m) }
func (*IndexMeta) ProtoMessage()    {}
func (*IndexMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1d6a79d693ba681, []int{6}
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

func (m *IndexMeta) GetStatus() IndexStatus {
	if m != nil {
		return m.Status
	}
	return IndexStatus_NONE
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
	proto.RegisterEnum("milvus.proto.service.IndexStatus", IndexStatus_name, IndexStatus_value)
	proto.RegisterType((*BuildIndexRequest)(nil), "milvus.proto.service.BuildIndexRequest")
	proto.RegisterType((*BuildIndexResponse)(nil), "milvus.proto.service.BuildIndexResponse")
	proto.RegisterType((*GetIndexFilePathsRequest)(nil), "milvus.proto.service.GetIndexFilePathsRequest")
	proto.RegisterType((*GetIndexFilePathsResponse)(nil), "milvus.proto.service.GetIndexFilePathsResponse")
	proto.RegisterType((*DescribleIndexRequest)(nil), "milvus.proto.service.DescribleIndexRequest")
	proto.RegisterType((*DescribleIndexResponse)(nil), "milvus.proto.service.DescribleIndexResponse")
	proto.RegisterType((*IndexMeta)(nil), "milvus.proto.service.IndexMeta")
}

func init() { proto.RegisterFile("index_builder.proto", fileDescriptor_c1d6a79d693ba681) }

var fileDescriptor_c1d6a79d693ba681 = []byte{
	// 626 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0xae, 0xe3, 0xf4, 0x90, 0x49, 0x5a, 0x25, 0xdb, 0xff, 0x47, 0x26, 0xa8, 0x52, 0x08, 0x17,
	0x44, 0x1c, 0x12, 0x91, 0x72, 0xd3, 0xdb, 0x36, 0x69, 0x1b, 0x21, 0xd2, 0xc8, 0xa6, 0x5c, 0x70,
	0x13, 0xf9, 0x30, 0x90, 0x95, 0xd6, 0x87, 0x7a, 0xd7, 0x15, 0xed, 0x73, 0xf0, 0x02, 0x3c, 0x06,
	0x12, 0x8f, 0xc2, 0xc3, 0x20, 0xaf, 0x37, 0xa5, 0x6e, 0x5d, 0xc5, 0x12, 0xe2, 0xd2, 0xb3, 0xdf,
	0xf7, 0xcd, 0x7c, 0x33, 0x3b, 0x6b, 0xd8, 0xa5, 0x81, 0x87, 0x5f, 0xe7, 0x4e, 0x42, 0x99, 0x87,
	0x71, 0x3f, 0x8a, 0x43, 0x11, 0x92, 0xff, 0x7c, 0xca, 0x2e, 0x13, 0x9e, 0x7d, 0xf5, 0x39, 0xc6,
	0x97, 0xd4, 0xc5, 0x76, 0xc3, 0x0d, 0x7d, 0x3f, 0x0c, 0xb2, 0x68, 0xf7, 0xa7, 0x06, 0xad, 0xc3,
	0x94, 0x35, 0x49, 0x05, 0x4c, 0xbc, 0x48, 0x90, 0x0b, 0xb2, 0x07, 0xe0, 0xd9, 0xc2, 0x9e, 0x47,
	0xb6, 0x58, 0x70, 0xa3, 0xd2, 0xd1, 0x7b, 0x35, 0xb3, 0x96, 0x46, 0x66, 0x69, 0x80, 0x1c, 0x42,
	0x5d, 0x5c, 0x45, 0x38, 0x8f, 0xec, 0xd8, 0xf6, 0xb9, 0xa1, 0x77, 0xf4, 0x5e, 0x7d, 0xf8, 0xb4,
	0x9f, 0x4b, 0xa7, 0xb2, 0xbc, 0xc3, 0xab, 0x8f, 0x36, 0x4b, 0x70, 0x66, 0xd3, 0xd8, 0x84, 0x94,
	0x35, 0x93, 0x24, 0x32, 0x82, 0x46, 0x56, 0xb3, 0x12, 0xa9, 0x96, 0x15, 0xa9, 0x4b, 0x5a, 0xa6,
	0xd2, 0x75, 0x81, 0xdc, 0xae, 0x9e, 0x47, 0x61, 0xc0, 0x91, 0xec, 0xc3, 0x06, 0x17, 0xb6, 0x48,
	0xb8, 0xa1, 0x75, 0xb4, 0x5e, 0x7d, 0xf8, 0xa4, 0x50, 0xd5, 0x92, 0x10, 0x53, 0x41, 0x89, 0x01,
	0x9b, 0x52, 0x79, 0x32, 0x32, 0x2a, 0x1d, 0xad, 0xa7, 0x9b, 0xcb, 0xcf, 0xee, 0x5b, 0x30, 0x4e,
	0x50, 0xc8, 0x14, 0xc7, 0x94, 0xa1, 0xec, 0xc1, 0xb2, 0x53, 0xb7, 0x58, 0x5a, 0x9e, 0xf5, 0x4d,
	0x83, 0xc7, 0x05, 0xb4, 0x7f, 0x52, 0x22, 0xe9, 0x41, 0x33, 0xeb, 0xe6, 0x67, 0xca, 0x50, 0x8d,
	0x4d, 0x97, 0x63, 0xdb, 0xa1, 0xb9, 0x02, 0xba, 0x6f, 0xe0, 0xff, 0x11, 0x72, 0x37, 0xa6, 0x0e,
	0xc3, 0xdc, 0xcc, 0x1f, 0x76, 0xf2, 0xbd, 0x02, 0x8f, 0xee, 0x72, 0xfe, 0xc6, 0xc6, 0xcd, 0xe8,
	0x15, 0x35, 0xf5, 0xb2, 0x73, 0x77, 0xf4, 0xea, 0xba, 0xf6, 0x65, 0x3e, 0x25, 0x90, 0x8d, 0xde,
	0xba, 0xd7, 0x0c, 0x3d, 0xdf, 0x8c, 0x3d, 0x00, 0x0c, 0x2e, 0x12, 0x9c, 0x0b, 0xea, 0xa3, 0x51,
	0x95, 0x87, 0x35, 0x19, 0xf9, 0x40, 0x7d, 0x24, 0xcf, 0x60, 0x9b, 0xbb, 0x0b, 0xf4, 0x12, 0xa6,
	0x10, 0xeb, 0x12, 0xd1, 0x58, 0x06, 0x25, 0xa8, 0x0f, 0xbb, 0x72, 0x99, 0xe6, 0x6e, 0xe8, 0x47,
	0x0c, 0x85, 0x82, 0x6e, 0x48, 0x68, 0x4b, 0x1e, 0x1d, 0xa9, 0x93, 0x14, 0xdf, 0xfd, 0x51, 0x81,
	0x9a, 0x2c, 0xf5, 0x3d, 0x0a, 0x9b, 0x1c, 0xe4, 0xda, 0x52, 0xca, 0xdb, 0xea, 0x19, 0xe7, 0x6d,
	0xe9, 0x2b, 0x6d, 0x55, 0xcb, 0xdb, 0x5a, 0x7f, 0xc0, 0x16, 0x39, 0x00, 0x3d, 0xc6, 0x0b, 0x69,
	0xbb, 0x3e, 0x7c, 0x5e, 0xec, 0xe2, 0xde, 0xf3, 0x61, 0xa6, 0x9c, 0xc2, 0x2b, 0xb9, 0x59, 0x74,
	0x25, 0x5f, 0x1c, 0x41, 0xfd, 0x56, 0x27, 0xc8, 0x16, 0x54, 0xa7, 0x67, 0xd3, 0x71, 0x73, 0x8d,
	0x34, 0x60, 0xeb, 0x7c, 0x3a, 0xb1, 0xac, 0xf3, 0xf1, 0xa8, 0xa9, 0x91, 0x1d, 0x80, 0xc9, 0x74,
	0x66, 0x9e, 0x9d, 0x98, 0x63, 0xcb, 0x6a, 0x56, 0xd2, 0xd3, 0xe3, 0xc9, 0x74, 0x62, 0x9d, 0x8e,
	0x47, 0x4d, 0x7d, 0xf8, 0xab, 0x02, 0x2d, 0xa9, 0x22, 0xcb, 0xb1, 0xb2, 0xda, 0x88, 0x0d, 0xf0,
	0xa7, 0x3c, 0x52, 0xd6, 0x40, 0xbb, 0xb7, 0x1a, 0x98, 0x2d, 0x40, 0x77, 0x8d, 0x30, 0xd8, 0x56,
	0xcb, 0x91, 0xed, 0x06, 0x79, 0x59, 0x4c, 0x2e, 0xdc, 0xba, 0xf6, 0xab, 0x72, 0xe0, 0x9b, 0x6c,
	0x97, 0xd0, 0xba, 0xf7, 0xa8, 0x90, 0x7e, 0xb1, 0xc8, 0x43, 0x8f, 0x56, 0x7b, 0x50, 0x1a, 0xbf,
	0xcc, 0x7b, 0x78, 0xfa, 0xe9, 0xf8, 0x0b, 0x15, 0x8b, 0xc4, 0x49, 0x77, 0x7a, 0x70, 0x4d, 0x19,
	0xa3, 0xd7, 0x02, 0xdd, 0xc5, 0x20, 0x53, 0x7a, 0xed, 0x51, 0x2e, 0x62, 0xea, 0x24, 0x02, 0xbd,
	0x01, 0x0d, 0x04, 0xc6, 0x81, 0xcd, 0x06, 0x52, 0x7e, 0x20, 0xa7, 0xad, 0xfe, 0x4c, 0x91, 0xe3,
	0x6c, 0xc8, 0xe8, 0xfe, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0xa2, 0xa4, 0x63, 0xb3, 0x06,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IndexBuildServiceClient is the client API for IndexBuildService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IndexBuildServiceClient interface {
	//*
	// @brief This method is used to create collection
	//
	// @param CollectionSchema, use to provide collection information to be created.
	//
	// @return Status
	BuildIndex(ctx context.Context, in *BuildIndexRequest, opts ...grpc.CallOption) (*BuildIndexResponse, error)
	DescribeIndex(ctx context.Context, in *DescribleIndexRequest, opts ...grpc.CallOption) (*DescribleIndexResponse, error)
	GetIndexFilePaths(ctx context.Context, in *GetIndexFilePathsRequest, opts ...grpc.CallOption) (*GetIndexFilePathsResponse, error)
}

type indexBuildServiceClient struct {
	cc *grpc.ClientConn
}

func NewIndexBuildServiceClient(cc *grpc.ClientConn) IndexBuildServiceClient {
	return &indexBuildServiceClient{cc}
}

func (c *indexBuildServiceClient) BuildIndex(ctx context.Context, in *BuildIndexRequest, opts ...grpc.CallOption) (*BuildIndexResponse, error) {
	out := new(BuildIndexResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.service.IndexBuildService/BuildIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexBuildServiceClient) DescribeIndex(ctx context.Context, in *DescribleIndexRequest, opts ...grpc.CallOption) (*DescribleIndexResponse, error) {
	out := new(DescribleIndexResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.service.IndexBuildService/DescribeIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexBuildServiceClient) GetIndexFilePaths(ctx context.Context, in *GetIndexFilePathsRequest, opts ...grpc.CallOption) (*GetIndexFilePathsResponse, error) {
	out := new(GetIndexFilePathsResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.service.IndexBuildService/GetIndexFilePaths", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexBuildServiceServer is the server API for IndexBuildService service.
type IndexBuildServiceServer interface {
	//*
	// @brief This method is used to create collection
	//
	// @param CollectionSchema, use to provide collection information to be created.
	//
	// @return Status
	BuildIndex(context.Context, *BuildIndexRequest) (*BuildIndexResponse, error)
	DescribeIndex(context.Context, *DescribleIndexRequest) (*DescribleIndexResponse, error)
	GetIndexFilePaths(context.Context, *GetIndexFilePathsRequest) (*GetIndexFilePathsResponse, error)
}

// UnimplementedIndexBuildServiceServer can be embedded to have forward compatible implementations.
type UnimplementedIndexBuildServiceServer struct {
}

func (*UnimplementedIndexBuildServiceServer) BuildIndex(ctx context.Context, req *BuildIndexRequest) (*BuildIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildIndex not implemented")
}
func (*UnimplementedIndexBuildServiceServer) DescribeIndex(ctx context.Context, req *DescribleIndexRequest) (*DescribleIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeIndex not implemented")
}
func (*UnimplementedIndexBuildServiceServer) GetIndexFilePaths(ctx context.Context, req *GetIndexFilePathsRequest) (*GetIndexFilePathsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndexFilePaths not implemented")
}

func RegisterIndexBuildServiceServer(s *grpc.Server, srv IndexBuildServiceServer) {
	s.RegisterService(&_IndexBuildService_serviceDesc, srv)
}

func _IndexBuildService_BuildIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexBuildServiceServer).BuildIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.service.IndexBuildService/BuildIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexBuildServiceServer).BuildIndex(ctx, req.(*BuildIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexBuildService_DescribeIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribleIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexBuildServiceServer).DescribeIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.service.IndexBuildService/DescribeIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexBuildServiceServer).DescribeIndex(ctx, req.(*DescribleIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexBuildService_GetIndexFilePaths_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIndexFilePathsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexBuildServiceServer).GetIndexFilePaths(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.service.IndexBuildService/GetIndexFilePaths",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexBuildServiceServer).GetIndexFilePaths(ctx, req.(*GetIndexFilePathsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IndexBuildService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "milvus.proto.service.IndexBuildService",
	HandlerType: (*IndexBuildServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuildIndex",
			Handler:    _IndexBuildService_BuildIndex_Handler,
		},
		{
			MethodName: "DescribeIndex",
			Handler:    _IndexBuildService_DescribeIndex_Handler,
		},
		{
			MethodName: "GetIndexFilePaths",
			Handler:    _IndexBuildService_GetIndexFilePaths_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "index_builder.proto",
}
