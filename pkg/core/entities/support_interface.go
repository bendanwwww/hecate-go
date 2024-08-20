package entities

import (
	"context"
	mapset "github.com/deckarep/golang-set"
)

type ICondition[T any] interface {
	// CanExecute check the node can be executed
	CanExecute(ctx context.Context, param *T) bool
}

type IChoose[T any] interface {
	// Choose choose the node's next nodes function
	Choose(ctx context.Context, param *T) mapset.Set
}
