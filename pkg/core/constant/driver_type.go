package constant

// BuilderType builder graph type
type BuilderType string

// SchedulerType scheduler type
type SchedulerType string

// ActuatorType actuator type
type ActuatorType string

const (
	/** default builder graph driver */
	DefaultBuilder BuilderType = "DefaultMap"
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
