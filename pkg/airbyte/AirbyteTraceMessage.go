package airbyte

// trace message: a message to communicate information about the status and performance of a
// connector
type AirbyteTraceMessage struct {
	// the time in ms that the message was emitted
	EmittedAt float64 `json:"emitted_at"`
	// error trace message: the error object
	Error *AirbyteErrorTraceMessage `json:"error,omitempty"`
	// Estimate trace message: a guess at how much data will be produced in this sync
	Estimate *AirbyteEstimateTraceMessage `json:"estimate,omitempty"`
	// the type of trace message
	Type TraceType `json:"type"`
}
