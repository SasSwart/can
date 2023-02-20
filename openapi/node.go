package openapi

import "fmt"

type Metadata map[string]string
type Traversable interface {
	// Tree Traversable
	getChildren() map[string]Traversable
	setChild(i string, t Traversable)
	GetParent() Traversable
	setParent(parent Traversable)

	// Node Attributes
	GetName() string
	setName(name string)
	getBasePath() string
	getRef() string
	setRenderer(r Renderer)
	getRenderer() Renderer
	GetOutputFile() string
	GetMetadata() Metadata
	SetMetadata(metadata Metadata)
}

type TraversalFunc func(key string, parent, child Traversable) (Traversable, error)

type node struct {
	basePath string
	parent   Traversable
	name     string
	renderer Renderer
	metadata Metadata
}

const (
	errNotImplemented = " not implemented by composed type"
	errCastFail       = " cast failed"
)

func (n *node) SetMetadata(metadata Metadata) {
	if n.GetParent() == nil {
		n.metadata = metadata
		return
	}
	n.GetParent().SetMetadata(metadata)

}

var _ Traversable = &node{}

func (n *node) GetMetadata() Metadata {
	if n.GetParent() == nil {
		return n.metadata
	}
	return n.GetParent().GetMetadata()
}

func (n *node) getChildren() map[string]Traversable {
	panic("(n *node) getChildren():" + errNotImplemented)
}

func (n *node) setChild(_ string, _ Traversable) {
	panic("(n *node) setChild():" + errNotImplemented)
}

func (n *node) GetParent() Traversable {
	return n.parent
}

func (n *node) setParent(parent Traversable) {
	n.parent = parent
}

// getBasePath recurses up the parental ladder until it's overridden by the *OpenAPI method
func (n *node) getBasePath() string {
	if n.parent == nil {
		return n.basePath
	}
	return n.parent.GetParent().getBasePath()
}

// TODO this function can do without it's overrides
func (n *node) GetOutputFile() string {
	return n.getRenderer().getOutputFile(n)
}

func (n *node) getRef() string {
	panic("(n *node) getRef():" + errNotImplemented)
	return ""
}

func (n *node) GetName() string {
	name := n.parent.GetName() + n.getRenderer().sanitiseName(n.name)
	return name
}

func (n *node) setName(name string) {
	n.name = name
}

func (n *node) setRenderer(r Renderer) {
	if n.parent == nil {
		n.renderer = r
		return
	}
	if n.parent.getRenderer() == nil {
		n.parent.setRenderer(r)
	}
}

func (n *node) getRenderer() Renderer {
	if n.parent == nil {
		return n.renderer
	}
	return n.parent.getRenderer()
}

// TraverseRecursor is an auxiliary function to Traverse that initiates a recursive loop over a tree of Traversable
// structs, applying a given function at every step.
func TraverseRecursor[T Traversable](node T, f TraversalFunc) (T, error) {
	children := node.getChildren()
	for i := range children {
		child := children[i]
		if child == nil {
			continue
		}
		// Update Child Node
		newChild, err := f(i, node, child)
		if err != nil {
			return node, err
		}
		node.setChild(i, newChild)

		if newChild == nil {
			continue
		}
		_, err = TraverseRecursor(newChild, f)
		if err != nil {
			return node, err
		}
	}

	return node, nil
}

// Traverse takes a Traversable node and enters into a recursive loop (TraverseRecursor) that applies a given function
// to the node.
func Traverse[T Traversable](node T, f TraversalFunc) (T, error) {
	if f == nil {
		return node, fmt.Errorf("no traversal function supplied")
	}

	result, err := f("", nil, node)
	if err != nil {
		return node, fmt.Errorf("traversal function[%#v] error: %w", f, err)
	}

	node, ok := result.(T)
	if !ok {
		return node, fmt.Errorf("function parameter f should return the same type that is was given")
	}

	return TraverseRecursor(node, f)
}
