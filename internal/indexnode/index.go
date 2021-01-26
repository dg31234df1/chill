package indexnode

/*

#cgo CFLAGS: -I${SRCDIR}/../core/output/include

#cgo LDFLAGS: -L${SRCDIR}/../core/output/lib -lmilvus_indexbuilder -Wl,-rpath=${SRCDIR}/../core/output/lib

#include <stdlib.h>	// free
#include "segcore/collection_c.h"
#include "indexbuilder/index_c.h"

*/
import "C"

import (
	"errors"
	"fmt"
	"strconv"
	"unsafe"

	"github.com/golang/protobuf/proto"
	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/indexcgopb"
	"github.com/zilliztech/milvus-distributed/internal/storage"
)

// TODO: use storage.Blob instead later
type Blob = storage.Blob

// just for debugging
type QueryResult interface {
	Delete() error
	NQ() int64
	TOPK() int64
	IDs() []int64
	Distances() []float32
}

type CQueryResult struct {
	ptr C.CIndexQueryResult
}

type CFunc func() C.CStatus

func TryCatch(fn CFunc) error {
	status := fn()
	errorCode := status.error_code
	if errorCode != 0 {
		errorMsg := C.GoString(status.error_msg)
		defer C.free(unsafe.Pointer(status.error_msg))
		return errors.New("error code = " + strconv.Itoa(int(errorCode)) + ", error msg = " + errorMsg)
	}
	return nil
}

func CreateQueryResult() (QueryResult, error) {
	var ptr C.CIndexQueryResult
	fn := func() C.CStatus {
		return C.CreateQueryResult(&ptr)
	}
	err := TryCatch(fn)
	if err != nil {
		return nil, err
	}
	return &CQueryResult{
		ptr: ptr,
	}, nil
}

func (qs *CQueryResult) Delete() error {
	fn := func() C.CStatus {
		return C.DeleteQueryResult(qs.ptr)
	}
	return TryCatch(fn)
}

func (qs *CQueryResult) NQ() int64 {
	return int64(C.NqOfQueryResult(qs.ptr))
}

func (qs *CQueryResult) TOPK() int64 {
	return int64(C.TopkOfQueryResult(qs.ptr))
}

func (qs *CQueryResult) IDs() []int64 {
	nq := qs.NQ()
	topk := qs.TOPK()

	if nq <= 0 || topk <= 0 {
		return []int64{}
	}

	// TODO: how could we avoid memory copy every time when this called
	ids := make([]int64, nq*topk)
	C.GetIdsOfQueryResult(qs.ptr, (*C.int64_t)(&ids[0]))

	return ids
}

func (qs *CQueryResult) Distances() []float32 {
	nq := qs.NQ()
	topk := qs.TOPK()

	if nq <= 0 || topk <= 0 {
		return []float32{}
	}

	// TODO: how could we avoid memory copy every time when this called
	distances := make([]float32, nq*topk)
	C.GetDistancesOfQueryResult(qs.ptr, (*C.float)(&distances[0]))

	return distances
}

type Index interface {
	Serialize() ([]*Blob, error)
	Load([]*Blob) error
	BuildFloatVecIndexWithoutIds(vectors []float32) error
	BuildBinaryVecIndexWithoutIds(vectors []byte) error
	QueryOnFloatVecIndex(vectors []float32) (QueryResult, error)
	QueryOnBinaryVecIndex(vectors []byte) (QueryResult, error)
	QueryOnFloatVecIndexWithParam(vectors []float32, params map[string]string) (QueryResult, error)
	QueryOnBinaryVecIndexWithParam(vectors []byte, params map[string]string) (QueryResult, error)
	Delete() error
}

type CIndex struct {
	indexPtr C.CIndex
}

func (index *CIndex) Serialize() ([]*Blob, error) {
	/*
		CStatus
		SerializeToSlicedBuffer(CIndex index, int32_t* buffer_size, char** res_buffer);
	*/

	var cDumpedSlicedBuffer *C.char
	var bufferSize int32
	status := C.SerializeToSlicedBuffer(index.indexPtr, (*C.int32_t)(unsafe.Pointer(&bufferSize)), &cDumpedSlicedBuffer)
	errorCode := status.error_code
	if errorCode != 0 {
		errorMsg := C.GoString(status.error_msg)
		defer C.free(unsafe.Pointer(status.error_msg))
		return nil, errors.New("SerializeToSlicedBuffer failed, C runtime error detected, error code = " + strconv.Itoa(int(errorCode)) + ", error msg = " + errorMsg)
	}

	defer C.free(unsafe.Pointer(cDumpedSlicedBuffer))

	dumpedSlicedBuffer := C.GoBytes(unsafe.Pointer(cDumpedSlicedBuffer), (C.int32_t)(bufferSize))
	var blobs indexcgopb.BinarySet
	err := proto.Unmarshal(dumpedSlicedBuffer, &blobs)
	if err != nil {
		return nil, err
	}

	ret := make([]*Blob, 0)
	for _, data := range blobs.Datas {
		ret = append(ret, &Blob{Key: data.Key, Value: data.Value})
	}

	return ret, nil
}

