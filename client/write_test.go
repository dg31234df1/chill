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

package client

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"
	"github.com/milvus-io/milvus/client/v2/entity"
	"github.com/milvus-io/milvus/pkg/util/merr"
	"github.com/samber/lo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type WriteSuite struct {
	MockSuiteBase

	schema    *entity.Schema
	schemaDyn *entity.Schema
}

func (s *WriteSuite) SetupSuite() {
	s.MockSuiteBase.SetupSuite()
	s.schema = entity.NewSchema().
		WithField(entity.NewField().WithName("id").WithDataType(entity.FieldTypeInt64).WithIsPrimaryKey(true)).
		WithField(entity.NewField().WithName("vector").WithDataType(entity.FieldTypeFloatVector).WithDim(128))

	s.schemaDyn = entity.NewSchema().WithDynamicFieldEnabled(true).
		WithField(entity.NewField().WithName("id").WithDataType(entity.FieldTypeInt64).WithIsPrimaryKey(true)).
		WithField(entity.NewField().WithName("vector").WithDataType(entity.FieldTypeFloatVector).WithDim(128))
}

func (s *WriteSuite) TestInsert() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s.Run("success", func() {
		collName := fmt.Sprintf("coll_%s", s.randString(6))
		partName := fmt.Sprintf("part_%s", s.randString(6))
		s.setupCache(collName, s.schema)

		s.mock.EXPECT().Insert(mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, ir *milvuspb.InsertRequest) (*milvuspb.MutationResult, error) {
			s.Equal(collName, ir.GetCollectionName())
			s.Equal(partName, ir.GetPartitionName())
			s.Require().Len(ir.GetFieldsData(), 2)
			s.EqualValues(3, ir.GetNumRows())
			return &milvuspb.MutationResult{
				Status: merr.Success(),
			}, nil
		}).Once()

		err := s.client.Insert(ctx, NewColumnBasedInsertOption(collName).
			WithFloatVectorColumn("vector", 128, lo.RepeatBy(3, func(i int) []float32 {
				return lo.RepeatBy(128, func(i int) float32 { return rand.Float32() })
			})).
			WithInt64Column("id", []int64{1, 2, 3}).WithPartition(partName))
		s.NoError(err)
	})

	s.Run("dynamic_schema", func() {
		collName := fmt.Sprintf("coll_%s", s.randString(6))
		partName := fmt.Sprintf("part_%s", s.randString(6))
		s.setupCache(collName, s.schemaDyn)

		s.mock.EXPECT().Insert(mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, ir *milvuspb.InsertRequest) (*milvuspb.MutationResult, error) {
			s.Equal(collName, ir.GetCollectionName())
			s.Equal(partName, ir.GetPartitionName())
			s.Require().Len(ir.GetFieldsData(), 3)
			s.EqualValues(3, ir.GetNumRows())
			return &milvuspb.MutationResult{
				Status: merr.Success(),
			}, nil
		}).Once()

		err := s.client.Insert(ctx, NewColumnBasedInsertOption(collName).
			WithFloatVectorColumn("vector", 128, lo.RepeatBy(3, func(i int) []float32 {
				return lo.RepeatBy(128, func(i int) float32 { return rand.Float32() })
			})).
			WithVarcharColumn("extra", []string{"a", "b", "c"}).
			WithInt64Column("id", []int64{1, 2, 3}).WithPartition(partName))
		s.NoError(err)
	})

	s.Run("bad_input", func() {
		collName := fmt.Sprintf("coll_%s", s.randString(6))
		s.setupCache(collName, s.schema)

		type badCase struct {
			tag   string
			input InsertOption
		}

		cases := []badCase{
			{
				tag:   "missing_column",
				input: NewColumnBasedInsertOption(collName).WithInt64Column("id", []int64{1}),
			},
			{
				tag: "row_count_not_match",
				input: NewColumnBasedInsertOption(collName).WithInt64Column("id", []int64{1}).
					WithFloatVectorColumn("vector", 128, lo.RepeatBy(3, func(i int) []float32 {
						return lo.RepeatBy(128, func(i int) float32 { return rand.Float32() })
					})),
			},
			{
				tag: "duplicated_columns",
				input: NewColumnBasedInsertOption(collName).
					WithInt64Column("id", []int64{1}).
					WithInt64Column("id", []int64{2}).
					WithFloatVectorColumn("vector", 128, lo.RepeatBy(1, func(i int) []float32 {
						return lo.RepeatBy(128, func(i int) float32 { return rand.Float32() })
					})),
			},
			{
				tag: "different_data_type",
				input: NewColumnBasedInsertOption(collName).
					WithVarcharColumn("id", []string{"1"}).
					WithFloatVectorColumn("vector", 128, lo.RepeatBy(1, func(i int) []float32 {
						return lo.RepeatBy(128, func(i int) float32 { return rand.Float32() })
					})),
			},
		}

		for _, tc := range cases {
			s.Run(tc.tag, func() {
				err := s.client.Insert(ctx, tc.input)
				s.Error(err)
			})
		}
	})

	s.Run("failure", func() {
		collName := fmt.Sprintf("coll_%s", s.randString(6))
		s.setupCache(collName, s.schema)

		s.mock.EXPECT().Insert(mock.Anything, mock.Anything).Return(nil, merr.WrapErrServiceInternal("mocked")).Once()

		err := s.client.Insert(ctx, NewColumnBasedInsertOption(collName).
			WithFloatVectorColumn("vector", 128, lo.RepeatBy(3, func(i int) []float32 {
				return lo.RepeatBy(128, func(i int) float32 { return rand.Float32() })
			})).
			WithInt64Column("id", []int64{1, 2, 3}))
		s.Error(err)
	})
}

