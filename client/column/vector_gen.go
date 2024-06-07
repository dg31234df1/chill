// Code generated by go generate; DO NOT EDIT
// This file is generated by go generated
package column

import (
	"fmt"

	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	"github.com/milvus-io/milvus/client/v2/entity"

	"github.com/cockroachdb/errors"
)

// ColumnBinaryVector generated columns type for BinaryVector
type ColumnBinaryVector struct {
	ColumnBase
	name   string
	dim    int
	values [][]byte
}

// Name returns column name
func (c *ColumnBinaryVector) Name() string {
	return c.name
}

// Type returns column entity.FieldType
func (c *ColumnBinaryVector) Type() entity.FieldType {
	return entity.FieldTypeBinaryVector
}

// Len returns column data length
func (c *ColumnBinaryVector) Len() int {
	return len(c.values)
}

// Dim returns vector dimension
func (c *ColumnBinaryVector) Dim() int {
	return c.dim
}

// Get returns values at index as interface{}.
func (c *ColumnBinaryVector) Get(idx int) (interface{}, error) {
	if idx < 0 || idx >= c.Len() {
		return nil, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// AppendValue append value into column
func (c *ColumnBinaryVector) AppendValue(i interface{}) error {
	v, ok := i.([]byte)
	if !ok {
		return fmt.Errorf("invalid type, expected []byte, got %T", i)
	}
	c.values = append(c.values, v)

	return nil
}

// Data returns column data
func (c *ColumnBinaryVector) Data() [][]byte {
	return c.values
}

// FieldData return column data mapped to schemapb.FieldData
func (c *ColumnBinaryVector) FieldData() *schemapb.FieldData {
	fd := &schemapb.FieldData{
		Type:      schemapb.DataType_BinaryVector,
		FieldName: c.name,
	}

	data := make([]byte, 0, len(c.values)*c.dim)

	for _, vector := range c.values {
		data = append(data, vector...)
	}

	fd.Field = &schemapb.FieldData_Vectors{
		Vectors: &schemapb.VectorField{
			Dim: int64(c.dim),

			Data: &schemapb.VectorField_BinaryVector{
				BinaryVector: data,
			},
		},
	}
	return fd
}

// NewColumnBinaryVector auto generated constructor
func NewColumnBinaryVector(name string, dim int, values [][]byte) *ColumnBinaryVector {
	return &ColumnBinaryVector{
		name:   name,
		dim:    dim,
		values: values,
	}
}

// ColumnFloatVector generated columns type for FloatVector
type ColumnFloatVector struct {
	ColumnBase
	name   string
	dim    int
	values [][]float32
}

// Name returns column name
func (c *ColumnFloatVector) Name() string {
	return c.name
}

// Type returns column entity.FieldType
func (c *ColumnFloatVector) Type() entity.FieldType {
	return entity.FieldTypeFloatVector
}

// Len returns column data length
func (c *ColumnFloatVector) Len() int {
	return len(c.values)
}

// Dim returns vector dimension
func (c *ColumnFloatVector) Dim() int {
	return c.dim
}

// Get returns values at index as interface{}.
func (c *ColumnFloatVector) Get(idx int) (interface{}, error) {
	if idx < 0 || idx >= c.Len() {
		return nil, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// AppendValue append value into column
func (c *ColumnFloatVector) AppendValue(i interface{}) error {
	v, ok := i.([]float32)
	if !ok {
		return fmt.Errorf("invalid type, expected []float32, got %T", i)
	}
	c.values = append(c.values, v)

	return nil
}

// Data returns column data
func (c *ColumnFloatVector) Data() [][]float32 {
	return c.values
}

// FieldData return column data mapped to schemapb.FieldData
func (c *ColumnFloatVector) FieldData() *schemapb.FieldData {
	fd := &schemapb.FieldData{
		Type:      schemapb.DataType_FloatVector,
		FieldName: c.name,
	}

	data := make([]float32, 0, len(c.values)*c.dim)

	for _, vector := range c.values {
		data = append(data, vector...)
	}

	fd.Field = &schemapb.FieldData_Vectors{
		Vectors: &schemapb.VectorField{
			Dim: int64(c.dim),

			Data: &schemapb.VectorField_FloatVector{
				FloatVector: &schemapb.FloatArray{
					Data: data,
				},
			},
		},
	}
	return fd
}

// NewColumnFloatVector auto generated constructor
func NewColumnFloatVector(name string, dim int, values [][]float32) *ColumnFloatVector {
	return &ColumnFloatVector{
		name:   name,
		dim:    dim,
		values: values,
	}
}

// ColumnFloat16Vector generated columns type for Float16Vector
type ColumnFloat16Vector struct {
	ColumnBase
	name   string
	dim    int
	values [][]byte
}

// Name returns column name
func (c *ColumnFloat16Vector) Name() string {
	return c.name
}

// Type returns column entity.FieldType
func (c *ColumnFloat16Vector) Type() entity.FieldType {
	return entity.FieldTypeFloat16Vector
}

// Len returns column data length
func (c *ColumnFloat16Vector) Len() int {
	return len(c.values)
}

// Dim returns vector dimension
func (c *ColumnFloat16Vector) Dim() int {
	return c.dim
}

// Get returns values at index as interface{}.
func (c *ColumnFloat16Vector) Get(idx int) (interface{}, error) {
	if idx < 0 || idx >= c.Len() {
		return nil, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// AppendValue append value into column
func (c *ColumnFloat16Vector) AppendValue(i interface{}) error {
	v, ok := i.([]byte)
	if !ok {
		return fmt.Errorf("invalid type, expected []byte, got %T", i)
	}
	c.values = append(c.values, v)

	return nil
}

// Data returns column data
func (c *ColumnFloat16Vector) Data() [][]byte {
	return c.values
}

// FieldData return column data mapped to schemapb.FieldData
func (c *ColumnFloat16Vector) FieldData() *schemapb.FieldData {
	fd := &schemapb.FieldData{
		Type:      schemapb.DataType_Float16Vector,
		FieldName: c.name,
	}

	data := make([]byte, 0, len(c.values)*c.dim*2)

	for _, vector := range c.values {
		data = append(data, vector...)
	}

	fd.Field = &schemapb.FieldData_Vectors{
		Vectors: &schemapb.VectorField{
			Dim: int64(c.dim),

			Data: &schemapb.VectorField_Float16Vector{
				Float16Vector: data,
			},
		},
	}
	return fd
}

// NewColumnFloat16Vector auto generated constructor
func NewColumnFloat16Vector(name string, dim int, values [][]byte) *ColumnFloat16Vector {
	return &ColumnFloat16Vector{
		name:   name,
		dim:    dim,
		values: values,
	}
}

// ColumnBFloat16Vector generated columns type for BFloat16Vector
type ColumnBFloat16Vector struct {
	ColumnBase
	name   string
	dim    int
	values [][]byte
}

// Name returns column name
func (c *ColumnBFloat16Vector) Name() string {
	return c.name
}

// Type returns column entity.FieldType
func (c *ColumnBFloat16Vector) Type() entity.FieldType {
	return entity.FieldTypeBFloat16Vector
}

// Len returns column data length
func (c *ColumnBFloat16Vector) Len() int {
	return len(c.values)
}

// Dim returns vector dimension
func (c *ColumnBFloat16Vector) Dim() int {
	return c.dim
}

// Get returns values at index as interface{}.
func (c *ColumnBFloat16Vector) Get(idx int) (interface{}, error) {
	if idx < 0 || idx >= c.Len() {
		return nil, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// AppendValue append value into column
func (c *ColumnBFloat16Vector) AppendValue(i interface{}) error {
	v, ok := i.([]byte)
	if !ok {
		return fmt.Errorf("invalid type, expected []byte, got %T", i)
	}
	c.values = append(c.values, v)

	return nil
}

// Data returns column data
func (c *ColumnBFloat16Vector) Data() [][]byte {
	return c.values
}

// FieldData return column data mapped to schemapb.FieldData
func (c *ColumnBFloat16Vector) FieldData() *schemapb.FieldData {
	fd := &schemapb.FieldData{
		Type:      schemapb.DataType_BFloat16Vector,
		FieldName: c.name,
	}

	data := make([]byte, 0, len(c.values)*c.dim*2)

	for _, vector := range c.values {
		data = append(data, vector...)
	}

	fd.Field = &schemapb.FieldData_Vectors{
		Vectors: &schemapb.VectorField{
			Dim: int64(c.dim),

			Data: &schemapb.VectorField_Bfloat16Vector{
				Bfloat16Vector: data,
			},
		},
	}
	return fd
}

// NewColumnBFloat16Vector auto generated constructor
func NewColumnBFloat16Vector(name string, dim int, values [][]byte) *ColumnBFloat16Vector {
	return &ColumnBFloat16Vector{
		name:   name,
		dim:    dim,
		values: values,
	}
}
