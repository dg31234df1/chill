// Copyright (C) 2019-2020 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License.
package dataservice

import (
	"sync"

	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/proto/commonpb"
	"github.com/milvus-io/milvus/internal/proto/datapb"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type cluster struct {
	mu             sync.RWMutex
	ctx            context.Context
	dataManager    *clusterNodeManager
	sessionManager sessionManager
	posProvider    positionProvider

	startupPolicy    clusterStartupPolicy
	registerPolicy   dataNodeRegisterPolicy
	unregisterPolicy dataNodeUnregisterPolicy
	assginPolicy     channelAssignPolicy
}

type clusterOption struct {
	apply func(c *cluster)
}

func withStartupPolicy(p clusterStartupPolicy) clusterOption {
	return clusterOption{
		apply: func(c *cluster) { c.startupPolicy = p },
	}
}

func withRegisterPolicy(p dataNodeRegisterPolicy) clusterOption {
	return clusterOption{
		apply: func(c *cluster) { c.registerPolicy = p },
	}
}

func withUnregistorPolicy(p dataNodeUnregisterPolicy) clusterOption {
	return clusterOption{
		apply: func(c *cluster) { c.unregisterPolicy = p },
	}
}

func withAssignPolicy(p channelAssignPolicy) clusterOption {
	return clusterOption{
		apply: func(c *cluster) { c.assginPolicy = p },
	}
}

func defaultStartupPolicy() clusterStartupPolicy {
	return newWatchRestartsStartupPolicy()
}

func defaultRegisterPolicy() dataNodeRegisterPolicy {
	return newDoNothingRegisterPolicy()
}

func defaultUnregisterPolicy() dataNodeUnregisterPolicy {
	return newDoNothingUnregisterPolicy()
}

func defaultAssignPolicy() channelAssignPolicy {
	return newAssignAllPolicy()
}

func newCluster(ctx context.Context, dataManager *clusterNodeManager, sessionManager sessionManager, posProvider positionProvider, opts ...clusterOption) *cluster {
	c := &cluster{
		ctx:              ctx,
		sessionManager:   sessionManager,
		dataManager:      dataManager,
		posProvider:      posProvider,
		startupPolicy:    defaultStartupPolicy(),
		registerPolicy:   defaultRegisterPolicy(),
		unregisterPolicy: defaultUnregisterPolicy(),
		assginPolicy:     defaultAssignPolicy(),
	}
	for _, opt := range opts {
		opt.apply(c)
	}

	return c
}

func (c *cluster) startup(dataNodes []*datapb.DataNodeInfo) error {
	deltaChange := c.dataManager.updateCluster(dataNodes)
	nodes := c.dataManager.getDataNodes(false)
	rets := c.startupPolicy.apply(nodes, deltaChange)
	c.dataManager.updateDataNodes(rets)
	rets = c.watch(rets)
	c.dataManager.updateDataNodes(rets)
	return nil
}

func (c *cluster) watch(nodes []*datapb.DataNodeInfo) []*datapb.DataNodeInfo {
	for _, n := range nodes {
		uncompletes := make([]vchannel, 0, len(nodes))
		for _, ch := range n.Channels {
			if ch.State == datapb.ChannelWatchState_Uncomplete {
				uncompletes = append(uncompletes, vchannel{
					CollectionID: ch.CollectionID,
					DmlChannel:   ch.Name,
					DdlChannel:   c.posProvider.GetDdlChannel(),
				})
			}
		}
		vchanInfos, err := c.posProvider.GetVChanPositions(uncompletes)
		if err != nil {
			log.Warn("get vchannel position failed", zap.Error(err))
			continue
		}
		cli, err := c.sessionManager.getOrCreateSession(n.Address)
		if err != nil {
			log.Warn("get session failed", zap.String("addr", n.Address), zap.Error(err))
			continue
		}
		req := &datapb.WatchDmChannelsRequest{
			Base: &commonpb.MsgBase{
				SourceID: Params.NodeID,
			},
			Vchannels: vchanInfos,
		}
		resp, err := cli.WatchDmChannels(c.ctx, req)
		if err != nil {
			log.Warn("watch dm channel failed", zap.String("addr", n.Address), zap.Error(err))
			continue
		}
		if resp.ErrorCode != commonpb.ErrorCode_Success {
			log.Warn("watch channels failed", zap.String("address", n.Address), zap.Error(err))
			continue
		}
		for _, ch := range n.Channels {
			if ch.State == datapb.ChannelWatchState_Uncomplete {
				ch.State = datapb.ChannelWatchState_Complete
			}
		}
	}
	return nodes
}

func (c *cluster) register(n *datapb.DataNodeInfo) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.dataManager.register(n)
	cNodes := c.dataManager.getDataNodes(true)
	rets := c.registerPolicy.apply(cNodes, n)
	c.dataManager.updateDataNodes(rets)
	rets = c.watch(rets)
	c.dataManager.updateDataNodes(rets)
}

