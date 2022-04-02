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

package typeutil

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/proto/schemapb"
	"go.uber.org/zap"
)

func GetMaxLengthOfVarLengthField(fieldSchema *schemapb.FieldSchema) (int, error) {
	maxLength := 0
	var err error

	paramsMap := make(map[string]string)
	for _, p := range fieldSchema.TypeParams {
		paramsMap[p.Key] = p.Value
	}

	maxLengthPerRowKey := "max_length_per_row"

	switch fieldSchema.DataType {
	case schemapb.DataType_VarChar:
		maxLengthPerRowValue, ok := paramsMap[maxLengthPerRowKey]
		if !ok {
			return 0, fmt.Errorf("the max_length_per_row was not specified, field type is %s", fieldSchema.DataType.String())
		}
		maxLength, err = strconv.Atoi(maxLengthPerRowValue)
		if err != nil {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("field %s is not a variable-length type", fieldSchema.DataType.String())
	}

	return maxLength, nil
}

// EstimateSizePerRecord returns the estimate size of a record in a collection
func EstimateSizePerRecord(schema *schemapb.CollectionSchema) (int, error) {
	res := 0
	for _, fs := range schema.Fields {
		switch fs.DataType {
		case schemapb.DataType_Bool, schemapb.DataType_Int8:
			res++
		case schemapb.DataType_Int16:
			res += 2
		case schemapb.DataType_Int32, schemapb.DataType_Float:
			res += 4
		case schemapb.DataType_Int64, schemapb.DataType_Double:
			res += 8
		case schemapb.DataType_VarChar:
			maxLengthPerRow, err := GetMaxLengthOfVarLengthField(fs)
			if err != nil {
				return 0, err
			}
			res += maxLengthPerRow
		case schemapb.DataType_BinaryVector:
			for _, kv := range fs.TypeParams {
				if kv.Key == "dim" {
					v, err := strconv.Atoi(kv.Value)
					if err != nil {
						return -1, err
					}
					res += v / 8
					break
				}
			}
		case schemapb.DataType_FloatVector:
			for _, kv := range fs.TypeParams {
				if kv.Key == "dim" {
					v, err := strconv.Atoi(kv.Value)
					if err != nil {
						return -1, err
					}
					res += v * 4
					break
				}
			}
		}
	}
	return res, nil
}

func EstimateEntitySize(fieldsData []*schemapb.FieldData, rowOffset int) (int, error) {
	res := 0
	for _, fs := range fieldsData {
		switch fs.GetType() {
		case schemapb.DataType_Bool, schemapb.DataType_Int8:
			res++
		case schemapb.DataType_Int16:
			res += 2
		case schemapb.DataType_Int32, schemapb.DataType_Float:
			res += 4
		case schemapb.DataType_Int64, schemapb.DataType_Double:
			res += 8
		case schemapb.DataType_VarChar:
			if rowOffset >= len(fs.GetScalars().GetStringData().GetData()) {
				return 0, fmt.Errorf("offset out range of field datas")
			}
			//TODO:: check len(varChar) <= maxLengthPerRow
			res += len(fs.GetScalars().GetStringData().Data[rowOffset])
		case schemapb.DataType_BinaryVector:
			res += int(fs.GetVectors().GetDim())
		case schemapb.DataType_FloatVector:
			res += int(fs.GetVectors().GetDim() * 4)
		}
	}
	return res, nil
}

// SchemaHelper provides methods to get the schema of fields
type SchemaHelper struct {
	schema           *schemapb.CollectionSchema
	nameOffset       map[string]int
	idOffset         map[int64]int
	primaryKeyOffset int
}

// CreateSchemaHelper returns a new SchemaHelper object
func CreateSchemaHelper(schema *schemapb.CollectionSchema) (*SchemaHelper, error) {
	if schema == nil {
		return nil, errors.New("schema is nil")
	}
	schemaHelper := SchemaHelper{schema: schema, nameOffset: make(map[string]int), idOffset: make(map[int64]int), primaryKeyOffset: -1}
	for offset, field := range schema.Fields {
		if _, ok := schemaHelper.nameOffset[field.Name]; ok {
			return nil, errors.New("duplicated fieldName: " + field.Name)
		}
		if _, ok := schemaHelper.idOffset[field.FieldID]; ok {
			return nil, errors.New("duplicated fieldID: " + strconv.FormatInt(field.FieldID, 10))
		}
		schemaHelper.nameOffset[field.Name] = offset
		schemaHelper.idOffset[field.FieldID] = offset
		if field.IsPrimaryKey {
			if schemaHelper.primaryKeyOffset != -1 {
				return nil, errors.New("primary key is not unique")
			}
			schemaHelper.primaryKeyOffset = offset
		}
	}
	return &schemaHelper, nil
}

// GetPrimaryKeyField returns the schema of the primary key
func (helper *SchemaHelper) GetPrimaryKeyField() (*schemapb.FieldSchema, error) {
	if helper.primaryKeyOffset == -1 {
		return nil, fmt.Errorf("failed to get primary key field: no primary in schema")
	}
	return helper.schema.Fields[helper.primaryKeyOffset], nil
}

// GetFieldFromName is used to find the schema by field name
func (helper *SchemaHelper) GetFieldFromName(fieldName string) (*schemapb.FieldSchema, error) {
	offset, ok := helper.nameOffset[fieldName]
	if !ok {
		return nil, fmt.Errorf("failed to get field schema by name: fieldName(%s) not found", fieldName)
	}
	return helper.schema.Fields[offset], nil
}

// GetFieldFromID returns the schema of specified field
func (helper *SchemaHelper) GetFieldFromID(fieldID int64) (*schemapb.FieldSchema, error) {
	offset, ok := helper.idOffset[fieldID]
	if !ok {
		return nil, fmt.Errorf("fieldID(%d) not found", fieldID)
	}
	return helper.schema.Fields[offset], nil
}

// GetVectorDimFromID returns the dimension of specified field
func (helper *SchemaHelper) GetVectorDimFromID(fieldID int64) (int, error) {
	sch, err := helper.GetFieldFromID(fieldID)
	if err != nil {
		return 0, err
	}
	if !IsVectorType(sch.DataType) {
		return 0, fmt.Errorf("field type = %s not has dim", schemapb.DataType_name[int32(sch.DataType)])
	}
	for _, kv := range sch.TypeParams {
		if kv.Key == "dim" {
			dim, err := strconv.Atoi(kv.Value)
			if err != nil {
				return 0, err
			}
			return dim, nil
		}
	}
	return 0, fmt.Errorf("fieldID(%d) not has dim", fieldID)
}

// IsVectorType returns true if input is a vector type, otherwise false
func IsVectorType(dataType schemapb.DataType) bool {
	switch dataType {
	case schemapb.DataType_FloatVector, schemapb.DataType_BinaryVector:
		return true
	default:
		return false
	}
}

// IsIntegerType returns true if input is an integer type, otherwise false
func IsIntegerType(dataType schemapb.DataType) bool {
	switch dataType {
	case schemapb.DataType_Int8, schemapb.DataType_Int16,
		schemapb.DataType_Int32, schemapb.DataType_Int64:
		return true
	default:
		return false
	}
}

// IsFloatingType returns true if input is a floating type, otherwise false
func IsFloatingType(dataType schemapb.DataType) bool {
	switch dataType {
	case schemapb.DataType_Float, schemapb.DataType_Double:
		return true
	default:
		return false
	}
}

// IsBoolType returns true if input is a bool type, otherwise false
func IsBoolType(dataType schemapb.DataType) bool {
	switch dataType {
	case schemapb.DataType_Bool:
		return true
	default:
		return false
	}
}

// IsStringType returns true if input is a varChar type, otherwise false
func IsStringType(dataType schemapb.DataType) bool {
	switch dataType {
	case schemapb.DataType_VarChar:
		return true
	default:
		return false
	}
}

// AppendFieldData appends fields data of specified index from src to dst
func AppendFieldData(dst []*schemapb.FieldData, src []*schemapb.FieldData, idx int64) {
	for i, fieldData := range src {
		switch fieldType := fieldData.Field.(type) {
		case *schemapb.FieldData_Scalars:
			if dst[i] == nil || dst[i].GetScalars() == nil {
				dst[i] = &schemapb.FieldData{
					Type:      fieldData.Type,
					FieldName: fieldData.FieldName,
					FieldId:   fieldData.FieldId,
					Field: &schemapb.FieldData_Scalars{
						Scalars: &schemapb.ScalarField{},
					},
				}
			}
			dstScalar := dst[i].GetScalars()
			switch srcScalar := fieldType.Scalars.Data.(type) {
			case *schemapb.ScalarField_BoolData:
				if dstScalar.GetBoolData() == nil {
					dstScalar.Data = &schemapb.ScalarField_BoolData{
						BoolData: &schemapb.BoolArray{
							Data: []bool{srcScalar.BoolData.Data[idx]},
						},
					}
				} else {
					dstScalar.GetBoolData().Data = append(dstScalar.GetBoolData().Data, srcScalar.BoolData.Data[idx])
				}
			case *schemapb.ScalarField_IntData:
				if dstScalar.GetIntData() == nil {
					dstScalar.Data = &schemapb.ScalarField_IntData{
						IntData: &schemapb.IntArray{
							Data: []int32{srcScalar.IntData.Data[idx]},
						},
					}
				} else {
					dstScalar.GetIntData().Data = append(dstScalar.GetIntData().Data, srcScalar.IntData.Data[idx])
				}
			case *schemapb.ScalarField_LongData:
				if dstScalar.GetLongData() == nil {
					dstScalar.Data = &schemapb.ScalarField_LongData{
						LongData: &schemapb.LongArray{
							Data: []int64{srcScalar.LongData.Data[idx]},
						},
					}
				} else {
					dstScalar.GetLongData().Data = append(dstScalar.GetLongData().Data, srcScalar.LongData.Data[idx])
				}
			case *schemapb.ScalarField_FloatData:
				if dstScalar.GetFloatData() == nil {
					dstScalar.Data = &schemapb.ScalarField_FloatData{
						FloatData: &schemapb.FloatArray{
							Data: []float32{srcScalar.FloatData.Data[idx]},
						},
					}
				} else {
					dstScalar.GetFloatData().Data = append(dstScalar.GetFloatData().Data, srcScalar.FloatData.Data[idx])
				}
			case *schemapb.ScalarField_DoubleData:
				if dstScalar.GetDoubleData() == nil {
					dstScalar.Data = &schemapb.ScalarField_DoubleData{
						DoubleData: &schemapb.DoubleArray{
							Data: []float64{srcScalar.DoubleData.Data[idx]},
						},
					}
				} else {
					dstScalar.GetDoubleData().Data = append(dstScalar.GetDoubleData().Data, srcScalar.DoubleData.Data[idx])
				}
			case *schemapb.ScalarField_StringData:
				if dstScalar.GetStringData() == nil {
					dstScalar.Data = &schemapb.ScalarField_StringData{
						StringData: &schemapb.StringArray{
							Data: []string{srcScalar.StringData.Data[idx]},
						},
					}
				} else {
					dstScalar.GetStringData().Data = append(dstScalar.GetStringData().Data, srcScalar.StringData.Data[idx])
				}
			default:
				log.Error("Not supported field type", zap.String("field type", fieldData.Type.String()))
			}
		case *schemapb.FieldData_Vectors:
			dim := fieldType.Vectors.Dim
			if dst[i] == nil || dst[i].GetVectors() == nil {
				dst[i] = &schemapb.FieldData{
					Type:      fieldData.Type,
					FieldName: fieldData.FieldName,
					FieldId:   fieldData.FieldId,
					Field: &schemapb.FieldData_Vectors{
						Vectors: &schemapb.VectorField{
							Dim: dim,
						},
					},
				}
			}
			dstVector := dst[i].GetVectors()
			switch srcVector := fieldType.Vectors.Data.(type) {
			case *schemapb.VectorField_BinaryVector:
				if dstVector.GetBinaryVector() == nil {
					srcToCopy := srcVector.BinaryVector[idx*(dim/8) : (idx+1)*(dim/8)]
					dstVector.Data = &schemapb.VectorField_BinaryVector{
						BinaryVector: make([]byte, len(srcToCopy)),
					}
					copy(dstVector.Data.(*schemapb.VectorField_BinaryVector).BinaryVector, srcToCopy)
				} else {
					dstBinaryVector := dstVector.Data.(*schemapb.VectorField_BinaryVector)
					dstBinaryVector.BinaryVector = append(dstBinaryVector.BinaryVector, srcVector.BinaryVector[idx*(dim/8):(idx+1)*(dim/8)]...)
				}
			case *schemapb.VectorField_FloatVector:
				if dstVector.GetFloatVector() == nil {
					srcToCopy := srcVector.FloatVector.Data[idx*dim : (idx+1)*dim]
					dstVector.Data = &schemapb.VectorField_FloatVector{
						FloatVector: &schemapb.FloatArray{
							Data: make([]float32, len(srcToCopy)),
						},
					}
					copy(dstVector.Data.(*schemapb.VectorField_FloatVector).FloatVector.Data, srcToCopy)
				} else {
					dstVector.GetFloatVector().Data = append(dstVector.GetFloatVector().Data, srcVector.FloatVector.Data[idx*dim:(idx+1)*dim]...)
				}
			default:
				log.Error("Not supported field type", zap.String("field type", fieldData.Type.String()))
			}
		}
	}
}

// GetPrimaryFieldSchema get primary field schema from collection schema
func GetPrimaryFieldSchema(schema *schemapb.CollectionSchema) (*schemapb.FieldSchema, error) {
	for _, fieldSchema := range schema.Fields {
		if fieldSchema.IsPrimaryKey {
			return fieldSchema, nil
		}
	}

	return nil, errors.New("primary field is not found")
}

// GetPrimaryFieldData get primary field data from all field data inserted from sdk
func GetPrimaryFieldData(datas []*schemapb.FieldData, primaryFieldSchema *schemapb.FieldSchema) (*schemapb.FieldData, error) {
	primaryFieldName := primaryFieldSchema.Name

	var primaryFieldData *schemapb.FieldData
	for _, field := range datas {
		if field.FieldName == primaryFieldName {
			if primaryFieldSchema.AutoID {
				return nil, fmt.Errorf("autoID field %v does not require data", primaryFieldName)
			}
			primaryFieldData = field
		}
	}

	if primaryFieldData == nil {
		return nil, fmt.Errorf("can't find data for primary field %v", primaryFieldName)
	}

	return primaryFieldData, nil
}

func AppendIDs(dst *schemapb.IDs, src *schemapb.IDs, idx int) {
	switch src.IdField.(type) {
	case *schemapb.IDs_IntId:
		if dst.GetIdField() == nil {
			dst.IdField = &schemapb.IDs_IntId{
				IntId: &schemapb.LongArray{
					Data: []int64{src.GetIntId().Data[idx]},
				},
			}
		} else {
			dst.GetIntId().Data = append(dst.GetIntId().Data, src.GetIntId().Data[idx])
		}
	case *schemapb.IDs_StrId:
		if dst.GetIdField() == nil {
			dst.IdField = &schemapb.IDs_StrId{
				StrId: &schemapb.StringArray{
					Data: []string{src.GetStrId().Data[idx]},
				},
			}
		} else {
			dst.GetStrId().Data = append(dst.GetStrId().Data, src.GetStrId().Data[idx])
		}
	default:
		//TODO
	}
}

func GetSizeOfIDs(data *schemapb.IDs) int {
	result := 0
	if data.IdField == nil {
		return result
	}

	switch data.GetIdField().(type) {
	case *schemapb.IDs_IntId:
		result = len(data.GetIntId().GetData())
	case *schemapb.IDs_StrId:
		result = len(data.GetStrId().GetData())
	default:
		//TODO::
	}

	return result
}

func IsPrimaryFieldType(dataType schemapb.DataType) bool {
	if dataType == schemapb.DataType_Int64 || dataType == schemapb.DataType_VarChar {
		return true
	}

	return false
}
