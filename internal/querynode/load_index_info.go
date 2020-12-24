package querynode

/*
#cgo CFLAGS: -I${SRCDIR}/../core/output/include
#cgo LDFLAGS: -L${SRCDIR}/../core/output/lib -lmilvus_segcore -Wl,-rpath=${SRCDIR}/../core/output/lib

#include "segcore/load_index_c.h"

*/
import "C"
import (
	"errors"
	"strconv"
	"unsafe"
)

type LoadIndexInfo struct {
	cLoadIndexInfo C.CLoadIndexInfo
}

func NewLoadIndexInfo() (*LoadIndexInfo, error) {
	var cLoadIndexInfo C.CLoadIndexInfo
	status := C.NewLoadIndexInfo(&cLoadIndexInfo)
	errorCode := status.error_code

	if errorCode != 0 {
		errorMsg := C.GoString(status.error_msg)
		defer C.free(unsafe.Pointer(status.error_msg))
		return nil, errors.New("NewLoadIndexInfo failed, C runtime error detected, error code = " + strconv.Itoa(int(errorCode)) + ", error msg = " + errorMsg)
	}
	return &LoadIndexInfo{cLoadIndexInfo: cLoadIndexInfo}, nil
}

func (li *LoadIndexInfo) AppendIndexParam(indexKey string, indexValue string) error {
	cIndexKey := C.CString(indexKey)
	cIndexValue := C.CString(indexValue)
	status := C.AppendIndexParam(li.cLoadIndexInfo, cIndexKey, cIndexValue)
	errorCode := status.error_code

	if errorCode != 0 {
		errorMsg := C.GoString(status.error_msg)
		defer C.free(unsafe.Pointer(status.error_msg))
		return errors.New("AppendIndexParam failed, C runtime error detected, error code = " + strconv.Itoa(int(errorCode)) + ", error msg = " + errorMsg)
	}
	return nil
}

func (li *LoadIndexInfo) AppendFieldInfo(fieldName string, fieldID int64) error {
	cFieldName := C.CString(fieldName)
	cFieldID := C.long(fieldID)
	status := C.AppendFieldInfo(li.cLoadIndexInfo, cFieldName, cFieldID)
	errorCode := status.error_code

	if errorCode != 0 {
		errorMsg := C.GoString(status.error_msg)
		defer C.free(unsafe.Pointer(status.error_msg))
		return errors.New("AppendFieldInfo failed, C runtime error detected, error code = " + strconv.Itoa(int(errorCode)) + ", error msg = " + errorMsg)
	}
	return nil
}

func (li *LoadIndexInfo) AppendIndex(bytesIndex [][]byte, indexKeys []string) error {
	var cBinarySet C.CBinarySet
	status := C.NewBinarySet(&cBinarySet)

	errorCode := status.error_code
	if errorCode != 0 {
		errorMsg := C.GoString(status.error_msg)
		defer C.free(unsafe.Pointer(status.error_msg))
		return errors.New("newBinarySet failed, C runtime error detected, error code = " + strconv.Itoa(int(errorCode)) + ", error msg = " + errorMsg)
	}

	for i, byteIndex := range bytesIndex {
		indexPtr := unsafe.Pointer(&byteIndex[0])
		indexLen := C.long(len(byteIndex))
		indexKey := C.CString(indexKeys[i])
		status = C.AppendBinaryIndex(cBinarySet, indexPtr, indexLen, indexKey)
		errorCode = status.error_code
		if errorCode != 0 {
			break
		}
	}
	if errorCode != 0 {
		errorMsg := C.GoString(status.error_msg)
		defer C.free(unsafe.Pointer(status.error_msg))
		return errors.New("AppendBinaryIndex failed, C runtime error detected, error code = " + strconv.Itoa(int(errorCode)) + ", error msg = " + errorMsg)
	}

	status = C.AppendIndex(li.cLoadIndexInfo, cBinarySet)
	errorCode = status.error_code
	if errorCode != 0 {
		errorMsg := C.GoString(status.error_msg)
		defer C.free(unsafe.Pointer(status.error_msg))
		return errors.New("AppendIndex failed, C runtime error detected, error code = " + strconv.Itoa(int(errorCode)) + ", error msg = " + errorMsg)
	}

	return nil
}
