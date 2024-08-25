package api

import (
	zagEntities "github.com/bendanwwww/hecate-go/pkg/core/entities"
)

type BuilderNode[T any] interface {
	// Name set node's name
	Name(name string)
	// Operator set node's operator
	Operator(nodeOperator zagEntities.NodeOperator[T])
	// Next set node's next node
	Next(isWeak bool, nodes ...zagEntities.NodeWrapper[T])
	// WeakNexts set node's next weak node group
	WeakNexts(groupName string, nodes ...zagEntities.NodeWrapper[T])
	// NodeChooseFunc set node's choose next nodes function
	NodeChooseFunc(nodeChoose zagEntities.IChoose[any])
	// NeedFinishSignal  set node need send sign when executed
	NeedFinishSignal()
	// NodeRunTimeout set node's timeout duration
	NodeRunTimeout(nodeTimeout int64)
	// NodeMustRun set node must be executed
	NodeMustRun()
}
