package main

import (
	"fmt"
	"github.com/sasswart/gin-in-a-can/config"
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/render"
	"github.com/sasswart/gin-in-a-can/tree"
	"os"
)

var Renderer *render.Engine

func main() {
	config := config.Data{}
	err := config.Load()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Reading API specification from \"%s\"\n", config.AbsOpenAPIPath)
	apiSpec, err := openapi.LoadAPISpec(config.AbsOpenAPIPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	Renderer = render.Engine{}.New(render.GinRenderer{}, config)

	apiSpec.SetMetadata(map[string]string{
		// TODO this doesn't look right
		"package": config.Generator.BasePackageName,
	})

	renderNode := Renderer.BuildRenderNode()
	_, err = tree.Traverse(apiSpec, renderNode)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
