package tree

import (
	"fmt"
	"github.com/sasswart/gin-in-a-can/config"
	"time"
)

type Traverser interface {
	GetChildren() map[string]NodeTraverser
	SetChild(i string, t NodeTraverser)

	GetParent() NodeTraverser
	SetParent(parent NodeTraverser)
}

type TraversalFunc func(key string, parent, child NodeTraverser) (NodeTraverser, error)

// traverseRecursor is an auxiliary function to Traverse that initiates a recursive loop over a tree of NodeTraverser
// structs, applying a given function at every step.
func traverseRecursor[T NodeTraverser](node T, f TraversalFunc) (T, error) {
	if config.Debug {
		t := time.Now()
		fmt.Println("traverseRecursor timer start")
		defer fmt.Println("traverseRecursor" + time.Since(t).String())
	}
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
		_, err = traverseRecursor(newChild, f)
		if err != nil {
			return node, err
		}
	}

	return node, nil
}

// Traverse takes a NodeTraverser Node and enters into a recursive loop (traverseRecursor) that applies a given function
// to the Node.
func Traverse[T NodeTraverser](node T, f TraversalFunc) (T, error) {
	if config.Debug {
		t := time.Now()
		fmt.Println("Traverse timer start")
		defer fmt.Println("Traverse" + time.Since(t).String())
	}
	if f == nil {
		return node, fmt.Errorf("`Traverse`:: no traversal function supplied")
	}

	result, err := f("", nil, node)
	if err != nil {
		return node, fmt.Errorf("`Traverse`:: traversal function[%p] error: %w", f, err)
	}

	node, ok := result.(T)
	if !ok {
		return node, fmt.Errorf("`Traverse`:: function parameter f should return the same type that is was given")
	}
	// TODO An error is being swallowed here
	rec, err := traverseRecursor(node, f)
	if err != nil {
		return node, err
	}
	return rec, nil
}
