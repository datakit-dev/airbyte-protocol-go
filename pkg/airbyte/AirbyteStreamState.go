package airbyte

type AirbyteStreamState struct {
	StreamDescriptor StreamDescriptor       `json:"stream_descriptor"`
	StreamState      map[string]interface{} `json:"stream_state,omitempty"`
}
