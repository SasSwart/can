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
	Value         interface{} `yaml:"value"` // TODO: does this imply an explicit need for reflective interpretation of data while marshalling?
	ExternalValue string      `yaml:"externalValue"`
}

// TODO: work out ABNF expressions for the resolution of $ref strings for below structs

// Header is a programmatic representation of the Header object defined here:https://swagger.io/specification/#header-object
type Header struct{}

// SecurityScheme is a programmatic representation of the SecurityScheme object defined here: https://swagger.io/specification/#security-scheme-object
type SecurityScheme struct{}

// Link is a programmatic representation of the Link object defined here: https://swagger.io/specification/#link-object
type Link struct{}

// Callback is a programmatic representation of the Callback object defined here: https://swagger.io/specification/#callback-object
type Callback struct{}
