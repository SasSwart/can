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

	newSchema, _ := Traverse(&schema, resolveRefs)

	if newSchema.(*Schema).Ref != absoluteRef {
		t.Log("Schema reference was not correctly resolved to an absolute path")
		t.Fail()
	}

	if newSchema.(*Schema).Items.Ref != absoluteSubRef {
		t.Log("Recursive Schema reference was not correctly resolved to an absolute path")
		t.Fail()
	}
}
