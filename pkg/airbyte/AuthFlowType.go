package airbyte

type AuthFlowType string

const (
	AuthFlowTypeOauth20 AuthFlowType = "oauth2.0"
	Oauth10             AuthFlowType = "oauth1.0"
)
