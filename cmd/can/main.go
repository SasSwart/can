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
	fmt.Printf("can %s\n", config.SemVer)
	cfg := config.Data{}
	err := cfg.Load()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if config.Debug {
		fmt.Printf("Reading API specification from \"%s\"\n", cfg.GetOpenAPIFilepath())
	}
	apiSpec, err := openapi.LoadFromYaml(cfg.GetOpenAPIFilepath())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Setup appropriate renderer via the `strategy` design pattern
	err = setRenderStrategy(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	apiSpec.SetMetadata(tree.Metadata{
		"package": cfg.Template.BasePackageName,
	})

	_, err = tree.Traverse(apiSpec, Renderer.BuildRenderNode())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func setRenderStrategy(cfg config.Data) error {
	switch cfg.Template.Name {
	case "go-gin", "go-client":
		e := render.Engine{}
		r := &golang.Renderer{Base: &render.Base{}}
		r.SetTemplateFuncMap(nil)
		Renderer = e.With(r, cfg)
		return nil
	case "openapi-3":
		return fmt.Errorf("openapi-3 renderer not implemented yet")
	default:
		return fmt.Errorf("%s is not a valid template name. Could not instantiate renderer", cfg.Template.Name)
	}
}
