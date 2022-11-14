package openapi

type traversable interface {
	getChildren() map[string]traversable
	setChild(i string, t traversable)
}

type traversalFunc func(parent, child traversable) (traversable, error)

// traverse takes a traversable node and applies some function to the node within the tree. It recursively calls itself and fails early when an error is thrown
func traverse(node traversable, f traversalFunc) (traversable, error) {
	children := node.getChildren()
	for i, child := range children {
		newChild, err := f(node, child)
		if err != nil {
			return nil, err
		}
		node.setChild(i, newChild)
		_, err = traverse(newChild, f)
		if err != nil {
			return nil, err
		}
	}

	return node, nil
}

type childContainerMap struct {
	mapContainer map[string]traversable
}

func (c childContainerMap) get(i string) traversable {
	return c.mapContainer[i]
}
func (c childContainerMap) set(i string, child traversable) {
	c.mapContainer[i] = child
}
func (c childContainerMap) len() int {
	return len(c.mapContainer)
}

type childContainerSlice struct {
	sliceContainer []traversable
}

func (c childContainerSlice) get(i int) traversable {
	return c.sliceContainer[i]
}
func (c childContainerSlice) set(i int, child traversable) {
	c.sliceContainer[i] = child
}
func (c childContainerSlice) len() int {
	return len(c.sliceContainer)
}
