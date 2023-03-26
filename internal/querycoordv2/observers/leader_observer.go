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

package observers

import (
	"context"
	"sync"
	"time"

	"github.com/milvus-io/milvus-proto/go-api/commonpb"
	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/proto/querypb"
	"github.com/milvus-io/milvus/internal/querycoordv2/meta"
	"github.com/milvus-io/milvus/internal/querycoordv2/session"
	"github.com/milvus-io/milvus/internal/querycoordv2/utils"
	"github.com/milvus-io/milvus/internal/util/commonpbutil"
	"go.uber.org/zap"
)

const (
	interval   = 1 * time.Second
	RPCTimeout = 3 * time.Second
)

// LeaderObserver is to sync the distribution with leader
type LeaderObserver struct {
	wg      sync.WaitGroup
	closeCh chan struct{}
	dist    *meta.DistributionManager
	meta    *meta.Meta
	target  *meta.TargetManager
	broker  meta.Broker
	cluster session.Cluster

	stopOnce sync.Once
}

func (o *LeaderObserver) Start(ctx context.Context) {
	o.wg.Add(1)
	go func() {
		defer o.wg.Done()
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-o.closeCh:
				log.Info("stop leader observer")
				return
			case <-ctx.Done():
				log.Info("stop leader observer due to ctx done")
				return
			case <-ticker.C:
				o.observe(ctx)
			}
		}
	}()
}

func (o *LeaderObserver) Stop() {
	o.stopOnce.Do(func() {
		close(o.closeCh)
		o.wg.Wait()
	})
}

func (o *LeaderObserver) observe(ctx context.Context) {
	o.observeSegmentsDist(ctx)
}

func (o *LeaderObserver) observeSegmentsDist(ctx context.Context) {
	collectionIDs := o.meta.CollectionManager.GetAll()
	for _, cid := range collectionIDs {
		o.observeCollection(ctx, cid)
	}
}

func (o *LeaderObserver) observeCollection(ctx context.Context, collection int64) {
	replicas := o.meta.ReplicaManager.GetByCollection(collection)
	for _, replica := range replicas {
		leaders := o.dist.ChannelDistManager.GetShardLeadersByReplica(replica)
		for ch, leaderID := range leaders {
			leaderView := o.dist.LeaderViewManager.GetLeaderShardView(leaderID, ch)
			if leaderView == nil {
				continue
			}
			dists := o.dist.SegmentDistManager.GetByShardWithReplica(ch, replica)
			needLoaded, needRemoved := o.findNeedLoadedSegments(leaderView, dists),
				o.findNeedRemovedSegments(leaderView, dists)
			o.sync(ctx, leaderView, append(needLoaded, needRemoved...))
		}
	}
}

func (o *LeaderObserver) findNeedLoadedSegments(leaderView *meta.LeaderView, dists []*meta.Segment) []*querypb.SyncAction {
	ret := make([]*querypb.SyncAction, 0)
	dists = utils.FindMaxVersionSegments(dists)
	for _, s := range dists {
		version, ok := leaderView.Segments[s.GetID()]
		currentTarget := o.target.GetHistoricalSegment(s.CollectionID, s.GetID(), meta.CurrentTarget)
		existInCurrentTarget := currentTarget != nil
		existInNextTarget := o.target.GetHistoricalSegment(s.CollectionID, s.GetID(), meta.NextTarget) != nil

		if !existInCurrentTarget && !existInNextTarget {
			continue
		}

		if !ok || version.GetVersion() < s.Version { // Leader misses this segment
			ctx := context.Background()
			resp, err := o.broker.GetSegmentInfo(ctx, s.GetID())
			if err != nil || len(resp.GetInfos()) == 0 {
				log.Warn("failed to get segment info from DataCoord", zap.Error(err))
				continue
			}
			segment := resp.GetInfos()[0]
			loadInfo := utils.PackSegmentLoadInfo(segment, nil)

			// Fix the leader view with lacks of delta logs
			if existInCurrentTarget && s.LastDeltaTimestamp < currentTarget.GetDmlPosition().GetTimestamp() {
				log.Info("Fix QueryNode delta logs lag",
					zap.Int64("nodeID", s.Node),
					zap.Int64("collectionID", s.GetCollectionID()),
					zap.Int64("partitionID", s.GetPartitionID()),
					zap.Int64("segmentID", s.GetID()),
					zap.Uint64("segmentDeltaTimestamp", s.LastDeltaTimestamp),
					zap.Uint64("channelTimestamp", currentTarget.GetDmlPosition().GetTimestamp()),
				)

				ret = append(ret, &querypb.SyncAction{
					Type:        querypb.SyncType_Amend,
					PartitionID: s.GetPartitionID(),
					SegmentID:   s.GetID(),
					NodeID:      s.Node,
					Version:     s.Version,
					Info:        loadInfo,
				})
			}

			ret = append(ret, &querypb.SyncAction{
				Type:        querypb.SyncType_Set,
				PartitionID: s.GetPartitionID(),
				SegmentID:   s.GetID(),
				NodeID:      s.Node,
				Version:     s.Version,
				Info:        loadInfo,
			})
		}
	}
	return ret
}

