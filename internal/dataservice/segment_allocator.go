package dataservice

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/zilliztech/milvus-distributed/internal/proto/datapb"

	"github.com/zilliztech/milvus-distributed/internal/util/typeutil"

	"github.com/zilliztech/milvus-distributed/internal/util/tsoutil"
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
	OpenSegment(segmentInfo *datapb.SegmentInfo) error
	// AllocSegment allocate rows and record the allocation.
	AllocSegment(collectionID UniqueID, partitionID UniqueID, channelName string, requestRows int) (UniqueID, int, Timestamp, error)
	// GetSealedSegments get all sealed segment.
	GetSealedSegments() ([]UniqueID, error)
	// SealSegment set segment sealed, the segment will not be allocated anymore.
	SealSegment(segmentID UniqueID) error
	// DropSegment drop the segment from allocator.
	DropSegment(segmentID UniqueID)
	// ExpireAllocations check all allocations' expire time and remove the expired allocation.
	ExpireAllocations(timeTick Timestamp) error
	// SealAllSegments get all opened segment ids of collection. return success and failed segment ids
	SealAllSegments(collectionID UniqueID)
	// IsAllocationsExpired check all allocations of segment expired.
	IsAllocationsExpired(segmentID UniqueID, ts Timestamp) (bool, error)
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

func (allocator *segmentAllocator) OpenSegment(segmentInfo *datapb.SegmentInfo) error {
	allocator.mu.Lock()
	defer allocator.mu.Unlock()
	if _, ok := allocator.segments[segmentInfo.SegmentID]; ok {
		return fmt.Errorf("segment %d already exist", segmentInfo.SegmentID)
	}
	totalRows, err := allocator.estimateTotalRows(segmentInfo.CollectionID)
	if err != nil {
		return err
	}
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

func (allocator *segmentAllocator) AllocSegment(collectionID UniqueID,
	partitionID UniqueID, channelName string, requestRows int) (segID UniqueID, retCount int, expireTime Timestamp, err error) {
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

func (allocator *segmentAllocator) GetSealedSegments() ([]UniqueID, error) {
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

func (allocator *segmentAllocator) SealSegment(segmentID UniqueID) error {
	allocator.mu.Lock()
	defer allocator.mu.Unlock()
	status, ok := allocator.segments[segmentID]
	if !ok {
		return nil
	}
	status.sealed = true
	return nil
}

func (allocator *segmentAllocator) DropSegment(segmentID UniqueID) {
	allocator.mu.Lock()
	defer allocator.mu.Unlock()
	delete(allocator.segments, segmentID)
}

func (allocator *segmentAllocator) ExpireAllocations(timeTick Timestamp) error {
	allocator.mu.Lock()
	defer allocator.mu.Unlock()
	for _, segStatus := range allocator.segments {
		for i := 0; i < len(segStatus.allocations); i++ {
			if timeTick < segStatus.allocations[i].expireTime {
				continue
			}
			segStatus.allocations = append(segStatus.allocations[:i], segStatus.allocations[i+1:]...)
			i--
		}
	}
	return nil
}

func (allocator *segmentAllocator) IsAllocationsExpired(segmentID UniqueID, ts Timestamp) (bool, error) {
	allocator.mu.RLock()
	defer allocator.mu.RUnlock()
	status, ok := allocator.segments[segmentID]
	if !ok {
		return false, fmt.Errorf("segment %d not found", segmentID)
	}
	return status.lastExpireTime <= ts, nil
}

func (allocator *segmentAllocator) SealAllSegments(collectionID UniqueID) {
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
