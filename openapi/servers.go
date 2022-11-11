package openapi

// Server is a programmatic representation of the OpenApi Document object defined here: https://swagger.io/specification/#openapi-object
type Server struct {
	Url         string `yaml:"url"`
	Description string `yaml:"description"`
	Variables   map[string]ServerVariable
}

type ServerVariable struct {
	Enum        []string `yaml:"enum"`
	Default     string   `yaml:"default"`
	Description string   `yaml:"description"`
}
