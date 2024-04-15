package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/milvus-io/milvus-proto/go-api/v2/rgpb"
	"github.com/milvus-io/milvus/internal/proto/querypb"
)

func TestResourceGroup(t *testing.T) {
	cfg := &rgpb.ResourceGroupConfig{
		Requests: &rgpb.ResourceGroupLimit{
			NodeNum: 1,
		},
		Limits: &rgpb.ResourceGroupLimit{
			NodeNum: 2,
		},
		TransferFrom: []*rgpb.ResourceGroupTransfer{{
			ResourceGroup: "rg2",
		}},
		TransferTo: []*rgpb.ResourceGroupTransfer{{
			ResourceGroup: "rg3",
		}},
	}
	rg := NewResourceGroup("rg1", cfg)
	cfg2 := rg.GetConfig()
	assert.Equal(t, cfg.Requests.NodeNum, cfg2.Requests.NodeNum)

	assertion := func() {
		assert.Equal(t, "rg1", rg.GetName())
		assert.Empty(t, rg.GetNodes())
		assert.Zero(t, rg.NodeNum())
		assert.Zero(t, rg.OversizedNumOfNodes())
		assert.Zero(t, rg.RedundantNumOfNodes())
		assert.Equal(t, 1, rg.MissingNumOfNodes())
		assert.Equal(t, 2, rg.ReachLimitNumOfNodes())
		assert.True(t, rg.HasFrom("rg2"))
		assert.False(t, rg.HasFrom("rg3"))
		assert.True(t, rg.HasTo("rg3"))
		assert.False(t, rg.HasTo("rg2"))
		assert.False(t, rg.ContainNode(1))
		assert.Error(t, rg.MeetRequirement())
	}
	assertion()

	// Test Txn
	mrg := rg.CopyForWrite()
	cfg = &rgpb.ResourceGroupConfig{
		Requests: &rgpb.ResourceGroupLimit{
			NodeNum: 2,
		},
		Limits: &rgpb.ResourceGroupLimit{
			NodeNum: 3,
		},
		TransferFrom: []*rgpb.ResourceGroupTransfer{{
			ResourceGroup: "rg3",
		}},
		TransferTo: []*rgpb.ResourceGroupTransfer{{
			ResourceGroup: "rg2",
		}},
	}
	mrg.UpdateConfig(cfg)

	// nothing happens before commit.
	assertion()

	rg = mrg.ToResourceGroup()
	assertion = func() {
		assert.Equal(t, "rg1", rg.GetName())
		assert.Empty(t, rg.GetNodes())
		assert.Zero(t, rg.NodeNum())
		assert.Zero(t, rg.OversizedNumOfNodes())
		assert.Zero(t, rg.RedundantNumOfNodes())
		assert.Equal(t, 2, rg.MissingNumOfNodes())
		assert.Equal(t, 3, rg.ReachLimitNumOfNodes())
		assert.True(t, rg.HasFrom("rg3"))
		assert.False(t, rg.HasFrom("rg2"))
		assert.True(t, rg.HasTo("rg2"))
		assert.False(t, rg.HasTo("rg3"))
		assert.False(t, rg.ContainNode(1))
		assert.Error(t, rg.MeetRequirement())
	}
	assertion()

	// Test AddNode
	mrg = rg.CopyForWrite()
	mrg.AssignNode(1)
	mrg.AssignNode(1)
	assertion()
	rg = mrg.ToResourceGroup()

	assertion = func() {
		assert.Equal(t, "rg1", rg.GetName())
		assert.ElementsMatch(t, []int64{1}, rg.GetNodes())
		assert.Equal(t, 1, rg.NodeNum())
		assert.Zero(t, rg.OversizedNumOfNodes())
		assert.Zero(t, rg.RedundantNumOfNodes())
		assert.Equal(t, 1, rg.MissingNumOfNodes())
		assert.Equal(t, 2, rg.ReachLimitNumOfNodes())
		assert.True(t, rg.HasFrom("rg3"))
		assert.False(t, rg.HasFrom("rg2"))
		assert.True(t, rg.HasTo("rg2"))
		assert.False(t, rg.HasTo("rg3"))
		assert.True(t, rg.ContainNode(1))
		assert.Error(t, rg.MeetRequirement())
	}
	assertion()

	// Test AddNode until meet requirement.
	mrg = rg.CopyForWrite()
	mrg.AssignNode(2)
	assertion()
	rg = mrg.ToResourceGroup()

	assertion = func() {
		assert.Equal(t, "rg1", rg.GetName())
		assert.ElementsMatch(t, []int64{1, 2}, rg.GetNodes())
		assert.Equal(t, 2, rg.NodeNum())
		assert.Zero(t, rg.OversizedNumOfNodes())
		assert.Zero(t, rg.RedundantNumOfNodes())
		assert.Equal(t, 0, rg.MissingNumOfNodes())
		assert.Equal(t, 1, rg.ReachLimitNumOfNodes())
		assert.True(t, rg.HasFrom("rg3"))
		assert.False(t, rg.HasFrom("rg2"))
		assert.True(t, rg.HasTo("rg2"))
		assert.False(t, rg.HasTo("rg3"))
		assert.True(t, rg.ContainNode(1))
		assert.True(t, rg.ContainNode(2))
		assert.NoError(t, rg.MeetRequirement())
	}
	assertion()

	// Test AddNode until exceed requirement.
	mrg = rg.CopyForWrite()
	mrg.AssignNode(3)
	mrg.AssignNode(4)
	assertion()
	rg = mrg.ToResourceGroup()

	assertion = func() {
		assert.Equal(t, "rg1", rg.GetName())
		assert.ElementsMatch(t, []int64{1, 2, 3, 4}, rg.GetNodes())
		assert.Equal(t, 4, rg.NodeNum())
		assert.Equal(t, 2, rg.OversizedNumOfNodes())
		assert.Equal(t, 1, rg.RedundantNumOfNodes())
		assert.Equal(t, 0, rg.MissingNumOfNodes())
		assert.Equal(t, 0, rg.ReachLimitNumOfNodes())
		assert.True(t, rg.HasFrom("rg3"))
		assert.False(t, rg.HasFrom("rg2"))
		assert.True(t, rg.HasTo("rg2"))
		assert.False(t, rg.HasTo("rg3"))
		assert.True(t, rg.ContainNode(1))
		assert.True(t, rg.ContainNode(2))
		assert.True(t, rg.ContainNode(3))
		assert.True(t, rg.ContainNode(4))
		assert.Error(t, rg.MeetRequirement())
	}
	assertion()

	// Test UnassignNode.
	mrg = rg.CopyForWrite()
	mrg.UnassignNode(3)
	assertion()
	rg = mrg.ToResourceGroup()
	rgMeta := rg.GetMeta()
	assert.Equal(t, 3, len(rgMeta.Nodes))
	assert.Equal(t, "rg1", rgMeta.Name)
	assert.Equal(t, "rg3", rgMeta.Config.TransferFrom[0].ResourceGroup)
	assert.Equal(t, "rg2", rgMeta.Config.TransferTo[0].ResourceGroup)
	assert.Equal(t, int32(2), rgMeta.Config.Requests.NodeNum)
	assert.Equal(t, int32(3), rgMeta.Config.Limits.NodeNum)

	assertion2 := func(rg *ResourceGroup) {
		assert.Equal(t, "rg1", rg.GetName())
		assert.ElementsMatch(t, []int64{1, 2, 4}, rg.GetNodes())
		assert.Equal(t, 3, rg.NodeNum())
		assert.Equal(t, 1, rg.OversizedNumOfNodes())
		assert.Equal(t, 0, rg.RedundantNumOfNodes())
		assert.Equal(t, 0, rg.MissingNumOfNodes())
		assert.Equal(t, 0, rg.ReachLimitNumOfNodes())
		assert.True(t, rg.HasFrom("rg3"))
		assert.False(t, rg.HasFrom("rg2"))
		assert.True(t, rg.HasTo("rg2"))
		assert.False(t, rg.HasTo("rg3"))
		assert.True(t, rg.ContainNode(1))
		assert.True(t, rg.ContainNode(2))
		assert.False(t, rg.ContainNode(3))
		assert.True(t, rg.ContainNode(4))
		assert.NoError(t, rg.MeetRequirement())
	}
	assertion2(rg)

	// snapshot do not change the original resource group.
	snapshot := rg.Snapshot()
	assertion2(snapshot)
	snapshot.cfg = nil
	snapshot.name = "rg2"
	snapshot.nodes = nil
	assertion2(rg)
}