func (index *CIndex) Load(blobs []*Blob) error {
	binarySet := &indexcgopb.BinarySet{Datas: make([]*indexcgopb.Binary, 0)}
	for _, blob := range blobs {
		binarySet.Datas = append(binarySet.Datas, &indexcgopb.Binary{Key: blob.Key, Value: blob.Value})
	}

	datas, err := proto.Marshal(binarySet)
	if err != nil {
		return err
	}

	/*
		CStatus
		LoadFromSlicedBuffer(CIndex index, const char* serialized_sliced_blob_buffer, int32_t size);
	*/
	status := C.LoadFromSlicedBuffer(index.indexPtr, (*C.char)(unsafe.Pointer(&datas[0])), (C.int32_t)(len(datas)))
	errorCode := status.error_code
	if errorCode != 0 {
		errorMsg := C.GoString(status.error_msg)
		defer C.free(unsafe.Pointer(status.error_msg))
		return errors.New("BuildFloatVecIndexWithoutIds failed, C runtime error detected, error code = " + strconv.Itoa(int(errorCode)) + ", error msg = " + errorMsg)
	}
	return nil
}

func (index *CIndex) BuildFloatVecIndexWithoutIds(vectors []float32) error {
	/*
		CStatus
		BuildFloatVecIndexWithoutIds(CIndex index, int64_t float_value_num, const float* vectors);
	*/
	fmt.Println("before BuildFloatVecIndexWithoutIds")
	status := C.BuildFloatVecIndexWithoutIds(index.indexPtr, (C.int64_t)(len(vectors)), (*C.float)(&vectors[0]))
	errorCode := status.error_code
	fmt.Println("BuildFloatVecIndexWithoutIds error code: ", errorCode)
	if errorCode != 0 {
		errorMsg := C.GoString(status.error_msg)
		fmt.Println("BuildFloatVecIndexWithoutIds error msg: ", errorMsg)
		defer C.free(unsafe.Pointer(status.error_msg))
		return errors.New("BuildFloatVecIndexWithoutIds failed, C runtime error detected, error code = " + strconv.Itoa(int(errorCode)) + ", error msg = " + errorMsg)
	}
	return nil
}

func (index *CIndex) BuildBinaryVecIndexWithoutIds(vectors []byte) error {
	/*
		CStatus
		BuildBinaryVecIndexWithoutIds(CIndex index, int64_t data_size, const uint8_t* vectors);
	*/
	status := C.BuildBinaryVecIndexWithoutIds(index.indexPtr, (C.int64_t)(len(vectors)), (*C.uint8_t)(&vectors[0]))
	errorCode := status.error_code
	if errorCode != 0 {
		errorMsg := C.GoString(status.error_msg)
		defer C.free(unsafe.Pointer(status.error_msg))
		return errors.New(" failed, C runtime error detected, error code = " + strconv.Itoa(int(errorCode)) + ", error msg = " + errorMsg)
	}
	return nil
}

func (index *CIndex) Delete() error {
	/*
		void
		DeleteIndex(CIndex index);
	*/
	C.DeleteIndex(index.indexPtr)
	// TODO: check if index.indexPtr will be released by golang, though it occupies little memory
	// C.free(index.indexPtr)
	return nil
}

