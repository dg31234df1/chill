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

package segments

/*
#cgo pkg-config: milvus_segcore

#include "segcore/collection_c.h"
#include "segcore/segment_c.h"
*/
import "C"

import (
	"fmt"
	"sync"

	"github.com/milvus-io/milvus/internal/proto/querypb"
	"github.com/milvus-io/milvus/pkg/eventlog"
	"github.com/milvus-io/milvus/pkg/metrics"
	"github.com/milvus-io/milvus/pkg/util/paramtable"
	. "github.com/milvus-io/milvus/pkg/util/typeutil"
)

type SegmentFilter func(segment Segment) bool

func WithPartition(partitionID UniqueID) SegmentFilter {
	return func(segment Segment) bool {
		return segment.Partition() == partitionID
	}
}

func WithChannel(channel string) SegmentFilter {
	return func(segment Segment) bool {
		return segment.Shard() == channel
	}
}

func WithType(typ SegmentType) SegmentFilter {
	return func(segment Segment) bool {
		return segment.Type() == typ
	}
}

func WithID(id int64) SegmentFilter {
	return func(segment Segment) bool {
		return segment.ID() == id
	}
}

type actionType int32

const (
	removeAction actionType = iota
	addAction
)

type Manager struct {
	Collection CollectionManager
	Segment    SegmentManager
}

func NewManager() *Manager {
	return &Manager{
		Collection: NewCollectionManager(),
		Segment:    NewSegmentManager(),
	}
}

type SegmentManager interface {
	// Put puts the given segments in,
	// and increases the ref count of the corresponding collection,
	// dup segments will not increase the ref count
	Put(segmentType SegmentType, segments ...Segment)
	Get(segmentID UniqueID) Segment
	GetWithType(segmentID UniqueID, typ SegmentType) Segment
	GetBy(filters ...SegmentFilter) []Segment
	GetSealed(segmentID UniqueID) Segment
	GetGrowing(segmentID UniqueID) Segment
	Empty() bool

	// Remove removes the given segment,
	// and decreases the ref count of the corresponding collection,
	// will not decrease the ref count if the given segment not exists
	Remove(segmentID UniqueID, scope querypb.DataScope) (int, int)
	RemoveBy(filters ...SegmentFilter) (int, int)
	Clear()
}

var _ SegmentManager = (*segmentManager)(nil)

// Manager manages all collections and segments
type segmentManager struct {
	mu sync.RWMutex // guards all

	growingSegments map[UniqueID]Segment
	sealedSegments  map[UniqueID]Segment
}

func NewSegmentManager() *segmentManager {
	return &segmentManager{
		growingSegments: make(map[int64]Segment),
		sealedSegments:  make(map[int64]Segment),
	}
}

func (mgr *segmentManager) Put(segmentType SegmentType, segments ...Segment) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()

	targetMap := mgr.growingSegments
	switch segmentType {
	case SegmentTypeGrowing:
		targetMap = mgr.growingSegments
	case SegmentTypeSealed:
		targetMap = mgr.sealedSegments
	default:
		panic("unexpected segment type")
	}

	for _, segment := range segments {
		if _, ok := targetMap[segment.ID()]; ok {
			continue
		}
		eventlog.Record(eventlog.NewRawEvt(eventlog.Level_Info, fmt.Sprintf("Segment %d[%d] loaded", segment.ID(), segment.Collection())))
		targetMap[segment.ID()] = segment
		metrics.QueryNodeNumSegments.WithLabelValues(
			fmt.Sprint(paramtable.GetNodeID()),
			fmt.Sprint(segment.Collection()),
			fmt.Sprint(segment.Partition()),
			segment.Type().String(),
			fmt.Sprint(len(segment.Indexes())),
		).Inc()
		if segment.RowNum() > 0 {
			metrics.QueryNodeNumEntities.WithLabelValues(
				fmt.Sprint(paramtable.GetNodeID()),
				fmt.Sprint(segment.Collection()),
				fmt.Sprint(segment.Partition()),
				segment.Type().String(),
				fmt.Sprint(len(segment.Indexes())),
			).Add(float64(segment.RowNum()))
		}
	}
	mgr.updateMetric()
}

func (mgr *segmentManager) Get(segmentID UniqueID) Segment {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()

	if segment, ok := mgr.growingSegments[segmentID]; ok {
		return segment
	} else if segment, ok = mgr.sealedSegments[segmentID]; ok {
		return segment
	}

	return nil
}

func (mgr *segmentManager) GetWithType(segmentID UniqueID, typ SegmentType) Segment {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()

	switch typ {
	case SegmentTypeSealed:
		return mgr.sealedSegments[segmentID]
	case SegmentTypeGrowing:
		return mgr.growingSegments[segmentID]
	default:
		return nil
	}
}

