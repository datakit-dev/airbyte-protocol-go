package airbyte

// log message: any kind of logging you want the platform to know about.
type AirbyteLogMessage struct {
	// log level
	Level Level `json:"level"`
	// log message
	Message string `json:"message"`
	// an optional stack trace if the log message corresponds to an exception
	StackTrace *string `json:"stack_trace,omitempty"`
}
