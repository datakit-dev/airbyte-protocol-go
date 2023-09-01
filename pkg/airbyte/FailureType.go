package airbyte

// The type of error
type FailureType string

const (
	ConfigError FailureType = "config_error"
	SystemError FailureType = "system_error"
)
