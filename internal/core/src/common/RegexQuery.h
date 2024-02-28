// Copyright (C) 2019-2020 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License

#pragma once

#include <string>
#include <regex>

#include "common/EasyAssert.h"

namespace milvus {
std::string
ReplaceUnescapedChars(const std::string& input,
                      char src,
                      const std::string& replacement);

std::string
TranslatePatternMatchToRegex(const std::string& pattern);

struct PatternMatchTranslator {
    template <typename T>
    inline std::string
    operator()(const T& pattern) {
        PanicInfo(OpTypeInvalid,
                  "pattern matching is only supported on string type");
    }
};

template <>
inline std::string
PatternMatchTranslator::operator()<std::string>(const std::string& pattern) {
    return TranslatePatternMatchToRegex(pattern);
}

struct RegexMatcher {
    template <typename T>
    inline bool
    operator()(const std::regex& reg, const T& operand) {
        return false;
    }
};

template <>
inline bool
RegexMatcher::operator()<std::string>(const std::regex& reg,
                                      const std::string& operand) {
    return std::regex_match(operand, reg);
}

template <>
inline bool
RegexMatcher::operator()<std::string_view>(const std::regex& reg,
                                           const std::string_view& operand) {
    return std::regex_match(operand.begin(), operand.end(), reg);
}
}  // namespace milvus
