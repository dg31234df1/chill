package dataservice

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/zilliztech/milvus-distributed/internal/log"
	"github.com/zilliztech/milvus-distributed/internal/util/trace"
	"github.com/zilliztech/milvus-distributed/internal/util/tsoutil"
	"github.com/zilliztech/milvus-distributed/internal/util/typeutil"

	"github.com/zilliztech/milvus-distributed/internal/proto/datapb"
)

type errRemainInSufficient struct {
	requestRows int
}

func newErrRemainInSufficient(requestRows int) errRemainInSufficient {
	return errRemainInSufficient{requestRows: requestRows}
}

func (err errRemainInSufficient) Error() string {
	return "segment remaining is insufficient for" + strconv.Itoa(err.requestRows)
}

// segmentAllocator is used to allocate rows for segments and record the allocations.
type segmentAllocatorInterface interface {
	// OpenSegment add the segment to allocator and set it allocatable
	OpenSegment(ctx context.Context, segmentInfo *datapb.SegmentInfo) error
	// AllocSegment allocate rows and record the allocation.
	AllocSegment(ctx context.Context, collectionID UniqueID, partitionID UniqueID, channelName string, requestRows int) (UniqueID, int, Timestamp, error)
	// GetSealedSegments get all sealed segment.
	GetSealedSegments(ctx context.Context) ([]UniqueID, error)
	// SealSegment set segment sealed, the segment will not be allocated anymore.
	SealSegment(ctx context.Context, segmentID UniqueID) error
	// DropSegment drop the segment from allocator.
	DropSegment(ctx context.Context, segmentID UniqueID)
	// ExpireAllocations check all allocations' expire time and remove the expired allocation.
	ExpireAllocations(ctx context.Context, timeTick Timestamp) error
	// SealAllSegments get all opened segment ids of collection. return success and failed segment ids
	SealAllSegments(ctx context.Context, collectionID UniqueID)
	// IsAllocationsExpired check all allocations of segment expired.
	IsAllocationsExpired(ctx context.Context, segmentID UniqueID, ts Timestamp) (bool, error)
}

type segmentStatus struct {
	id             UniqueID
	collectionID   UniqueID
	partitionID    UniqueID
	total          int
	sealed         bool
	lastExpireTime Timestamp
	allocations    []*allocation
	insertChannel  string
}
type allocation struct {
	rowNums    int
	expireTime Timestamp
}
type segmentAllocator struct {
	mt                     *meta
	segments               map[UniqueID]*segmentStatus //segment id -> status
	segmentExpireDuration  int64
	segmentThreshold       float64
	segmentThresholdFactor float64
	mu                     sync.RWMutex
	allocator              allocatorInterface
}

func newSegmentAllocator(meta *meta, allocator allocatorInterface) *segmentAllocator {
	segmentAllocator := &segmentAllocator{
		mt:                     meta,
		segments:               make(map[UniqueID]*segmentStatus),
		segmentExpireDuration:  Params.SegIDAssignExpiration,
		segmentThreshold:       Params.SegmentSize * 1024 * 1024,
		segmentThresholdFactor: Params.SegmentSizeFactor,
		allocator:              allocator,
	}
	return segmentAllocator
}

func (allocator *segmentAllocator) OpenSegment(ctx context.Context, segmentInfo *datapb.SegmentInfo) error {
	sp, _ := trace.StartSpanFromContext(ctx)
	defer sp.Finish()
	allocator.mu.Lock()
	defer allocator.mu.Unlock()
	if _, ok := allocator.segments[segmentInfo.SegmentID]; ok {
		return fmt.Errorf("segment %d already exist", segmentInfo.SegmentID)
	}
	totalRows, err := allocator.estimateTotalRows(segmentInfo.CollectionID)
	if err != nil {
		return err
	}
	log.Debug("dataservice: estimateTotalRows: ",
		zap.Int64("CollectionID", segmentInfo.CollectionID),
		zap.Int64("SegmentID", segmentInfo.SegmentID),
		zap.Int("Rows", totalRows))
	allocator.segments[segmentInfo.SegmentID] = &segmentStatus{
		id:             segmentInfo.SegmentID,
		collectionID:   segmentInfo.CollectionID,
		partitionID:    segmentInfo.PartitionID,
		total:          totalRows,
		sealed:         false,
		lastExpireTime: 0,
		insertChannel:  segmentInfo.InsertChannel,
	}
	return nil
}

func (allocator *segmentAllocator) AllocSegment(ctx context.Context, collectionID UniqueID,
	partitionID UniqueID, channelName string, requestRows int) (segID UniqueID, retCount int, expireTime Timestamp, err error) {
	sp, _ := trace.StartSpanFromContext(ctx)
	defer sp.Finish()
	allocator.mu.Lock()
	defer allocator.mu.Unlock()

	for _, segStatus := range allocator.segments {
		if segStatus.sealed || segStatus.collectionID != collectionID || segStatus.partitionID != partitionID ||
			segStatus.insertChannel != channelName {
			continue
		}
		var success bool
		success, err = allocator.alloc(segStatus, requestRows)
		if err != nil {
			return
		}
		if !success {
			continue
		}
		segID = segStatus.id
		retCount = requestRows
		expireTime = segStatus.lastExpireTime
		return
	}

	err = newErrRemainInSufficient(requestRows)
	return
}

