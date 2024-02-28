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

#include <gtest/gtest.h>

#include "common/RegexQuery.h"

TEST(TranslatePatternMatchToRegexTest, SimplePatternWithPercent) {
    std::string pattern = "abc%";
    std::string result = milvus::TranslatePatternMatchToRegex(pattern);
    EXPECT_EQ(result, "abc.*");
}

TEST(TranslatePatternMatchToRegexTest, PatternWithUnderscore) {
    std::string pattern = "a_c";
    std::string result = milvus::TranslatePatternMatchToRegex(pattern);
    EXPECT_EQ(result, "a.c");
}

TEST(TranslatePatternMatchToRegexTest, PatternWithSpecialCharacters) {
    std::string pattern = "a\\%b\\_c";
    std::string result = milvus::TranslatePatternMatchToRegex(pattern);
    EXPECT_EQ(result, "a\\%b\\_c");
}

TEST(TranslatePatternMatchToRegexTest,
     PatternWithMultiplePercentAndUnderscore) {
    std::string pattern = "%a_b%";
    std::string result = milvus::TranslatePatternMatchToRegex(pattern);
    EXPECT_EQ(result, ".*a.b.*");
}

TEST(TranslatePatternMatchToRegexTest, PatternWithRegexChar) {
    std::string pattern = "abc*def.ghi+";
    std::string result = milvus::TranslatePatternMatchToRegex(pattern);
    EXPECT_EQ(result, "abc\\*def\\.ghi\\+");
}

TEST(PatternMatchTranslatorTest, InvalidTypeTest) {
    using namespace milvus;
    PatternMatchTranslator translator;

    ASSERT_ANY_THROW(translator(123));
    ASSERT_ANY_THROW(translator(3.14));
    ASSERT_ANY_THROW(translator(true));
}

TEST(PatternMatchTranslatorTest, StringTypeTest) {
    using namespace milvus;
    PatternMatchTranslator translator;

    std::string pattern1 = "abc";
    std::string pattern2 = "xyz";
    std::string pattern3 = "%a_b%";

    EXPECT_EQ(translator(pattern1), "abc");
    EXPECT_EQ(translator(pattern2), "xyz");
    EXPECT_EQ(translator(pattern3), ".*a.b.*");
}

TEST(RegexMatcherTest, DefaultBehaviorTest) {
    using namespace milvus;
    RegexMatcher matcher;
    std::regex pattern("Hello.*");

    int operand1 = 123;
    double operand2 = 3.14;
    bool operand3 = true;

    EXPECT_FALSE(matcher(pattern, operand1));
    EXPECT_FALSE(matcher(pattern, operand2));
    EXPECT_FALSE(matcher(pattern, operand3));
}

TEST(RegexMatcherTest, StringMatchTest) {
    using namespace milvus;
    RegexMatcher matcher;
    std::regex pattern("Hello.*");

    std::string str1 = "Hello, World!";
    std::string str2 = "Hi there!";
    std::string str3 = "Hello, OpenAI!";

    EXPECT_TRUE(matcher(pattern, str1));
    EXPECT_FALSE(matcher(pattern, str2));
    EXPECT_TRUE(matcher(pattern, str3));
}

TEST(RegexMatcherTest, StringViewMatchTest) {
    using namespace milvus;
    RegexMatcher matcher;
    std::regex pattern("Hello.*");

    std::string_view str1 = "Hello, World!";
    std::string_view str2 = "Hi there!";
    std::string_view str3 = "Hello, OpenAI!";

    EXPECT_TRUE(matcher(pattern, str1));
    EXPECT_FALSE(matcher(pattern, str2));
    EXPECT_TRUE(matcher(pattern, str3));
}
