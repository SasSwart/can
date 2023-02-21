package tree

import "fmt"

type Traverser interface {
	GetChildren() map[string]NodeTraverser
	SetChild(i string, t NodeTraverser)

	GetParent() NodeTraverser
	SetParent(parent NodeTraverser)
}

type TraversalFunc func(key string, parent, child NodeTraverser) (NodeTraverser, error)

// TraverseRecursor is an auxiliary function to Traverse that initiates a recursive loop over a tree of NodeTraverser
// structs, applying a given function at every step.
func TraverseRecursor[T NodeTraverser](node T, f TraversalFunc) (T, error) {
	children := node.GetChildren()
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
		node.SetChild(i, newChild)

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

// Traverse takes a NodeTraverser Node and enters into a recursive loop (TraverseRecursor) that applies a given function
// to the Node.
func Traverse[T NodeTraverser](node T, f TraversalFunc) (T, error) {
	if f == nil {
		return node, fmt.Errorf("no traversal function supplied")
	}

	result, err := f("", nil, node)
	if err != nil {
		return node, fmt.Errorf("traversal function[%p] error: %w", f, err)
	}

	node, ok := result.(T)
	if !ok {
		return node, fmt.Errorf("function parameter f should return the same type that is was given")
	}

	return TraverseRecursor(node, f) // An error is being swallowed here
}
