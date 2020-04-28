// Copyright (C) 2019-2020 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License.

#pragma once

#include "server/web_impl/dto/Dto.h"
#include "server/web_impl/dto/StatusDto.hpp"
#include "server/web_impl/Constants.h"

namespace milvus {
namespace server {
namespace web {

#include OATPP_CODEGEN_BEGIN(DTO)


class TableRequestDto : public oatpp::data::mapping::type::Object {
 DTO_INIT(TableRequestDto, Object)

    DTO_FIELD(String, collection_name, "collection_name");
    DTO_FIELD(Int64, dimension, "dimension");
    DTO_FIELD(Int64, index_file_size, "index_file_size") = VALUE_TABLE_INDEX_FILE_SIZE_DEFAULT;
    DTO_FIELD(String, metric_type, "metric_type") = VALUE_TABLE_METRIC_TYPE_DEFAULT;
};

class TableFieldsDto : public oatpp::data::mapping::type::Object {
 DTO_INIT(TableFieldsDto, Object)

    DTO_FIELD(String, collection_name);
    DTO_FIELD(Int64, dimension);
    DTO_FIELD(Int64, index_file_size);
    DTO_FIELD(String, metric_type);
    DTO_FIELD(Int64, count);
    DTO_FIELD(String, index);
    DTO_FIELD(String, index_params);
};

class TableListDto : public OObject {
 DTO_INIT(TableListDto, Object)

    DTO_FIELD(List<String>::ObjectWrapper, collection_names);
};

class TableListFieldsDto : public OObject {
    DTO_INIT(TableListFieldsDto, Object)

    DTO_FIELD(List<TableFieldsDto::ObjectWrapper>::ObjectWrapper, collections);
    DTO_FIELD(Int64, count) = 0L;
};

#include OATPP_CODEGEN_END(DTO)

} // namespace web
} // namespace server
} // namespace milvus