func (s *WriteSuite) TestUpsert() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s.Run("success", func() {
		collName := fmt.Sprintf("coll_%s", s.randString(6))
		partName := fmt.Sprintf("part_%s", s.randString(6))
		s.setupCache(collName, s.schema)

		s.mock.EXPECT().Upsert(mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, ur *milvuspb.UpsertRequest) (*milvuspb.MutationResult, error) {
			s.Equal(collName, ur.GetCollectionName())
			s.Equal(partName, ur.GetPartitionName())
			s.Require().Len(ur.GetFieldsData(), 2)
			s.EqualValues(3, ur.GetNumRows())
			return &milvuspb.MutationResult{
				Status: merr.Success(),
			}, nil
		}).Once()

		err := s.client.Upsert(ctx, NewColumnBasedInsertOption(collName).
			WithFloatVectorColumn("vector", 128, lo.RepeatBy(3, func(i int) []float32 {
				return lo.RepeatBy(128, func(i int) float32 { return rand.Float32() })
			})).
			WithInt64Column("id", []int64{1, 2, 3}).WithPartition(partName))
		s.NoError(err)
	})

	s.Run("dynamic_schema", func() {
		collName := fmt.Sprintf("coll_%s", s.randString(6))
		partName := fmt.Sprintf("part_%s", s.randString(6))
		s.setupCache(collName, s.schemaDyn)

		s.mock.EXPECT().Upsert(mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, ur *milvuspb.UpsertRequest) (*milvuspb.MutationResult, error) {
			s.Equal(collName, ur.GetCollectionName())
			s.Equal(partName, ur.GetPartitionName())
			s.Require().Len(ur.GetFieldsData(), 3)
			s.EqualValues(3, ur.GetNumRows())
			return &milvuspb.MutationResult{
				Status: merr.Success(),
			}, nil
		}).Once()

		err := s.client.Upsert(ctx, NewColumnBasedInsertOption(collName).
			WithFloatVectorColumn("vector", 128, lo.RepeatBy(3, func(i int) []float32 {
				return lo.RepeatBy(128, func(i int) float32 { return rand.Float32() })
			})).
			WithVarcharColumn("extra", []string{"a", "b", "c"}).
			WithInt64Column("id", []int64{1, 2, 3}).WithPartition(partName))
		s.NoError(err)
	})

	s.Run("bad_input", func() {
		collName := fmt.Sprintf("coll_%s", s.randString(6))
		s.setupCache(collName, s.schema)

		type badCase struct {
			tag   string
			input UpsertOption
		}

		cases := []badCase{
			{
				tag:   "missing_column",
				input: NewColumnBasedInsertOption(collName).WithInt64Column("id", []int64{1}),
			},
			{
				tag: "row_count_not_match",
				input: NewColumnBasedInsertOption(collName).WithInt64Column("id", []int64{1}).
					WithFloatVectorColumn("vector", 128, lo.RepeatBy(3, func(i int) []float32 {
						return lo.RepeatBy(128, func(i int) float32 { return rand.Float32() })
					})),
			},
			{
				tag: "duplicated_columns",
				input: NewColumnBasedInsertOption(collName).
					WithInt64Column("id", []int64{1}).
					WithInt64Column("id", []int64{2}).
					WithFloatVectorColumn("vector", 128, lo.RepeatBy(1, func(i int) []float32 {
						return lo.RepeatBy(128, func(i int) float32 { return rand.Float32() })
					})),
			},
			{
				tag: "different_data_type",
				input: NewColumnBasedInsertOption(collName).
					WithVarcharColumn("id", []string{"1"}).
					WithFloatVectorColumn("vector", 128, lo.RepeatBy(1, func(i int) []float32 {
						return lo.RepeatBy(128, func(i int) float32 { return rand.Float32() })
					})),
			},
		}

		for _, tc := range cases {
			s.Run(tc.tag, func() {
				err := s.client.Upsert(ctx, tc.input)
				s.Error(err)
			})
		}
	})

	s.Run("failure", func() {
		collName := fmt.Sprintf("coll_%s", s.randString(6))
		s.setupCache(collName, s.schema)

		s.mock.EXPECT().Upsert(mock.Anything, mock.Anything).Return(nil, merr.WrapErrServiceInternal("mocked")).Once()

		err := s.client.Upsert(ctx, NewColumnBasedInsertOption(collName).
			WithFloatVectorColumn("vector", 128, lo.RepeatBy(3, func(i int) []float32 {
				return lo.RepeatBy(128, func(i int) float32 { return rand.Float32() })
			})).
			WithInt64Column("id", []int64{1, 2, 3}))
		s.Error(err)
	})
}

