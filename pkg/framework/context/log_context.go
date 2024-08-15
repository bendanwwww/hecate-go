package context

import (
	"context"
	"github.com/bendanwwww/hecate-go/pkg/framework/entities"
)

type LogContext struct {
	context.Context
	// graph scenes name
	ScenesName string
	// trace id
	RequestId string
	// node id
	NodeId int
	// last node id
	FromNodeId int
	// node name
	NodeName string
	// last node name
	FromNodeName string
	// node scheduler state
	NodeSchedulerState string
	// node actuator state
	NodeActuatorState string
	// node actuator result
	NodeActuatorResult string
	// need print log
	NeedPrintLog bool
	// need upload log
	NeedWriteUpload bool
	// graph runtime more than x ms, must upload log
	MustWriteTime int64
	// log buffer array
	LogBuffer *[][]*entities.LogInfo
	// graph start execution time
	StartExecutionTime int64
}

func NewLogContextWithFirstNodeInfo(ctx context.Context, requestId string,
	scenesName string, startExecutionTime int64,
	nodeId int, nodeName string, needPrintLog bool, needWriteUpload bool, mustWriteTime int64,
	logBuffer *[][]*entities.LogInfo) *LogContext {
	return NewLogContext(ctx,
		requestId, scenesName, startExecutionTime,
		nodeId, nodeName, -1, "start", needPrintLog, needWriteUpload, mustWriteTime, logBuffer)
}

func NewLogContext(ctx context.Context, requestId string,
	scenesName string, startExecutionTime int64,
	nodeId int, nodeName string, fromNodeId int, fromNodeName string,
	needPrintLog bool, needWriteUpload bool, mustWriteTime int64,
	logBuffer *[][]*entities.LogInfo) *LogContext {
	return &LogContext{
		Context:            ctx,
		RequestId:          requestId,
		ScenesName:         scenesName,
		StartExecutionTime: startExecutionTime,
		NodeId:             nodeId,
		FromNodeId:         fromNodeId,
		NodeName:           nodeName,
		FromNodeName:       fromNodeName,
		NodeSchedulerState: "None",
		NodeActuatorState:  "None",
		NodeActuatorResult: "None",
		NeedPrintLog:       needPrintLog,
		NeedWriteUpload:    needWriteUpload,
		MustWriteTime:      mustWriteTime,
		LogBuffer:          logBuffer,
	}
}
