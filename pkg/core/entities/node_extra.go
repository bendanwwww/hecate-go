package entities

type NodeContext struct {
	// node name
	NodeName string
	// the node below groups' name list
	NodeGroupNameList []string
}

func NewNodeContext(nodeName string, nodeGroupNameList []string) *NodeContext {
	return &NodeContext{
		NodeName:          nodeName,
		NodeGroupNameList: nodeGroupNameList,
	}
}