func (o *LeaderObserver) findNeedRemovedSegments(leaderView *meta.LeaderView, dists []*meta.Segment) []*querypb.SyncAction {
	ret := make([]*querypb.SyncAction, 0)
	distMap := make(map[int64]struct{})
	for _, s := range dists {
		distMap[s.GetID()] = struct{}{}
	}
	for sid := range leaderView.Segments {
		_, ok := distMap[sid]
		existInCurrentTarget := o.target.GetHistoricalSegment(leaderView.CollectionID, sid, meta.CurrentTarget) != nil
		existInNextTarget := o.target.GetHistoricalSegment(leaderView.CollectionID, sid, meta.NextTarget) != nil
		if ok || existInCurrentTarget || existInNextTarget {
			continue
		}
		log.Debug("leader observer append a segment to remove:", zap.Int64("collectionID", leaderView.CollectionID),
			zap.String("Channel", leaderView.Channel), zap.Int64("leaderViewID", leaderView.ID),
			zap.Int64("segmentID", sid), zap.Bool("distMap_exist", ok),
			zap.Bool("existInCurrentTarget", existInCurrentTarget),
			zap.Bool("existInNextTarget", existInNextTarget))
		ret = append(ret, &querypb.SyncAction{
			Type:      querypb.SyncType_Remove,
			SegmentID: sid,
		})
	}
	return ret
}

func (o *LeaderObserver) sync(ctx context.Context, leaderView *meta.LeaderView, diffs []*querypb.SyncAction) {
	if len(diffs) == 0 {
		return
	}

	log := log.With(
		zap.Int64("leaderID", leaderView.ID),
		zap.Int64("collectionID", leaderView.CollectionID),
		zap.String("channel", leaderView.Channel),
	)
	req := &querypb.SyncDistributionRequest{
		Base: commonpbutil.NewMsgBase(
			commonpbutil.WithMsgType(commonpb.MsgType_SyncDistribution),
		),
		CollectionID: leaderView.CollectionID,
		Channel:      leaderView.Channel,
		Actions:      diffs,
	}
	resp, err := o.cluster.SyncDistribution(ctx, leaderView.ID, req)
	if err != nil {
		log.Error("failed to sync distribution", zap.Error(err))
		return
	}

	if resp.ErrorCode != commonpb.ErrorCode_Success {
		log.Error("failed to sync distribution", zap.String("reason", resp.GetReason()))
	}
}

func NewLeaderObserver(
	dist *meta.DistributionManager,
	meta *meta.Meta,
	targetMgr *meta.TargetManager,
	broker meta.Broker,
	cluster session.Cluster,
) *LeaderObserver {
	return &LeaderObserver{
		closeCh: make(chan struct{}),
		dist:    dist,
		meta:    meta,
		target:  targetMgr,
		broker:  broker,
		cluster: cluster,
	}
}
