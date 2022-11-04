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

package proxy

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/milvus-io/milvus-proto/go-api/commonpb"
	"github.com/milvus-io/milvus/internal/metrics"
	"github.com/milvus-io/milvus/internal/proto/rootcoordpb"
	"github.com/milvus-io/milvus/internal/util/commonpbutil"
	"github.com/milvus-io/milvus/internal/util/paramtable"
	"github.com/milvus-io/milvus/internal/util/timerecord"
)

// timestampAllocator implements tsoAllocator.
type timestampAllocator struct {
	ctx    context.Context
	tso    timestampAllocatorInterface
	peerID UniqueID
}

// newTimestampAllocator creates a new timestampAllocator
func newTimestampAllocator(ctx context.Context, tso timestampAllocatorInterface, peerID UniqueID) (*timestampAllocator, error) {
	a := &timestampAllocator{
		ctx:    ctx,
		peerID: peerID,
		tso:    tso,
	}
	return a, nil
}

func (ta *timestampAllocator) alloc(count uint32) ([]Timestamp, error) {
	tr := timerecord.NewTimeRecorder("applyTimestamp")
	ctx, cancel := context.WithTimeout(ta.ctx, 5*time.Second)
	req := &rootcoordpb.AllocTimestampRequest{
		Base: commonpbutil.NewMsgBase(
			commonpbutil.WithMsgType(commonpb.MsgType_RequestTSO),
			commonpbutil.WithMsgID(0),
			commonpbutil.WithTimeStamp(0),
			commonpbutil.WithSourceID(ta.peerID),
		),
		Count: count,
	}

	resp, err := ta.tso.AllocTimestamp(ctx, req)
	defer func() {
		cancel()
		metrics.ProxyApplyTimestampLatency.WithLabelValues(strconv.FormatInt(paramtable.GetNodeID(), 10)).Observe(float64(tr.ElapseSpan().Milliseconds()))
	}()

	if err != nil {
		return nil, fmt.Errorf("syncTimestamp Failed:%w", err)
	}
	if resp.Status.ErrorCode != commonpb.ErrorCode_Success {
		return nil, fmt.Errorf("syncTimeStamp Failed:%s", resp.Status.Reason)
	}
	start, cnt := resp.Timestamp, resp.Count
	ret := make([]Timestamp, cnt)
	for i := uint32(0); i < cnt; i++ {
		ret[i] = start + uint64(i)
	}

	return ret, nil
}

// AllocOne allocates a timestamp.
func (ta *timestampAllocator) AllocOne() (Timestamp, error) {
	ret, err := ta.alloc(1)
	if err != nil {
		return 0, err
	}
	return ret[0], nil
}
