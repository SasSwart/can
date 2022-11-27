package openapi

type RequestBody struct {
	node
	Description string
	Content     map[string]*MediaType
	Required    bool
}

//func (r *RequestBody) GetName() string {
//	return r.name
//}

//func (r *RequestBody) getRef() string {
//	panic("(r *RequestBody) getRef() This should never be called") // Refs are in media type -> schemas
//	return ""
//}

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
	panic("(r *RequestBody) setChild borked")
}
