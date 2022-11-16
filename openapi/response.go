package openapi

type Response struct {
	operationChildNode
	Description string            `yaml:"description"`
	Headers     map[string]Header // can also be a $ref
	Content     map[string]MediaType
	Links       map[string]Link // can also be a $ref
}

func (r *Response) getBasePath() string {
	return r.parent.getBasePath()
}

func (r *Response) GetName() string {
	return r.parent.GetName() + r.name
}

func (r *Response) getRef() string {
	return ""
}

var _ refContainer = &Response{}

func (r *Response) getChildren() map[string]Traversable {
	children := map[string]Traversable{}
	for name, mediaType := range r.Content {
		children[name] = &mediaType
	}
	return children
}

func (r *Response) setChild(i string, t Traversable) {
	mediaType, _ := t.(*MediaType)
	r.Content[i] = *mediaType
}
