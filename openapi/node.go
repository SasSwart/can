package openapi

type Traversable interface {
	getChildren() map[string]Traversable
	setChild(i string, t Traversable)
	GetParent() Traversable
	setParent(parent Traversable)
	GetName() string
	setName(name string)
	getBasePath() string
	getRef() string
	setRenderer(r Renderer)
	getRenderer() Renderer
	GetOutputFile() string
}

type TraversalFunc func(key string, parent, child Traversable) (Traversable, error)

type node struct {
	basePath string
	parent   Traversable
	name     string
	renderer Renderer
	ref      string
}

var _ Traversable = &node{}

func (n *node) getChildren() map[string]Traversable {
	panic("not implemented by composed type")
}

func (n *node) setChild(_ string, _ Traversable) {
	panic("not implemented by composed type")
}

func (n *node) GetParent() Traversable {
	return n.parent
}

func (n *node) setParent(parent Traversable) {
	n.parent = parent
}

// Recurses up the parental ladder until it's overridden by the *OpenAPI method
func (n *node) getBasePath() string {
	return n.parent.getBasePath()
}

func (n *node) GetOutputFile() string {
	return n.getRenderer().getOutputFile(n)
}

func (n *node) getRef() string {
	return n.ref
}

func (n *node) GetName() string {
	name := n.parent.GetName() + n.renderer.sanitiseName(n.name)
	return name
}

func (n *node) setName(name string) {
	n.name = name
}

func (n *node) setRenderer(r Renderer) {
	n.renderer = r
}

func (n *node) getRenderer() Renderer {
	return n.renderer
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
			// TODO should child node pass down modified base path for easy resolution?
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

func Dig(node Traversable, key ...string) Traversable {
	if len(key) == 0 {
		return node
	}
	return Dig(node.getChildren()[key[0]], key[1:]...)
}