func (c *cluster) unregister(n *datapb.DataNodeInfo) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.sessionManager.releaseSession(n.Address)
	c.dataManager.unregister(n)
	cNodes := c.dataManager.getDataNodes(true)
	rets := c.unregisterPolicy.apply(cNodes, n)
	c.dataManager.updateDataNodes(rets)
	rets = c.watch(rets)
	c.dataManager.updateDataNodes(rets)
}

func (c *cluster) watchIfNeeded(channel string, collectionID UniqueID) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cNodes := c.dataManager.getDataNodes(true)
	rets := c.assginPolicy.apply(cNodes, channel, collectionID)
	c.dataManager.updateDataNodes(rets)
	rets = c.watch(rets)
	c.dataManager.updateDataNodes(rets)
}

func (c *cluster) flush(segments []*datapb.SegmentInfo) {
	log.Debug("prepare to flush", zap.Any("segments", segments))
	c.mu.Lock()
	defer c.mu.Unlock()

	m := make(map[string]map[UniqueID][]UniqueID) // channel-> map[collectionID]segmentIDs

	for _, seg := range segments {
		if _, ok := m[seg.InsertChannel]; !ok {
			m[seg.InsertChannel] = make(map[UniqueID][]UniqueID)
		}

		m[seg.InsertChannel][seg.CollectionID] = append(m[seg.InsertChannel][seg.CollectionID], seg.ID)
	}

	dataNodes := c.dataManager.getDataNodes(true)

	channel2Node := make(map[string]string)
	for _, node := range dataNodes {
		for _, chstatus := range node.Channels {
			channel2Node[chstatus.Name] = node.Address
		}
	}

	for ch, coll2seg := range m {
		node, ok := channel2Node[ch]
		if !ok {
			continue
		}
		cli, err := c.sessionManager.getOrCreateSession(node)
		if err != nil {
			log.Warn("get session failed", zap.String("addr", node), zap.Error(err))
			continue
		}
		for coll, segs := range coll2seg {
			req := &datapb.FlushSegmentsRequest{
				Base: &commonpb.MsgBase{
					MsgType:  commonpb.MsgType_Flush,
					SourceID: Params.NodeID,
				},
				CollectionID: coll,
				SegmentIDs:   segs,
			}
			resp, err := cli.FlushSegments(c.ctx, req)
			if err != nil {
				log.Warn("flush segment failed", zap.String("addr", node), zap.Error(err))
				continue
			}
			if resp.ErrorCode != commonpb.ErrorCode_Success {
				log.Warn("flush segment failed", zap.String("dataNode", node), zap.Error(err))
				continue
			}
			log.Debug("flush segments succeed", zap.Any("segmentIDs", segs))
		}
	}
}

func (c *cluster) releaseSessions() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.sessionManager.release()
}
