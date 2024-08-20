package entities

type ZagNode[T any] interface {
	GetNodeId() int
	GetHeadNodeIndex() int
	Access() bool
	GetNodeArrayOffset() int
	GetNodeName() string
	GetGroupIndexList() []int
	IsHead() bool
	IsTail() bool
	GetNodeOperator() NodeOperator[T]
	GetNodeChoose() IChoose[T]
	GetNodeFinishSignal() bool
	GetNodeTimeout() int64
	MustRun() bool
	SetNodeId(id int)
	SetHeadNodeIndex(headNodeIndex int)
	SetAccess(access bool)
	SetNodeArrayOffset(nodeArrayOffset int)
	SetNodeName(nodeName string)
	AddGroup(groupIndex int)
	SetHeadState(state bool)
	SetTailState(state bool)
	SetNodeOperator(nodeOperator NodeOperator[T])
	SetNodeChoose(nodeChoose IChoose[T])
	SetNodeFinishSignal(needFinishSignal bool)
	SetNodeTimeout(nodeTimeout int64)
	SetMustRunState()
}

type ZagNodeField[T any] struct {
	/** node id */
	nodeId int
	/** head node index in array. only useful with head node. only use it in build function */
	headNodeIndex int
	/** sign a node is access. only use it in build function */
	access bool
	/** the node offset in array. only use it in build function */
	nodeArrayOffset int
	/** node name */
	nodeName string
	/** node's below groups  */
	nodeGroupList []int
	/** the node is a head in the map */
	isHead bool
	/** the node is a tail in the node */
	isTail bool
	/** node operator */
	nodeOperator NodeOperator[T]
	/** node choose function  */
	nodeChoose IChoose[T]
	/** the node send a sign when executed */
	nodeFinishSignal bool
	/** node execute timeout duration  */
	nodeTimeout int64
	/** the node must be executed */
	nodeMustRun bool
}

func (zagNode *ZagNodeField[T]) GetNodeId() int {
	return zagNode.nodeId
}

func (zagNode *ZagNodeField[T]) GetHeadNodeIndex() int {
	return zagNode.headNodeIndex
}

func (zagNode *ZagNodeField[T]) Access() bool {
	return zagNode.access
}

func (zagNode *ZagNodeField[T]) GetNodeArrayOffset() int {
	return zagNode.nodeArrayOffset
}

func (zagNode *ZagNodeField[T]) GetNodeName() string {
	return zagNode.nodeName
}

func (zagNode *ZagNodeField[T]) GetGroupIndexList() []int {
	return zagNode.nodeGroupList
}

func (zagNode *ZagNodeField[T]) IsHead() bool {
	return zagNode.isHead
}

func (zagNode *ZagNodeField[T]) IsTail() bool {
	return zagNode.isTail
}

func (zagNode *ZagNodeField[T]) GetNodeOperator() NodeOperator[T] {
	return zagNode.nodeOperator
}

func (zagNode *ZagNodeField[T]) GetNodeChoose() IChoose[T] {
	return zagNode.nodeChoose
}

func (zagNode *ZagNodeField[T]) GetNodeFinishSignal() bool {
	return zagNode.nodeFinishSignal
}

func (zagNode *ZagNodeField[T]) GetNodeTimeout() int64 {
	return zagNode.nodeTimeout
}

func (zagNode *ZagNodeField[T]) MustRun() bool {
	return zagNode.nodeMustRun
}

func (zagNode *ZagNodeField[T]) SetNodeId(id int) {
	zagNode.nodeId = id
}

func (zagNode *ZagNodeField[T]) SetHeadNodeIndex(headNodeIndex int) {
	zagNode.headNodeIndex = headNodeIndex
}

func (zagNode *ZagNodeField[T]) SetAccess(access bool) {
	zagNode.access = access
}

func (zagNode *ZagNodeField[T]) SetNodeArrayOffset(nodeArrayOffset int) {
	zagNode.nodeArrayOffset = nodeArrayOffset
}

func (zagNode *ZagNodeField[T]) SetNodeName(nodeName string) {
	zagNode.nodeName = nodeName
}

func (zagNode *ZagNodeField[T]) AddGroup(groupIndex int) {
	zagNode.nodeGroupList = append(zagNode.nodeGroupList, groupIndex)
}

func (zagNode *ZagNodeField[T]) SetHeadState(state bool) {
	zagNode.isHead = state
}

func (zagNode *ZagNodeField[T]) SetTailState(state bool) {
	zagNode.isTail = state
}

func (zagNode *ZagNodeField[T]) SetNodeOperator(nodeOperator NodeOperator[T]) {
	zagNode.nodeOperator = nodeOperator
}

func (zagNode *ZagNodeField[T]) SetNodeChoose(nodeChoose IChoose[T]) {
	zagNode.nodeChoose = nodeChoose
}

func (zagNode *ZagNodeField[T]) SetNodeFinishSignal(needFinishSignal bool) {
	zagNode.nodeFinishSignal = needFinishSignal
}

func (zagNode *ZagNodeField[T]) SetNodeTimeout(nodeTimeout int64) {
	zagNode.nodeTimeout = nodeTimeout
}

func (zagNode *ZagNodeField[T]) SetMustRunState() {
	zagNode.nodeMustRun = true
}

var _ ZagNode[any] = (*ZagNodeField[any])(nil)
