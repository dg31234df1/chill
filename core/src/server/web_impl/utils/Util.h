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

#include <string>
#include <vector>

#include "db/Types.h"
#include "server/web_impl/Types.h"
#include "utils/Status.h"

namespace milvus {
namespace server {
namespace web {

Status
CopyRowRecords(const OList<OList<OFloat32>::ObjectWrapper>::ObjectWrapper& records, std::vector<float>& vectors);

Status
CopyBinRowRecords(const OList<OList<OInt64>::ObjectWrapper>::ObjectWrapper& records, std::vector<uint8_t>& vectors);

Status
WebRequestHandler::ParseQueryInteger(const OQueryParams& query_params, const std::string& key, int64_t& value,
                                     bool nullable) {
    auto query = query_params.get(key.c_str());
    if (nullptr != query.get() && query->getSize() > 0) {
        std::string value_str = query->std_str();
        if (!ValidationUtil::ValidateStringIsNumber(value_str).ok()) {
            return Status(ILLEGAL_QUERY_PARAM,
                          "Query param \'offset\' is illegal, only non-negative integer supported");
        }

        value = std::stol(value_str);
    } else if (!nullable) {
        return Status(QUERY_PARAM_LOSS, "Query param \"" + key + "\" is required");
    }

    return Status::OK();
}

Status
WebRequestHandler::ParseQueryStr(const OQueryParams& query_params, const std::string& key, std::string& value,
                                 bool nullable) {
    auto query = query_params.get(key.c_str());
    if (nullptr != query.get() && query->getSize() > 0) {
        value = query->std_str();
    } else if (!nullable) {
        return Status(QUERY_PARAM_LOSS, "Query param \"" + key + "\" is required");
    }

    return Status::OK();
}

Status
WebRequestHandler::ParseQueryBool(const OQueryParams& query_params, const std::string& key, bool& value,
                                  bool nullable) {
    auto query = query_params.get(key.c_str());
    if (nullptr != query.get() && query->getSize() > 0) {
        std::string value_str = query->std_str();
        if (!ValidationUtil::ValidateStringIsBool(value_str).ok()) {
            return Status(ILLEGAL_QUERY_PARAM, "Query param \'all_required\' must be a bool");
        }
        value = value_str == "True" || value_str == "true";
        return Status::OK();
    }

    if (!nullable) {
        return Status(QUERY_PARAM_LOSS, "Query param \"" + key + "\" is required");
    }

    return Status::OK();
}

}  // namespace web
}  // namespace server
}  // namespace milvus
