package datanode

import (
	"context"

	"github.com/zilliztech/milvus-distributed/internal/proto/datapb"
)

type Client struct {
	ctx context.Context
	// GOOSE TODO: add DataNodeClient
}

func (c *Client) Init() error {
	panic("implement me")
}

func (c *Client) Start() error {
	panic("implement me")
}

func (c *Client) Stop() error {
	panic("implement me")
}

func (c *Client) WatchDmChannels(datapb.WatchDmChannelRequest, error) {
	panic("implement me")
}

func (c *Client) FlushSegment() (datapb.FlushSegRequest, error) {
	panic("implement me")
}
