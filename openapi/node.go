package openapi

type comparable interface {
	string | int
}
type childContainer[I comparable] interface {
	get(I) traversable
	set(I, traversable)
	len() int
}

type traversable interface {
	getParent() traversable
	getChildren() childContainer[comparable]
	setChild(i int, t traversable)
}

func traverse(n traversable, f func(traversable) traversable) traversable {
	children := n.getChildren()
	for i := 0; i < children.len(); i++ {
		child := children.get(i)
		newChild := f(child)
		n.setChild(i, newChild)

		traverse(newChild, f)
	}

	return n
}

type childContainerMap[T comparable] struct {
	mapContainer map[T]traversable
}

func (c childContainerMap[T]) get(i T) traversable {
	return c.mapContainer[i]
}
func (c childContainerMap[T]) set(i T, child traversable) {
	c.mapContainer[i] = child
}
func (c childContainerMap[T]) len() int {
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
