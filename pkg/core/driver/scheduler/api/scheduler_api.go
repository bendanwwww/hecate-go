package api

import (
	"github.com/bendanwwww/hecate-go/pkg/core/driver/builder/api"
	driverContext "github.com/bendanwwww/hecate-go/pkg/framework/context"
)

type SchedulerApi[T any] interface {
	Run(ctx *driverContext.RuntimeContext[T], zagMap api.BuilderApi[T]) chan string
}
