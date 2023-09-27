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

package etcdkv

import (
	"fmt"
	"os"
	"path"
	"sort"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	clientv3 "go.etcd.io/etcd/client/v3"
	"golang.org/x/exp/maps"

	"github.com/milvus-io/milvus/internal/kv/predicates"
	"github.com/milvus-io/milvus/pkg/util/etcd"
	"github.com/milvus-io/milvus/pkg/util/funcutil"
	"github.com/milvus-io/milvus/pkg/util/paramtable"
)

var Params = paramtable.Get()

func TestMain(m *testing.M) {
	paramtable.Init()
	code := m.Run()
	os.Exit(code)
}

type EtcdKVSuite struct {
	suite.Suite

	rootPath string
	etcdCli  *clientv3.Client
	etcdKV   *etcdKV
}

func (s *EtcdKVSuite) SetupSuite() {
	etcdCli, err := etcd.GetEtcdClient(
		Params.EtcdCfg.UseEmbedEtcd.GetAsBool(),
		Params.EtcdCfg.EtcdUseSSL.GetAsBool(),
		Params.EtcdCfg.Endpoints.GetAsStrings(),
		Params.EtcdCfg.EtcdTLSCert.GetValue(),
		Params.EtcdCfg.EtcdTLSKey.GetValue(),
		Params.EtcdCfg.EtcdTLSCACert.GetValue(),
		Params.EtcdCfg.EtcdTLSMinVersion.GetValue())
	s.Require().NoError(err)

	s.etcdCli = etcdCli
}

func (s *EtcdKVSuite) TearDownSuite() {
	if s.etcdCli != nil {
		s.etcdCli.Close()
	}
}

func (s *EtcdKVSuite) SetupTest() {
	s.rootPath = path.Join("unittest/etcdkv", funcutil.RandomString(8))
	s.etcdKV = NewEtcdKV(s.etcdCli, s.rootPath)
}

func (s *EtcdKVSuite) TearDownTest() {
	s.etcdKV.RemoveWithPrefix("")
	s.etcdKV.Close()
}

func (s *EtcdKVSuite) TestSaveLoad() {
	etcdKV := s.etcdKV
	saveAndLoadTests := []struct {
		key   string
		value string
	}{
		{"test1", "value1"},
		{"test2", "value2"},
		{"test1/a", "value_a"},
		{"test1/b", "value_b"},
	}

	for i, test := range saveAndLoadTests {
		if i < 4 {
			err := etcdKV.Save(test.key, test.value)
			s.Require().NoError(err)
		}

		val, err := etcdKV.Load(test.key)
		s.Require().NoError(err)
		s.Equal(test.value, val)
	}

	invalidLoadTests := []struct {
		invalidKey string
	}{
		{"t"},
		{"a"},
		{"test1a"},
	}

	for _, test := range invalidLoadTests {
		val, err := etcdKV.Load(test.invalidKey)
		s.Error(err)
		s.Zero(val)
	}

	loadPrefixTests := []struct {
		prefix string

		expectedKeys   []string
		expectedValues []string
		expectedError  error
	}{
		{"test", []string{
			etcdKV.GetPath("test1"),
			etcdKV.GetPath("test2"),
			etcdKV.GetPath("test1/a"),
			etcdKV.GetPath("test1/b"),
		}, []string{"value1", "value2", "value_a", "value_b"}, nil},
		{"test1", []string{
			etcdKV.GetPath("test1"),
			etcdKV.GetPath("test1/a"),
			etcdKV.GetPath("test1/b"),
		}, []string{"value1", "value_a", "value_b"}, nil},
		{"test2", []string{etcdKV.GetPath("test2")}, []string{"value2"}, nil},
		{"", []string{
			etcdKV.GetPath("test1"),
			etcdKV.GetPath("test2"),
			etcdKV.GetPath("test1/a"),
			etcdKV.GetPath("test1/b"),
		}, []string{"value1", "value2", "value_a", "value_b"}, nil},
		{"test1/a", []string{etcdKV.GetPath("test1/a")}, []string{"value_a"}, nil},
		{"a", []string{}, []string{}, nil},
		{"root", []string{}, []string{}, nil},
		{"/etcd/test/root", []string{}, []string{}, nil},
	}

	for _, test := range loadPrefixTests {
		actualKeys, actualValues, err := etcdKV.LoadWithPrefix(test.prefix)
		s.ElementsMatch(test.expectedKeys, actualKeys)
		s.ElementsMatch(test.expectedValues, actualValues)
		s.Equal(test.expectedError, err)
	}

	removeTests := []struct {
		validKey   string
		invalidKey string
	}{
		{"test1", "abc"},
		{"test1/a", "test1/lskfjal"},
		{"test1/b", "test1/b"},
		{"test2", "-"},
	}

	for _, test := range removeTests {
		err := etcdKV.Remove(test.validKey)
		s.NoError(err)

		_, err = etcdKV.Load(test.validKey)
		s.Error(err)

		err = etcdKV.Remove(test.validKey)
		s.NoError(err)
		err = etcdKV.Remove(test.invalidKey)
		s.NoError(err)
	}
}

