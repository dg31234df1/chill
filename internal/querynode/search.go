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
	"fmt"
	"sync"

	"github.com/cockroachdb/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"

	"github.com/milvus-io/milvus-proto/go-api/commonpb"
	"github.com/milvus-io/milvus/pkg/log"
	"github.com/milvus-io/milvus/pkg/metrics"
	"github.com/milvus-io/milvus/pkg/util/paramtable"
	"github.com/milvus-io/milvus/pkg/util/timerecord"
	"github.com/milvus-io/milvus/pkg/util/typeutil"
)

// searchOnSegments performs search on listed segments
// all segment ids are validated before calling this function
func searchSegments(ctx context.Context, replica ReplicaInterface, segType segmentType, searchReq *searchRequest, segIDs []UniqueID) ([]*SearchResult, error) {
	var (
		// results variables
		resultCh = make(chan *SearchResult, len(segIDs))
		errs     = make([]error, len(segIDs))
		wg       sync.WaitGroup

		// For log only
		mu                   sync.Mutex
		segmentsWithoutIndex []UniqueID
	)

	searchLabel := metrics.SealedSegmentLabel
	if segType == commonpb.SegmentState_Growing {
		searchLabel = metrics.GrowingSegmentLabel
	}

	// calling segment search in goroutines
	for i, segID := range segIDs {
		wg.Add(1)
		go func(segID UniqueID, i int) {
			defer wg.Done()
			ctx, span := otel.Tracer(typeutil.QueryNodeRole).Start(ctx, "Search-Segment")
			span.SetAttributes(attribute.String("segmentType", searchLabel))
			span.SetAttributes(attribute.Int64("segmentID", segID))
			defer span.End()

			seg, err := replica.getSegmentByID(segID, segType)
			if err != nil {
				if errors.Is(err, ErrSegmentNotFound) {
					return
				}
				log.Error(err.Error()) // should not happen but still ignore it since the result is still correct
				return
			}

			if !seg.hasLoadIndexForIndexedField(searchReq.searchFieldID) {
				mu.Lock()
				segmentsWithoutIndex = append(segmentsWithoutIndex, segID)
				mu.Unlock()
			}
			// record search time
			tr := timerecord.NewTimeRecorder("searchOnSegments")
			searchResult, err := seg.search(ctx, searchReq)
			errs[i] = err
			resultCh <- searchResult
			// update metrics
			metrics.QueryNodeSQSegmentLatency.WithLabelValues(fmt.Sprint(paramtable.GetNodeID()),
				metrics.SearchLabel, searchLabel).Observe(float64(tr.ElapseSpan().Milliseconds()))
		}(segID, i)
	}
	wg.Wait()
	close(resultCh)

	searchResults := make([]*SearchResult, 0, len(segIDs))
	for result := range resultCh {
		searchResults = append(searchResults, result)
	}

	for _, err := range errs {
		if err != nil {
			deleteSearchResults(searchResults)
			return nil, err
		}
	}

	if len(segmentsWithoutIndex) > 0 {
		log.Ctx(ctx).Info("search growing/sealed segments without indexes", zap.Int64s("segmentIDs", segmentsWithoutIndex))
	}

	return searchResults, nil
}

// search will search on the historical segments the target segments in historical.
// if segIDs is not specified, it will search on all the historical segments speficied by partIDs.
// if segIDs is specified, it will only search on the segments specified by the segIDs.
// if partIDs is empty, it means all the partitions of the loaded collection or all the partitions loaded.
func searchHistorical(ctx context.Context, replica ReplicaInterface, searchReq *searchRequest, collID UniqueID, partIDs []UniqueID, segIDs []UniqueID) ([]*SearchResult, []UniqueID, []UniqueID, error) {
	var err error
	var searchResults []*SearchResult
	var searchSegmentIDs []UniqueID
	var searchPartIDs []UniqueID
	searchPartIDs, searchSegmentIDs, err = validateOnHistoricalReplica(ctx, replica, collID, partIDs, segIDs)
	if err != nil {
		return searchResults, searchSegmentIDs, searchPartIDs, err
	}
	searchResults, err = searchSegments(ctx, replica, segmentTypeSealed, searchReq, searchSegmentIDs)
	return searchResults, searchPartIDs, searchSegmentIDs, err
}

// searchStreaming will search all the target segments in streaming
// if partIDs is empty, it means all the partitions of the loaded collection or all the partitions loaded.
func searchStreaming(ctx context.Context, replica ReplicaInterface, searchReq *searchRequest, collID UniqueID, partIDs []UniqueID, vChannel Channel) ([]*SearchResult, []UniqueID, []UniqueID, error) {
	var err error
	var searchResults []*SearchResult
	var searchPartIDs []UniqueID
	var searchSegmentIDs []UniqueID

	searchPartIDs, searchSegmentIDs, err = validateOnStreamReplica(ctx, replica, collID, partIDs, vChannel)
	if err != nil {
		return searchResults, searchSegmentIDs, searchPartIDs, err
	}
	searchResults, err = searchSegments(ctx, replica, segmentTypeGrowing, searchReq, searchSegmentIDs)
	return searchResults, searchPartIDs, searchSegmentIDs, err
}
