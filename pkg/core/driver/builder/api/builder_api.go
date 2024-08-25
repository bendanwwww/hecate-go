package api

import (
	zagEntities "github.com/bendanwwww/hecate-go/pkg/core/entities"
	"github.com/bendanwwww/hecate-go/pkg/framework/tools"
)

type Builder[T any] interface {
	// GetMapScenes get the map's scenes
	GetMapScenes() string
	// GetNodeNumber get node number in the map
	GetNodeNumber() int
	// GetNodeGroupNumber get node group number in the map
	GetNodeGroupNumber() int
	// GetEdgeNumber get edge number in the map
	GetEdgeNumber() int
	// GetHeads get all head nodes in the map
	GetHeads() []int
	// GetTails get all tail nodes in the map
	GetTails() []int
	// GetVirtualEndNode get virtual end node
	GetVirtualEndNode() int
	// GetNodeIndex get node index
	GetNodeIndex(node zagEntities.ZagNode[T]) int
	// GetNodeByIndex get node by index
	GetNodeByIndex(index int) zagEntities.ZagNode[T]
	// GetGroupByGroupIndex get node group by group index
	GetGroupByGroupIndex(groupIndex int) zagEntities.ZagNodeGroup
	// HasNextNode check the node has next node
	HasNextNode(index int) bool
	// HasCustomGroup check the node is below at least one custom node group
	HasCustomGroup(index int) bool
	// GetNextNodeByIndex get the node next nodes
	// return an array. res[i][0] is the next node index, res[i][1] is the node and the next node's edge index.
	GetNextNodeByIndex(index int) [][]int
	// GetNodeStrongDependBitByIndex get the node binary
	GetNodeStrongDependBitByIndex(index int) *tools.BigBinary
	// GetNodeWeakDependBitByIndex get the node weak edge group binary
	GetNodeWeakDependBitByIndex(index int) []tools.BigBinary
	// NodeCanRun check the node can be executed
	NodeCanRun(index int, nodeBinary *tools.BigBinary) bool
	// NodeNeedCompleteSignal the node send a sign when executed
	NodeNeedCompleteSignal(index int) bool
}