func (s *WriteSuite) TestDelete() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s.Run("success", func() {
		collName := fmt.Sprintf("coll_%s", s.randString(6))
		partName := fmt.Sprintf("part_%s", s.randString(6))

		type testCase struct {
			tag        string
			input      DeleteOption
			expectExpr string
		}

		cases := []testCase{
			{
				tag:        "raw_expr",
				input:      NewDeleteOption(collName).WithPartition(partName).WithExpr("id > 100"),
				expectExpr: "id > 100",
			},
			{
				tag:        "int_ids",
				input:      NewDeleteOption(collName).WithPartition(partName).WithInt64IDs("id", []int64{1, 2, 3}),
				expectExpr: "id in [1,2,3]",
			},
			{
				tag:        "str_ids",
				input:      NewDeleteOption(collName).WithPartition(partName).WithStringIDs("id", []string{"a", "b", "c"}),
				expectExpr: `id in ["a","b","c"]`,
			},
		}

		for _, tc := range cases {
			s.Run(tc.tag, func() {
				s.mock.EXPECT().Delete(mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, dr *milvuspb.DeleteRequest) (*milvuspb.MutationResult, error) {
					s.Equal(collName, dr.GetCollectionName())
					s.Equal(partName, dr.GetPartitionName())
					s.Equal(tc.expectExpr, dr.GetExpr())
					return &milvuspb.MutationResult{
						Status: merr.Success(),
					}, nil
				}).Once()
				err := s.client.Delete(ctx, tc.input)
				s.NoError(err)
			})
		}
	})
}

func TestWrite(t *testing.T) {
	suite.Run(t, new(WriteSuite))
}
