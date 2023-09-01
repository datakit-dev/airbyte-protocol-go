package airbyte

type DestinationSyncMode string

const (
	Append      DestinationSyncMode = "append"
	AppendDedup DestinationSyncMode = "append_dedup"
	Overwrite   DestinationSyncMode = "overwrite"
)