func (s *EtcdKVSuite) TestSaveAndLoadBytes() {
	etcdKV := s.etcdKV

	saveAndLoadTests := []struct {
		key   string
		value string
	}{
		{"test1", "value1"},
		{"test2", "value2"},
		{"test1/a", "value_a"},
		{"test1/b", "value_b"},
	}

	for i, test := range saveAndLoadTests {
		if i < 4 {
			err := etcdKV.SaveBytes(test.key, []byte(test.value))
			s.Require().NoError(err)
		}

		val, err := etcdKV.LoadBytes(test.key)
		s.NoError(err)
		s.Equal(test.value, string(val))
	}

	invalidLoadTests := []struct {
		invalidKey string
	}{
		{"t"},
		{"a"},
		{"test1a"},
	}

	for _, test := range invalidLoadTests {
		val, err := etcdKV.LoadBytes(test.invalidKey)
		s.Error(err)
		s.Zero(string(val))
	}

	loadPrefixTests := []struct {
		prefix string

		expectedKeys   []string
		expectedValues []string
		expectedError  error
	}{
		{"test", []string{
			etcdKV.GetPath("test1"),
			etcdKV.GetPath("test2"),
			etcdKV.GetPath("test1/a"),
			etcdKV.GetPath("test1/b"),
		}, []string{"value1", "value2", "value_a", "value_b"}, nil},
		{"test1", []string{
			etcdKV.GetPath("test1"),
			etcdKV.GetPath("test1/a"),
			etcdKV.GetPath("test1/b"),
		}, []string{"value1", "value_a", "value_b"}, nil},
		{"test2", []string{etcdKV.GetPath("test2")}, []string{"value2"}, nil},
		{"", []string{
			etcdKV.GetPath("test1"),
			etcdKV.GetPath("test2"),
			etcdKV.GetPath("test1/a"),
			etcdKV.GetPath("test1/b"),
		}, []string{"value1", "value2", "value_a", "value_b"}, nil},
		{"test1/a", []string{etcdKV.GetPath("test1/a")}, []string{"value_a"}, nil},
		{"a", []string{}, []string{}, nil},
		{"root", []string{}, []string{}, nil},
		{"/etcd/test/root", []string{}, []string{}, nil},
	}

	for _, test := range loadPrefixTests {
		actualKeys, actualValues, err := etcdKV.LoadBytesWithPrefix(test.prefix)
		actualStringValues := make([]string, len(actualValues))
		for i := range actualValues {
			actualStringValues[i] = string(actualValues[i])
		}
		s.ElementsMatch(test.expectedKeys, actualKeys)
		s.ElementsMatch(test.expectedValues, actualStringValues)
		s.Equal(test.expectedError, err)

		actualKeys, actualValues, versions, err := etcdKV.LoadBytesWithPrefix2(test.prefix)
		actualStringValues = make([]string, len(actualValues))
		for i := range actualValues {
			actualStringValues[i] = string(actualValues[i])
		}
		s.ElementsMatch(test.expectedKeys, actualKeys)
		s.ElementsMatch(test.expectedValues, actualStringValues)
		s.NotZero(versions)
		s.Equal(test.expectedError, err)
	}

	removeTests := []struct {
		validKey   string
		invalidKey string
	}{
		{"test1", "abc"},
		{"test1/a", "test1/lskfjal"},
		{"test1/b", "test1/b"},
		{"test2", "-"},
	}

	for _, test := range removeTests {
		err := etcdKV.Remove(test.validKey)
		s.NoError(err)

		_, err = etcdKV.Load(test.validKey)
		s.Error(err)

		err = etcdKV.Remove(test.validKey)
		s.NoError(err)
		err = etcdKV.Remove(test.invalidKey)
		s.NoError(err)
	}
}

