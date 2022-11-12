package openapi

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// readRef takes a reference and attempts to unmarshal it's content into the struct being passed as `i`.
func readRef(filename string, i interface{}) error {
	fmt.Printf("Reading reference: %s\n", filename)
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("unable to resolve Path Reference:\n%w", err)
	}

	// TODO: schema.go passes in reference to base struct for unmarshalling. This should instead be the ref object within the struct.
	// TODO: look to reduce complexity. Possibly return values instead of OO style manipulation.
	err = yaml.Unmarshal(content, i)
	if err != nil {
		return fmt.Errorf("unable to unmarshal reference file:\n%w", err)
	}

	return nil
}
