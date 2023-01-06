package openapi

// RequestBody is a programmatic representation of the RequestBody object defined here: https://swagger.io/specification/#request-body-object
type RequestBody struct {
	node
	Ref         string
	Description string
	Content     map[string]*MediaType
	Required    bool
}

func (r *RequestBody) GetName() string {
	return r.parent.GetName() + r.name
}

func (r *RequestBody) getRef() string {
	return r.Ref
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
	if content, ok := t.(*MediaType); ok {
		r.Content[i] = content
		return
	}
	panic("(r *RequestBody) setChild(): " + errCastFail)
}
