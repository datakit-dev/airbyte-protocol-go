package airbyte

// Additional and optional specification object to describe what an 'advanced' Auth flow
// would need to function.
// - A connector should be able to fully function with the configuration as described by the
// ConnectorSpecification in a 'basic' mode.
// - The 'advanced' mode provides easier UX for the user with UI improvements and
// automations. However, this requires further setup on the
// server side by instance or workspace admins beforehand. The trade-off is that the user
// does not have to provide as many technical
// inputs anymore and the auth process is faster and easier to complete.
type AdvancedAuth struct {
	AuthFlowType             *AuthFlowType             `json:"auth_flow_type,omitempty"`
	OauthConfigSpecification *OAuthConfigSpecification `json:"oauth_config_specification,omitempty"`
	// Json Path to a field in the connectorSpecification that should exist for the advanced
	// auth to be applicable.
	PredicateKey []string `json:"predicate_key,omitempty"`
	// Value of the predicate_key fields for the advanced auth to be applicable.
	PredicateValue *string `json:"predicate_value,omitempty"`
}
