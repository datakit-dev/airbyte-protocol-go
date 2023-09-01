package airbyte

// connector config orchestrator message: the updated config for the platform to store for
// this connector
type AirbyteControlConnectorConfigMessage struct {
	// the config items from this connector's spec to update
	Config map[string]interface{} `json:"config"`
}
