package airbyte

type AirbyteGlobalState struct {
	SharedState  map[string]interface{} `json:"shared_state,omitempty"`
	StreamStates []AirbyteStreamState   `json:"stream_states"`
}
