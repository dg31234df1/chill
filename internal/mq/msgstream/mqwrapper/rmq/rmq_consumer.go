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

package rmq

import (
	"sync"
	"sync/atomic"

	"github.com/milvus-io/milvus/internal/mq/mqimpl/rocksmq/client"

	"github.com/milvus-io/milvus/internal/mq/msgstream/mqwrapper"
)

// Consumer is a client that used to consume messages from rocksmq
type Consumer struct {
	c          client.Consumer
	msgChannel chan mqwrapper.Message
	closeCh    chan struct{}
	once       sync.Once
	skip       int32
	wg         sync.WaitGroup
}

// Subscription returns the subscription name of this consumer
func (rc *Consumer) Subscription() string {
	return rc.c.Subscription()
}

// Chan returns a channel to read messages from rocksmq
func (rc *Consumer) Chan() <-chan mqwrapper.Message {
	if rc.msgChannel == nil {
		rc.once.Do(func() {
			rc.msgChannel = make(chan mqwrapper.Message, 256)
			rc.wg.Add(1)
			go func() {
				defer rc.wg.Done()
				for { //nolint:gosimple
					select {
					case msg, ok := <-rc.c.Chan():
						if !ok {
							close(rc.msgChannel)
							return
						}
						skip := atomic.LoadInt32(&rc.skip)
						if skip != 1 {
							rc.msgChannel <- &rmqMessage{msg: msg}
						} else {
							atomic.StoreInt32(&rc.skip, 0)
						}
					case <-rc.closeCh:
						close(rc.msgChannel)
						rc.c.Close()
						return
					}
				}
			}()
		})
	}
	return rc.msgChannel
}

// Seek is used to seek the position in rocksmq topic
func (rc *Consumer) Seek(id mqwrapper.MessageID, inclusive bool) error {
	msgID := id.(*rmqID).messageID
	// skip the first message when consume
	if !inclusive {
		atomic.StoreInt32(&rc.skip, 1)
	}
	return rc.c.Seek(msgID)
}

// Ack is used to ask a rocksmq message
func (rc *Consumer) Ack(message mqwrapper.Message) {
}

// Close is used to free the resources of this consumer
func (rc *Consumer) Close() {
	close(rc.closeCh)
	rc.wg.Wait()
}

func (rc *Consumer) GetLatestMsgID() (mqwrapper.MessageID, error) {
	msgID, err := rc.c.GetLatestMsgID()
	return &rmqID{messageID: msgID}, err
}