func (s *EtcdKVSuite) TestLoadBytesWithRevision() {
	etcdKV := s.etcdKV

	prepareKV := []struct {
		inKey   string
		inValue string
	}{
		{"a", "a_version1"},
		{"b", "b_version2"},
		{"a", "a_version3"},
		{"c", "c_version4"},
		{"a/suba", "a_version5"},
	}

	for _, test := range prepareKV {
		err := etcdKV.SaveBytes(test.inKey, []byte(test.inValue))
		s.NoError(err)
	}

	loadWithRevisionTests := []struct {
		inKey string

		expectedKeyNo  int
		expectedValues []string
	}{
		{"a", 2, []string{"a_version3", "a_version5"}},
		{"b", 1, []string{"b_version2"}},
		{"c", 1, []string{"c_version4"}},
	}

	for _, test := range loadWithRevisionTests {
		keys, values, revision, err := etcdKV.LoadBytesWithRevision(test.inKey)
		s.NoError(err)
		s.Equal(test.expectedKeyNo, len(keys))
		stringValues := make([]string, len(values))
		for i := range values {
			stringValues[i] = string(values[i])
		}
		s.ElementsMatch(test.expectedValues, stringValues)
		s.NotZero(revision)
	}
}

func (s *EtcdKVSuite) TestMultiSaveAndMultiLoad() {
	etcdKV := s.etcdKV
	multiSaveTests := map[string]string{
		"key_1":      "value_1",
		"key_2":      "value_2",
		"key_3/a":    "value_3a",
		"multikey_1": "multivalue_1",
		"multikey_2": "multivalue_2",
		"_":          "other",
	}

	err := etcdKV.MultiSave(multiSaveTests)
	s.Require().NoError(err)
	for k, v := range multiSaveTests {
		actualV, err := etcdKV.Load(k)
		s.NoError(err)
		s.Equal(v, actualV)
	}

	multiLoadTests := []struct {
		inputKeys      []string
		expectedValues []string
	}{
		{[]string{"key_1"}, []string{"value_1"}},
		{[]string{"key_1", "key_2", "key_3/a"}, []string{"value_1", "value_2", "value_3a"}},
		{[]string{"multikey_1", "multikey_2"}, []string{"multivalue_1", "multivalue_2"}},
		{[]string{"_"}, []string{"other"}},
	}

	for _, test := range multiLoadTests {
		vs, err := etcdKV.MultiLoad(test.inputKeys)
		s.NoError(err)
		s.Equal(test.expectedValues, vs)
	}

	invalidMultiLoad := []struct {
		invalidKeys    []string
		expectedValues []string
	}{
		{[]string{"a", "key_1"}, []string{"", "value_1"}},
		{[]string{".....", "key_1"}, []string{"", "value_1"}},
		{[]string{"*********"}, []string{""}},
		{[]string{"key_1", "1"}, []string{"value_1", ""}},
	}

	for _, test := range invalidMultiLoad {
		vs, err := etcdKV.MultiLoad(test.invalidKeys)
		s.Error(err)
		s.Equal(test.expectedValues, vs)
	}

	removeWithPrefixTests := []string{
		"key_1",
		"multi",
	}

	for _, k := range removeWithPrefixTests {
		err = etcdKV.RemoveWithPrefix(k)
		s.NoError(err)

		ks, vs, err := etcdKV.LoadWithPrefix(k)
		s.Empty(ks)
		s.Empty(vs)
		s.NoError(err)
	}

	multiRemoveTests := []string{
		"key_2",
		"key_3/a",
		"multikey_2",
		"_",
	}

	err = etcdKV.MultiRemove(multiRemoveTests)
	s.NoError(err)

	ks, vs, err := etcdKV.LoadWithPrefix("")
	s.NoError(err)
	s.Empty(ks)
	s.Empty(vs)

	multiSaveAndRemoveTests := []struct {
		multiSaves   map[string]string
		multiRemoves []string
	}{
		{map[string]string{"key_1": "value_1"}, []string{}},
		{map[string]string{"key_2": "value_2"}, []string{"key_1"}},
		{map[string]string{"key_3/a": "value_3a"}, []string{"key_2"}},
		{map[string]string{"multikey_1": "multivalue_1"}, []string{}},
		{map[string]string{"multikey_2": "multivalue_2"}, []string{"multikey_1", "key_3/a"}},
		{make(map[string]string), []string{"multikey_2"}},
	}
	for _, test := range multiSaveAndRemoveTests {
		err = etcdKV.MultiSaveAndRemove(test.multiSaves, test.multiRemoves)
		s.NoError(err)
	}

	ks, vs, err = etcdKV.LoadWithPrefix("")
	s.NoError(err)
	s.Empty(ks)
	s.Empty(vs)
}

