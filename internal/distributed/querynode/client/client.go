package grpcquerynodeclient

import (
	"context"
	"time"

	"google.golang.org/grpc"

	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/milvuspb"
	"github.com/zilliztech/milvus-distributed/internal/proto/querypb"
)

const (
	RPCConnectionTimeout = 30 * time.Second
	Retry                = 3
)

type Client struct {
	ctx        context.Context
	grpcClient querypb.QueryNodeClient
	conn       *grpc.ClientConn
	addr       string
}

func NewClient(address string) *Client {
	return &Client{
		addr: address,
	}
}

func (c *Client) Init() error {
	tracer := opentracing.GlobalTracer()
	ctx, cancel := context.WithTimeout(context.Background(), RPCConnectionTimeout)
	defer cancel()
	var err error
	for i := 0; i < Retry*100; i++ {
		if c.conn, err = grpc.DialContext(ctx, c.addr, grpc.WithInsecure(), grpc.WithBlock(),
			grpc.WithUnaryInterceptor(
				otgrpc.OpenTracingClientInterceptor(tracer)),
			grpc.WithStreamInterceptor(
				otgrpc.OpenTracingStreamClientInterceptor(tracer))); err == nil {
			break
		}
	}
	if err != nil {
		return err
	}
	c.grpcClient = querypb.NewQueryNodeClient(c.conn)
	return nil
}

func (c *Client) Start() error {
	return nil
}

func (c *Client) Stop() error {
	return c.conn.Close()
}

func (c *Client) GetComponentStates(ctx context.Context) (*internalpb.ComponentStates, error) {
	return c.grpcClient.GetComponentStates(ctx, &internalpb.GetComponentStatesRequest{})
}

func (c *Client) GetTimeTickChannel(ctx context.Context) (*milvuspb.StringResponse, error) {
	return c.grpcClient.GetTimeTickChannel(ctx, &internalpb.GetTimeTickChannelRequest{})
}

func (c *Client) GetStatisticsChannel(ctx context.Context) (*milvuspb.StringResponse, error) {
	return c.grpcClient.GetStatisticsChannel(ctx, &internalpb.GetStatisticsChannelRequest{})
}

func (c *Client) AddQueryChannel(ctx context.Context, req *querypb.AddQueryChannelRequest) (*commonpb.Status, error) {
	return c.grpcClient.AddQueryChannel(ctx, req)
}

func (c *Client) RemoveQueryChannel(ctx context.Context, req *querypb.RemoveQueryChannelRequest) (*commonpb.Status, error) {
	return c.grpcClient.RemoveQueryChannel(ctx, req)
}

func (c *Client) WatchDmChannels(ctx context.Context, req *querypb.WatchDmChannelsRequest) (*commonpb.Status, error) {
	return c.grpcClient.WatchDmChannels(ctx, req)
}

func (c *Client) LoadSegments(ctx context.Context, req *querypb.LoadSegmentsRequest) (*commonpb.Status, error) {
	return c.grpcClient.LoadSegments(ctx, req)
}

func (c *Client) ReleaseCollection(ctx context.Context, req *querypb.ReleaseCollectionRequest) (*commonpb.Status, error) {
	return c.grpcClient.ReleaseCollection(ctx, req)
}

func (c *Client) ReleasePartitions(ctx context.Context, req *querypb.ReleasePartitionsRequest) (*commonpb.Status, error) {
	return c.grpcClient.ReleasePartitions(ctx, req)
}

func (c *Client) ReleaseSegments(ctx context.Context, req *querypb.ReleaseSegmentsRequest) (*commonpb.Status, error) {
	return c.grpcClient.ReleaseSegments(ctx, req)
}

func (c *Client) GetSegmentInfo(ctx context.Context, req *querypb.GetSegmentInfoRequest) (*querypb.GetSegmentInfoResponse, error) {
	return c.grpcClient.GetSegmentInfo(ctx, req)
}
