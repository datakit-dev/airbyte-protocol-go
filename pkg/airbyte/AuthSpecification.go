package airbyte

// deprecated, switching to advanced_auth instead
type AuthSpecification struct {
	AuthType *AuthType `json:"auth_type,omitempty"`
	// If the connector supports OAuth, this field should be non-null.
	Oauth2Specification *OAuth2Specification `json:"oauth2Specification,omitempty"`
}
