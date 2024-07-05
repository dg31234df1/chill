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

package indexparamcheck

import (
	"fmt"
	"math"
	"strings"

	"github.com/cockroachdb/errors"

	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	"github.com/milvus-io/milvus/pkg/common"
)

type baseChecker struct{}

func (c baseChecker) CheckTrain(params map[string]string) error {
	// vector dimension should be checked on collection creation. this is just some basic check
	isSparse := false
	if val, exist := params[common.IsSparseKey]; exist {
		val = strings.ToLower(val)
		if val != "true" && val != "false" {
			return fmt.Errorf("invalid is_sparse value: %s, must be true or false", val)
		}
		if val == "true" {
			isSparse = true
		}
	}
	if isSparse {
		if !CheckStrByValues(params, Metric, SparseMetrics) {
			return fmt.Errorf("metric type not found or not supported for sparse float vectors, supported: %v", SparseMetrics)
		}
	} else {
		// we do not check dim for sparse
		if !CheckIntByRange(params, DIM, 1, math.MaxInt) {
			return fmt.Errorf("failed to check vector dimension, should be larger than 0 and smaller than math.MaxInt")
		}
	}
	return nil
}

// CheckValidDataType check whether the field data type is supported for the index type
func (c baseChecker) CheckValidDataType(field *schemapb.FieldSchema) error {
	return nil
}

func (c baseChecker) SetDefaultMetricTypeIfNotExist(m map[string]string, dType schemapb.DataType) {}

func (c baseChecker) StaticCheck(params map[string]string) error {
	return errors.New("unsupported index type")
}

func newBaseChecker() IndexChecker {
	return &baseChecker{}
}
