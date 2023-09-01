package airbyte

// record message: the record
type AirbyteRecordMessage struct {
	// record data
	Data map[string]interface{} `json:"data"`
	// when the data was emitted from the source. epoch in millisecond.
	EmittedAt int64 `json:"emitted_at"`
	// namespace the data is associated with
	Namespace *string `json:"namespace,omitempty"`
	// stream the data is associated with
	Stream string `json:"stream"`
}
