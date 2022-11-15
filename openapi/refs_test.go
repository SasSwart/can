package openapi

import (
	"path/filepath"
	"testing"
)

func TestReadRef(t *testing.T) {
	testSchema1 := "./fixtures/sub_schema.yml"
	testSchema2, _ := filepath.Abs(testSchema1)
	testSchema3 := "./fixtures/sub_schema.yaml"
	testSchema4, _ := filepath.Abs(testSchema3)

	tables := []struct {
		ref      string
		expected string
	}{
		{ref: testSchema1, expected: "This is a sub schema"}, // relative
		{ref: testSchema2, expected: "This is a sub schema"}, // absolute
		{ref: testSchema3, expected: "This is a sub schema"}, // relative
		{ref: testSchema4, expected: "This is a sub schema"}, // absolute
	}

	type testStruct struct {
		Description string `yaml:"description"`
	}

	for _, table := range tables {
		i := testStruct{}
		err := readRef(table.ref, &i)

		if err != nil {
			t.Logf("readRef failed on %s", table.ref)
			t.Fail()
		}

		if i.Description != table.expected {
			t.Errorf("")
		}
	}
}
