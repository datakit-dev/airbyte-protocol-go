package airbyte

// The type of state the other fields represent. Is set to LEGACY, the state data should be
// read from the `data` field for backwards compatibility. If not set, assume the state
// object is type LEGACY. GLOBAL means that the state should be read from `global` and means
// that it represents the state for all the streams. It contains one shared state and
// individual stream states. PER_STREAM means that the state should be read from `stream`.
// The state present in this field correspond to the isolated state of the associated stream
// description.
type AirbyteStateType string

const (
	AirbyteStateTypeSTREAM AirbyteStateType = "STREAM"
	Global                 AirbyteStateType = "GLOBAL"
	Legacy                 AirbyteStateType = "LEGACY"
)
