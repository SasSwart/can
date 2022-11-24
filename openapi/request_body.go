package openapi

type RequestBody struct {
	node
	Description string
	Content     map[string]*MediaType
	Required    bool
}

func (r *RequestBody) getRef() string {
	return ""
}

var _ Traversable = &RequestBody{}

func (r *RequestBody) getChildren() map[string]Traversable {
	children := map[string]Traversable{}
	for name := range r.Content {
		mediaType := r.Content[name]
		children[name] = mediaType
	}
	return children
}

func (r *RequestBody) setChild(i string, t Traversable) {
	// TODO: handle this error
	content, _ := t.(*MediaType)
	r.Content[i] = content
}
