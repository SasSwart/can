package openapi

type Traversable interface {
	getChildren() map[string]Traversable
	setChild(i string, t Traversable)
}

type refContainer interface {
	Traversable
	getBasePath() string
	getRef() string
	GetName() string
}

type TraversalFunc func(key string, parent, child Traversable) (Traversable, error)

type node[T Traversable] struct {
	parent T
	name   string
}

type refContainerNode node[refContainer]

func (n refContainerNode) GetName() string {
	return n.parent.GetName() + n.name
}

// Traverse takes a Traversable node and applies some function to the node within the tree. It recursively calls itself and fails early when an error is thrown
func Traverse(node Traversable, f TraversalFunc) (Traversable, error) {
	children := node.getChildren()
	for i, child := range children {
		// Update Child Node
		newChild, err := f(i, node, child)
		if err != nil {
			return nil, err
		}
		node.setChild(i, newChild)

		// Traverse Child Node
		_, err = Traverse(newChild, f)
		if err != nil {
			return nil, err
		}
	}

	return node, nil
}
