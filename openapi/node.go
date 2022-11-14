package openapi

type traversable interface {
	getParent() traversable
	getChildren() map[T]traversable
	setChild(i T, t traversable)
}

type node[p, c traversable[comparable]] struct {
	parent   p
	children []c
}

//func (n node[p]) ResolveRefs() error {
//	return nil
//}

func traverse[T comparable](n traversable[T], f func(traversable[any]) traversable[any]) traversable[T] {
	children := n.getChildren()
	for i, child := range children {
		newChild := f(child)
		n.setChild(i, newChild)

		traverse(newChild, f)
	}

	return n
}

func (n node[p, c]) getChildren() []c {
	return n.children
}

func (n node[p, c]) setChild(i int, child c) {
	n.children[i] = child
}
