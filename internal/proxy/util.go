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

package proxy

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/proto/commonpb"
	"github.com/milvus-io/milvus/internal/util/retry"
)

func GetPulsarConfig(protocol, ip, port, url string) (map[string]interface{}, error) {
	var resp *http.Response
	var err error

	getResp := func() error {
		log.Debug("proxy util", zap.String("url", protocol+"://"+ip+":"+port+url))
		resp, err = http.Get(protocol + "://" + ip + ":" + port + url)
		return err
	}

	err = retry.Do(context.TODO(), getResp, retry.Attempts(10), retry.Sleep(time.Second))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ret := make(map[string]interface{})
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GetAttrByKeyFromRepeatedKV(key string, kvs []*commonpb.KeyValuePair) (string, error) {
	for _, kv := range kvs {
		if kv.Key == key {
			return kv.Value, nil
		}
	}

	return "", errors.New("key " + key + " not found")
}
