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

package querynode

import (
	"context"
	"testing"

	"github.com/milvus-io/milvus/internal/util/flowgraph"
)

func TestServiceTimeNode_Operate(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	genServiceTimeNode := func() *serviceTimeNode {
		tSafe := newTSafeReplica()
		tSafe.addTSafe(defaultDMLChannel)

		node := newServiceTimeNode(ctx,
			tSafe,
			defaultCollectionID,
			defaultDMLChannel)
		return node
	}

	t.Run("test operate of loadTypeCollection", func(t *testing.T) {
		node := genServiceTimeNode()
		msg := &serviceTimeMsg{
			timeRange: TimeRange{
				timestampMin: 0,
				timestampMax: 1000,
			},
		}
		in := []flowgraph.Msg{msg}
		node.Operate(in)
	})

	t.Run("test operate of loadTypePartition", func(t *testing.T) {
		node := genServiceTimeNode()
		msg := &serviceTimeMsg{
			timeRange: TimeRange{
				timestampMin: 0,
				timestampMax: 1000,
			},
		}
		in := []flowgraph.Msg{msg}
		node.Operate(in)
	})

	t.Run("test invalid msg type", func(t *testing.T) {
		node := genServiceTimeNode()
		msg := &serviceTimeMsg{
			timeRange: TimeRange{
				timestampMin: 0,
				timestampMax: 1000,
			},
		}
		in := []flowgraph.Msg{msg, msg}
		node.Operate(in)
	})

	t.Run("test no tSafe", func(t *testing.T) {
		node := genServiceTimeNode()
		node.tSafeReplica.removeTSafe(defaultDMLChannel)
		msg := &serviceTimeMsg{
			timeRange: TimeRange{
				timestampMin: 0,
				timestampMax: 1000,
			},
		}
		in := []flowgraph.Msg{msg, msg}
		node.Operate(in)
	})
}
