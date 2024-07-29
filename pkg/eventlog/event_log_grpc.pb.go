// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.4
// source: event_log.proto

package eventlog

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	EventLogService_Listen_FullMethodName = "/milvus.proto.eventlog.EventLogService/Listen"
)

// EventLogServiceClient is the client API for EventLogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventLogServiceClient interface {
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (EventLogService_ListenClient, error)
}

type eventLogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventLogServiceClient(cc grpc.ClientConnInterface) EventLogServiceClient {
	return &eventLogServiceClient{cc}
}

func (c *eventLogServiceClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (EventLogService_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventLogService_ServiceDesc.Streams[0], EventLogService_Listen_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &eventLogServiceListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventLogService_ListenClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventLogServiceListenClient struct {
	grpc.ClientStream
}

func (x *eventLogServiceListenClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventLogServiceServer is the server API for EventLogService service.
// All implementations should embed UnimplementedEventLogServiceServer
// for forward compatibility
type EventLogServiceServer interface {
	Listen(*ListenRequest, EventLogService_ListenServer) error
}

// UnimplementedEventLogServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEventLogServiceServer struct {
}

func (UnimplementedEventLogServiceServer) Listen(*ListenRequest, EventLogService_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}

// UnsafeEventLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventLogServiceServer will
// result in compilation errors.
type UnsafeEventLogServiceServer interface {
	mustEmbedUnimplementedEventLogServiceServer()
}

func RegisterEventLogServiceServer(s grpc.ServiceRegistrar, srv EventLogServiceServer) {
	s.RegisterService(&EventLogService_ServiceDesc, srv)
}

func _EventLogService_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventLogServiceServer).Listen(m, &eventLogServiceListenServer{stream})
}

type EventLogService_ListenServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type eventLogServiceListenServer struct {
	grpc.ServerStream
}

func (x *eventLogServiceListenServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

// EventLogService_ServiceDesc is the grpc.ServiceDesc for EventLogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventLogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "milvus.proto.eventlog.EventLogService",
	HandlerType: (*EventLogServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _EventLogService_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "event_log.proto",
}
