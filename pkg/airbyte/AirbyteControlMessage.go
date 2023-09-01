package airbyte

// connector config message: a message to communicate an updated configuration from a
// connector that should be persisted
type AirbyteControlMessage struct {
	// connector config orchestrator message: the updated config for the platform to store for
	// this connector
	ConnectorConfig *AirbyteControlConnectorConfigMessage `json:"connectorConfig,omitempty"`
	// the time in ms that the message was emitted
	EmittedAt float64 `json:"emitted_at"`
	// the type of orchestrator message
	Type OrchestratorType `json:"type"`
}
