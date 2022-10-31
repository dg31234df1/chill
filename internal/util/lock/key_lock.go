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

package lock

import (
	"sync"

	"github.com/milvus-io/milvus/internal/log"
	"go.uber.org/zap"
)

type RefLock struct {
	mutex      sync.RWMutex
	refCounter int
}

func (m *RefLock) ref() {
	m.refCounter++
}

func (m *RefLock) unref() {
	m.refCounter--
}

func newRefLock() *RefLock {
	c := RefLock{
		sync.RWMutex{},
		0,
	}
	return &c
}

type KeyLock struct {
	keyLocksMutex sync.Mutex
	refLocks      map[string]*RefLock
}

func NewKeyLock() *KeyLock {
	keyLock := KeyLock{
		refLocks: make(map[string]*RefLock),
	}
	return &keyLock
}

func (k *KeyLock) Lock(key string) {
	k.keyLocksMutex.Lock()
	// update the key map
	if keyLock, ok := k.refLocks[key]; ok {
		keyLock.ref()

		k.keyLocksMutex.Unlock()
		keyLock.mutex.Lock()
	} else {
		newKLock := newRefLock()
		newKLock.mutex.Lock()
		k.refLocks[key] = newKLock
		newKLock.ref()

		k.keyLocksMutex.Unlock()
		return
	}
}

func (k *KeyLock) Unlock(lockedKey string) {
	k.keyLocksMutex.Lock()
	defer k.keyLocksMutex.Unlock()
	keyLock, ok := k.refLocks[lockedKey]
	if !ok {
		log.Warn("Unlocking non-existing key", zap.String("key", lockedKey))
		return
	}
	keyLock.unref()
	if keyLock.refCounter == 0 {
		delete(k.refLocks, lockedKey)
	}
	keyLock.mutex.Unlock()
}

func (k *KeyLock) RLock(key string) {
	k.keyLocksMutex.Lock()
	// update the key map
	if keyLock, ok := k.refLocks[key]; ok {
		keyLock.ref()

		k.keyLocksMutex.Unlock()
		keyLock.mutex.RLock()
	} else {
		newKLock := newRefLock()
		newKLock.mutex.RLock()
		k.refLocks[key] = newKLock
		newKLock.ref()

		k.keyLocksMutex.Unlock()
		return
	}
}

func (k *KeyLock) RUnlock(lockedKey string) {
	k.keyLocksMutex.Lock()
	defer k.keyLocksMutex.Unlock()
	keyLock, ok := k.refLocks[lockedKey]
	if !ok {
		log.Warn("Unlocking non-existing key", zap.String("key", lockedKey))
		return
	}
	keyLock.unref()
	if keyLock.refCounter == 0 {
		delete(k.refLocks, lockedKey)
	}
	keyLock.mutex.RUnlock()
}

func (k *KeyLock) size() int {
	k.keyLocksMutex.Lock()
	defer k.keyLocksMutex.Unlock()
	return len(k.refLocks)
}
