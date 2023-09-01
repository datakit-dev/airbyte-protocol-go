package airbyte

type Status string

const (
	Failed    Status = "FAILED"
	Succeeded Status = "SUCCEEDED"
)
