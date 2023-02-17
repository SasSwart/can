package info

// Info is a programmatic representation of the Info object defined here: https://swagger.io/specification/#info-object
type Info struct {
	Title          string `yaml:"title"`
	Description    string `yaml:"description"`
	TermsOfService string `yaml:"termsOfService"`
	Version        string `yaml:"version"`
	Contact        struct {
		Name  string `yaml:"name"`
		Url   string `yaml:"url"`
		Email string `yaml:"email"`
	}
	License struct {
		Name string `yaml:"name"`
		Url  string `yaml:"url"`
	}
}
