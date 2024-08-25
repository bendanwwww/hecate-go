package impl

import (
	"github.com/bendanwwww/hecate-go/pkg/core/driver/actuator/api"
	builderApi "github.com/bendanwwww/hecate-go/pkg/core/driver/builder/api"
	schedulerApi "github.com/bendanwwww/hecate-go/pkg/core/driver/scheduler/api"
	driverContext "github.com/bendanwwww/hecate-go/pkg/framework/context"
)

type DefaultSchedulerImpl[T any] struct {
	executor api.Actuator[T]
}

func NewDefaultScheduler[T any](executor api.Actuator[T]) *DefaultSchedulerImpl[T] {
	return &DefaultSchedulerImpl[T]{
		executor: executor,
	}
}

func (d DefaultSchedulerImpl[T]) Run(ctx *driverContext.RuntimeContext[T], builder builderApi.Builder[T]) chan string {
	//TODO implement me
	panic("implement me")
}

var _ schedulerApi.Scheduler[any] = (*DefaultSchedulerImpl[any])(nil)
