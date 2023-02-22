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
	cfg := config.Data{}
	err := cfg.Load()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Reading API specification from \"%s\"\n", cfg.AbsOpenAPIPath)
	apiSpec, err := openapi.LoadAPISpec(cfg.AbsOpenAPIPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	Renderer = render.Engine{}.New(render.GinRenderer{}, cfg)

	apiSpec.SetMetadata(map[string]string{
		// TODO this doesn't look right
		"package": cfg.Generator.BasePackageName,
	})

	renderNode := Renderer.BuildRenderNode()
	_, err = tree.Traverse(apiSpec, renderNode)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
