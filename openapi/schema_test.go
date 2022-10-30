package openapi

import (
	"path/filepath"
	"testing"
)

func TestSchema_ResolveRefs(t *testing.T) {
	basePath := "../"

	ref := "openapi/fixtures/test_schema.yaml"
	absoluteRef, _ := filepath.Abs(filepath.Join(basePath, ref))

	subRef := "openapi/fixtures/sub_schema.yml"
	absoluteSubRef, _ := filepath.Abs(filepath.Join(basePath, subRef))

	schema := Schema{
		Ref: ref,
	}

	schema.ResolveRefs(basePath)

	if schema.Ref != absoluteRef {
		t.Log("Schema reference was not correctly resolved to an absolute path")
		t.Fail()
	}

	if schema.Items.Ref != absoluteSubRef {
		t.Log("Recursive Schema reference was not correctly resolved to an absolute path")
		t.Fail()
	}
}
