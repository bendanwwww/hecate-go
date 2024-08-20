package constant

// MapType builder graph type
type MapType string

// SchedulerType scheduler type
type SchedulerType string

// ActuatorType actuator type
type ActuatorType string

const (
	/** default builder graph driver */
	DefaultMap MapType = "DefaultMap"
)

const (
	/** simple scheduler driver, just use it to test */
	SimpleScheduler SchedulerType = "SimpleScheduler"
	/** default scheduler */
	DefaultScheduler SchedulerType = "DefaultScheduler"
)

const (
	/** default actuator */
	DefaultActuator ActuatorType = "DefaultActuator"
)
