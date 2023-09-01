package airbyte

// Message type
type Type string

const (
	Catalog          Type = "CATALOG"
	ConnectionStatus Type = "CONNECTION_STATUS"
	Control          Type = "CONTROL"
	Log              Type = "LOG"
	Record           Type = "RECORD"
	Spec             Type = "SPEC"
	State            Type = "STATE"
	TypeTRACE        Type = "TRACE"
)