func (s *EtcdKVSuite) TestMultiSaveBytesAndMultiLoadBytes() {
	etcdKV := s.etcdKV
	multiSaveTests := map[string]string{
		"key_1":      "value_1",
		"key_2":      "value_2",
		"key_3/a":    "value_3a",
		"multikey_1": "multivalue_1",
		"multikey_2": "multivalue_2",
		"_":          "other",
	}

	multiSaveBytesTests := make(map[string][]byte)
	for k, v := range multiSaveTests {
		multiSaveBytesTests[k] = []byte(v)
	}

	err := etcdKV.MultiSaveBytes(multiSaveBytesTests)
	s.Require().NoError(err)
	for k, v := range multiSaveTests {
		actualV, err := etcdKV.LoadBytes(k)
		s.NoError(err)
		s.Equal(v, string(actualV))
	}

	multiLoadTests := []struct {
		inputKeys      []string
		expectedValues []string
	}{
		{[]string{"key_1"}, []string{"value_1"}},
		{[]string{"key_1", "key_2", "key_3/a"}, []string{"value_1", "value_2", "value_3a"}},
		{[]string{"multikey_1", "multikey_2"}, []string{"multivalue_1", "multivalue_2"}},
		{[]string{"_"}, []string{"other"}},
	}

	for _, test := range multiLoadTests {
		vs, err := etcdKV.MultiLoadBytes(test.inputKeys)
		stringVs := make([]string, len(vs))
		for i := range vs {
			stringVs[i] = string(vs[i])
		}
		s.NoError(err)
		s.Equal(test.expectedValues, stringVs)
	}

	invalidMultiLoad := []struct {
		invalidKeys    []string
		expectedValues []string
	}{
		{[]string{"a", "key_1"}, []string{"", "value_1"}},
		{[]string{".....", "key_1"}, []string{"", "value_1"}},
		{[]string{"*********"}, []string{""}},
		{[]string{"key_1", "1"}, []string{"value_1", ""}},
	}

	for _, test := range invalidMultiLoad {
		vs, err := etcdKV.MultiLoadBytes(test.invalidKeys)
		stringVs := make([]string, len(vs))
		for i := range vs {
			stringVs[i] = string(vs[i])
		}
		s.Error(err)
		s.Equal(test.expectedValues, stringVs)
	}

	removeWithPrefixTests := []string{
		"key_1",
		"multi",
	}

	for _, k := range removeWithPrefixTests {
		err = etcdKV.RemoveWithPrefix(k)
		s.NoError(err)

		ks, vs, err := etcdKV.LoadBytesWithPrefix(k)
		s.Empty(ks)
		s.Empty(vs)
		s.NoError(err)
	}

	multiRemoveTests := []string{
		"key_2",
		"key_3/a",
		"multikey_2",
		"_",
	}

	err = etcdKV.MultiRemove(multiRemoveTests)
	s.NoError(err)

	ks, vs, err := etcdKV.LoadBytesWithPrefix("")
	s.NoError(err)
	s.Empty(ks)
	s.Empty(vs)

	multiSaveAndRemoveTests := []struct {
		multiSaves   map[string][]byte
		multiRemoves []string
	}{
		{map[string][]byte{"key_1": []byte("value_1")}, []string{}},
		{map[string][]byte{"key_2": []byte("value_2")}, []string{"key_1"}},
		{map[string][]byte{"key_3/a": []byte("value_3a")}, []string{"key_2"}},
		{map[string][]byte{"multikey_1": []byte("multivalue_1")}, []string{}},
		{map[string][]byte{"multikey_2": []byte("multivalue_2")}, []string{"multikey_1", "key_3/a"}},
		{make(map[string][]byte), []string{"multikey_2"}},
	}

	for _, test := range multiSaveAndRemoveTests {
		err = etcdKV.MultiSaveBytesAndRemove(test.multiSaves, test.multiRemoves)
		s.NoError(err)
	}

	ks, vs, err = etcdKV.LoadBytesWithPrefix("")
	s.NoError(err)
	s.Empty(ks)
	s.Empty(vs)
}

