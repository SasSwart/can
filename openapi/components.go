package openapi

// Components is a programmatic representation of the Components object defined here: https://swagger.io/specification/#components-object
type Components struct {
	Schemas         map[string]Schema         // can also be a $ref
	Responses       map[string]Response       // can also be a $ref
	Parameters      map[string]Parameter      // can also be a $ref
	Examples        map[string]Example        // can also be a $ref
	RequestBodies   map[string]RequestBody    // can also be a $ref
	Headers         map[string]Header         // can also be a $ref
	SecuritySchemes map[string]SecurityScheme // can also be a $ref
	Links           map[string]Link           // can also be a $ref
	Callbacks       map[string]Callback       // can also be a $ref
}

// Example is a programmatic representation of the Example object defined here: https://swagger.io/specification/#components-object
type Example struct {
	Summary       string      `yaml:"summary"`
	Description   string      `yaml:"description"`
	Value         interface{} `yaml:"value"`
	ExternalValue string      `yaml:"externalValue"`
}

// TODO: work out ABNF expressions for the resolution of $ref strings for below structs

// Header is a programmatic representation of the Header object defined here:https://swagger.io/specification/#header-object
type Header struct {
	Description     string `yaml:"description"`
	Required        bool   `yaml:"required"`
	Deprecated      bool   `yaml:"deprecated"`
	AllowEmptyValue bool   `yaml:"allowEmptyValue"`
	Schema          Schema
}

// SecurityScheme is a programmatic representation of the SecurityScheme object defined here: https://swagger.io/specification/#security-scheme-object
type SecurityScheme struct {
	Type             string `yaml:"type"`
	Description      string `yaml:"description"`
	Name             string `yaml:"name"`
	In               string `yaml:"in"`
	Scheme           string `yaml:"scheme"`
	BearerFormat     string `yaml:"BearerFormat"`
	Flows            OAuthFlows
	OpenIdConnectUrl string `yaml:"openIdConnectUrl"`
}

// Link is a programmatic representation of the Link object defined here: https://swagger.io/specification/#link-object
type Link struct {
	OperationRef string `yaml:"operationRef"`
	OperationId  string `yaml:"operationId"`
	Parameters   map[string]interface{}
	RequestBody  interface{}
	Description  string `yaml:"description"`
	Server       Server
}

// Callback is a programmatic representation of the Callback object defined here: https://swagger.io/specification/#callback-object
type Callback struct{}

// OAuthFlows is a programmatic representation of the OAuthFlows object defined here: https://swagger.io/specification/#oauth-flow-object
// Allows configuration of the following types of OAuth flows: implicit, password, clientCredentials, authorizationCode
// TODO: make sure that the keys listed above are tested in parent map containing OAuthFlows
type OAuthFlows struct {
	AuthorizationUrl string `yaml:"authorizationUrl"`
	TokenUrl         string `yaml:"tokenUrl"`
	RefreshUrl       string `yaml:"refreshUrl"`
	Scopes           map[string]string
}
