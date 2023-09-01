package airbyte

// Estimates are either per-stream (STREAM) or for the entire sync (SYNC). STREAM is
// preferred, and requires the source to count how many records are about to be emitted
// per-stream (e.g. there will be 100 rows from this table emitted). For the rare source
// which cannot tell which stream a record belongs to before reading (e.g. CDC databases),
// SYNC estimates can be emitted. Sources should not emit both STREAM and SOURCE estimates
// within a sync.
type EstimateType string

const (
	EstimateTypeSTREAM EstimateType = "STREAM"
	Sync               EstimateType = "SYNC"
)
