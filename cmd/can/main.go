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

func main() {
	fmt.Printf("can %s\n", config.SemVer)
	cfg := mustLoadConfig()
	if config.Debug {
		fmt.Printf("Reading API specification from \"%s\"\n", cfg.GetOpenAPIFilepath())
	}
	apiSpec := mustLoadOpenApiFile(cfg.GetOpenAPIFilepath())

	engine := render.NewEngine(cfg)

	// Setup appropriate renderer via the `strategy` design pattern
	mustSetStrategy(engine, cfg.Template.Strategy)

	apiSpec.SetMetadata(tree.Metadata{
		"package": cfg.Template.BasePackageName,
	})

	if _, err := tree.Traverse(apiSpec, engine.Render); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func mustSetStrategy(engine render.Engine, strategy string) {
	err := setStrategy(engine, strategy)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func mustLoadOpenApiFile(path string) *openapi.OpenAPI {
	// TODO: What if JSON?
	apiSpec, err := openapi.LoadFromYaml(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return apiSpec
}

// setStrategy creates and applies the renderer for the given strategy. An error is returned if the strategy is invalid
func setStrategy(e render.Engine, strategy string) error {
	var r render.Renderer
	switch strategy {
	case "go":
		r = &golang.Renderer{}
		r.SetTemplateFuncMap(golang.DefaultFuncMap())
		e.SetRenderer(r)
		return nil
	}
	return fmt.Errorf("%s render strategy not implemented yet", strategy)
}

func mustLoadConfig() config.Data {
	cfg := config.Data{}
	err := cfg.Load()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cfg
}
