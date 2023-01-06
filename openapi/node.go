package openapi

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
	GetMetadata() map[string]string
	SetMetadata(metadata map[string]string)
}

type TraversalFunc func(key string, parent, child Traversable) (Traversable, error)

type node struct {
	basePath string
	parent   Traversable
	name     string
	renderer Renderer
}

const (
	errNotImplemented = " not implemented by composed type"
	errCastFail       = " cast failed"
)

func (n *node) SetMetadata(metadata map[string]string) {
	n.parent.SetMetadata(metadata)
}

var _ Traversable = &node{}

func (n *node) GetMetadata() map[string]string {
	return n.parent.GetMetadata()
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
	return n.parent.getBasePath()
}

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
	n.renderer = r
}

func (n *node) getRenderer() Renderer {
	return n.parent.getRenderer()
}

// Traverse takes a Traversable node and applies some function to the node within the tree. It recursively calls itself and fails early when an error is thrown
func Traverse(node Traversable, f TraversalFunc) (Traversable, error) {
	if node == nil || f == nil {
		return node, nil
	}
	var recurse func(node Traversable, f TraversalFunc) (Traversable, error)
	recurse = func(node Traversable, f TraversalFunc) (Traversable, error) {
		children := node.getChildren()
		for i := range children {
			child := children[i]
			if child == nil {
				continue
			}
			// Update Child Node
			newChild, err := f(i, node, child)
			if err != nil {
				return nil, err
			}
			node.setChild(i, newChild)

			if newChild == nil {
				continue
			}
			_, err = recurse(newChild, f)
			if err != nil {
				return nil, err
			}
		}

		return node, nil
	}
	node, err := f("", nil, node)
	if err != nil {
		return nil, err
	}

	return recurse(node, f)
}
