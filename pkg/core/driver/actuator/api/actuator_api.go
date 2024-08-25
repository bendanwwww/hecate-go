package api

import "context"

type Actuator[T any] interface {
	// Execute TODO context type
	Execute(ctx *context.Context)
}
