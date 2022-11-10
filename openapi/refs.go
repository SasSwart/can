package openapi

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func readRef(filename string, i interface{}) error {
	fmt.Printf("Reading reference: %s\n", filename)
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("unable to resolve PathItem Reference:\n%w", err)
	}

	err = yaml.Unmarshal(content, i)
	if err != nil {
		return fmt.Errorf("unable to unmarshal reference file:\n%w", err)
	}

	return nil
}
