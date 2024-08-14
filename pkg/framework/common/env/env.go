package env

import (
	"github.com/bendanwwww/hecate-go/pkg/framework/constant"
	"net"
	"os"
	"strconv"
)

// GetEnv get environment by name
func GetEnv(name string) string {
	str := os.Getenv(name)
	if len(str) == 0 {
		return ""
	}
	return str
}

// GetEnvInt get environment of type integer
func GetEnvInt(name string, defaultValue int) int {
	str := os.Getenv(name)
	if len(str) == 0 {
		return defaultValue
	}
	v64, err := strconv.ParseInt(str, 10, 64)
	v := int(v64)
	if err != nil {
		return defaultValue
	}
	return v
}

// GetEnvBool get environment of type boolean
func GetEnvBool(name string, defaultValue bool) bool {
	str := os.Getenv(name)
	if len(str) == 0 {
		return defaultValue
	}
	v, err := strconv.ParseBool(str)
	if err != nil {
		return defaultValue
	}
	return v
}

// GetLogLevel get log level
func GetLogLevel() int {
	if constant.LogLevel != nil {
		return *constant.LogLevel
	}
	level := GetEnvInt(constant.FlowLogLevel, 0)
	constant.LogLevel = &level
	return level
}

// SetLogLevel set log level
func SetLogLevel(level int) {
	constant.LogLevel = &level
}

// GetPid get process pid
func GetPid() int {
	return os.Getpid()
}

// GetHostname get hostname
func GetHostname() string {
	var hostname = GetEnv(constant.EnvHostName)
	if hostname != "" {
		return hostname
	}
	hostname, err := os.Hostname()
	if err != nil {
		return constant.Unknown
	}
	return hostname
}

// GetFirstNotNullMAC get mac
func GetFirstNotNullMAC() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return constant.Unknown
	}
	for _, item := range interfaces {
		if item.HardwareAddr != nil {
			return item.HardwareAddr.String()
		}
	}
	return constant.Unknown
}

// GetUniquelyId get unique id
// if hostname not equals 'localhost', then return hostname, otherwise, return mac address.
func GetUniquelyId() string {
	datacenterId := GetHostname()
	if datacenterId == constant.LocalHost {
		datacenterId = GetFirstNotNullMAC()
	}
	return datacenterId
}
