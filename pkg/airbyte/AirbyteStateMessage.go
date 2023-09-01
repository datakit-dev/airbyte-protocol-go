package airbyte

// schema message: the state. Must be the last message produced. The platform uses this
// information
type AirbyteStateMessage struct {
	// (Deprecated) the state data
	Data   map[string]interface{} `json:"data,omitempty"`
	Global *AirbyteGlobalState    `json:"global,omitempty"`
	Stream *AirbyteStreamState    `json:"stream,omitempty"`
	Type   *AirbyteStateType      `json:"type,omitempty"`
}
