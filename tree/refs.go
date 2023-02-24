package tree

import (
	"fmt"
	"github.com/sasswart/gin-in-a-can/config"
	"github.com/sasswart/gin-in-a-can/errors"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

// ReadRef takes a reference and attempts to unmarshal it's content into the struct being passed as `i`.
// As it happens, this ref is contained within the struct that is being unmarshalled into.
func readRef(absFilename string, n NodeTraverser) error {
	if errors.Debug { // this can be a particularly noisy Printf call
		fmt.Printf("[%s]::Reading reference: %s\n", config.SemVer, absFilename)
	}
	content, err := os.ReadFile(absFilename)
	if err != nil {
		return fmt.Errorf("unable to resolve Reference: %w", err)
	}

	err = yaml.Unmarshal(content, n)
	if err != nil {
		return fmt.Errorf("unable to unmarshal reference file:\n%w", err)
	}

	return nil
}

// ResolveRefs calls readRef on references with the ref path modified appropriately for it's use
func ResolveRefs(key string, parent, node NodeTraverser) (NodeTraverser, error) {
	node.SetParent(parent)
	if node.GetParent() != nil {
		node.SetName(key) // Don't set the root name as that's already been done by this point
	}
	nodeRef := node.GetRef()
	if nodeRef != "" {
		openapiBasePath := node.GetBasePath()
		ref := filepath.Base(node.GetRef())
		path := filepath.Join(openapiBasePath, ref)
		err := readRef(path, node)
		if err != nil {
			return nil, fmt.Errorf("Unable to read reference:\n%w", err)
		}
	}
	return node, nil
}
