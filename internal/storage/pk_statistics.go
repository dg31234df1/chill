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

package storage

import (
	"fmt"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/cockroachdb/errors"

	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	"github.com/milvus-io/milvus/pkg/common"
)

// pkStatistics contains pk field statistic information
type PkStatistics struct {
	PkFilter *bloom.BloomFilter //  bloom filter of pk inside a segment
	MinPK    PrimaryKey         //	minimal pk value, shortcut for checking whether a pk is inside this segment
	MaxPK    PrimaryKey         //  maximal pk value, same above
}

// update set pk min/max value if input value is beyond former range.
func (st *PkStatistics) UpdateMinMax(pk PrimaryKey) error {
	if st == nil {
		return errors.New("nil pk statistics")
	}
	if st.MinPK == nil {
		st.MinPK = pk
	} else if st.MinPK.GT(pk) {
		st.MinPK = pk
	}

	if st.MaxPK == nil {
		st.MaxPK = pk
	} else if st.MaxPK.LT(pk) {
		st.MaxPK = pk
	}

	return nil
}

func (st *PkStatistics) UpdatePKRange(ids FieldData) error {
	switch pks := ids.(type) {
	case *Int64FieldData:
		buf := make([]byte, 8)
		for _, pk := range pks.Data {
			id := NewInt64PrimaryKey(pk)
			err := st.UpdateMinMax(id)
			if err != nil {
				return err
			}
			common.Endian.PutUint64(buf, uint64(pk))
			st.PkFilter.Add(buf)
		}
	case *StringFieldData:
		for _, pk := range pks.Data {
			id := NewVarCharPrimaryKey(pk)
			err := st.UpdateMinMax(id)
			if err != nil {
				return err
			}
			st.PkFilter.AddString(pk)
		}
	default:
		return fmt.Errorf("invalid data type for primary key: %T", ids)
	}
	return nil
}

func (st *PkStatistics) PkExist(pk PrimaryKey) bool {
	// empty pkStatics
	if st.MinPK == nil || st.MaxPK == nil || st.PkFilter == nil {
		return false
	}
	// check pk range first, ugly but key it for now
	if st.MinPK.GT(pk) || st.MaxPK.LT(pk) {
		return false
	}

	// if in range, check bloom filter
	switch pk.Type() {
	case schemapb.DataType_Int64:
		buf := make([]byte, 8)
		int64Pk := pk.(*Int64PrimaryKey)
		common.Endian.PutUint64(buf, uint64(int64Pk.Value))
		return st.PkFilter.Test(buf)
	case schemapb.DataType_VarChar:
		varCharPk := pk.(*VarCharPrimaryKey)
		return st.PkFilter.TestString(varCharPk.Value)
	default:
		// TODO::
	}
	// no idea, just make it as false positive
	return true
}

// Locations returns a list of hash locations representing a data item.
func Locations(pk PrimaryKey, k uint) []uint64 {
	switch pk.Type() {
	case schemapb.DataType_Int64:
		buf := make([]byte, 8)
		int64Pk := pk.(*Int64PrimaryKey)
		common.Endian.PutUint64(buf, uint64(int64Pk.Value))
		return bloom.Locations(buf, k)
	case schemapb.DataType_VarChar:
		varCharPk := pk.(*VarCharPrimaryKey)
		return bloom.Locations([]byte(varCharPk.Value), k)
	default:
		// TODO::
	}
	return nil
}

func (st *PkStatistics) TestLocations(pk PrimaryKey, locs []uint64) bool {
	// empty pkStatics
	if st.MinPK == nil || st.MaxPK == nil || st.PkFilter == nil {
		return false
	}

	// check bf first, TestLocation just do some bitset compute, cost is cheaper
	if !st.PkFilter.TestLocations(locs) {
		return false
	}

	// check pk range first, ugly but key it for now
	return st.MinPK.LE(pk) && st.MaxPK.GE(pk)
}

func (st *PkStatistics) TestLocationCache(lc LocationsCache) bool {
	// empty pkStatics
	if st.MinPK == nil || st.MaxPK == nil || st.PkFilter == nil {
		return false
	}

	// check bf first, TestLocation just do some bitset compute, cost is cheaper
	if !st.PkFilter.TestLocations(lc.Locations(st.PkFilter.K())) {
		return false
	}

	// check pk range after
	return st.MinPK.LE(lc.pk) && st.MaxPK.GE(lc.pk)
}

// LocationsCache is a helper struct caching pk bloom filter locations.
// Note that this helper is not concurrent safe and shall be used in same goroutine.
type LocationsCache struct {
	pk        PrimaryKey
	locations map[uint][]uint64
}

func (lc LocationsCache) Locations(k uint) []uint64 {
	locs, ok := lc.locations[k]
	if ok {
		return locs
	}
	locs = Locations(lc.pk, k)
	lc.locations[k] = locs
	return locs
}

func NewLocationsCache(pk PrimaryKey) LocationsCache {
	return LocationsCache{
		pk:        pk,
		locations: make(map[uint][]uint64),
	}
}