func (s *EtcdKVSuite) TestTxnWithPredicates() {
	etcdKV := s.etcdKV

	prepareKV := map[string]string{
		"lease1": "1",
		"lease2": "2",
	}

	err := etcdKV.MultiSave(prepareKV)
	s.Require().NoError(err)

	badPredicate := predicates.NewMockPredicate(s.T())
	badPredicate.EXPECT().Type().Return(0)
	badPredicate.EXPECT().Target().Return(predicates.PredTargetValue)

	multiSaveAndRemovePredTests := []struct {
		tag           string
		multiSave     map[string]string
		preds         []predicates.Predicate
		expectSuccess bool
	}{
		{"predicate_ok", map[string]string{"a": "b"}, []predicates.Predicate{predicates.ValueEqual("lease1", "1")}, true},
		{"predicate_fail", map[string]string{"a": "b"}, []predicates.Predicate{predicates.ValueEqual("lease1", "2")}, false},
		{"bad_predicate", map[string]string{"a": "b"}, []predicates.Predicate{badPredicate}, false},
	}

	for _, test := range multiSaveAndRemovePredTests {
		s.Run(test.tag, func() {
			err := etcdKV.MultiSaveAndRemove(test.multiSave, nil, test.preds...)
			if test.expectSuccess {
				s.NoError(err)
			} else {
				s.Error(err)
			}
			err = etcdKV.MultiSaveAndRemoveWithPrefix(test.multiSave, nil, test.preds...)
			if test.expectSuccess {
				s.NoError(err)
			} else {
				s.Error(err)
			}
		})
	}
}

func (s *EtcdKVSuite) TestMultiSaveAndRemoveWithPrefix() {
	etcdKV := s.etcdKV

	prepareTests := map[string]string{
		"x/abc/1": "1",
		"x/abc/2": "2",
		"x/def/1": "10",
		"x/def/2": "20",
		"x/den/1": "100",
		"x/den/2": "200",
	}

	// MultiSaveAndRemoveWithPrefix
	err := etcdKV.MultiSave(prepareTests)
	s.Require().NoError(err)
	multiSaveAndRemoveWithPrefixTests := []struct {
		multiSave map[string]string
		prefix    []string

		loadPrefix         string
		lengthBeforeRemove int
		lengthAfterRemove  int
	}{
		{map[string]string{}, []string{"x/abc", "x/def", "x/den"}, "x", 6, 0},
		{map[string]string{"y/a": "vvv", "y/b": "vvv"}, []string{}, "y", 0, 2},
		{map[string]string{"y/c": "vvv"}, []string{}, "y", 2, 3},
		{map[string]string{"p/a": "vvv"}, []string{"y/a", "y"}, "y", 3, 0},
		{map[string]string{}, []string{"p"}, "p", 1, 0},
	}

	for _, test := range multiSaveAndRemoveWithPrefixTests {
		k, _, err := etcdKV.LoadWithPrefix(test.loadPrefix)
		s.NoError(err)
		s.Equal(test.lengthBeforeRemove, len(k))

		err = etcdKV.MultiSaveAndRemoveWithPrefix(test.multiSave, test.prefix)
		s.NoError(err)

		k, _, err = etcdKV.LoadWithPrefix(test.loadPrefix)
		s.NoError(err)
		s.Equal(test.lengthAfterRemove, len(k))
	}
}

func (s *EtcdKVSuite) TestWatch() {
	etcdKV := s.etcdKV

	ch := etcdKV.Watch("x")
	resp := <-ch
	s.True(resp.Created)

	ch = etcdKV.WatchWithPrefix("x")
	resp = <-ch
	s.True(resp.Created)
}

