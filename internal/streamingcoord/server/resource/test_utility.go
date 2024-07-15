//go:build test
// +build test

package resource

import (
	"github.com/milvus-io/milvus/internal/streamingnode/client/manager"
)

// OptStreamingManagerClient provides streaming manager client to the resource.
func OptStreamingManagerClient(c manager.ManagerClient) optResourceInit {
	return func(r *resourceImpl) {
		r.streamingNodeManagerClient = c
	}
}

// InitForTest initializes the singleton of resources for test.
func InitForTest(opts ...optResourceInit) {
	r = &resourceImpl{}
	for _, opt := range opts {
		opt(r)
	}
}
