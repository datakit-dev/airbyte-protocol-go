package airbyte

// Estimate trace message: a guess at how much data will be produced in this sync
type AirbyteEstimateTraceMessage struct {
	// The estimated number of bytes to be emitted by this sync for this stream
	ByteEstimate *int64 `json:"byte_estimate,omitempty"`
	// The name of the stream
	Name string `json:"name"`
	// The namespace of the stream
	Namespace *string `json:"namespace,omitempty"`
	// The estimated number of rows to be emitted by this sync for this stream
	RowEstimate *int64 `json:"row_estimate,omitempty"`
	// Estimates are either per-stream (STREAM) or for the entire sync (SYNC). STREAM is
	// preferred, and requires the source to count how many records are about to be emitted
	// per-stream (e.g. there will be 100 rows from this table emitted). For the rare source
	// which cannot tell which stream a record belongs to before reading (e.g. CDC databases),
	// SYNC estimates can be emitted. Sources should not emit both STREAM and SOURCE estimates
	// within a sync.
	Type EstimateType `json:"type"`
}
