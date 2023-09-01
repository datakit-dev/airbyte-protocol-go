package airbyte

// Airbyte connection status
type AirbyteConnectionStatus struct {
	Message *string `json:"message,omitempty"`
	Status  Status  `json:"status"`
}
