package context

import (
	"context"
	"fmt"
	"github.com/bendanwwww/hecate-go/pkg/framework/constant"
	LogEntities "github.com/bendanwwww/hecate-go/pkg/framework/entities"
	"github.com/bendanwwww/hecate-go/pkg/framework/tools"
	"time"
)

type RuntimeContext[T any] struct {
	context.Context
	// graph scenes name
	ScenesName string
	// trace id
	RequestId string
	// business context
	ServingContext *T
	// graph start execution time
	ExecutionStartTime int64
	// graph execution timeout duration
	GraphTimeout int64
	// need print log
	NeedPrintLog bool
	// need upload log
	NeedWriteSimpleLog bool
	// graph runtime more than x ms, must upload log
	MustWriteTime int64
	// node execute finish sign
	ReqChan chan string
	// node's log buffer array
	// array index: node id
	LogBuffer [][]*LogEntities.LogInfo
}

func NewRuntimeContext[T any](ctx context.Context,
	scenesName string,
	logPrintSwitch bool, logMustWriteTime int64, logSamplingSwitch bool, logSamplingRate float64,
	servingContext *T, nodeSize int, graphTimeout int64) *RuntimeContext[T] {
	requestContext := &RuntimeContext[T]{
		Context:            ctx,
		ScenesName:         scenesName,
		ServingContext:     servingContext,
		ExecutionStartTime: time.Now().UnixMilli(),
		GraphTimeout:       graphTimeout,
		ReqChan:            make(chan string, 10),
		NeedPrintLog:       logPrintSwitch,
		MustWriteTime:      logMustWriteTime,
		LogBuffer:          make([][]*LogEntities.LogInfo, nodeSize),
	}
	// if trace id in context, then use it, otherwise, call tools.GetNextId() make one.
	if ctx.Value(constant.TraceIdKey) == nil {
		requestContext.RequestId = tools.GetNextId()
		ctx = context.WithValue(ctx, constant.TraceIdKey, requestContext.RequestId)
	} else {
		requestContext.RequestId = fmt.Sprint(ctx.Value(constant.TraceIdKey))
	}
	// check if needed sampling log
	if logSamplingSwitch {
		requestContext.NeedWriteSimpleLog = tools.RandFloat64()*100 < logSamplingRate
	}
	return requestContext
}
