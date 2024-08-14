package log

import (
	"context"
	"fmt"
	"github.com/bendanwwww/hecate-go/pkg/framework/common/env"
	"github.com/bendanwwww/hecate-go/pkg/framework/constant"
	"github.com/bendanwwww/hecate-go/pkg/framework/entities"
	"github.com/bendanwwww/hecate-go/pkg/framework/tools"

	logContext "github.com/bendanwwww/hecate-go/pkg/framework/context"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasttemplate"
	"os"
	"strings"
)

type LogPlaceholder string

type LogType string

/** log placeholder  */
const (
	// ZagLogType log type
	ZagLogType LogPlaceholder = "{{logType}}"
	// RequestId trace id
	RequestId LogPlaceholder = "{{requestId}}"
	// BusinessName business name
	BusinessName LogPlaceholder = "{{businessName}}"
	// GraphName graph name
	GraphName LogPlaceholder = "{{graphName}}"
	// GraphVersion graph version
	GraphVersion LogPlaceholder = "{{graphVersion}}"
	// NodeId node id
	NodeId LogPlaceholder = "{{nodeId}}"
	// FromNodeId last node id
	FromNodeId LogPlaceholder = "{{fromNodeId}}"
	// NodeName node name
	NodeName LogPlaceholder = "{{nodeName}}"
	// FromNodeName last node name
	FromNodeName LogPlaceholder = "{{fromNodeName}}"
	// NodeSchedulerState node scheduler state
	NodeSchedulerState LogPlaceholder = "{{nodeSchedulerState}}"
	// NodeActuatorState node actuator state
	NodeActuatorState LogPlaceholder = "{{nodeActuatorState}}"
	// NodeActuatorResult node actuator result (success error...)
	NodeActuatorResult LogPlaceholder = "{{nodeActuatorResult}}"
	// MessageInfo log text
	MessageInfo LogPlaceholder = "{{messageInfo}}"
)

const (
	System       LogType = "sys"
	Business     LogType = "biz"
	Data         LogType = "data"
	Runtime      LogType = "runtime"
	GraphRuntime LogType = "graphRuntime"
	GraphTimeout LogType = "graphTimeout"
)

type HecateLog struct {
	/** log format */
	Format *fasttemplate.Template
	/** log level */
	LogLevel int
	/** log type */
	LogType LogType
}

