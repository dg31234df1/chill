package reader

/*

#cgo CFLAGS: -I../core/include

#cgo LDFLAGS: -L../core/lib -lmilvus_dog_segment -Wl,-rpath=../core/lib

#include "collection_c.h"
#include "partition_c.h"
#include "segment_c.h"

*/
import "C"
import (
	"github.com/czs007/suvlim/errors"
	schema "github.com/czs007/suvlim/pkg/message"
	"strconv"
	"unsafe"
)

const SegmentLifetime = 20000

const (
	SegmentOpened = 0
	SegmentClosed = 1
)

type Segment struct {
	SegmentPtr       C.CSegmentBase
	SegmentId        int64
	SegmentCloseTime uint64
}

func (s *Segment) GetStatus() int {
	/*C.IsOpened
	bool
	IsOpened(CSegmentBase c_segment);
	*/
	var isOpened = C.IsOpened(s.SegmentPtr)
	if isOpened {
		return SegmentOpened
	} else {
		return SegmentClosed
	}
}

func (s *Segment) GetRowCount() int64 {
	/*C.GetRowCount
	long int
	GetRowCount(CSegmentBase c_segment);
	*/
	var rowCount = C.GetRowCount(s.SegmentPtr)
	return int64(rowCount)
}

func (s *Segment) GetDeletedCount() int64 {
	/*C.GetDeletedCount
	long int
	GetDeletedCount(CSegmentBase c_segment);
	*/
	var deletedCount = C.GetDeletedCount(s.SegmentPtr)
	return int64(deletedCount)
}

func (s *Segment) Close() error {
	/*C.Close
	int
	Close(CSegmentBase c_segment);
	*/
	var status = C.Close(s.SegmentPtr)
	if status != 0 {
		return errors.New("Close segment failed, error code = " + strconv.Itoa(int(status)))
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////
func (s *Segment) SegmentPreInsert(numOfRecords int) int64 {
	/*C.PreInsert
	long int
	PreInsert(CSegmentBase c_segment, long int size);
	*/
	var offset = C.PreInsert(s.SegmentPtr, C.long(int64(numOfRecords)))

	return int64(offset)
}

func (s *Segment) SegmentPreDelete(numOfRecords int) int64 {
	/*C.PreDelete
	long int
	PreDelete(CSegmentBase c_segment, long int size);
	*/
	var offset = C.PreDelete(s.SegmentPtr, C.long(int64(numOfRecords)))

	return int64(offset)
}

func (s *Segment) SegmentInsert(offset int64, entityIDs *[]int64, timestamps *[]uint64, records *[][]byte) error {
	/*C.Insert
	int
	Insert(CSegmentBase c_segment,
	           long int reserved_offset,
	           signed long int size,
	           const long* primary_keys,
	           const unsigned long* timestamps,
	           void* raw_data,
	           int sizeof_per_row,
	           signed long int count);
	*/
	// Blobs to one big blob
	var rawData []byte
	for i := 0; i < len(*records); i++ {
		copy(rawData, (*records)[i])
	}

	var cOffset = C.long(offset)
	var cNumOfRows = C.long(len(*entityIDs))
	var cEntityIdsPtr = (*C.long)(&(*entityIDs)[0])
	var cTimestampsPtr = (*C.ulong)(&(*timestamps)[0])
	var cSizeofPerRow = C.int(len((*records)[0]))
	var cRawDataVoidPtr = unsafe.Pointer(&rawData[0])

	var status = C.Insert(s.SegmentPtr,
							cOffset,
							cNumOfRows,
							cEntityIdsPtr,
							cTimestampsPtr,
		cRawDataVoidPtr,
							cSizeofPerRow,
							cNumOfRows)

	if status != 0 {
		return errors.New("Insert failed, error code = " + strconv.Itoa(int(status)))
	}

	return nil
}

func (s *Segment) SegmentDelete(offset int64, entityIDs *[]int64, timestamps *[]uint64) error {
	/*C.Delete
	int
	Delete(CSegmentBase c_segment,
	           long int reserved_offset,
	           long size,
	           const long* primary_keys,
	           const unsigned long* timestamps);
	*/
	var cOffset = C.long(offset)
	var cSize = C.long(len(*entityIDs))
	var cEntityIdsPtr = (*C.long)(&(*entityIDs)[0])
	var cTimestampsPtr = (*C.ulong)(&(*timestamps)[0])

	var status = C.Delete(s.SegmentPtr, cOffset, cSize, cEntityIdsPtr, cTimestampsPtr)

	if status != 0 {
		return errors.New("Delete failed, error code = " + strconv.Itoa(int(status)))
	}

	return nil
}

func (s *Segment) SegmentSearch(queryString string, timestamp uint64, vectorRecord *schema.VectorRowRecord) (*SearchResult, error) {
	/*C.Search
	int
	Search(CSegmentBase c_segment,
	           void* fake_query,
	           unsigned long timestamp,
	           long int* result_ids,
	           float* result_distances);
	*/
	// TODO: get top-k's k from queryString
	const TopK = 1

	resultIds := make([]int64, TopK)
	resultDistances := make([]float32, TopK)

	var cQueryPtr = unsafe.Pointer(nil)
	var cTimestamp = C.ulong(timestamp)
	var cResultIds = (*C.long)(&resultIds[0])
	var cResultDistances = (*C.float)(&resultDistances[0])

	var status = C.Search(s.SegmentPtr, cQueryPtr, cTimestamp, cResultIds, cResultDistances)

	if status != 0 {
		return nil, errors.New("Search failed, error code = " + strconv.Itoa(int(status)))
	}

	return &SearchResult{ResultIds: resultIds, ResultDistances: resultDistances}, nil
}
