package entities

import "context"

type NodeOperator[T any] interface {
	// NodeCondition node execute condition
	NodeCondition(ctx context.Context, param *T) bool
	// Before before node execute, this function will be run
	Before(ctx context.Context)
	// Execute node execute function
	Execute(ctx context.Context, param *T)
	// After after node execute, this function will be run
	After(ctx context.Context)
	// Error when the node execute error, this function will be run
	Error(ctx context.Context, param *T, result *NodeOperatorResult)
}
