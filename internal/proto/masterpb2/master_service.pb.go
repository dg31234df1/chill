// Code generated by protoc-gen-go. DO NOT EDIT.
// source: master_service.proto

package masterpb2

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	commonpb "github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	internalpb2 "github.com/zilliztech/milvus-distributed/internal/proto/internalpb2"
	milvuspb "github.com/zilliztech/milvus-distributed/internal/proto/milvuspb"
	servicepb "github.com/zilliztech/milvus-distributed/internal/proto/servicepb"
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

type IDRequest struct {
	Base                 *internalpb2.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Count                uint32               `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *IDRequest) Reset()         { *m = IDRequest{} }
func (m *IDRequest) String() string { return proto.CompactTextString(m) }
func (*IDRequest) ProtoMessage()    {}
func (*IDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a501bcba839fe29, []int{0}
}

func (m *IDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDRequest.Unmarshal(m, b)
}
func (m *IDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDRequest.Marshal(b, m, deterministic)
}
func (m *IDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDRequest.Merge(m, src)
}
func (m *IDRequest) XXX_Size() int {
	return xxx_messageInfo_IDRequest.Size(m)
}
func (m *IDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IDRequest proto.InternalMessageInfo

func (m *IDRequest) GetBase() *internalpb2.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *IDRequest) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type IDResponse struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	ID                   int64            `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Count                uint32           `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *IDResponse) Reset()         { *m = IDResponse{} }
func (m *IDResponse) String() string { return proto.CompactTextString(m) }
func (*IDResponse) ProtoMessage()    {}
func (*IDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a501bcba839fe29, []int{1}
}

func (m *IDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDResponse.Unmarshal(m, b)
}
func (m *IDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDResponse.Marshal(b, m, deterministic)
}
func (m *IDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDResponse.Merge(m, src)
}
func (m *IDResponse) XXX_Size() int {
	return xxx_messageInfo_IDResponse.Size(m)
}
func (m *IDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IDResponse proto.InternalMessageInfo

func (m *IDResponse) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *IDResponse) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *IDResponse) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type TsoRequest struct {
	Base                 *internalpb2.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Count                uint32               `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TsoRequest) Reset()         { *m = TsoRequest{} }
func (m *TsoRequest) String() string { return proto.CompactTextString(m) }
func (*TsoRequest) ProtoMessage()    {}
func (*TsoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a501bcba839fe29, []int{2}
}

func (m *TsoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TsoRequest.Unmarshal(m, b)
}
func (m *TsoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TsoRequest.Marshal(b, m, deterministic)
}
func (m *TsoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TsoRequest.Merge(m, src)
}
func (m *TsoRequest) XXX_Size() int {
	return xxx_messageInfo_TsoRequest.Size(m)
}
func (m *TsoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TsoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TsoRequest proto.InternalMessageInfo

func (m *TsoRequest) GetBase() *internalpb2.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *TsoRequest) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type TsoResponse struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Timestamp            uint64           `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Count                uint32           `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TsoResponse) Reset()         { *m = TsoResponse{} }
func (m *TsoResponse) String() string { return proto.CompactTextString(m) }
func (*TsoResponse) ProtoMessage()    {}
func (*TsoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a501bcba839fe29, []int{3}
}

func (m *TsoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TsoResponse.Unmarshal(m, b)
}
func (m *TsoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TsoResponse.Marshal(b, m, deterministic)
}
func (m *TsoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TsoResponse.Merge(m, src)
}
func (m *TsoResponse) XXX_Size() int {
	return xxx_messageInfo_TsoResponse.Size(m)
}
func (m *TsoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TsoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TsoResponse proto.InternalMessageInfo

func (m *TsoResponse) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *TsoResponse) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TsoResponse) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*IDRequest)(nil), "milvus.proto.master.IDRequest")
	proto.RegisterType((*IDResponse)(nil), "milvus.proto.master.IDResponse")
	proto.RegisterType((*TsoRequest)(nil), "milvus.proto.master.TsoRequest")
	proto.RegisterType((*TsoResponse)(nil), "milvus.proto.master.TsoResponse")
}

