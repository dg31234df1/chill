package querynode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition_Segments(t *testing.T) {
	node := newQueryNodeMock()
	collectionID := UniqueID(0)
	initTestMeta(t, node, collectionID, 0)

	collection, err := node.replica.getCollectionByID(collectionID)
	assert.NoError(t, err)

	partitions := collection.Partitions()
	targetPartition := (*partitions)[0]

	const segmentNum = 3
	for i := 0; i < segmentNum; i++ {
		err := node.replica.addSegment2(UniqueID(i), targetPartition.partitionTag, collection.ID(), segTypeGrowing)
		assert.NoError(t, err)
	}

	segments := targetPartition.Segments()
	assert.Equal(t, segmentNum+1, len(*segments))
}

func TestPartition_newPartition(t *testing.T) {
	partitionTag := "default"
	partition := newPartition2(partitionTag)
	assert.Equal(t, partition.partitionTag, partitionTag)
}
