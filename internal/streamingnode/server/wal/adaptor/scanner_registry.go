package adaptor

import (
	"fmt"

	"github.com/milvus-io/milvus/internal/streamingnode/server/wal"
	"github.com/milvus-io/milvus/internal/util/streamingutil/util"
	"github.com/milvus-io/milvus/pkg/streaming/util/types"
)

type scannerRegistry struct {
	channel     types.PChannelInfo
	idAllocator *util.IDAllocator
}

// AllocateScannerName a scanner name for a scanner.
// The scanner name should be persistent on meta for garbage clean up.
func (m *scannerRegistry) AllocateScannerName() (string, error) {
	name := m.newSubscriptionName()
	// TODO: persistent the subscription name on meta.
	return name, nil
}

func (m *scannerRegistry) RegisterNewScanner(string, wal.Scanner) {
}

// newSubscriptionName generates a new subscription name.
func (m *scannerRegistry) newSubscriptionName() string {
	id := m.idAllocator.Allocate()
	return fmt.Sprintf("%s/%d/%d", m.channel.Name, m.channel.Term, id)
}