type HecateLogApi interface {
	// SetLevel set log level
	SetLevel(logLevel int)
	/*** print log ***/
	Infof(ctx context.Context, format string, args ...interface{})
	Debugf(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Fatalf(ctx context.Context, format string, args ...interface{})
}

func NewHecateLog(format string, logLevel int, logType LogType) *HecateLog {
	t, _ := fasttemplate.NewTemplate(format, "{{", "}}")
	return &HecateLog{
		Format:   t,
		LogLevel: logLevel,
		LogType:  logType,
	}
}

func (hecateLog *HecateLog) SetLevel(logLevel int) {
	if logLevel > 0 {
		hecateLog.LogLevel = logLevel
		logrus.SetLevel(logrus.Level(logLevel))
		_ = os.Setenv(constant.FlowLogLevel, tools.Int2String(logLevel))
		env.SetLogLevel(logLevel)
	}
}

func (hecateLog *HecateLog) Infof(ctx context.Context, format string, args ...interface{}) {
	// check log level
	if logrus.GetLevel() < logrus.InfoLevel || hecateLog.LogLevel < InfoLevel {
		return
	}
	// write log buffer
	hecateLog.writeUpload(ctx, fmt.Sprintf(businessLogFormat(ctx, hecateLog, format), args...))
	// if need print log
	if hecateLog.needPrintLog(ctx) {
		logrus.Infof(businessLogFormat(ctx, hecateLog, format), args...)
	}
}

func (hecateLog *HecateLog) Debugf(ctx context.Context, format string, args ...interface{}) {
	// check log level
	if logrus.GetLevel() < logrus.DebugLevel || hecateLog.LogLevel < DebugLevel {
		return
	}
	// write log buffer
	hecateLog.writeUpload(ctx, fmt.Sprintf(businessLogFormat(ctx, hecateLog, format), args...))
	// if need print log
	if hecateLog.needPrintLog(ctx) {
		logrus.Debugf(businessLogFormat(ctx, hecateLog, format), args...)
	}
}

func (hecateLog *HecateLog) Errorf(ctx context.Context, format string, args ...interface{}) {
	// check log level
	if logrus.GetLevel() < logrus.ErrorLevel || hecateLog.LogLevel < ErrorLevel {
		return
	}
	// write log buffer
	hecateLog.writeUpload(ctx, fmt.Sprintf(businessLogFormat(ctx, hecateLog, format), args...))
	// if need print log
	if hecateLog.needPrintLog(ctx) {
		logrus.Errorf(businessLogFormat(ctx, hecateLog, format), args...)
	}
}

func (hecateLog *HecateLog) Warnf(ctx context.Context, format string, args ...interface{}) {
	// check log level
	if logrus.GetLevel() < logrus.WarnLevel || hecateLog.LogLevel < WarnLevel {
		return
	}
	// write log buffer
	hecateLog.writeUpload(ctx, fmt.Sprintf(businessLogFormat(ctx, hecateLog, format), args...))
	// if need print log
	if hecateLog.needPrintLog(ctx) {
		logrus.Warnf(businessLogFormat(ctx, hecateLog, format), args...)
	}
}

func (hecateLog *HecateLog) Fatalf(ctx context.Context, format string, args ...interface{}) {
	// check log level
	if logrus.GetLevel() < logrus.FatalLevel || hecateLog.LogLevel < FatalLevel {
		return
	}
	// write log buffer
	hecateLog.writeUpload(ctx, fmt.Sprintf(businessLogFormat(ctx, hecateLog, format), args...))
	// if need print log
	if hecateLog.needPrintLog(ctx) {
		logrus.Fatalf(businessLogFormat(ctx, hecateLog, format), args...)
	}
}

func (hecateLog *HecateLog) needPrintLog(ctx context.Context) bool {
	switch c := ctx.(type) {
	case *logContext.LogContext:
		return c.NeedPrintLog
	default:
		return false
	}
}

func (hecateLog *HecateLog) writeUpload(ctx context.Context, log string) {
	switch c := ctx.(type) {
	case *logContext.LogContext:
		if len((*c.LogBuffer)[c.NodeId]) == 0 {
			(*c.LogBuffer)[c.NodeId] = make([]*entities.LogInfo, 0, 100)
		}
		nodeLog := entities.NewLogInfo(string(hecateLog.LogType), c.NodeName, log)
		(*c.LogBuffer)[c.NodeId] = append((*c.LogBuffer)[c.NodeId], nodeLog)
	default:
		return
	}

}

func businessLogFormat(ctx context.Context, hecateLog *HecateLog, msgFormat string) string {
	switch c := ctx.(type) {
	case *logContext.LogContext:
		return hecateLog.Format.ExecuteString(map[string]interface{}{
			"logType":            string(hecateLog.LogType),
			"requestId":          strings.ReplaceAll(c.RequestId, " ", "_"),
			"businessName":       strings.ReplaceAll(c.BusinessName, " ", "_"),
			"graphName":          strings.ReplaceAll(c.GraphName, " ", "_"),
			"graphVersion":       strings.ReplaceAll(c.Version, " ", "_"),
			"nodeId":             tools.Int2String(c.NodeId),
			"fromNodeId":         tools.Int2String(c.FromNodeId),
			"nodeName":           strings.ReplaceAll(c.NodeName, " ", "_"),
			"fromNodeName":       strings.ReplaceAll(c.FromNodeName, " ", "_"),
			"nodeSchedulerState": strings.ReplaceAll(c.NodeSchedulerState, " ", "_"),
			"nodeActuatorState":  strings.ReplaceAll(c.NodeActuatorState, " ", "_"),
			"nodeActuatorResult": strings.ReplaceAll(c.NodeActuatorResult, " ", "/"),
			"messageInfo":        msgFormat,
		})
	default:
		return msgFormat
	}
}

var _ HecateLogApi = (*HecateLog)(nil)
