package entities

type ResultState string

const (
	// NodeInit init
	NodeInit ResultState = "init"
	// NodeSuccess success
	NodeSuccess ResultState = "success"
	// NodeTimeout timeout
	NodeTimeout ResultState = "timeout"
	// NodeSkip skip
	NodeSkip ResultState = "skip"
	// ActuatorLimit concurrency limit
	ActuatorLimit ResultState = "concurrencyLimit"
	// NodeDefaultError error
	NodeDefaultError ResultState = "error"
)

type NodeOperatorResult struct {
	/** node execute result state */
	ResultState ResultState
	/** error info */
	Ex error
}
