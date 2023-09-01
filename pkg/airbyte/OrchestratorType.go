package airbyte

// the type of orchestrator message
type OrchestratorType string

const (
	ConnectorConfig OrchestratorType = "CONNECTOR_CONFIG"
)
