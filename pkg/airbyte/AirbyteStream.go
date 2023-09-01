package airbyte

type AirbyteStream struct {
	// Path to the field that will be used to determine if a record is new or modified since the
	// last sync. If not provided by the source, the end user will have to specify the
	// comparable themselves.
	DefaultCursorField []string `json:"default_cursor_field,omitempty"`
	// Stream schema using Json Schema specs.
	JSONSchema map[string]interface{} `json:"json_schema"`
	// Stream's name.
	Name string `json:"name"`
	// Optional Source-defined namespace. Currently only used by JDBC destinations to determine
	// what schema to write to. Airbyte streams from the same sources should have the same
	// namespace.
	Namespace *string `json:"namespace,omitempty"`
	// If the source defines the cursor field, then any other cursor field inputs will be
	// ignored. If it does not, either the user_provided one is used, or the default one is used
	// as a backup.
	SourceDefinedCursor *bool `json:"source_defined_cursor,omitempty"`
	// If the source defines the primary key, paths to the fields that will be used as a primary
	// key. If not provided by the source, the end user will have to specify the primary key
	// themselves.
	SourceDefinedPrimaryKey [][]string `json:"source_defined_primary_key,omitempty"`
	// List of sync modes supported by this stream.
	SupportedSyncModes []SyncMode `json:"supported_sync_modes"`
}
