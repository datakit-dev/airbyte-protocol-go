package airbyte

// the type of trace message
type TraceType string

const (
	Estimate       TraceType = "ESTIMATE"
	TraceTypeERROR TraceType = "ERROR"
)
