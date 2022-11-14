package openapi

import "fmt"

type RequestBody struct {
	Description string
	Content     map[string]MediaType
	Required    bool
}

func (r *RequestBody) Render() error {
	if r.Content == nil {
		return nil
	}
	fmt.Println("Rendering API Request Body")

	for _, mediaType := range r.Content {
		err := mediaType.Render()
		if err != nil {
			return err
		}
	}
	return nil
}