func (s *EtcdKVSuite) TestRevisionBytes() {
	etcdKV := s.etcdKV

	revisionTests := []struct {
		inKey       string
		fistValue   []byte
		secondValue []byte
	}{
		{"a", []byte("v1"), []byte("v11")},
		{"y", []byte("v2"), []byte("v22")},
		{"z", []byte("v3"), []byte("v33")},
	}

	for _, test := range revisionTests {
		err := etcdKV.SaveBytes(test.inKey, test.fistValue)
		s.Require().NoError(err)

		_, _, revision, _ := etcdKV.LoadBytesWithRevision(test.inKey)
		ch := etcdKV.WatchWithRevision(test.inKey, revision+1)

		err = etcdKV.SaveBytes(test.inKey, test.secondValue)
		s.Require().NoError(err)

		resp := <-ch
		s.Equal(1, len(resp.Events))
		s.Equal(string(test.secondValue), string(resp.Events[0].Kv.Value))
		s.Equal(revision+1, resp.Header.Revision)
	}

	success, err := etcdKV.CompareVersionAndSwapBytes("a/b/c", 0, []byte("1"))
	s.NoError(err)
	s.True(success)

	value, err := etcdKV.LoadBytes("a/b/c")
	s.NoError(err)
	s.Equal(string(value), "1")

	success, err = etcdKV.CompareVersionAndSwapBytes("a/b/c", 0, []byte("1"))
	s.NoError(err)
	s.False(success)
}

func TestEtcdKV(t *testing.T) {
	suite.Run(t, new(EtcdKVSuite))
}

func Test_WalkWithPagination(t *testing.T) {
	etcdCli, err := etcd.GetEtcdClient(
		Params.EtcdCfg.UseEmbedEtcd.GetAsBool(),
		Params.EtcdCfg.EtcdUseSSL.GetAsBool(),
		Params.EtcdCfg.Endpoints.GetAsStrings(),
		Params.EtcdCfg.EtcdTLSCert.GetValue(),
		Params.EtcdCfg.EtcdTLSKey.GetValue(),
		Params.EtcdCfg.EtcdTLSCACert.GetValue(),
		Params.EtcdCfg.EtcdTLSMinVersion.GetValue())
	defer etcdCli.Close()
	assert.NoError(t, err)

	rootPath := "/etcd/test/root/pagination"
	etcdKV := NewEtcdKV(etcdCli, rootPath)

	defer etcdKV.Close()
	defer etcdKV.RemoveWithPrefix("")

	kvs := map[string]string{
		"A/100":    "v1",
		"AA/100":   "v2",
		"AB/100":   "v3",
		"AB/2/100": "v4",
		"B/100":    "v5",
	}

	err = etcdKV.MultiSave(kvs)
	assert.NoError(t, err)
	for k, v := range kvs {
		actualV, err := etcdKV.Load(k)
		assert.NoError(t, err)
		assert.Equal(t, v, actualV)
	}

	t.Run("apply function error ", func(t *testing.T) {
		err = etcdKV.WalkWithPrefix("A", 5, func(key []byte, value []byte) error {
			return errors.New("error")
		})
		assert.Error(t, err)
	})

	t.Run("get with non-exist prefix ", func(t *testing.T) {
		err = etcdKV.WalkWithPrefix("non-exist-prefix", 5, func(key []byte, value []byte) error {
			return nil
		})
		assert.NoError(t, err)
	})

	t.Run("with different pagination", func(t *testing.T) {
		testFn := func(pagination int) {
			expected := map[string]string{
				"A/100":    "v1",
				"AA/100":   "v2",
				"AB/100":   "v3",
				"AB/2/100": "v4",
			}

			expectedSortedKey := maps.Keys(expected)
			sort.Strings(expectedSortedKey)

			ret := make(map[string]string)
			actualSortedKey := make([]string, 0)

			err = etcdKV.WalkWithPrefix("A", pagination, func(key []byte, value []byte) error {
				k := string(key)
				k = k[len(rootPath)+1:]
				ret[k] = string(value)
				actualSortedKey = append(actualSortedKey, k)
				return nil
			})

			assert.NoError(t, err)
			assert.Equal(t, expected, ret, fmt.Errorf("pagination: %d", pagination))
			assert.Equal(t, expectedSortedKey, actualSortedKey, fmt.Errorf("pagination: %d", pagination))
		}

		testFn(-100)
		testFn(-1)
		testFn(0)
		testFn(1)
		testFn(5)
		testFn(100)
	})
}

