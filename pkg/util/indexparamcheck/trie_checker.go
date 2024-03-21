package indexparamcheck

import (
	"fmt"

	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	"github.com/milvus-io/milvus/pkg/util/typeutil"
)

// TRIEChecker checks if a TRIE index can be built.
type TRIEChecker struct {
	scalarIndexChecker
}

func (c *TRIEChecker) CheckTrain(params map[string]string) error {
	return c.scalarIndexChecker.CheckTrain(params)
}

func (c *TRIEChecker) CheckValidDataType(dType schemapb.DataType) error {
	if !typeutil.IsStringType(dType) {
		return fmt.Errorf("TRIE are only supported on varchar field")
	}
	return nil
}

func newTRIEChecker() *TRIEChecker {
	return &TRIEChecker{}
}
