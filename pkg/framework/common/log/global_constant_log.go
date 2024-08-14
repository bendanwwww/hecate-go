package log

import (
	"fmt"
	"github.com/bendanwwww/hecate-go/pkg/framework/tools"
	"github.com/sirupsen/logrus"
)

var (
	InfoLevel, _  = tools.String2Int(logrus.InfoLevel.String())
	DebugLevel, _ = tools.String2Int(logrus.DebugLevel.String())
	WarnLevel, _  = tools.String2Int(logrus.WarnLevel.String())
	ErrorLevel, _ = tools.String2Int(logrus.ErrorLevel.String())
	FatalLevel, _ = tools.String2Int(logrus.FatalLevel.String())

	// default log format

	// DefaultNodeLog default log, normally used for report system log, log placeholder is [sys]
	// [logType]|[requestId]|[graphName]|[fromNodeId:fromNodeName][nodeId:nodeName][nodeSchedulerState-nodeActuatorState:nodeActuatorResult] logInfo
	DefaultNodeLog = NewHecateLog(fmt.Sprintf("[%s]|[%s]|[%s]|[%s:%s][%s:%s][%s-%s:%s] %s",
		ZagLogType, RequestId, GraphName, FromNodeId, FromNodeName, NodeId, NodeName, NodeSchedulerState, NodeActuatorState, NodeActuatorResult, MessageInfo),
		InfoLevel, System)

	// BizNodeLog business log, normally used for report node's business logic, log placeholder is [biz]
	// [logType]|[requestId]|[graphName]|[nodeId:nodeName] logInfo
	BizNodeLog = NewHecateLog(fmt.Sprintf("[%s]|[%s]|[%s]|[%s:%s] %s",
		ZagLogType, RequestId, GraphName, NodeId, NodeName, MessageInfo),
		InfoLevel, Business)

	// DataNodeLog data log, normally used for report node's data information, log placeholder is [data]
	// [logType]|[requestId]|[graphName]|[nodeId:nodeName] logInfo
	DataNodeLog = NewHecateLog(fmt.Sprintf("[%s]|[%s]|[%s]|[%s:%s] %s",
		ZagLogType, RequestId, GraphName, NodeId, NodeName, MessageInfo),
		InfoLevel, Data)

	// RuntimeNodeLog node's execution time log, log placeholder is [runtime]
	// [logType]|[requestId]|[graphName]|[nodeId:nodeName] logInfo
	RuntimeNodeLog = NewHecateLog(fmt.Sprintf("[%s]|[%s]|[%s]|[%s:%s] %s",
		ZagLogType, RequestId, GraphName, NodeId, NodeName, MessageInfo),
		InfoLevel, Runtime)
)
