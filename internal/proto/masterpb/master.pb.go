// Code generated by protoc-gen-go. DO NOT EDIT.
// source: master.proto

package masterpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	commonpb "github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	internalpb "github.com/zilliztech/milvus-distributed/internal/proto/internalpb"
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

func init() { proto.RegisterFile("master.proto", fileDescriptor_f9c348dec43a6705) }

var fileDescriptor_f9c348dec43a6705 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x51, 0xaf, 0xd2, 0x30,
	0x14, 0xc7, 0x79, 0xba, 0xc6, 0x86, 0xcb, 0xf5, 0xd6, 0x37, 0x7c, 0xf1, 0xee, 0xc9, 0x80, 0x6c,
	0x46, 0xbf, 0x80, 0xc2, 0x1e, 0x20, 0xd1, 0x84, 0x00, 0x2f, 0x6a, 0x0c, 0x76, 0xa3, 0x81, 0xc6,
	0x6e, 0x9d, 0x3d, 0x67, 0x98, 0xf0, 0xe1, 0xfc, 0x6c, 0x66, 0x1b, 0xdd, 0x56, 0xa1, 0x88, 0xf7,
	0x8d, 0xb6, 0xff, 0xf3, 0xfb, 0x73, 0xce, 0xf9, 0x67, 0xa4, 0x9b, 0x30, 0x40, 0xae, 0xfd, 0x4c,
	0x2b, 0x54, 0xf4, 0x79, 0x22, 0xe4, 0x3e, 0x87, 0xea, 0xe4, 0x57, 0x4f, 0xfd, 0x6e, 0xac, 0x92,
	0x44, 0xa5, 0xd5, 0x65, 0x9f, 0x8a, 0x14, 0xb9, 0x4e, 0x99, 0x5c, 0x27, 0xb0, 0x3d, 0xde, 0xdd,
	0x03, 0xd7, 0x7b, 0x11, 0xf3, 0xe6, 0xea, 0xed, 0xef, 0xa7, 0xe4, 0xe6, 0x53, 0x59, 0x4f, 0x19,
	0x79, 0x36, 0xd1, 0x9c, 0x21, 0x9f, 0x28, 0x29, 0x79, 0x8c, 0x42, 0xa5, 0xd4, 0xf7, 0x2d, 0x27,
	0xc3, 0xf4, 0xff, 0x16, 0x2e, 0xf8, 0xcf, 0x9c, 0x03, 0xf6, 0x5f, 0xd8, 0xfa, 0xe3, 0x3f, 0x5a,
	0x22, 0xc3, 0x1c, 0xbc, 0x0e, 0xfd, 0x46, 0x7a, 0xa1, 0x56, 0x59, 0xcb, 0xe0, 0xb5, 0xc3, 0xc0,
	0x96, 0x5d, 0x89, 0x8f, 0xc8, 0xed, 0x94, 0x41, 0x8b, 0x3e, 0x74, 0xd0, 0x2d, 0x95, 0x81, 0x7b,
	0xb6, 0xf8, 0x38, 0x2b, 0x7f, 0xac, 0x94, 0x5c, 0x70, 0xc8, 0x54, 0x0a, 0xdc, 0xeb, 0xd0, 0x9c,
	0xd0, 0x90, 0x43, 0xac, 0x45, 0xd4, 0x9e, 0xd3, 0x1b, 0x57, 0x1b, 0x27, 0x52, 0xe3, 0x36, 0x3c,
	0xef, 0xd6, 0x08, 0xab, 0xd2, 0xac, 0xf8, 0xe9, 0x75, 0xe8, 0x0f, 0x72, 0xb7, 0xdc, 0xa9, 0x5f,
	0xcd, 0x33, 0x38, 0x47, 0x67, 0xeb, 0x8c, 0xdf, 0xab, 0xf3, 0x7e, 0x4b, 0xd4, 0x22, 0xdd, 0x7e,
	0x14, 0x80, 0xad, 0x1e, 0xd7, 0xe4, 0xae, 0x5a, 0xf0, 0x9c, 0x69, 0x14, 0x65, 0x83, 0xa3, 0x8b,
	0x41, 0xa8, 0x75, 0x57, 0x2e, 0xea, 0x2b, 0xb9, 0x2d, 0x16, 0xdc, 0xe0, 0x87, 0x17, 0x62, 0xf0,
	0xbf, 0xf0, 0xef, 0xa4, 0x3b, 0x65, 0xd0, 0xb0, 0x07, 0xee, 0x10, 0x9c, 0xa0, 0xaf, 0xcb, 0x80,
	0x26, 0xf7, 0x66, 0xb1, 0x8d, 0x4d, 0xf0, 0x8f, 0x08, 0x9c, 0x78, 0x0d, 0xce, 0x7b, 0xd5, 0x3a,
	0x3b, 0x00, 0x82, 0xf4, 0x8a, 0xc5, 0xd6, 0xaf, 0xe0, 0x9c, 0x99, 0x25, 0x7b, 0xcc, 0xfa, 0x3f,
	0x93, 0xde, 0x07, 0x29, 0x55, 0xbc, 0x12, 0x09, 0x07, 0x64, 0x49, 0x46, 0x1f, 0x1c, 0x56, 0x2b,
	0x50, 0x8e, 0xc9, 0xd9, 0x92, 0x1a, 0x3d, 0x27, 0x4f, 0x4a, 0xf4, 0x2c, 0xa4, 0x2f, 0x1d, 0x05,
	0xb3, 0xd0, 0x20, 0x1f, 0x2e, 0x28, 0x0c, 0x71, 0x3c, 0xfe, 0xf2, 0x7e, 0x2b, 0x70, 0x97, 0x47,
	0x45, 0x0e, 0x82, 0x83, 0x90, 0x52, 0x1c, 0x90, 0xc7, 0xbb, 0xa0, 0xaa, 0x1d, 0x6d, 0x04, 0xa0,
	0x16, 0x51, 0x8e, 0x7c, 0x13, 0x18, 0x42, 0x50, 0x02, 0x83, 0xea, 0xbb, 0x99, 0x45, 0xd1, 0x4d,
	0x79, 0x7e, 0xf7, 0x27, 0x00, 0x00, 0xff, 0xff, 0xba, 0x9e, 0x0e, 0x5d, 0x65, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MasterClient is the client API for Master service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MasterClient interface {
	//*
	// @brief This method is used to create collection
	//
	// @param CreateCollectionRequest, use to provide collection information to be created.
	//
	// @return Status
	CreateCollection(ctx context.Context, in *internalpb.CreateCollectionRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	//*
	// @brief This method is used to delete collection.
	//
	// @param DropCollectionRequest, collection name is going to be deleted.
	//
	// @return Status
	DropCollection(ctx context.Context, in *internalpb.DropCollectionRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	//*
	// @brief This method is used to test collection existence.
	//
	// @param HasCollectionRequest, collection name is going to be tested.
	//
	// @return BoolResponse
	HasCollection(ctx context.Context, in *internalpb.HasCollectionRequest, opts ...grpc.CallOption) (*servicepb.BoolResponse, error)
	//*
	// @brief This method is used to get collection schema.
	//
	// @param DescribeCollectionRequest, target collection name.
	//
	// @return CollectionSchema
	DescribeCollection(ctx context.Context, in *internalpb.DescribeCollectionRequest, opts ...grpc.CallOption) (*servicepb.CollectionDescription, error)
	//*
	// @brief This method is used to list all collections.
	//
	// @return StringListResponse, collection name list
	ShowCollections(ctx context.Context, in *internalpb.ShowCollectionRequest, opts ...grpc.CallOption) (*servicepb.StringListResponse, error)
	//*
	// @brief This method is used to create partition
	//
	// @return Status
	CreatePartition(ctx context.Context, in *internalpb.CreatePartitionRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	//*
	// @brief This method is used to drop partition
	//
	// @return Status
	DropPartition(ctx context.Context, in *internalpb.DropPartitionRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	//*
	// @brief This method is used to test partition existence.
	//
	// @return BoolResponse
	HasPartition(ctx context.Context, in *internalpb.HasPartitionRequest, opts ...grpc.CallOption) (*servicepb.BoolResponse, error)
	//*
	// @brief This method is used to get basic partition infomation.
	//
	// @return PartitionDescription
	DescribePartition(ctx context.Context, in *internalpb.DescribePartitionRequest, opts ...grpc.CallOption) (*servicepb.PartitionDescription, error)
	//*
	// @brief This method is used to show partition information
	//
	// @param ShowPartitionRequest, target collection name.
	//
	// @return StringListResponse
	ShowPartitions(ctx context.Context, in *internalpb.ShowPartitionRequest, opts ...grpc.CallOption) (*servicepb.StringListResponse, error)
	AllocTimestamp(ctx context.Context, in *internalpb.TsoRequest, opts ...grpc.CallOption) (*internalpb.TsoResponse, error)
	AllocID(ctx context.Context, in *internalpb.IDRequest, opts ...grpc.CallOption) (*internalpb.IDResponse, error)
}

type masterClient struct {
	cc *grpc.ClientConn
}

func NewMasterClient(cc *grpc.ClientConn) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) CreateCollection(ctx context.Context, in *internalpb.CreateCollectionRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/CreateCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) DropCollection(ctx context.Context, in *internalpb.DropCollectionRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/DropCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) HasCollection(ctx context.Context, in *internalpb.HasCollectionRequest, opts ...grpc.CallOption) (*servicepb.BoolResponse, error) {
	out := new(servicepb.BoolResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/HasCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) DescribeCollection(ctx context.Context, in *internalpb.DescribeCollectionRequest, opts ...grpc.CallOption) (*servicepb.CollectionDescription, error) {
	out := new(servicepb.CollectionDescription)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/DescribeCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ShowCollections(ctx context.Context, in *internalpb.ShowCollectionRequest, opts ...grpc.CallOption) (*servicepb.StringListResponse, error) {
	out := new(servicepb.StringListResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/ShowCollections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) CreatePartition(ctx context.Context, in *internalpb.CreatePartitionRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/CreatePartition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) DropPartition(ctx context.Context, in *internalpb.DropPartitionRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/DropPartition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) HasPartition(ctx context.Context, in *internalpb.HasPartitionRequest, opts ...grpc.CallOption) (*servicepb.BoolResponse, error) {
	out := new(servicepb.BoolResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/HasPartition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) DescribePartition(ctx context.Context, in *internalpb.DescribePartitionRequest, opts ...grpc.CallOption) (*servicepb.PartitionDescription, error) {
	out := new(servicepb.PartitionDescription)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/DescribePartition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ShowPartitions(ctx context.Context, in *internalpb.ShowPartitionRequest, opts ...grpc.CallOption) (*servicepb.StringListResponse, error) {
	out := new(servicepb.StringListResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/ShowPartitions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) AllocTimestamp(ctx context.Context, in *internalpb.TsoRequest, opts ...grpc.CallOption) (*internalpb.TsoResponse, error) {
	out := new(internalpb.TsoResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/AllocTimestamp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) AllocID(ctx context.Context, in *internalpb.IDRequest, opts ...grpc.CallOption) (*internalpb.IDResponse, error) {
	out := new(internalpb.IDResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.master.Master/AllocID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServer is the server API for Master service.
type MasterServer interface {
	//*
	// @brief This method is used to create collection
	//
	// @param CreateCollectionRequest, use to provide collection information to be created.
	//
	// @return Status
	CreateCollection(context.Context, *internalpb.CreateCollectionRequest) (*commonpb.Status, error)
	//*
	// @brief This method is used to delete collection.
	//
	// @param DropCollectionRequest, collection name is going to be deleted.
	//
	// @return Status
	DropCollection(context.Context, *internalpb.DropCollectionRequest) (*commonpb.Status, error)
	//*
	// @brief This method is used to test collection existence.
	//
	// @param HasCollectionRequest, collection name is going to be tested.
	//
	// @return BoolResponse
	HasCollection(context.Context, *internalpb.HasCollectionRequest) (*servicepb.BoolResponse, error)
	//*
	// @brief This method is used to get collection schema.
	//
	// @param DescribeCollectionRequest, target collection name.
	//
	// @return CollectionSchema
	DescribeCollection(context.Context, *internalpb.DescribeCollectionRequest) (*servicepb.CollectionDescription, error)
	//*
	// @brief This method is used to list all collections.
	//
	// @return StringListResponse, collection name list
	ShowCollections(context.Context, *internalpb.ShowCollectionRequest) (*servicepb.StringListResponse, error)
	//*
	// @brief This method is used to create partition
	//
	// @return Status
	CreatePartition(context.Context, *internalpb.CreatePartitionRequest) (*commonpb.Status, error)
	//*
	// @brief This method is used to drop partition
	//
	// @return Status
	DropPartition(context.Context, *internalpb.DropPartitionRequest) (*commonpb.Status, error)
	//*
	// @brief This method is used to test partition existence.
	//
	// @return BoolResponse
	HasPartition(context.Context, *internalpb.HasPartitionRequest) (*servicepb.BoolResponse, error)
	//*
	// @brief This method is used to get basic partition infomation.
	//
	// @return PartitionDescription
	DescribePartition(context.Context, *internalpb.DescribePartitionRequest) (*servicepb.PartitionDescription, error)
	//*
	// @brief This method is used to show partition information
	//
	// @param ShowPartitionRequest, target collection name.
	//
	// @return StringListResponse
	ShowPartitions(context.Context, *internalpb.ShowPartitionRequest) (*servicepb.StringListResponse, error)
	AllocTimestamp(context.Context, *internalpb.TsoRequest) (*internalpb.TsoResponse, error)
	AllocID(context.Context, *internalpb.IDRequest) (*internalpb.IDResponse, error)
}

// UnimplementedMasterServer can be embedded to have forward compatible implementations.
type UnimplementedMasterServer struct {
}

func (*UnimplementedMasterServer) CreateCollection(ctx context.Context, req *internalpb.CreateCollectionRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCollection not implemented")
}
func (*UnimplementedMasterServer) DropCollection(ctx context.Context, req *internalpb.DropCollectionRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropCollection not implemented")
}
func (*UnimplementedMasterServer) HasCollection(ctx context.Context, req *internalpb.HasCollectionRequest) (*servicepb.BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasCollection not implemented")
}
func (*UnimplementedMasterServer) DescribeCollection(ctx context.Context, req *internalpb.DescribeCollectionRequest) (*servicepb.CollectionDescription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeCollection not implemented")
}
func (*UnimplementedMasterServer) ShowCollections(ctx context.Context, req *internalpb.ShowCollectionRequest) (*servicepb.StringListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowCollections not implemented")
}
func (*UnimplementedMasterServer) CreatePartition(ctx context.Context, req *internalpb.CreatePartitionRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePartition not implemented")
}
func (*UnimplementedMasterServer) DropPartition(ctx context.Context, req *internalpb.DropPartitionRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropPartition not implemented")
}
func (*UnimplementedMasterServer) HasPartition(ctx context.Context, req *internalpb.HasPartitionRequest) (*servicepb.BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasPartition not implemented")
}
func (*UnimplementedMasterServer) DescribePartition(ctx context.Context, req *internalpb.DescribePartitionRequest) (*servicepb.PartitionDescription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePartition not implemented")
}
func (*UnimplementedMasterServer) ShowPartitions(ctx context.Context, req *internalpb.ShowPartitionRequest) (*servicepb.StringListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowPartitions not implemented")
}
func (*UnimplementedMasterServer) AllocTimestamp(ctx context.Context, req *internalpb.TsoRequest) (*internalpb.TsoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocTimestamp not implemented")
}
func (*UnimplementedMasterServer) AllocID(ctx context.Context, req *internalpb.IDRequest) (*internalpb.IDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocID not implemented")
}

func RegisterMasterServer(s *grpc.Server, srv MasterServer) {
	s.RegisterService(&_Master_serviceDesc, srv)
}

func _Master_CreateCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.CreateCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).CreateCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/CreateCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).CreateCollection(ctx, req.(*internalpb.CreateCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_DropCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.DropCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).DropCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/DropCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).DropCollection(ctx, req.(*internalpb.DropCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_HasCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.HasCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).HasCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/HasCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).HasCollection(ctx, req.(*internalpb.HasCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_DescribeCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.DescribeCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).DescribeCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/DescribeCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).DescribeCollection(ctx, req.(*internalpb.DescribeCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ShowCollections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.ShowCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ShowCollections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/ShowCollections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ShowCollections(ctx, req.(*internalpb.ShowCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_CreatePartition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.CreatePartitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).CreatePartition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/CreatePartition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).CreatePartition(ctx, req.(*internalpb.CreatePartitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_DropPartition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.DropPartitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).DropPartition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/DropPartition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).DropPartition(ctx, req.(*internalpb.DropPartitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_HasPartition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.HasPartitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).HasPartition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/HasPartition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).HasPartition(ctx, req.(*internalpb.HasPartitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_DescribePartition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.DescribePartitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).DescribePartition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/DescribePartition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).DescribePartition(ctx, req.(*internalpb.DescribePartitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ShowPartitions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.ShowPartitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ShowPartitions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/ShowPartitions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ShowPartitions(ctx, req.(*internalpb.ShowPartitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_AllocTimestamp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.TsoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).AllocTimestamp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/AllocTimestamp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).AllocTimestamp(ctx, req.(*internalpb.TsoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_AllocID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).AllocID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.master.Master/AllocID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).AllocID(ctx, req.(*internalpb.IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Master_serviceDesc = grpc.ServiceDesc{
	ServiceName: "milvus.proto.master.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCollection",
			Handler:    _Master_CreateCollection_Handler,
		},
		{
			MethodName: "DropCollection",
			Handler:    _Master_DropCollection_Handler,
		},
		{
			MethodName: "HasCollection",
			Handler:    _Master_HasCollection_Handler,
		},
		{
			MethodName: "DescribeCollection",
			Handler:    _Master_DescribeCollection_Handler,
		},
		{
			MethodName: "ShowCollections",
			Handler:    _Master_ShowCollections_Handler,
		},
		{
			MethodName: "CreatePartition",
			Handler:    _Master_CreatePartition_Handler,
		},
		{
			MethodName: "DropPartition",
			Handler:    _Master_DropPartition_Handler,
		},
		{
			MethodName: "HasPartition",
			Handler:    _Master_HasPartition_Handler,
		},
		{
			MethodName: "DescribePartition",
			Handler:    _Master_DescribePartition_Handler,
		},
		{
			MethodName: "ShowPartitions",
			Handler:    _Master_ShowPartitions_Handler,
		},
		{
			MethodName: "AllocTimestamp",
			Handler:    _Master_AllocTimestamp_Handler,
		},
		{
			MethodName: "AllocID",
			Handler:    _Master_AllocID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master.proto",
}