func (mgr *segmentManager) GetBy(filters ...SegmentFilter) []Segment {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()

	ret := make([]Segment, 0)
	for _, segment := range mgr.growingSegments {
		if filter(segment, filters...) {
			ret = append(ret, segment)
		}
	}

	for _, segment := range mgr.sealedSegments {
		if filter(segment, filters...) {
			ret = append(ret, segment)
		}
	}
	return ret
}

func filter(segment Segment, filters ...SegmentFilter) bool {
	for _, filter := range filters {
		if !filter(segment) {
			return false
		}
	}
	return true
}

func (mgr *segmentManager) GetSealed(segmentID UniqueID) Segment {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()

	if segment, ok := mgr.sealedSegments[segmentID]; ok {
		return segment
	}

	return nil
}

func (mgr *segmentManager) GetGrowing(segmentID UniqueID) Segment {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()

	if segment, ok := mgr.growingSegments[segmentID]; ok {
		return segment
	}

	return nil
}

func (mgr *segmentManager) Empty() bool {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()

	return len(mgr.growingSegments)+len(mgr.sealedSegments) == 0
}

// returns true if the segment exists,
// false otherwise
func (mgr *segmentManager) Remove(segmentID UniqueID, scope querypb.DataScope) (int, int) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()

	var removeGrowing, removeSealed int
	switch scope {
	case querypb.DataScope_Streaming:
		if remove(segmentID, mgr.growingSegments) {
			removeGrowing = 1
		}

	case querypb.DataScope_Historical:
		if remove(segmentID, mgr.sealedSegments) {
			removeSealed = 1
		}

	case querypb.DataScope_All:
		if remove(segmentID, mgr.growingSegments) {
			removeGrowing = 1
		}
		if remove(segmentID, mgr.sealedSegments) {
			removeSealed = 1
		}
	}

	mgr.updateMetric()
	return removeGrowing, removeSealed
}

func (mgr *segmentManager) RemoveBy(filters ...SegmentFilter) (int, int) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()

	var removeGrowing, removeSealed int
	for id, segment := range mgr.growingSegments {
		if filter(segment, filters...) && remove(id, mgr.growingSegments) {
			removeGrowing++
		}
	}

	for id, segment := range mgr.sealedSegments {
		if filter(segment, filters...) && remove(id, mgr.sealedSegments) {
			removeSealed++
		}
	}

	mgr.updateMetric()
	return removeGrowing, removeSealed
}

func (mgr *segmentManager) Clear() {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()

	for id := range mgr.growingSegments {
		remove(id, mgr.growingSegments)
	}

	for id := range mgr.sealedSegments {
		remove(id, mgr.sealedSegments)
	}
	mgr.updateMetric()
}

func (mgr *segmentManager) updateMetric() {
	// update collection and partiation metric
	var collections, partiations = make(Set[int64]), make(Set[int64])
	for _, seg := range mgr.growingSegments {
		collections.Insert(seg.Collection())
		partiations.Insert(seg.Partition())
	}
	for _, seg := range mgr.sealedSegments {
		collections.Insert(seg.Collection())
		partiations.Insert(seg.Partition())
	}
	metrics.QueryNodeNumCollections.WithLabelValues(fmt.Sprint(paramtable.GetNodeID())).Set(float64(collections.Len()))
	metrics.QueryNodeNumPartitions.WithLabelValues(fmt.Sprint(paramtable.GetNodeID())).Set(float64(partiations.Len()))
}

// returns true if the segment exists,
// false otherwise
func remove(segmentID int64, container map[int64]Segment) bool {
	segment, ok := container[segmentID]
	if !ok {
		return false
	}
	delete(container, segmentID)

	rowNum := segment.RowNum()
	DeleteSegment(segment.(*LocalSegment))

	metrics.QueryNodeNumSegments.WithLabelValues(
		fmt.Sprint(paramtable.GetNodeID()),
		fmt.Sprint(segment.Collection()),
		fmt.Sprint(segment.Partition()),
		segment.Type().String(),
		fmt.Sprint(len(segment.Indexes())),
	).Dec()
	if rowNum > 0 {
		metrics.QueryNodeNumEntities.WithLabelValues(
			fmt.Sprint(paramtable.GetNodeID()),
			fmt.Sprint(segment.Collection()),
			fmt.Sprint(segment.Partition()),
			segment.Type().String(),
			fmt.Sprint(len(segment.Indexes())),
		).Sub(float64(rowNum))
	}
	return true
}