func TestResourceGroupMeta(t *testing.T) {
	rgMeta := &querypb.ResourceGroup{
		Name:     "rg1",
		Capacity: 1,
		Nodes:    []int64{1, 2},
	}
	rg := NewResourceGroupFromMeta(rgMeta)
	assert.Equal(t, "rg1", rg.GetName())
	assert.ElementsMatch(t, []int64{1, 2}, rg.GetNodes())
	assert.Equal(t, 2, rg.NodeNum())
	assert.Equal(t, 1, rg.OversizedNumOfNodes())
	assert.Equal(t, 1, rg.RedundantNumOfNodes())
	assert.Equal(t, 0, rg.MissingNumOfNodes())
	assert.Equal(t, 0, rg.ReachLimitNumOfNodes())
	assert.False(t, rg.HasFrom("rg3"))
	assert.False(t, rg.HasFrom("rg2"))
	assert.False(t, rg.HasTo("rg2"))
	assert.False(t, rg.HasTo("rg3"))
	assert.True(t, rg.ContainNode(1))
	assert.True(t, rg.ContainNode(2))
	assert.False(t, rg.ContainNode(3))
	assert.False(t, rg.ContainNode(4))
	assert.Error(t, rg.MeetRequirement())

	rgMeta = &querypb.ResourceGroup{
		Name:     "rg1",
		Capacity: 1,
		Nodes:    []int64{1, 2, 4},
		Config: &rgpb.ResourceGroupConfig{
			Requests: &rgpb.ResourceGroupLimit{
				NodeNum: 2,
			},
			Limits: &rgpb.ResourceGroupLimit{
				NodeNum: 3,
			},
			TransferFrom: []*rgpb.ResourceGroupTransfer{{
				ResourceGroup: "rg3",
			}},
			TransferTo: []*rgpb.ResourceGroupTransfer{{
				ResourceGroup: "rg2",
			}},
		},
	}
	rg = NewResourceGroupFromMeta(rgMeta)
	assert.Equal(t, "rg1", rg.GetName())
	assert.ElementsMatch(t, []int64{1, 2, 4}, rg.GetNodes())
	assert.Equal(t, 3, rg.NodeNum())
	assert.Equal(t, 1, rg.OversizedNumOfNodes())
	assert.Equal(t, 0, rg.RedundantNumOfNodes())
	assert.Equal(t, 0, rg.MissingNumOfNodes())
	assert.Equal(t, 0, rg.ReachLimitNumOfNodes())
	assert.True(t, rg.HasFrom("rg3"))
	assert.False(t, rg.HasFrom("rg2"))
	assert.True(t, rg.HasTo("rg2"))
	assert.False(t, rg.HasTo("rg3"))
	assert.True(t, rg.ContainNode(1))
	assert.True(t, rg.ContainNode(2))
	assert.False(t, rg.ContainNode(3))
	assert.True(t, rg.ContainNode(4))
	assert.NoError(t, rg.MeetRequirement())

	newMeta := rg.GetMeta()
	assert.Equal(t, int32(2), newMeta.Capacity)

	// Recover Default Resource Group.
	rgMeta = &querypb.ResourceGroup{
		Name:     DefaultResourceGroupName,
		Capacity: defaultResourceGroupCapacity,
		Nodes:    []int64{1, 2},
	}
	rg = NewResourceGroupFromMeta(rgMeta)
	assert.Equal(t, DefaultResourceGroupName, rg.GetName())
	assert.ElementsMatch(t, []int64{1, 2}, rg.GetNodes())
	assert.Equal(t, 2, rg.NodeNum())
	assert.Equal(t, 2, rg.OversizedNumOfNodes())
	assert.Equal(t, 0, rg.RedundantNumOfNodes())
	assert.Equal(t, 0, rg.MissingNumOfNodes())
	assert.Equal(t, int(defaultResourceGroupCapacity-2), rg.ReachLimitNumOfNodes())
	assert.False(t, rg.HasFrom("rg3"))
	assert.False(t, rg.HasFrom("rg2"))
	assert.False(t, rg.HasTo("rg2"))
	assert.False(t, rg.HasTo("rg3"))
	assert.True(t, rg.ContainNode(1))
	assert.True(t, rg.ContainNode(2))
	assert.False(t, rg.ContainNode(3))
	assert.False(t, rg.ContainNode(4))
	assert.NoError(t, rg.MeetRequirement())

	newMeta = rg.GetMeta()
	assert.Equal(t, defaultResourceGroupCapacity, newMeta.Capacity)

	// Recover Default Resource Group.
	rgMeta = &querypb.ResourceGroup{
		Name:  DefaultResourceGroupName,
		Nodes: []int64{1, 2},
		Config: &rgpb.ResourceGroupConfig{
			Requests: &rgpb.ResourceGroupLimit{
				NodeNum: 2,
			},
			Limits: &rgpb.ResourceGroupLimit{
				NodeNum: 3,
			},
			TransferFrom: []*rgpb.ResourceGroupTransfer{{
				ResourceGroup: "rg3",
			}},
			TransferTo: []*rgpb.ResourceGroupTransfer{{
				ResourceGroup: "rg2",
			}},
		},
	}
	rg = NewResourceGroupFromMeta(rgMeta)
	assert.Equal(t, DefaultResourceGroupName, rg.GetName())
	assert.ElementsMatch(t, []int64{1, 2}, rg.GetNodes())
	assert.Equal(t, 2, rg.NodeNum())
	assert.Equal(t, 0, rg.OversizedNumOfNodes())
	assert.Equal(t, 0, rg.RedundantNumOfNodes())
	assert.Equal(t, 0, rg.MissingNumOfNodes())
	assert.Equal(t, 1, rg.ReachLimitNumOfNodes())
	assert.True(t, rg.HasFrom("rg3"))
	assert.False(t, rg.HasFrom("rg2"))
	assert.True(t, rg.HasTo("rg2"))
	assert.False(t, rg.HasTo("rg3"))
	assert.True(t, rg.ContainNode(1))
	assert.True(t, rg.ContainNode(2))
	assert.False(t, rg.ContainNode(3))
	assert.False(t, rg.ContainNode(4))
	assert.NoError(t, rg.MeetRequirement())

	newMeta = rg.GetMeta()
	assert.Equal(t, int32(1000000), newMeta.Capacity)
}
