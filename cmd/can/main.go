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
	// TODO: MustLoadConfig
	cfg := config.Data{}
	err := cfg.Load()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if config.Debug {
		fmt.Printf("Reading API specification from \"%s\"\n", cfg.GetOpenAPIFilepath())
	}
	// TODO: What if JSON?
	// TODO: MustLoadOpenApiFile
	apiSpec, err := openapi.LoadFromYaml(cfg.GetOpenAPIFilepath())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	engine := render.NewEngine(cfg)

	// Setup appropriate renderer via the `strategy` design pattern
	// TODO: MustSetStrategy
	err = setStrategy(engine, cfg.Template.Strategy)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	apiSpec.SetMetadata(tree.Metadata{
		"package": cfg.Template.BasePackageName,
	})

	_, err = tree.Traverse(apiSpec, engine.Render)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
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
