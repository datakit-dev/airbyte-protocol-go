package airbyte

// AirbyteProtocol structs
type Coordinate struct {
	AirbyteMessage           *AirbyteMessage           `json:"airbyte_message,omitempty"`
	ConfiguredAirbyteCatalog *ConfiguredAirbyteCatalog `json:"configured_airbyte_catalog,omitempty"`
}
