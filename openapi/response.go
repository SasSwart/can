package openapi

type Response struct {
	Description string
	Content     map[string]MediaType
}

func (r *Response) ResolveRefs(basePath string, components *Components) error {
	for key, mediaType := range r.Content {
		err := mediaType.ResolveRefs(basePath, components)
		if err != nil {
			return err
		}
		r.Content[key] = mediaType
	}

	return nil
}
