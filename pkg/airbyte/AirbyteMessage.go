package airbyte

type AirbyteMessage struct {
	// catalog message: the catalog
	Catalog          *AirbyteCatalog          `json:"catalog,omitempty"`
	ConnectionStatus *AirbyteConnectionStatus `json:"connectionStatus,omitempty"`
	// connector config message: a message to communicate an updated configuration from a
	// connector that should be persisted
	Control *AirbyteControlMessage `json:"control,omitempty"`
	// log message: any kind of logging you want the platform to know about.
	Log *AirbyteLogMessage `json:"log,omitempty"`
	// record message: the record
	Record *AirbyteRecordMessage   `json:"record,omitempty"`
	Spec   *ConnectorSpecification `json:"spec,omitempty"`
	// schema message: the state. Must be the last message produced. The platform uses this
	// information
	State *AirbyteStateMessage `json:"state,omitempty"`
	// trace message: a message to communicate information about the status and performance of a
	// connector
	Trace *AirbyteTraceMessage `json:"trace,omitempty"`
	// Message type
	Type Type `json:"type"`
}
