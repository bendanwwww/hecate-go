package api

import (
	"github.com/bendanwwww/hecate-go/pkg/core/driver/builder/api"
	driverContext "github.com/bendanwwww/hecate-go/pkg/framework/context"
)

type Scheduler[T any] interface {
	Run(ctx *driverContext.RuntimeContext[T], zagMap api.Builder[T]) chan string
}
