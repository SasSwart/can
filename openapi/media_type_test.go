package openapi

import (
	"path/filepath"
	"testing"
)

func TestMediaType_ResolveRefs(t *testing.T) {
	basePath := "../"

	ref := "openapi/fixtures/test_schema.yaml"
	absoluteRef, _ := filepath.Abs(filepath.Join(basePath, ref))

	subRef := "openapi/fixtures/sub_schema.yml"
	absoluteSubRef, _ := filepath.Abs(filepath.Join(basePath, subRef))

	mediaType := MediaType{
		Schema: &Schema{
			Ref: ref,
		},
	}

	mediaType.ResolveRefs(basePath)

	if mediaType.Schema.Ref != absoluteRef {
		t.Log("Schema reference was not correctly resolved to an absolute path")
		t.Fail()
	}

	if mediaType.Schema.Items.Ref != absoluteSubRef {
		t.Log("Recursive Schema reference was not correctly resolved to an absolute path")
		t.Fail()
	}
}
