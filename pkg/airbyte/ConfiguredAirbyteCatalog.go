package airbyte

// Airbyte stream schema catalog
type ConfiguredAirbyteCatalog struct {
	Streams []ConfiguredAirbyteStream `json:"streams"`
}
