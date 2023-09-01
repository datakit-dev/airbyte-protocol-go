package airbyte

// Specification of a connector (source/destination)
type ConnectorSpecification struct {
	// Additional and optional specification object to describe what an 'advanced' Auth flow
	// would need to function.
	// - A connector should be able to fully function with the configuration as described by the
	// ConnectorSpecification in a 'basic' mode.
	// - The 'advanced' mode provides easier UX for the user with UI improvements and
	// automations. However, this requires further setup on the
	// server side by instance or workspace admins beforehand. The trade-off is that the user
	// does not have to provide as many technical
	// inputs anymore and the auth process is faster and easier to complete.
	AdvancedAuth *AdvancedAuth `json:"advanced_auth,omitempty"`
	// deprecated, switching to advanced_auth instead
	AuthSpecification *AuthSpecification `json:"authSpecification,omitempty"`
	ChangelogURL      *string            `json:"changelogUrl,omitempty"`
	// ConnectorDefinition specific blob. Must be a valid JSON string.
	ConnectionSpecification map[string]interface{} `json:"connectionSpecification"`
	DocumentationURL        *string                `json:"documentationUrl,omitempty"`
	// the Airbyte Protocol version supported by the connector. Protocol versioning uses SemVer.
	ProtocolVersion *string `json:"protocol_version,omitempty"`
	// List of destination sync modes supported by the connector
	SupportedDestinationSyncModes []DestinationSyncMode `json:"supported_destination_sync_modes,omitempty"`
	// If the connector supports DBT or not.
	SupportsDBT *bool `json:"supportsDBT,omitempty"`
	// (deprecated) If the connector supports incremental mode or not.
	SupportsIncremental *bool `json:"supportsIncremental,omitempty"`
	// If the connector supports normalization or not.
	SupportsNormalization *bool `json:"supportsNormalization,omitempty"`
}
