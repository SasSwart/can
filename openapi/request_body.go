package openapi

type RequestBody struct {
	Description string
	Content     map[string]MediaType
	Required    bool
}

func (r *RequestBody) ResolveRefs(basePath string, components *Components) error {
	if r.Content == nil {
		return nil
	}

	for m, mediaType := range r.Content {
		err := mediaType.ResolveRefs(basePath, components)
		if err != nil {
			return err
		}
		r.Content[m] = mediaType
	}
	return nil
}