func (allocator *segmentAllocator) alloc(segStatus *segmentStatus, numRows int) (bool, error) {
	totalOfAllocations := 0
	for _, allocation := range segStatus.allocations {
		totalOfAllocations += allocation.rowNums
	}
	segMeta, err := allocator.mt.GetSegment(segStatus.id)
	if err != nil {
		return false, err
	}
	free := segStatus.total - int(segMeta.NumRows) - totalOfAllocations
	log.Debug("dataservice::alloc: ",
		zap.Any("segMeta.NumRows", int(segMeta.NumRows)),
		zap.Any("totalOfAllocations", totalOfAllocations))
	if numRows > free {
		return false, nil
	}

	ts, err := allocator.allocator.allocTimestamp()
	if err != nil {
		return false, err
	}
	physicalTs, logicalTs := tsoutil.ParseTS(ts)
	expirePhysicalTs := physicalTs.Add(time.Duration(allocator.segmentExpireDuration) * time.Millisecond)
	expireTs := tsoutil.ComposeTS(expirePhysicalTs.UnixNano()/int64(time.Millisecond), int64(logicalTs))
	segStatus.lastExpireTime = expireTs
	segStatus.allocations = append(segStatus.allocations, &allocation{
		numRows,
		expireTs,
	})

	return true, nil
}

func (allocator *segmentAllocator) estimateTotalRows(collectionID UniqueID) (int, error) {
	collMeta, err := allocator.mt.GetCollection(collectionID)
	if err != nil {
		return -1, err
	}
	sizePerRecord, err := typeutil.EstimateSizePerRecord(collMeta.Schema)
	if err != nil {
		return -1, err
	}
	return int(allocator.segmentThreshold / float64(sizePerRecord)), nil
}

func (allocator *segmentAllocator) GetSealedSegments(ctx context.Context) ([]UniqueID, error) {
	sp, _ := trace.StartSpanFromContext(ctx)
	defer sp.Finish()
	allocator.mu.Lock()
	defer allocator.mu.Unlock()
	keys := make([]UniqueID, 0)
	for _, segStatus := range allocator.segments {
		if !segStatus.sealed {
			sealed, err := allocator.checkSegmentSealed(segStatus)
			if err != nil {
				return nil, err
			}
			segStatus.sealed = sealed
		}
		if segStatus.sealed {
			keys = append(keys, segStatus.id)
		}
	}
	return keys, nil
}

func (allocator *segmentAllocator) checkSegmentSealed(segStatus *segmentStatus) (bool, error) {
	segMeta, err := allocator.mt.GetSegment(segStatus.id)
	if err != nil {
		return false, err
	}
	return float64(segMeta.NumRows) >= allocator.segmentThresholdFactor*float64(segStatus.total), nil
}

func (allocator *segmentAllocator) SealSegment(ctx context.Context, segmentID UniqueID) error {
	sp, _ := trace.StartSpanFromContext(ctx)
	defer sp.Finish()
	allocator.mu.Lock()
	defer allocator.mu.Unlock()
	status, ok := allocator.segments[segmentID]
	if !ok {
		return nil
	}
	status.sealed = true
	return nil
}

func (allocator *segmentAllocator) DropSegment(ctx context.Context, segmentID UniqueID) {
	sp, _ := trace.StartSpanFromContext(ctx)
	defer sp.Finish()
	allocator.mu.Lock()
	defer allocator.mu.Unlock()
	delete(allocator.segments, segmentID)
}

func (allocator *segmentAllocator) ExpireAllocations(ctx context.Context, timeTick Timestamp) error {
	sp, _ := trace.StartSpanFromContext(ctx)
	defer sp.Finish()
	allocator.mu.Lock()
	defer allocator.mu.Unlock()
	for _, segStatus := range allocator.segments {
		for i := 0; i < len(segStatus.allocations); i++ {
			if timeTick < segStatus.allocations[i].expireTime {
				continue
			}
			log.Debug("dataservice::ExpireAllocations: ",
				zap.Any("segStatus.id", segStatus.id),
				zap.Any("segStatus.allocations.rowNums", segStatus.allocations[i].rowNums))
			segStatus.allocations = append(segStatus.allocations[:i], segStatus.allocations[i+1:]...)
			i--
		}
	}
	return nil
}

func (allocator *segmentAllocator) IsAllocationsExpired(ctx context.Context, segmentID UniqueID, ts Timestamp) (bool, error) {
	sp, _ := trace.StartSpanFromContext(ctx)
	defer sp.Finish()
	allocator.mu.RLock()
	defer allocator.mu.RUnlock()
	status, ok := allocator.segments[segmentID]
	if !ok {
		return false, fmt.Errorf("segment %d not found", segmentID)
	}
	return status.lastExpireTime <= ts, nil
}

func (allocator *segmentAllocator) SealAllSegments(ctx context.Context, collectionID UniqueID) {
	sp, _ := trace.StartSpanFromContext(ctx)
	defer sp.Finish()
	allocator.mu.Lock()
	defer allocator.mu.Unlock()
	for _, status := range allocator.segments {
		if status.collectionID == collectionID {
			if status.sealed {
				continue
			}
			status.sealed = true
		}
	}
}
