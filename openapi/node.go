package openapi

type Traversable interface {
	getChildren() map[string]Traversable
	setChild(i string, t Traversable)
}

// Traverse takes a Traversable node and applies some function to the node within the tree.
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
	for i := range children {
		child := children[i]
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

type childContainerMap struct {
	mapContainer map[string]Traversable
}

func (c childContainerMap) get(i string) Traversable {
	return c.mapContainer[i]
}
func (c childContainerMap) set(i string, child Traversable) {
	c.mapContainer[i] = child
}
func (c childContainerMap) len() int {
	return len(c.mapContainer)
}

type childContainerSlice struct {
	sliceContainer []Traversable
}

func (c childContainerSlice) get(i int) Traversable {
	return c.sliceContainer[i]
}
func (c childContainerSlice) set(i int, child Traversable) {
	c.sliceContainer[i] = child
}
func (c childContainerSlice) len() int {
	return len(c.sliceContainer)
}
