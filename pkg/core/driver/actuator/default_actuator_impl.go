package impl

import (
	"context"
	actuatorApi "github.com/bendanwwww/hecate-go/pkg/core/driver/actuator/api"
)

type DefaultActuatorImpl[T any] struct {
}

func NewDefaultActuator[T any](maxConcurrency int, usePool bool) *DefaultActuatorImpl[T] {
	executor := &DefaultActuatorImpl[T]{}
	return executor
}

func (d DefaultActuatorImpl[T]) Execute(ctx *context.Context) {
	//TODO implement me
	panic("implement me")
}

var _ actuatorApi.Actuator[any] = (*DefaultActuatorImpl[any])(nil)
