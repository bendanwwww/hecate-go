package impl

import (
	"github.com/bendanwwww/hecate-go/pkg/core/driver/builder/api"
	"github.com/bendanwwww/hecate-go/pkg/core/entities"
	"github.com/bendanwwww/hecate-go/pkg/framework/tools"
)

type DefaultBuilderImpl[T any] struct {
}

func NewDefaultBuilderImpl[T any](scenes string, nodeWrappers []entities.NodeWrapper[T]) *DefaultBuilderImpl[T] {
	zagMap := &DefaultBuilderImpl[T]{}
	return zagMap
}

func (d DefaultBuilderImpl[T]) GetMapScenes() string {
	//TODO implement me
	panic("implement me")
}

func (d DefaultBuilderImpl[T]) GetNodeNumber() int {
	//TODO implement me
	panic("implement me")
}

func (d DefaultBuilderImpl[T]) GetNodeGroupNumber() int {
	//TODO implement me
	panic("implement me")
}

func (d DefaultBuilderImpl[T]) GetEdgeNumber() int {
	//TODO implement me
	panic("implement me")
}

func (d DefaultBuilderImpl[T]) GetHeads() []int {
	//TODO implement me
	panic("implement me")
}

func (d DefaultBuilderImpl[T]) GetTails() []int {
	//TODO implement me
	panic("implement me")
}

func (d DefaultBuilderImpl[T]) GetVirtualEndNode() int {
	//TODO implement me
	panic("implement me")
}

func (d DefaultBuilderImpl[T]) GetNodeIndex(node entities.ZagNode[T]) int {
	//TODO implement me
	panic("implement me")
}

func (d DefaultBuilderImpl[T]) GetNodeByIndex(index int) entities.ZagNode[T] {
	//TODO implement me
	panic("implement me")
}

func (d DefaultBuilderImpl[T]) GetGroupByGroupIndex(groupIndex int) entities.ZagNodeGroup {
	//TODO implement me
	panic("implement me")
}

func (d DefaultBuilderImpl[T]) HasNextNode(index int) bool {
	//TODO implement me
	panic("implement me")
}

func (d DefaultBuilderImpl[T]) HasCustomGroup(index int) bool {
	//TODO implement me
	panic("implement me")
}

func (d DefaultBuilderImpl[T]) GetNextNodeByIndex(index int) [][]int {
	//TODO implement me
	panic("implement me")
}

func (d DefaultBuilderImpl[T]) GetNodeStrongDependBitByIndex(index int) *tools.BigBinary {
	//TODO implement me
	panic("implement me")
}

func (d DefaultBuilderImpl[T]) GetNodeWeakDependBitByIndex(index int) []tools.BigBinary {
	//TODO implement me
	panic("implement me")
}

func (d DefaultBuilderImpl[T]) NodeCanRun(index int, nodeBinary *tools.BigBinary) bool {
	//TODO implement me
	panic("implement me")
}

func (d DefaultBuilderImpl[T]) NodeNeedCompleteSignal(index int) bool {
	//TODO implement me
	panic("implement me")
}

var _ api.Builder[any] = (*DefaultBuilderImpl[any])(nil)
