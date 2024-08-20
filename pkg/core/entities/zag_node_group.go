package entities

type ZagNodeGroup interface {
	GetGroupId() int
	GetGroupName() string
	GetGroupTimeout() int64
	GetGroupHeadNodeNumber() int
	GetGroupNodeSize() int
	IsHead(index int) bool
	IsInGroup(index int) bool
	SetGroupId(id int)
	SetGroupName(name string)
	SetGroupTimeout(timeout int64)
	AddHeadNode(index int)
	RemoveHeadNode(index int)
	AddNode(index int)
}

type ZagNodeGroupField struct {
	/** group id */
	groupId int
	/** group name  */
	groupName string
	/** group timeout duration */
	groupTimeout int64
	/** the first nodes number in the group. the first node means no node is depended on this node in the group */
	groupHeadNodeNumber int
	/** group size */
	groupNodeSize int
	/** the first nodes list */
	groupHeadList []int
	/** nodes list */
	groupNodeList []int
}

func NewNodeGroup(nodeSize int) *ZagNodeGroupField {
	return &ZagNodeGroupField{
		groupHeadList: make([]int, nodeSize),
		groupNodeList: make([]int, nodeSize),
	}
}

func (zagNodeGroup *ZagNodeGroupField) GetGroupId() int {
	return zagNodeGroup.groupId
}

func (zagNodeGroup *ZagNodeGroupField) GetGroupName() string {
	return zagNodeGroup.groupName
}

func (zagNodeGroup *ZagNodeGroupField) GetGroupTimeout() int64 {
	return zagNodeGroup.groupTimeout
}

func (zagNodeGroup *ZagNodeGroupField) GetGroupHeadNodeNumber() int {
	return zagNodeGroup.groupHeadNodeNumber
}

func (zagNodeGroup *ZagNodeGroupField) GetGroupNodeSize() int {
	return zagNodeGroup.groupNodeSize
}

func (zagNodeGroup *ZagNodeGroupField) IsHead(index int) bool {
	return zagNodeGroup.groupHeadList[index] == 1
}

func (zagNodeGroup *ZagNodeGroupField) IsInGroup(index int) bool {
	return zagNodeGroup.groupNodeList[index] == 1
}

func (zagNodeGroup *ZagNodeGroupField) SetGroupId(id int) {
	zagNodeGroup.groupId = id
}

func (zagNodeGroup *ZagNodeGroupField) SetGroupName(name string) {
	zagNodeGroup.groupName = name
}

func (zagNodeGroup *ZagNodeGroupField) SetGroupTimeout(timeout int64) {
	zagNodeGroup.groupTimeout = timeout
}

func (zagNodeGroup *ZagNodeGroupField) AddHeadNode(index int) {
	if zagNodeGroup.groupHeadList[index] == 0 {
		zagNodeGroup.groupHeadNodeNumber += 1
	}
	zagNodeGroup.groupHeadList[index] = 1
}

func (zagNodeGroup *ZagNodeGroupField) RemoveHeadNode(index int) {
	if zagNodeGroup.groupHeadList[index] == 1 {
		zagNodeGroup.groupHeadNodeNumber -= 1
	}
	zagNodeGroup.groupHeadList[index] = 0
}

func (zagNodeGroup *ZagNodeGroupField) AddNode(index int) {
	if zagNodeGroup.groupNodeList[index] == 0 {
		zagNodeGroup.groupNodeSize += 1
	}
	zagNodeGroup.groupNodeList[index] = 1
}

var _ ZagNodeGroup = (*ZagNodeGroupField)(nil)
