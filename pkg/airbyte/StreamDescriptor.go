package airbyte

type StreamDescriptor struct {
	Name      string  `json:"name"`
	Namespace *string `json:"namespace,omitempty"`
}
