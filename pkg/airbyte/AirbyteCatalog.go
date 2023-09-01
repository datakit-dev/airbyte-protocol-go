package airbyte

// catalog message: the catalog
//
// Airbyte stream schema catalog
type AirbyteCatalog struct {
	Streams []AirbyteStream `json:"streams"`
}
