/*******************************************************************************
 * Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 ******************************************************************************/
#pragma once

#include "Meta.h"
#include "Options.h"

namespace zilliz {
namespace vecwise {
namespace engine {
namespace meta {

auto StoragePrototype(const std::string& path);

class DBMetaImpl : public Meta {
public:
    DBMetaImpl(const DBMetaOptions& options_);

    virtual Status CreateTable(TableSchema& table_schema) override;
    virtual Status DescribeTable(TableSchema& group_info_) override;
    virtual Status HasTable(const std::string& table_id, bool& has_or_not) override;

    virtual Status CreateTableFile(TableFileSchema& file_schema) override;
    virtual Status DropPartitionsByDates(const std::string& table_id,
            const DatesT& dates) override;

    virtual Status has_group_file(const std::string& table_id_,
                                  const std::string& file_id_,
                                  bool& has_or_not_) override;
    virtual Status get_group_file(const std::string& table_id_,
                                  const std::string& file_id_,
                                  TableFileSchema& group_file_info_) override;
    virtual Status update_group_file(TableFileSchema& group_file_) override;

    virtual Status get_group_files(const std::string& table_id_,
                                   const int date_delta_,
                                   TableFilesSchema& group_files_info_) override;

    virtual Status update_files(TableFilesSchema& files) override;

    virtual Status files_to_merge(const std::string& table_id,
            DatePartionedTableFilesSchema& files) override;

    virtual Status files_to_search(const std::string& table_id,
                                  const DatesT& partition,
                                  DatePartionedTableFilesSchema& files) override;

    virtual Status files_to_index(TableFilesSchema&) override;

    virtual Status archive_files() override;

    virtual Status size(long& result) override;

    virtual Status cleanup() override;

    virtual Status cleanup_ttl_files(uint16_t seconds) override;

    virtual Status drop_all() override;

    virtual Status count(const std::string& table_id, long& result) override;

    virtual ~DBMetaImpl();

private:
    Status NextFileId(std::string& file_id);
    Status NextGroupId(std::string& table_id);
    Status discard_files_of_size(long to_discard_size);
    std::string GetGroupPath(const std::string& table_id);
    std::string GetGroupDatePartitionPath(const std::string& table_id, DateT& date);
    void GetGroupFilePath(TableFileSchema& group_file);
    Status initialize();

    const DBMetaOptions _options;
}; // DBMetaImpl

} // namespace meta
} // namespace engine
} // namespace vecwise
} // namespace zilliz
