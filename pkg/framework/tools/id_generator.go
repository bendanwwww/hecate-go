package tools

import (
	"fmt"
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"github.com/bendanwwww/hecate-go/pkg/framework/common/env"
)

var (
	// if hostname not equals 'localhost', then use hostname to data center's id, otherwise, use mac address.
	// use pid to worker id.
	datacenterId    = env.GetUniquelyId()
	workerId        = env.GetPid()
	snowflakeObj, _ = snowflake.NewSnowflake(0, 0)
)

// GetNextId get unique id
func GetNextId() string {
	return fmt.Sprintf("%s-%d-%d", datacenterId, workerId, snowflakeObj.NextVal())
}
