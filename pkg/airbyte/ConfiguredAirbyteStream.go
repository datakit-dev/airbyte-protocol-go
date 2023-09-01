package airbyte

type ConfiguredAirbyteStream struct {
	// Path to the field that will be used to determine if a record is new or modified since the
	// last sync. This field is REQUIRED if `sync_mode` is `incremental`. Otherwise it is
	// ignored.
	CursorField         []string            `json:"cursor_field,omitempty"`
	DestinationSyncMode DestinationSyncMode `json:"destination_sync_mode"`
	// Paths to the fields that will be used as primary key. This field is REQUIRED if
	// `destination_sync_mode` is `*_dedup`. Otherwise it is ignored.
	PrimaryKey [][]string    `json:"primary_key,omitempty"`
	Stream     AirbyteStream `json:"stream"`
	SyncMode   SyncMode      `json:"sync_mode"`
}
