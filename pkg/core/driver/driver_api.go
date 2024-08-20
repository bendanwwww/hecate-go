package driver

import (
	"context"
	"github.com/bendanwwww/hecate-go/pkg/core/entities"
)

type ZagDriver[T any] interface {
	// AddNodes add node in the graph
	AddNodes(scenes string, node ...entities.NodeWrapper[T])
	// AddGroupNodes add node group in the graph
	AddGroupNodes(scenes string, groupName string, node ...entities.NodeWrapper[T])
	// AddGroupNodesWithTimeout add node group in the graph with timeout
	AddGroupNodesWithTimeout(scenes string, groupName string, groupTimeout int64, node ...entities.NodeWrapper[T])
	// BuildAll build all graph in driver
	BuildAll()
	// Build build simple graph
	Build(scenes string)
	// Run execute graph
	Run(ctx context.Context, servingContext *T, scenes string, timeout int64) (chan string, error)
	// Clear clear graph data
	Clear(scenes string)
	// ToString print graph information
	ToString(scenes string) string
}
