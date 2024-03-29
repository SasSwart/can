package tree

type Metadata map[string]string
type Node struct {
	basePath string
	parent   NodeTraverser
	metadata Metadata
	Name     string
}
type Noder interface {
	SetName(name string)
	GetName() []string

	GetBasePath() string
	SetBasePath(path string)

	GetRef() string
	SetMetadata(metadata Metadata)
	GetMetadata() Metadata
}

//
//// Noder Functions
//

func (n *Node) SetName(name string) {
	n.Name = name
}
func (n *Node) GetName() []string {
	if n.GetParent() == nil {
		return []string{n.Name}
	}
	return append(n.GetParent().GetName(), n.Name)
}

func (n *Node) SetBasePath(path string) {
	if n.GetParent() == nil {
		n.basePath = path
		return
	}
	n.GetParent().SetBasePath(path)
}

func (n *Node) GetBasePath() string {
	if n.GetParent() == nil {
		return n.basePath
	}
	return n.GetParent().GetBasePath()
}

// SetMetadata sets metadata for the root node of the tree
func (n *Node) SetMetadata(metadata Metadata) {
	if n.GetParent() == nil {
		if n.metadata == nil {
			n.metadata = make(Metadata, 2)
		}
		n.metadata = metadata
		return
	}
	n.GetParent().SetMetadata(metadata)
}

// GetMetadata reads metadata from the root node of the tree
func (n *Node) GetMetadata() Metadata {
	if n.GetParent() == nil {
		return n.metadata
	}
	return n.GetParent().GetMetadata()
}

//
//// Traverser Functions
//

func (n *Node) GetParent() NodeTraverser {
	return n.parent
}

func (n *Node) SetParent(parent NodeTraverser) {
	n.parent = parent
}
