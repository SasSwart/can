package openapi

type RequestBody struct {
	parent      refContainer
	Description string
	Content     map[string]MediaType
	Required    bool
}

func (r *RequestBody) getBasePath() string {
	return r.parent.getBasePath()
}

func (r *RequestBody) getRef() string {
	return ""
}

var _ refContainer = &RequestBody{}

func (r *RequestBody) getChildren() map[string]Traversable {
	children := map[string]Traversable{}
	for name, mediaType := range r.Content {
		children[name] = &mediaType
	}
	return children
}

func (r *RequestBody) setChild(i string, t Traversable) {
	//TODO implement me
	panic("implement me")
}
