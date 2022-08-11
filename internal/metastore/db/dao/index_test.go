package dao

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/milvus-io/milvus/internal/metastore/db/dbmodel"
	"github.com/stretchr/testify/assert"
)

func TestIndex_Get(t *testing.T) {
	var indexes = []*dbmodel.Index{
		{
			TenantID:     "",
			FieldID:      fieldID1,
			CollectionID: collID1,
			IndexID:      indexID1,
			IndexName:    "test_index_1",
			IndexParams:  "",
		},
	}

	// expectation
	mock.ExpectQuery("SELECT * FROM `indexes` WHERE tenant_id = ? AND collection_id = ? AND is_deleted = false").
		WithArgs(tenantID, collID1).
		WillReturnRows(
			sqlmock.NewRows([]string{"tenant_id", "field_id", "collection_id", "index_id", "index_name", "index_params"}).
				AddRow(indexes[0].TenantID, indexes[0].FieldID, indexes[0].CollectionID, indexes[0].IndexID, indexes[0].IndexName, indexes[0].IndexParams))

	// actual
	res, err := indexTestDb.Get(tenantID, collID1)
	assert.Nil(t, err)
	assert.Equal(t, indexes, res)
}

func TestIndex_Get_Error(t *testing.T) {
	// expectation
	mock.ExpectQuery("SELECT * FROM `indexes` WHERE tenant_id = ? AND collection_id = ? AND is_deleted = false").
		WithArgs(tenantID, collID1).
		WillReturnError(errors.New("test error"))

	// actual
	res, err := indexTestDb.Get(tenantID, collID1)
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestIndex_List(t *testing.T) {
	var indexResults = []*dbmodel.IndexResult{
		{
			FieldID:        fieldID1,
			CollectionID:   collID1,
			IndexID:        indexID1,
			IndexName:      "test_index_1",
			IndexParams:    "",
			SegmentID:      segmentID1,
			PartitionID:    partitionID1,
			EnableIndex:    false,
			IndexBuildID:   indexBuildID1,
			IndexSize:      1024,
			IndexFilePaths: "",
		},
	}

	// expectation
	mock.ExpectQuery("SELECT indexes.field_id AS field_id, indexes.collection_id AS collection_id, indexes.index_id AS index_id, indexes.index_name AS index_name, indexes.index_params AS index_params, segment_indexes.segment_id AS segment_id, segment_indexes.partition_id AS partition_id, segment_indexes.enable_index AS enable_index, segment_indexes.index_build_id AS index_build_id, segment_indexes.index_size AS index_size, segment_indexes.index_file_paths AS index_file_paths FROM `indexes` LEFT JOIN segment_indexes ON indexes.index_id = segment_indexes.index_id AND indexes.tenant_id = segment_indexes.tenant_id AND segment_indexes.tenant_id = ? AND segment_indexes.is_deleted = false WHERE indexes.is_deleted = false AND indexes.tenant_id = ?").
		WithArgs(tenantID, tenantID).
		WillReturnRows(
			sqlmock.NewRows([]string{"field_id", "collection_id", "index_id", "index_name", "index_params", "segment_id", "partition_id", "enable_index", "index_build_id", "index_size", "index_file_paths"}).
				AddRow(indexResults[0].FieldID, indexResults[0].CollectionID, indexResults[0].IndexID, indexResults[0].IndexName, indexResults[0].IndexParams, indexResults[0].SegmentID, indexResults[0].PartitionID, indexResults[0].EnableIndex, indexResults[0].IndexBuildID, indexResults[0].IndexSize, indexResults[0].IndexFilePaths))

	// actual
	res, err := indexTestDb.List(tenantID)
	assert.Nil(t, err)
	assert.Equal(t, indexResults, res)
}

func TestIndex_List_Error(t *testing.T) {
	// expectation
	mock.ExpectQuery("SELECT indexes.field_id AS field_id, indexes.collection_id AS collection_id, indexes.index_id AS index_id, indexes.index_name AS index_name, indexes.index_params AS index_params, segment_indexes.segment_id AS segment_id, segment_indexes.partition_id AS partition_id, segment_indexes.enable_index AS enable_index, segment_indexes.index_build_id AS index_build_id, segment_indexes.index_size AS index_size, segment_indexes.index_file_paths AS index_file_paths FROM `indexes` LEFT JOIN segment_indexes ON indexes.index_id = segment_indexes.index_id AND indexes.tenant_id = segment_indexes.tenant_id AND segment_indexes.tenant_id = ? AND segment_indexes.is_deleted = false WHERE indexes.is_deleted = false AND indexes.tenant_id = ?").
		WithArgs(tenantID, tenantID).
		WillReturnError(errors.New("test error"))

	// actual
	res, err := indexTestDb.List(tenantID)
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestIndex_Insert(t *testing.T) {
	var indexes = []*dbmodel.Index{
		{
			TenantID:     tenantID,
			FieldID:      fieldID1,
			CollectionID: collID1,
			IndexID:      indexID1,
			IndexName:    "test_index_1",
			IndexParams:  "",
			CreateTime:   uint64(1011),
			IsDeleted:    false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	// expectation
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `indexes` (`tenant_id`,`field_id`,`collection_id`,`index_id`,`index_name`,`index_params`,`create_time`,`is_deleted`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?)").
		WithArgs(indexes[0].TenantID, indexes[0].FieldID, indexes[0].CollectionID, indexes[0].IndexID, indexes[0].IndexName, indexes[0].IndexParams, indexes[0].CreateTime, indexes[0].IsDeleted, indexes[0].CreatedAt, indexes[0].UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// actual
	err := indexTestDb.Insert(indexes)
	assert.Nil(t, err)
}

func TestIndex_Insert_Error(t *testing.T) {
	var indexes = []*dbmodel.Index{
		{
			TenantID:     tenantID,
			FieldID:      fieldID1,
			CollectionID: collID1,
			IndexID:      indexID1,
			IndexName:    "test_index_1",
			IndexParams:  "",
			CreateTime:   uint64(1011),
			IsDeleted:    false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	// expectation
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `indexes` (`tenant_id`,`field_id`,`collection_id`,`index_id`,`index_name`,`index_params`,`create_time`,`is_deleted`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?)").
		WithArgs(indexes[0].TenantID, indexes[0].FieldID, indexes[0].CollectionID, indexes[0].IndexID, indexes[0].IndexName, indexes[0].IndexParams, indexes[0].CreateTime, indexes[0].IsDeleted, indexes[0].CreatedAt, indexes[0].UpdatedAt).
		WillReturnError(errors.New("test error"))
	mock.ExpectRollback()

	// actual
	err := indexTestDb.Insert(indexes)
	assert.Error(t, err)
}

func TestIndex_Update(t *testing.T) {
	var index = &dbmodel.Index{
		TenantID:    tenantID,
		IndexName:   "test_index_name_1",
		IndexID:     indexID1,
		IndexParams: "",
		IsDeleted:   true,
		CreateTime:  uint64(1111),
	}

	// expectation
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `indexes` SET `create_time`=?,`is_deleted`=?,`updated_at`=? WHERE tenant_id = ? AND collection_id = ? AND index_id = ?").
		WithArgs(index.CreateTime, index.IsDeleted, AnyTime{}, index.TenantID, index.CollectionID, index.IndexID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// actual
	err := indexTestDb.Update(index)
	assert.Nil(t, err)
}

func TestIndex_Update_Error(t *testing.T) {
	var index = &dbmodel.Index{
		TenantID:    tenantID,
		IndexName:   "test_index_name_1",
		IndexID:     indexID1,
		IndexParams: "",
		IsDeleted:   true,
		CreateTime:  uint64(1111),
	}

	// expectation
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `indexes` SET `create_time`=?,`is_deleted`=?,`updated_at`=? WHERE tenant_id = ? AND collection_id = ? AND index_id = ?").
		WithArgs(index.CreateTime, index.IsDeleted, AnyTime{}, index.TenantID, index.CollectionID, index.IndexID).
		WillReturnError(errors.New("test error"))
	mock.ExpectRollback()

	// actual
	err := indexTestDb.Update(index)
	assert.Error(t, err)
}

func TestIndex_MarkDeletedByCollID(t *testing.T) {
	// expectation
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `indexes` SET `is_deleted`=?,`updated_at`=? WHERE tenant_id = ? AND collection_id = ?").
		WithArgs(true, AnyTime{}, tenantID, collID1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// actual
	err := indexTestDb.MarkDeletedByCollectionID(tenantID, collID1)
	assert.Nil(t, err)
}

func TestIndex_MarkDeletedByCollID_Error(t *testing.T) {
	// expectation
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `indexes` SET `is_deleted`=?,`updated_at`=? WHERE tenant_id = ? AND collection_id = ?").
		WithArgs(true, AnyTime{}, tenantID, collID1).
		WillReturnError(errors.New("test error"))
	mock.ExpectRollback()

	// actual
	err := indexTestDb.MarkDeletedByCollectionID(tenantID, collID1)
	assert.Error(t, err)
}

func TestIndex_MarkDeletedByIdxID(t *testing.T) {
	// expectation
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `indexes` SET `is_deleted`=?,`updated_at`=? WHERE tenant_id = ? AND index_id = ?").
		WithArgs(true, AnyTime{}, tenantID, indexID1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// actual
	err := indexTestDb.MarkDeletedByIndexID(tenantID, indexID1)
	assert.Nil(t, err)
}

func TestIndex_MarkDeletedByIdxID_Error(t *testing.T) {
	// expectation
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `indexes` SET `is_deleted`=?,`updated_at`=? WHERE tenant_id = ? AND index_id = ?").
		WithArgs(true, AnyTime{}, tenantID, indexID1).
		WillReturnError(errors.New("test error"))
	mock.ExpectRollback()

	// actual
	err := indexTestDb.MarkDeletedByIndexID(tenantID, indexID1)
	assert.Error(t, err)
}
