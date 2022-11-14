package openapi

import "fmt"

type Response struct {
	Description string            `yaml:"description"`
	Headers     map[string]Header // can also be a $ref
	Content     map[string]MediaType
	Links       map[string]Link // can also be a $ref
}

func (r *Response) Render() error {
	fmt.Println("Rendering API Response")
	for _, mediaType := range r.Content {
		err := mediaType.Render()
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Response) GetSchemas(name string) (schemas map[string]Schema) {
	schemas = map[string]Schema{}
	for _, mediaType := range r.Content {
		for name, schema := range mediaType.GetSchemas(name) {
			schemas[name] = schema
		}
	}

	return schemas
}