func TestElapse(t *testing.T) {
	start := time.Now()
	isElapse := CheckElapseAndWarn(start, "err message")
	assert.Equal(t, isElapse, false)

	time.Sleep(2001 * time.Millisecond)
	isElapse = CheckElapseAndWarn(start, "err message")
	assert.Equal(t, isElapse, true)
}

func TestCheckValueSizeAndWarn(t *testing.T) {
	ret := CheckValueSizeAndWarn("k", "v")
	assert.False(t, ret)

	v := make([]byte, 1024000)
	ret = CheckValueSizeAndWarn("k", v)
	assert.True(t, ret)
}

func TestCheckTnxBytesValueSizeAndWarn(t *testing.T) {
	kvs := make(map[string][]byte, 0)
	kvs["k"] = []byte("v")
	ret := CheckTnxBytesValueSizeAndWarn(kvs)
	assert.False(t, ret)

	kvs["k"] = make([]byte, 1024000)
	ret = CheckTnxBytesValueSizeAndWarn(kvs)
	assert.True(t, ret)
}

func TestCheckTnxStringValueSizeAndWarn(t *testing.T) {
	kvs := make(map[string]string, 0)
	kvs["k"] = "v"
	ret := CheckTnxStringValueSizeAndWarn(kvs)
	assert.False(t, ret)

	kvs["k1"] = funcutil.RandomString(1024000)
	ret = CheckTnxStringValueSizeAndWarn(kvs)
	assert.True(t, ret)
}

func TestHas(t *testing.T) {
	etcdCli, err := etcd.GetEtcdClient(
		Params.EtcdCfg.UseEmbedEtcd.GetAsBool(),
		Params.EtcdCfg.EtcdUseSSL.GetAsBool(),
		Params.EtcdCfg.Endpoints.GetAsStrings(),
		Params.EtcdCfg.EtcdTLSCert.GetValue(),
		Params.EtcdCfg.EtcdTLSKey.GetValue(),
		Params.EtcdCfg.EtcdTLSCACert.GetValue(),
		Params.EtcdCfg.EtcdTLSMinVersion.GetValue())
	defer etcdCli.Close()
	assert.NoError(t, err)
	rootPath := "/etcd/test/root/has"
	kv := NewEtcdKV(etcdCli, rootPath)
	err = kv.RemoveWithPrefix("")
	require.NoError(t, err)

	defer kv.Close()
	defer kv.RemoveWithPrefix("")

	has, err := kv.Has("key1")
	assert.NoError(t, err)
	assert.False(t, has)

	err = kv.Save("key1", "value1")
	assert.NoError(t, err)

	has, err = kv.Has("key1")
	assert.NoError(t, err)
	assert.True(t, has)

	err = kv.Remove("key1")
	assert.NoError(t, err)

	has, err = kv.Has("key1")
	assert.NoError(t, err)
	assert.False(t, has)
}

func TestHasPrefix(t *testing.T) {
	etcdCli, err := etcd.GetEtcdClient(
		Params.EtcdCfg.UseEmbedEtcd.GetAsBool(),
		Params.EtcdCfg.EtcdUseSSL.GetAsBool(),
		Params.EtcdCfg.Endpoints.GetAsStrings(),
		Params.EtcdCfg.EtcdTLSCert.GetValue(),
		Params.EtcdCfg.EtcdTLSKey.GetValue(),
		Params.EtcdCfg.EtcdTLSCACert.GetValue(),
		Params.EtcdCfg.EtcdTLSMinVersion.GetValue())
	defer etcdCli.Close()
	assert.NoError(t, err)
	rootPath := "/etcd/test/root/hasprefix"
	kv := NewEtcdKV(etcdCli, rootPath)
	err = kv.RemoveWithPrefix("")
	require.NoError(t, err)

	defer kv.Close()
	defer kv.RemoveWithPrefix("")

	has, err := kv.HasPrefix("key")
	assert.NoError(t, err)
	assert.False(t, has)

	err = kv.Save("key1", "value1")
	assert.NoError(t, err)

	has, err = kv.HasPrefix("key")
	assert.NoError(t, err)
	assert.True(t, has)

	err = kv.Remove("key1")
	assert.NoError(t, err)

	has, err = kv.HasPrefix("key")
	assert.NoError(t, err)
	assert.False(t, has)
}
