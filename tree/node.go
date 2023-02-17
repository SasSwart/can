package tree

import (
	"fmt"
	"github.com/sasswart/gin-in-a-can/errors"
	"github.com/sasswart/gin-in-a-can/render"

	// TODO tree package shouldn't have to know about the root
	"github.com/sasswart/gin-in-a-can/openapi/root"
)

type Noder interface {
	GetName() string
	SetName(name string)

	GetBasePath() string

	GetRef() string

	GetOutputFile() string

	SetMetadata(metadata map[string]string)
	GetMetadata() map[string]string
}

type Traverser interface {
	GetChildren() map[string]NodeTraverser
	SetChild(i string, t NodeTraverser)

	GetParent() NodeTraverser
	SetParent(parent NodeTraverser)
}
type NodeTraverser interface {
	Traverser
	Noder
}

type TraversalFunc func(key string, parent, child NodeTraverser) (NodeTraverser, error)

type Node struct {
	basePath string
	parent   NodeTraverser
	name     string
	renderer render.Renderer
}

func (n *Node) SetMetadata(metadata map[string]string) {
	if n.parent.GetParent() == nil {
		// TODO tree package shouldn't have to know about the root
		top, ok := n.GetParent().(*root.Root)
		if ok {
			top.metadata = metadata
			return
		}
		panic("we should never get here as *Root should always be at the top of the tree")
	}
	n.parent.SetMetadata(metadata)
}

var _ NodeTraverser = &Node{}

func (n *Node) GetMetadata() map[string]string {
	// TODO tree package shouldn't have to know about the root
	if top, ok := n.parent.(*root.Root); ok {
		return top.metadata
	}
	return n.parent.GetMetadata()
}

func (n *Node) GetChildren() map[string]NodeTraverser {
	errors.Unimplemented("(n *Node) GetChildren()")
	return nil
}

func (n *Node) SetChild(_ string, _ NodeTraverser) {
	errors.Unimplemented("(n *Node) SetChild()")
}

func (n *Node) GetParent() NodeTraverser {
	return n.parent
}

func (n *Node) SetParent(parent NodeTraverser) {
	n.parent = parent
}

// GetBasePath recurses up the parental ladder until it's overridden by the *Root method
func (n *Node) GetBasePath() string {
	if n.parent == nil {
		return n.basePath
	}
	return n.parent.GetParent().GetBasePath()
}

//func (n *Node) GetOutputFile() string {
//	// TODO this function can do without it's overrides
//	return n.GetRenderer().GetOutputFile(n)
//}

func (n *Node) GetRef() string {
	errors.Unimplemented("(n *Node) GetRef()")
	return ""
}

func (n *Node) GetName() string {
	// TODO tree package shouldn't have to know about the renderer
	name := n.parent.GetName() + n.GetRenderer().SanitiseName(n.name)
	return name
}

func (n *Node) SetName(name string) {
	n.name = name
}

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

	return TraverseRecursor(node, f)
}
