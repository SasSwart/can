package openapi

import "strconv"

type Response struct {
	node
	Description string            `yaml:"description"`
	Headers     map[string]Header // can also be a $ref
	Content     map[string]MediaType
	Links       map[string]Link // can also be a $ref
}

//func (r *Response) GetName() string {
//	return r.parent.GetName() + r.renderer.sanitiseName(r.name) + "Response"
//}

//func (r *Response) getRef() string {
//	panic("Composed type should override this")
//	return ""
//}

var _ Traversable = &Response{}

func (r *Response) getChildren() map[string]Traversable {
	responses := map[string]Traversable{} // Where string is either `default` or an HTTP status code
	for name, mediaType := range r.Content {
		if _, err := strconv.Atoi(name); err != nil || name == "default" {
			responses[name] = &mediaType
		} else {
			panic("Response spec broken")
		}
	}
	return responses
}

func (r *Response) setChild(i string, t Traversable) {
	if _, err := strconv.Atoi(i); err != nil || i == "default" {
		if mediaType, ok := t.(*MediaType); ok {
			r.Content[i] = *mediaType
			return
		}
	}
	panic("Response spec broken")
}
