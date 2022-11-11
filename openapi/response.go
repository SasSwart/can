package openapi

import "fmt"

type Response struct {
	Description string
	Content     map[string]MediaType
}

func (r *Response) ResolveRefs(basePath string) error {
	for key, mediaType := range r.Content {
		err := mediaType.ResolveRefs(basePath)
		if err != nil {
			return err
		}
		r.Content[key] = mediaType
	}

	return nil
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
