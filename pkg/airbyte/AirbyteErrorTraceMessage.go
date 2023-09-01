package airbyte

// error trace message: the error object
type AirbyteErrorTraceMessage struct {
	// The type of error
	FailureType *FailureType `json:"failure_type,omitempty"`
	// The internal error that caused the failure
	InternalMessage *string `json:"internal_message,omitempty"`
	// A user-friendly message that indicates the cause of the error
	Message string `json:"message"`
	// The full stack trace of the error
	StackTrace *string `json:"stack_trace,omitempty"`
}
