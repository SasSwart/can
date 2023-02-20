package schema_test

import (
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"testing"
)

func TestOpenAPI_Schema_getChildren(t *testing.T) {
	t.Errorf("TODO")
}

func TestOpenAPI_Schema_IsRequired(t *testing.T) {
	nilSchema := &schema.Schema{
		Required: nil,
	}
	schemaWithRequiredName := &schema.Schema{
		Required: []string{"name"},
	}
	if nilSchema.IsRequired("name") {
		t.Fatalf("Name does not exist in nilSchema")
	}
	if !schemaWithRequiredName.IsRequired("name") {
		t.Fatalf("Name does exist in schemaWithRequiredName and was not found")
	}
}
