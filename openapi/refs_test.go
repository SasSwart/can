package openapi

import "testing"

func TestOpenAPI_readRef(t *testing.T) {
	// readRef takes a reference and attempts to unmarshal it's content into the struct being passed as `i`.
	// As it happens, this ref is contained within the struct that is being unmarshalled into.
	// TODO can `i` become a non-empty interface? perhaps [Node | Traversable]?
	//func readRef(absFilename string, i interface{}) error {

	t.Errorf("TODO")
}
