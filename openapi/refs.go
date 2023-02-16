package openapi

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// readRef takes a reference and attempts to unmarshal it's content into the struct being passed as `i`.
// As it happens, this ref is contained within the struct that is being unmarshalled into.
func readRef(absFilename string, i interface{}) error {
	if Debug { // this can be a particularly noisy Printf call
		fmt.Printf("Reading reference: %s\n", absFilename)
	}
	content, err := os.ReadFile(absFilename)
	if err != nil {
		return fmt.Errorf("unable to resolve Reference: %w", err)
	}

	err = yaml.Unmarshal(content, i)
	if err != nil {
		return fmt.Errorf("unable to unmarshal reference file:\n%w", err)
	}

	return nil
}
