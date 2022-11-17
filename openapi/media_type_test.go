package openapi

import (
	"reflect"
	"testing"
)

func getMediaType() *MediaType {
	openapi, _ := LoadOpenAPI(openapiFile)
	mt := openapi.getChildren()["/endpoint"].(*PathItem).getChildren()["post"].(*Operation).getChildren()["RequestBody"].(*RequestBody).getChildren()["application/json"]
	cast, ok := mt.(*MediaType)
	if ok {
		return cast
	}
	return nil
}

func TestMediaType_GetName(t *testing.T) {
	//mt := getMediaType()
	// TODO
	t.Fail()
}

func TestMediaType_GetChildren(t *testing.T) {
	mt := getMediaType()
	transversable := mt.getChildren()
	s, ok := transversable["Model"].(*Schema)
	if !ok {
		t.Logf("MediaType.getChildren() didn't return a *Schema")
		t.Fail()
	}
	if s == nil {
		t.Logf("MediaType.getChildren() returned a nil schema")
		t.Fail()
	}
}

func TestMediaType_SetChild(t *testing.T) {
	mt := getMediaType()
	transversable := mt.getChildren()
	s, _ := transversable["Model"].(*Schema)
	sOld := *s
	t.Logf("Original schema name: %v", s.name)
	newSchemaName := "NewSchema"
	mt.setChild("Model", &Schema{name: newSchemaName})
	transversable = mt.getChildren()
	sNew, _ := transversable["Model"].(*Schema)
	if reflect.DeepEqual(s, sNew) {
		t.Errorf("Child schema not set successfully")
		t.Fail()
	}
	if (&sOld).name == sNew.name {
		t.Errorf("Child schema not set successfully")
		t.Errorf("New Schema name: %v, Expected New Schema Name: %v", sNew.name, newSchemaName)
		t.Fail()
	}
}
