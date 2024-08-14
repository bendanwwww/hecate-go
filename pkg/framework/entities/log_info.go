package entities

import "time"

type LogInfo struct {
	// log type
	LogType string
	// node name
	NodeName string
	// log write time
	LogWriteTime int64
	// log text
	LogStr string
}

func NewLogInfo(logType string, nodeName string, logStr string) *LogInfo {
	return &LogInfo{
		LogType:      logType,
		NodeName:     nodeName,
		LogStr:       logStr,
		LogWriteTime: time.Now().UnixMilli(),
	}
}
