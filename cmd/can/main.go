package main

import (
	"fmt"
	"github.com/sasswart/gin-in-a-can/config"
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/render"
	golang "github.com/sasswart/gin-in-a-can/render/go"
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

	fmt.Printf("Reading API specification from \"%s\"\n", cfg.GetOpenAPIFilepath())
	apiSpec, err := openapi.LoadAPISpec(cfg.GetOpenAPIFilepath())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Setup appropriate renderer
	switch cfg.Template.Name {
	case "go-gin", "go-client":
		render.Engine{}.New(golang.Renderer{}, cfg)
	case "openapi-3":
		fmt.Printf("Openapi-3 renderer not implemented yet")
		os.Exit(1)
	default:
		fmt.Printf("%s is not a valid template name. Could not instantiate renderer", cfg.Template.Name)
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	apiSpec.SetMetadata(map[string]string{
		// TODO this doesn't look right
		"package": cfg.Generator.BasePackageName,
	})

	_, err = tree.Traverse(apiSpec, Renderer.BuildRenderNode())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