func init() { proto.RegisterFile("master_service.proto", fileDescriptor_9a501bcba839fe29) }

var fileDescriptor_9a501bcba839fe29 = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xed, 0x6e, 0xd3, 0x3c,
	0x14, 0xc7, 0xf7, 0xf6, 0xec, 0xd1, 0x4e, 0xd7, 0x0e, 0xcc, 0x04, 0x53, 0x40, 0xdb, 0xd4, 0x2f,
	0xac, 0xed, 0x48, 0xa4, 0xee, 0x0a, 0xd6, 0x46, 0xda, 0x2a, 0x31, 0x09, 0xa5, 0x1b, 0x12, 0x20,
	0x28, 0x49, 0x6a, 0xb5, 0x96, 0x92, 0xb8, 0xe4, 0x38, 0x1d, 0xda, 0x45, 0x73, 0x0d, 0xa8, 0x79,
	0x71, 0xe2, 0x2d, 0xa9, 0x82, 0xe0, 0xa3, 0x9d, 0x5f, 0x7e, 0xff, 0x1e, 0x9f, 0x93, 0x1a, 0x0e,
	0x7d, 0x1b, 0x05, 0x0d, 0x27, 0x48, 0xc3, 0x25, 0x73, 0xa9, 0xbe, 0x08, 0xb9, 0xe0, 0xe4, 0x85,
	0xcf, 0xbc, 0x65, 0x84, 0xc9, 0x4a, 0x4f, 0x10, 0x6d, 0xdf, 0xe5, 0xbe, 0xcf, 0x83, 0x64, 0x53,
	0xdb, 0x2f, 0x22, 0x5a, 0x8b, 0x05, 0x82, 0x86, 0x81, 0xed, 0xa5, 0xeb, 0xe7, 0xa9, 0x6f, 0xe2,
	0xe3, 0x2c, 0xd9, 0x6a, 0xdf, 0xc1, 0xde, 0xc8, 0xb4, 0xe8, 0x8f, 0x88, 0xa2, 0x20, 0x7d, 0xd8,
	0x71, 0x6c, 0xa4, 0x47, 0x9b, 0xa7, 0x9b, 0x67, 0x8d, 0xfe, 0xb1, 0xae, 0xe4, 0x49, 0xd7, 0x0d,
	0xce, 0x06, 0x36, 0x52, 0x2b, 0x66, 0xc9, 0x21, 0xfc, 0xe7, 0xf2, 0x28, 0x10, 0x47, 0x5b, 0xa7,
	0x9b, 0x67, 0x4d, 0x2b, 0x59, 0xb4, 0x67, 0x00, 0x2b, 0x2d, 0x2e, 0x78, 0x80, 0x94, 0x5c, 0xc0,
	0x2e, 0x0a, 0x5b, 0x44, 0x98, 0x9a, 0x5f, 0xab, 0xe6, 0xb4, 0x82, 0x71, 0x8c, 0x58, 0x29, 0x4a,
	0x5a, 0xb0, 0x35, 0x32, 0x63, 0xeb, 0xb6, 0xb5, 0x35, 0x32, 0xf3, 0xa0, 0xed, 0x62, 0xd0, 0x47,
	0x80, 0x5b, 0xe4, 0xff, 0xa4, 0x00, 0xc5, 0xbb, 0x84, 0x46, 0xec, 0xfd, 0x9b, 0x0a, 0xde, 0xc0,
	0x9e, 0x60, 0x3e, 0x45, 0x61, 0xfb, 0x8b, 0xb8, 0x90, 0x1d, 0x2b, 0xdf, 0x28, 0xcf, 0xed, 0xff,
	0x6a, 0x40, 0xf3, 0x26, 0xee, 0xec, 0x38, 0xe9, 0x15, 0x99, 0xc0, 0xb3, 0x61, 0x48, 0x6d, 0x41,
	0x87, 0xdc, 0xf3, 0xa8, 0x2b, 0x18, 0x0f, 0xc8, 0xb9, 0x1a, 0x9f, 0x2e, 0x1e, 0x63, 0xe9, 0xa9,
	0x68, 0xeb, 0x7e, 0x6c, 0x7b, 0x83, 0x7c, 0x81, 0x96, 0x19, 0xf2, 0x45, 0x41, 0xdf, 0x2d, 0xd5,
	0xab, 0x50, 0x4d, 0xf9, 0x77, 0x68, 0x5e, 0xdb, 0x58, 0x70, 0x77, 0x4a, 0xdd, 0x0a, 0x93, 0xa9,
	0xdb, 0x2a, 0x9a, 0x7d, 0x0c, 0x03, 0xce, 0xbd, 0xac, 0x31, 0xed, 0x0d, 0x72, 0x0f, 0xc4, 0xa4,
	0xe8, 0x86, 0xcc, 0x29, 0x9e, 0x90, 0x5e, 0x5e, 0xc2, 0x13, 0x30, 0xcb, 0x32, 0x6a, 0xf3, 0x32,
	0x78, 0x09, 0xaf, 0xae, 0xa8, 0xc8, 0x1f, 0xad, 0x6a, 0x66, 0x28, 0x98, 0x8b, 0xa4, 0x57, 0xde,
	0x1f, 0x05, 0xc5, 0x2c, 0xfa, 0xbc, 0x1e, 0x2c, 0x73, 0x3d, 0x38, 0x18, 0xcf, 0xf9, 0x7d, 0x0e,
	0x60, 0x45, 0xc3, 0x54, 0x2a, 0x8b, 0xeb, 0xd5, 0x62, 0x65, 0xda, 0x57, 0x38, 0x48, 0xe6, 0xea,
	0x83, 0x1d, 0x0a, 0x16, 0x9f, 0x6d, 0x6f, 0xcd, 0xf4, 0x49, 0xaa, 0xe6, 0x7c, 0x7c, 0x82, 0xe6,
	0x6a, 0xae, 0x72, 0x79, 0xa7, 0x72, 0xf6, 0xfe, 0x54, 0xfd, 0x0d, 0xf6, 0xaf, 0x6d, 0xcc, 0xcd,
	0x67, 0x55, 0x93, 0xf7, 0x44, 0x5c, 0x6f, 0xf0, 0x10, 0x5e, 0x5e, 0x51, 0x21, 0x5f, 0x2e, 0xb4,
	0xbf, 0xbc, 0x1d, 0x0a, 0x89, 0xeb, 0xdb, 0xf1, 0x98, 0x95, 0xa1, 0x0c, 0x5a, 0xab, 0x56, 0xc9,
	0xe7, 0x58, 0x71, 0x60, 0x0a, 0x94, 0x65, 0x75, 0xeb, 0xa0, 0x32, 0xea, 0x0e, 0x1a, 0x49, 0x4f,
	0x47, 0xc1, 0x94, 0xfe, 0x24, 0x6f, 0xd7, 0x74, 0x3d, 0x26, 0x6a, 0xb6, 0x65, 0x0e, 0xcd, 0xec,
	0xb3, 0x4a, 0xc4, 0x9d, 0xb5, 0x9f, 0x9e, 0xa2, 0xee, 0xd6, 0x41, 0x0b, 0x05, 0xb4, 0x2e, 0x3d,
	0x8f, 0xbb, 0xb7, 0xf2, 0x3f, 0xf7, 0x44, 0x2f, 0xb9, 0x42, 0xf5, 0xfc, 0x02, 0xd1, 0x4e, 0xab,
	0x01, 0xa9, 0x7d, 0x0f, 0xff, 0xc7, 0xda, 0x91, 0x49, 0x8e, 0x4b, 0x71, 0x79, 0xa1, 0x6a, 0x27,
	0x95, 0xcf, 0x33, 0xdb, 0x60, 0xf8, 0xf9, 0x72, 0xc6, 0xc4, 0x3c, 0x72, 0x56, 0x07, 0x65, 0x3c,
	0x30, 0xcf, 0x63, 0x0f, 0x82, 0xba, 0x73, 0x23, 0x79, 0xf3, 0xdd, 0x94, 0xa1, 0x08, 0x99, 0x13,
	0x09, 0x3a, 0x35, 0xb2, 0x1b, 0xcc, 0x88, 0x75, 0x46, 0xa2, 0x5b, 0x38, 0x7d, 0x67, 0x37, 0xde,
	0xb8, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xe1, 0xb5, 0xf4, 0x29, 0x38, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MasterServiceClient is the client API for MasterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MasterServiceClient interface {
	//*
	// @brief This method is used to create collection
	//
	// @param CreateCollectionRequest, use to provide collection information to be created.
	//
	// @return Status
	CreateCollection(ctx context.Context, in *milvuspb.CreateCollectionRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	//*
	// @brief This method is used to delete collection.
	//
	// @param DropCollectionRequest, collection name is going to be deleted.
	//
	// @return Status
	DropCollection(ctx context.Context, in *milvuspb.DropCollectionRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	//*
	// @brief This method is used to test collection existence.
	//
	// @param HasCollectionRequest, collection name is going to be tested.
	//
	// @return BoolResponse
	HasCollection(ctx context.Context, in *milvuspb.HasCollectionRequest, opts ...grpc.CallOption) (*servicepb.BoolResponse, error)
	//*
	// @brief This method is used to get collection schema.
	//
	// @param DescribeCollectionRequest, target collection name.
	//
	// @return CollectionSchema
	DescribeCollection(ctx context.Context, in *milvuspb.DescribeCollectionRequest, opts ...grpc.CallOption) (*milvuspb.DescribeCollectionResponse, error)
	GetCollectionStatistics(ctx context.Context, in *milvuspb.CollectionStatsRequest, opts ...grpc.CallOption) (*milvuspb.CollectionStatsResponse, error)
	//*
	// @brief This method is used to list all collections.
	//
	// @return StringListResponse, collection name list
	ShowCollections(ctx context.Context, in *milvuspb.ShowCollectionRequest, opts ...grpc.CallOption) (*milvuspb.ShowCollectionResponse, error)
	//*
	// @brief This method is used to create partition
	//
	// @return Status
	CreatePartition(ctx context.Context, in *milvuspb.CreatePartitionRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	//*
	// @brief This method is used to drop partition
	//
	// @return Status
	DropPartition(ctx context.Context, in *milvuspb.DropPartitionRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	//*
	// @brief This method is used to test partition existence.
	//
	// @return BoolResponse
	HasPartition(ctx context.Context, in *milvuspb.HasPartitionRequest, opts ...grpc.CallOption) (*servicepb.BoolResponse, error)
	GetPartitionStatistics(ctx context.Context, in *milvuspb.PartitionStatsRequest, opts ...grpc.CallOption) (*milvuspb.PartitionStatsResponse, error)
	//*
	// @brief This method is used to show partition information
	//
	// @param ShowPartitionRequest, target collection name.
	//
	// @return StringListResponse
	ShowPartitions(ctx context.Context, in *milvuspb.ShowPartitionRequest, opts ...grpc.CallOption) (*milvuspb.ShowPartitionResponse, error)
	CreateIndex(ctx context.Context, in *milvuspb.CreateIndexRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	DescribeIndex(ctx context.Context, in *milvuspb.DescribeIndexRequest, opts ...grpc.CallOption) (*milvuspb.DescribeIndexResponse, error)
	AllocTimestamp(ctx context.Context, in *TsoRequest, opts ...grpc.CallOption) (*TsoResponse, error)
	AllocID(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*IDResponse, error)
}

type masterServiceClient struct {
	cc *grpc.ClientConn
}

func NewMasterServiceClient(cc *grpc.ClientConn) MasterServiceClient {
	return &masterServiceClient{cc}
}

func (c *masterServiceClient) CreateCollection(ctx context.Context, in *milvuspb.CreateCollectionRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.MasterService/CreateCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) DropCollection(ctx context.Context, in *milvuspb.DropCollectionRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.MasterService/DropCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) HasCollection(ctx context.Context, in *milvuspb.HasCollectionRequest, opts ...grpc.CallOption) (*servicepb.BoolResponse, error) {
	out := new(servicepb.BoolResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.MasterService/HasCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) DescribeCollection(ctx context.Context, in *milvuspb.DescribeCollectionRequest, opts ...grpc.CallOption) (*milvuspb.DescribeCollectionResponse, error) {
	out := new(milvuspb.DescribeCollectionResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.MasterService/DescribeCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) GetCollectionStatistics(ctx context.Context, in *milvuspb.CollectionStatsRequest, opts ...grpc.CallOption) (*milvuspb.CollectionStatsResponse, error) {
	out := new(milvuspb.CollectionStatsResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.MasterService/GetCollectionStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) ShowCollections(ctx context.Context, in *milvuspb.ShowCollectionRequest, opts ...grpc.CallOption) (*milvuspb.ShowCollectionResponse, error) {
	out := new(milvuspb.ShowCollectionResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.MasterService/ShowCollections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) CreatePartition(ctx context.Context, in *milvuspb.CreatePartitionRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.MasterService/CreatePartition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) DropPartition(ctx context.Context, in *milvuspb.DropPartitionRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.MasterService/DropPartition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) HasPartition(ctx context.Context, in *milvuspb.HasPartitionRequest, opts ...grpc.CallOption) (*servicepb.BoolResponse, error) {
	out := new(servicepb.BoolResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.MasterService/HasPartition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) GetPartitionStatistics(ctx context.Context, in *milvuspb.PartitionStatsRequest, opts ...grpc.CallOption) (*milvuspb.PartitionStatsResponse, error) {
	out := new(milvuspb.PartitionStatsResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.MasterService/GetPartitionStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) ShowPartitions(ctx context.Context, in *milvuspb.ShowPartitionRequest, opts ...grpc.CallOption) (*milvuspb.ShowPartitionResponse, error) {
	out := new(milvuspb.ShowPartitionResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.MasterService/ShowPartitions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) CreateIndex(ctx context.Context, in *milvuspb.CreateIndexRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.MasterService/CreateIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) DescribeIndex(ctx context.Context, in *milvuspb.DescribeIndexRequest, opts ...grpc.CallOption) (*milvuspb.DescribeIndexResponse, error) {
	out := new(milvuspb.DescribeIndexResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.MasterService/DescribeIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) AllocTimestamp(ctx context.Context, in *TsoRequest, opts ...grpc.CallOption) (*TsoResponse, error) {
	out := new(TsoResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.MasterService/AllocTimestamp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) AllocID(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*IDResponse, error) {
	out := new(IDResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.MasterService/AllocID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServiceServer is the server API for MasterService service.
type MasterServiceServer interface {
	//*
	// @brief This method is used to create collection
	//
	// @param CreateCollectionRequest, use to provide collection information to be created.
	//
	// @return Status
	CreateCollection(context.Context, *milvuspb.CreateCollectionRequest) (*commonpb.Status, error)
	//*
	// @brief This method is used to delete collection.
	//
	// @param DropCollectionRequest, collection name is going to be deleted.
	//
	// @return Status
	DropCollection(context.Context, *milvuspb.DropCollectionRequest) (*commonpb.Status, error)
	//*
	// @brief This method is used to test collection existence.
	//
	// @param HasCollectionRequest, collection name is going to be tested.
	//
	// @return BoolResponse
	HasCollection(context.Context, *milvuspb.HasCollectionRequest) (*servicepb.BoolResponse, error)
	//*
	// @brief This method is used to get collection schema.
	//
	// @param DescribeCollectionRequest, target collection name.
	//
	// @return CollectionSchema
	DescribeCollection(context.Context, *milvuspb.DescribeCollectionRequest) (*milvuspb.DescribeCollectionResponse, error)
	GetCollectionStatistics(context.Context, *milvuspb.CollectionStatsRequest) (*milvuspb.CollectionStatsResponse, error)
	//*
	// @brief This method is used to list all collections.
	//
	// @return StringListResponse, collection name list
	ShowCollections(context.Context, *milvuspb.ShowCollectionRequest) (*milvuspb.ShowCollectionResponse, error)
	//*
	// @brief This method is used to create partition
	//
	// @return Status
	CreatePartition(context.Context, *milvuspb.CreatePartitionRequest) (*commonpb.Status, error)
	//*
	// @brief This method is used to drop partition
	//
	// @return Status
	DropPartition(context.Context, *milvuspb.DropPartitionRequest) (*commonpb.Status, error)
	//*
	// @brief This method is used to test partition existence.
	//
	// @return BoolResponse
	HasPartition(context.Context, *milvuspb.HasPartitionRequest) (*servicepb.BoolResponse, error)
	GetPartitionStatistics(context.Context, *milvuspb.PartitionStatsRequest) (*milvuspb.PartitionStatsResponse, error)
	//*
	// @brief This method is used to show partition information
	//
	// @param ShowPartitionRequest, target collection name.
	//
	// @return StringListResponse
	ShowPartitions(context.Context, *milvuspb.ShowPartitionRequest) (*milvuspb.ShowPartitionResponse, error)
	CreateIndex(context.Context, *milvuspb.CreateIndexRequest) (*commonpb.Status, error)
	DescribeIndex(context.Context, *milvuspb.DescribeIndexRequest) (*milvuspb.DescribeIndexResponse, error)
	AllocTimestamp(context.Context, *TsoRequest) (*TsoResponse, error)
	AllocID(context.Context, *IDRequest) (*IDResponse, error)
}

// UnimplementedMasterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMasterServiceServer struct {
}

func (*UnimplementedMasterServiceServer) CreateCollection(ctx context.Context, req *milvuspb.CreateCollectionRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCollection not implemented")
}
func (*UnimplementedMasterServiceServer) DropCollection(ctx context.Context, req *milvuspb.DropCollectionRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropCollection not implemented")
}
func (*UnimplementedMasterServiceServer) HasCollection(ctx context.Context, req *milvuspb.HasCollectionRequest) (*servicepb.BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasCollection not implemented")
}
func (*UnimplementedMasterServiceServer) DescribeCollection(ctx context.Context, req *milvuspb.DescribeCollectionRequest) (*milvuspb.DescribeCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeCollection not implemented")
}
func (*UnimplementedMasterServiceServer) GetCollectionStatistics(ctx context.Context, req *milvuspb.CollectionStatsRequest) (*milvuspb.CollectionStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollectionStatistics not implemented")
}
func (*UnimplementedMasterServiceServer) ShowCollections(ctx context.Context, req *milvuspb.ShowCollectionRequest) (*milvuspb.ShowCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowCollections not implemented")
}
func (*UnimplementedMasterServiceServer) CreatePartition(ctx context.Context, req *milvuspb.CreatePartitionRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePartition not implemented")
}
func (*UnimplementedMasterServiceServer) DropPartition(ctx context.Context, req *milvuspb.DropPartitionRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropPartition not implemented")
}
func (*UnimplementedMasterServiceServer) HasPartition(ctx context.Context, req *milvuspb.HasPartitionRequest) (*servicepb.BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasPartition not implemented")
}
func (*UnimplementedMasterServiceServer) GetPartitionStatistics(ctx context.Context, req *milvuspb.PartitionStatsRequest) (*milvuspb.PartitionStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartitionStatistics not implemented")
}
func (*UnimplementedMasterServiceServer) ShowPartitions(ctx context.Context, req *milvuspb.ShowPartitionRequest) (*milvuspb.ShowPartitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowPartitions not implemented")
}
func (*UnimplementedMasterServiceServer) CreateIndex(ctx context.Context, req *milvuspb.CreateIndexRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIndex not implemented")
}
func (*UnimplementedMasterServiceServer) DescribeIndex(ctx context.Context, req *milvuspb.DescribeIndexRequest) (*milvuspb.DescribeIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeIndex not implemented")
}
func (*UnimplementedMasterServiceServer) AllocTimestamp(ctx context.Context, req *TsoRequest) (*TsoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocTimestamp not implemented")
}
func (*UnimplementedMasterServiceServer) AllocID(ctx context.Context, req *IDRequest) (*IDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocID not implemented")
}

func RegisterMasterServiceServer(s *grpc.Server, srv MasterServiceServer) {
	s.RegisterService(&_MasterService_serviceDesc, srv)
}

func _MasterService_CreateCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(milvuspb.CreateCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).CreateCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.MasterService/CreateCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).CreateCollection(ctx, req.(*milvuspb.CreateCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_DropCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(milvuspb.DropCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).DropCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.MasterService/DropCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).DropCollection(ctx, req.(*milvuspb.DropCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_HasCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(milvuspb.HasCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).HasCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.MasterService/HasCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).HasCollection(ctx, req.(*milvuspb.HasCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_DescribeCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(milvuspb.DescribeCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).DescribeCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.MasterService/DescribeCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).DescribeCollection(ctx, req.(*milvuspb.DescribeCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_GetCollectionStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(milvuspb.CollectionStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).GetCollectionStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.MasterService/GetCollectionStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).GetCollectionStatistics(ctx, req.(*milvuspb.CollectionStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_ShowCollections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(milvuspb.ShowCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).ShowCollections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.MasterService/ShowCollections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).ShowCollections(ctx, req.(*milvuspb.ShowCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_CreatePartition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(milvuspb.CreatePartitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).CreatePartition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.MasterService/CreatePartition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).CreatePartition(ctx, req.(*milvuspb.CreatePartitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_DropPartition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(milvuspb.DropPartitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).DropPartition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.MasterService/DropPartition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).DropPartition(ctx, req.(*milvuspb.DropPartitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_HasPartition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(milvuspb.HasPartitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).HasPartition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.MasterService/HasPartition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).HasPartition(ctx, req.(*milvuspb.HasPartitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_GetPartitionStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(milvuspb.PartitionStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).GetPartitionStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.MasterService/GetPartitionStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).GetPartitionStatistics(ctx, req.(*milvuspb.PartitionStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_ShowPartitions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(milvuspb.ShowPartitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).ShowPartitions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.MasterService/ShowPartitions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).ShowPartitions(ctx, req.(*milvuspb.ShowPartitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_CreateIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(milvuspb.CreateIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).CreateIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.MasterService/CreateIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).CreateIndex(ctx, req.(*milvuspb.CreateIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_DescribeIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(milvuspb.DescribeIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).DescribeIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.MasterService/DescribeIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).DescribeIndex(ctx, req.(*milvuspb.DescribeIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_AllocTimestamp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TsoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).AllocTimestamp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.MasterService/AllocTimestamp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).AllocTimestamp(ctx, req.(*TsoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_AllocID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).AllocID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.MasterService/AllocID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).AllocID(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MasterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "milvus.proto.master.MasterService",
	HandlerType: (*MasterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCollection",
			Handler:    _MasterService_CreateCollection_Handler,
		},
		{
			MethodName: "DropCollection",
			Handler:    _MasterService_DropCollection_Handler,
		},
		{
			MethodName: "HasCollection",
			Handler:    _MasterService_HasCollection_Handler,
		},
		{
			MethodName: "DescribeCollection",
			Handler:    _MasterService_DescribeCollection_Handler,
		},
		{
			MethodName: "GetCollectionStatistics",
			Handler:    _MasterService_GetCollectionStatistics_Handler,
		},
		{
			MethodName: "ShowCollections",
			Handler:    _MasterService_ShowCollections_Handler,
		},
		{
			MethodName: "CreatePartition",
			Handler:    _MasterService_CreatePartition_Handler,
		},
		{
			MethodName: "DropPartition",
			Handler:    _MasterService_DropPartition_Handler,
		},
		{
			MethodName: "HasPartition",
			Handler:    _MasterService_HasPartition_Handler,
		},
		{
			MethodName: "GetPartitionStatistics",
			Handler:    _MasterService_GetPartitionStatistics_Handler,
		},
		{
			MethodName: "ShowPartitions",
			Handler:    _MasterService_ShowPartitions_Handler,
		},
		{
			MethodName: "CreateIndex",
			Handler:    _MasterService_CreateIndex_Handler,
		},
		{
			MethodName: "DescribeIndex",
			Handler:    _MasterService_DescribeIndex_Handler,
		},
		{
			MethodName: "AllocTimestamp",
			Handler:    _MasterService_AllocTimestamp_Handler,
		},
		{
			MethodName: "AllocID",
			Handler:    _MasterService_AllocID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master_service.proto",
}
