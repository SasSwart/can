package tree_test

import "testing"

func TestOpenAPI_ResolveRefs(t *testing.T) {
	// TODO test resolveRefs through public methods that call private methods
	//apiSpec := openapi.OpenAPI{
	//	Node: tree.Node{
	//		//basePath: filepath.Dir(test.AbsOpenAPI),
	//	},
	//	//Components: components.Components{},
	//	Paths: map[string]*path.Item{},
	//}
	//content, _ := os.ReadFile(test.AbsOpenAPI)
	//_ = yaml.Unmarshal(content, &apiSpec)
	//
	//newApi, err := tree.Traverse(&apiSpec, openapi.ResolveRefs)
	//
	//if err != nil {
	//	t.Errorf(err.Error()) // just bubbling up is enough here
	//}
	//if newApi == nil {
	//	t.Errorf("%s resulted in a nil api tree", test.OpenapiFile)
	//}
}
func TestOpenAPI_readRef(t *testing.T) {
	// ReadRef takes a reference and attempts to unmarshal it's content into the struct being passed as `i`.
	// As it happens, this ref is contained within the struct that is being unmarshalled into.
	//func readRef(absFilename string, n tree.NodeTraverser) error {
}
