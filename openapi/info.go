package openapi

type Info struct {
	Title       string
	Description string
	Version     string
	License     struct {
		Name string
		Url  string
	}
}
