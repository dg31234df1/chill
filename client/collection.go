// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"context"

	"github.com/cockroachdb/errors"
	"google.golang.org/grpc"

	"github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"
	"github.com/milvus-io/milvus/client/v2/entity"
	"github.com/milvus-io/milvus/pkg/util/merr"
)

// CreateCollection is the API for create a collection in Milvus.
func (c *Client) CreateCollection(ctx context.Context, option CreateCollectionOption, callOptions ...grpc.CallOption) error {
	req := option.Request()

	err := c.callService(func(milvusService milvuspb.MilvusServiceClient) error {
		resp, err := milvusService.CreateCollection(ctx, req, callOptions...)
		return merr.CheckRPCCall(resp, err)
	})
	if err != nil {
		return err
	}

	indexes := option.Indexes()
	for _, indexOption := range indexes {
		task, err := c.CreateIndex(ctx, indexOption, callOptions...)
		if err != nil {
			return err
		}
		err = task.Await(ctx)
		if err != nil {
			return nil
		}
	}

	if option.IsFast() {
		task, err := c.LoadCollection(ctx, NewLoadCollectionOption(req.GetCollectionName()))
		if err != nil {
			return err
		}
		return task.Await(ctx)
	}

	return nil
}

type ListCollectionOption interface {
	Request() *milvuspb.ShowCollectionsRequest
}

func (c *Client) ListCollections(ctx context.Context, option ListCollectionOption, callOptions ...grpc.CallOption) (collectionNames []string, err error) {
	req := option.Request()
	err = c.callService(func(milvusService milvuspb.MilvusServiceClient) error {
		resp, err := milvusService.ShowCollections(ctx, req, callOptions...)
		err = merr.CheckRPCCall(resp, err)
		if err != nil {
			return err
		}

		collectionNames = resp.GetCollectionNames()
		return nil
	})

	return collectionNames, err
}

func (c *Client) DescribeCollection(ctx context.Context, option *describeCollectionOption, callOptions ...grpc.CallOption) (collection *entity.Collection, err error) {
	req := option.Request()
	err = c.callService(func(milvusService milvuspb.MilvusServiceClient) error {
		resp, err := milvusService.DescribeCollection(ctx, req, callOptions...)
		err = merr.CheckRPCCall(resp, err)
		if err != nil {
			return err
		}

		collection = &entity.Collection{
			ID:               resp.GetCollectionID(),
			Schema:           entity.NewSchema().ReadProto(resp.GetSchema()),
			PhysicalChannels: resp.GetPhysicalChannelNames(),
			VirtualChannels:  resp.GetVirtualChannelNames(),
			ConsistencyLevel: entity.ConsistencyLevel(resp.ConsistencyLevel),
			ShardNum:         resp.GetShardsNum(),
			Properties:       entity.KvPairsMap(resp.GetProperties()),
		}
		collection.Name = collection.Schema.CollectionName
		return nil
	})

	return collection, err
}

func (c *Client) HasCollection(ctx context.Context, option HasCollectionOption, callOptions ...grpc.CallOption) (has bool, err error) {
	req := option.Request()
	err = c.callService(func(milvusService milvuspb.MilvusServiceClient) error {
		resp, err := milvusService.DescribeCollection(ctx, req, callOptions...)
		err = merr.CheckRPCCall(resp, err)
		if err != nil {
			// ErrCollectionNotFound for collection not exist
			if errors.Is(err, merr.ErrCollectionNotFound) {
				return nil
			}
			return err
		}
		has = true
		return nil
	})
	return has, err
}

func (c *Client) DropCollection(ctx context.Context, option DropCollectionOption, callOptions ...grpc.CallOption) error {
	req := option.Request()
	err := c.callService(func(milvusService milvuspb.MilvusServiceClient) error {
		resp, err := milvusService.DropCollection(ctx, req, callOptions...)
		return merr.CheckRPCCall(resp, err)
	})
	return err
}
