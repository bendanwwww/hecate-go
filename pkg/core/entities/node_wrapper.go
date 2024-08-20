package entities

import (
	"github.com/bendanwwww/hecate-go/pkg/framework/constant"
	mapset "github.com/deckarep/golang-set"
)

type NodeWrapper[T any] interface {
	GetZagNode() ZagNode[T]
	GetGroupNameList() []string
	GetGroupTimeout(groupName string) int64
	GetStrongNextNodes() mapset.Set
	GetWeakNextNodes() map[string]mapset.Set
	AddGroupName(groupName string)
	AddGroupNameWithTimeout(groupName string, timeout int64)
	SetStrongNextNodes(strongDependNodes ...NodeWrapper[T])
	SetWeakNextNodes(weakDependNodes ...NodeWrapper[T])
	SetWeakGroupNextNodes(groupName string, weakDependNodes ...NodeWrapper[T])
}

type NodeWrapperField[T any] struct {
	/** node base info */
	zagNode ZagNode[T]
	/** node group list */
	groupNames []string
	/** group timeout duration */
	groupTimeout map[string]int64
	/** the node next nodes */
	StrongNextNodes mapset.Set
	/** the node weak next nodes (key: weak group's name value: group's nodes set) */
	weakNextNodes map[string]mapset.Set
}

func NewZagNode[T any]() *NodeWrapperField[T] {
	var node ZagNode[T] = &ZagNodeField[T]{
		nodeId:        -1,
		headNodeIndex: -1,
		nodeTimeout:   -1,
	}
	return &NodeWrapperField[T]{
		zagNode:      node,
		groupNames:   []string{constant.DefaultKey},
		groupTimeout: map[string]int64{},
	}
}

func (nodeWrapper *NodeWrapperField[T]) Name(name string) {
	nodeWrapper.zagNode.SetNodeName(name)
}

func (nodeWrapper *NodeWrapperField[T]) Operator(nodeOperator NodeOperator[T]) {
	nodeWrapper.zagNode.SetNodeOperator(nodeOperator)
}

func (nodeWrapper *NodeWrapperField[T]) Next(isWeak bool, nodeWrappers ...NodeWrapper[T]) {
	if !isWeak {
		if nodeWrapper.StrongNextNodes == nil {
			nodeWrapper.StrongNextNodes = mapset.NewSet()
		}
		for _, item := range nodeWrappers {
			nodeWrapper.StrongNextNodes.Add(item)
		}
	} else {
		if nodeWrapper.weakNextNodes == nil {
			nodeWrapper.weakNextNodes = make(map[string]mapset.Set)
		}
		for _, item := range nodeWrappers {
			if nodeWrapper.weakNextNodes[constant.DefaultKey+"_group"] == nil {
				nodeWrapper.weakNextNodes[constant.DefaultKey+"_group"] = mapset.NewSet()
			}
			nodeWrapper.weakNextNodes[constant.DefaultKey+"_group"].Add(item)
		}
	}
}

func (nodeWrapper *NodeWrapperField[T]) WeakNexts(groupName string, nodeWrappers ...NodeWrapper[T]) {
	if nodeWrapper.weakNextNodes == nil {
		nodeWrapper.weakNextNodes = make(map[string]mapset.Set)
	}
	for _, item := range nodeWrappers {
		if nodeWrapper.weakNextNodes[groupName] == nil {
			nodeWrapper.weakNextNodes[groupName] = mapset.NewSet()
		}
		nodeWrapper.weakNextNodes[groupName].Add(item)
	}
}

func (nodeWrapper *NodeWrapperField[T]) NodeChooseFunc(nodeChoose IChoose[T]) {
	nodeWrapper.zagNode.SetNodeChoose(nodeChoose)
}

func (nodeWrapper *NodeWrapperField[T]) NeedFinishSignal() {
	nodeWrapper.zagNode.SetNodeFinishSignal(true)
}

func (nodeWrapper *NodeWrapperField[T]) NodeRunTimeout(nodeTimeout int64) {
	nodeWrapper.zagNode.SetNodeTimeout(nodeTimeout)
}

func (nodeWrapper *NodeWrapperField[T]) NodeMustRun() {
	nodeWrapper.zagNode.SetMustRunState()
}

func (nodeWrapper *NodeWrapperField[T]) GetZagNode() ZagNode[T] {
	return nodeWrapper.zagNode
}

func (nodeWrapper *NodeWrapperField[T]) GetGroupNameList() []string {
	return nodeWrapper.groupNames
}

func (nodeWrapper *NodeWrapperField[T]) GetGroupTimeout(groupName string) int64 {
	timeout, isIn := nodeWrapper.groupTimeout[groupName]
	if !isIn {
		return -1
	}
	return timeout
}

func (nodeWrapper *NodeWrapperField[T]) GetStrongNextNodes() mapset.Set {
	return nodeWrapper.StrongNextNodes
}

func (nodeWrapper *NodeWrapperField[T]) GetWeakNextNodes() map[string]mapset.Set {
	return nodeWrapper.weakNextNodes
}

func (nodeWrapper *NodeWrapperField[T]) AddGroupName(groupName string) {
	nodeWrapper.AddGroupNameWithTimeout(groupName, -1)
}

func (nodeWrapper *NodeWrapperField[T]) AddGroupNameWithTimeout(groupName string, timeout int64) {
	if len(nodeWrapper.groupNames) == 1 && nodeWrapper.groupNames[0] == constant.DefaultKey {
		nodeWrapper.groupNames = []string{}
	}
	nodeWrapper.groupNames = append(nodeWrapper.groupNames, groupName)
	nodeWrapper.groupTimeout[groupName] = timeout
}

func (nodeWrapper *NodeWrapperField[T]) SetStrongNextNodes(strongDependNodes ...NodeWrapper[T]) {
	if nodeWrapper.StrongNextNodes == nil {
		nodeWrapper.StrongNextNodes = mapset.NewSet()
	}
	for _, item := range strongDependNodes {
		nodeWrapper.StrongNextNodes.Add(item)
	}
}

func (nodeWrapper *NodeWrapperField[T]) SetWeakNextNodes(weakDependNodes ...NodeWrapper[T]) {
	if nodeWrapper.weakNextNodes == nil {
		nodeWrapper.weakNextNodes = map[string]mapset.Set{}
	}
	weakDepends := mapset.NewSet()
	for _, item := range weakDependNodes {
		weakDepends.Add(item)
	}
	nodeWrapper.weakNextNodes[constant.DefaultKey] = weakDepends
}

type AA struct {
	bs []byte
}

func (nodeWrapper *NodeWrapperField[T]) SetWeakGroupNextNodes(groupName string, weakDependNodes ...NodeWrapper[T]) {
	if nodeWrapper.weakNextNodes == nil {
		nodeWrapper.weakNextNodes = map[string]mapset.Set{}
	}
	weakDepends := mapset.NewSet()
	for _, item := range weakDependNodes {
		weakDepends.Add(item)
	}
	nodeWrapper.weakNextNodes[groupName] = weakDepends
}

var _ NodeWrapper[any] = (*NodeWrapperField[any])(nil)