func NewCIndex(typeParams, indexParams map[string]string) (Index, error) {
	protoTypeParams := &indexcgopb.TypeParams{
		Params: make([]*commonpb.KeyValuePair, 0),
	}
	for key, value := range typeParams {
		protoTypeParams.Params = append(protoTypeParams.Params, &commonpb.KeyValuePair{Key: key, Value: value})
	}
	typeParamsStr := proto.MarshalTextString(protoTypeParams)

	protoIndexParams := &indexcgopb.IndexParams{
		Params: make([]*commonpb.KeyValuePair, 0),
	}
	for key, value := range indexParams {
		protoIndexParams.Params = append(protoIndexParams.Params, &commonpb.KeyValuePair{Key: key, Value: value})
	}
	indexParamsStr := proto.MarshalTextString(protoIndexParams)

	typeParamsPointer := C.CString(typeParamsStr)
	indexParamsPointer := C.CString(indexParamsStr)

	/*
		CStatus
		CreateIndex(const char* serialized_type_params,
					const char* serialized_index_params,
					CIndex* res_index);
	*/
	var indexPtr C.CIndex
	fmt.Println("before create index ........................................")
	status := C.CreateIndex(typeParamsPointer, indexParamsPointer, &indexPtr)
	fmt.Println("after create index ........................................")
	errorCode := status.error_code
	fmt.Println("EEEEEEEEEEEEEEEEEEEEEEEEEE error code: ", errorCode)
	if errorCode != 0 {
		errorMsg := C.GoString(status.error_msg)
		fmt.Println("EEEEEEEEEEEEEEEEEEEEEEEEEE error msg: ", errorMsg)
		defer C.free(unsafe.Pointer(status.error_msg))
		return nil, errors.New(" failed, C runtime error detected, error code = " + strconv.Itoa(int(errorCode)) + ", error msg = " + errorMsg)
	}

	return &CIndex{
		indexPtr: indexPtr,
	}, nil
}

func (index *CIndex) QueryOnFloatVecIndex(vectors []float32) (QueryResult, error) {
	if len(vectors) <= 0 {
		return nil, errors.New("nq is zero")
	}
	res, err := CreateQueryResult()
	if err != nil {
		return nil, err
	}
	fn := func() C.CStatus {
		cRes, ok := res.(*CQueryResult)
		if !ok {
			// TODO: ugly here, fix me later
			panic("only CQueryResult is supported now!")
		}
		return C.QueryOnFloatVecIndex(index.indexPtr, (C.int64_t)(len(vectors)), (*C.float)(&vectors[0]), &cRes.ptr)
	}
	err = TryCatch(fn)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (index *CIndex) QueryOnBinaryVecIndex(vectors []byte) (QueryResult, error) {
	if len(vectors) <= 0 {
		return nil, errors.New("nq is zero")
	}
	res, err := CreateQueryResult()
	if err != nil {
		return nil, err
	}
	fn := func() C.CStatus {
		cRes, ok := res.(*CQueryResult)
		if !ok {
			// TODO: ugly here, fix me later
			panic("only CQueryResult is supported now!")
		}
		return C.QueryOnBinaryVecIndex(index.indexPtr, (C.int64_t)(len(vectors)), (*C.uint8_t)(&vectors[0]), &cRes.ptr)
	}
	err = TryCatch(fn)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (index *CIndex) QueryOnFloatVecIndexWithParam(vectors []float32, params map[string]string) (QueryResult, error) {
	if len(vectors) <= 0 {
		return nil, errors.New("nq is zero")
	}

	protoParams := &indexcgopb.MapParams{
		Params: make([]*commonpb.KeyValuePair, 0),
	}
	for key, value := range params {
		protoParams.Params = append(protoParams.Params, &commonpb.KeyValuePair{Key: key, Value: value})
	}
	paramsStr := proto.MarshalTextString(protoParams)
	paramsPointer := C.CString(paramsStr)

	res, err := CreateQueryResult()
	if err != nil {
		return nil, err
	}
	fn := func() C.CStatus {
		cRes, ok := res.(*CQueryResult)
		if !ok {
			// TODO: ugly here, fix me later
			panic("only CQueryResult is supported now!")
		}
		return C.QueryOnFloatVecIndexWithParam(index.indexPtr, (C.int64_t)(len(vectors)), (*C.float)(&vectors[0]), paramsPointer, &cRes.ptr)
	}
	err = TryCatch(fn)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (index *CIndex) QueryOnBinaryVecIndexWithParam(vectors []byte, params map[string]string) (QueryResult, error) {
	if len(vectors) <= 0 {
		return nil, errors.New("nq is zero")
	}

	protoParams := &indexcgopb.MapParams{
		Params: make([]*commonpb.KeyValuePair, 0),
	}
	for key, value := range params {
		protoParams.Params = append(protoParams.Params, &commonpb.KeyValuePair{Key: key, Value: value})
	}
	paramsStr := proto.MarshalTextString(protoParams)
	paramsPointer := C.CString(paramsStr)

	res, err := CreateQueryResult()
	if err != nil {
		return nil, err
	}
	fn := func() C.CStatus {
		cRes, ok := res.(*CQueryResult)
		if !ok {
			// TODO: ugly here, fix me later
			panic("only CQueryResult is supported now!")
		}
		return C.QueryOnBinaryVecIndexWithParam(index.indexPtr, (C.int64_t)(len(vectors)), (*C.uint8_t)(&vectors[0]), paramsPointer, &cRes.ptr)
	}
	err = TryCatch(fn)
	if err != nil {
		return nil, err
	}
	return res, nil
}
