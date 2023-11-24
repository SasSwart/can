package main

import (
	"bytes"
	"fmt"
	"github.com/sasswart/gin-in-a-can/config"
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/render"
	golang "github.com/sasswart/gin-in-a-can/render/go"
	"github.com/sasswart/gin-in-a-can/tree"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
)

var canCmd = &cobra.Command{
	Use:   "can",
	Short: "Generate code based on OpenAPI specifications",
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate files",
	Run: func(cmd *cobra.Command, args []string) {
		absCfgPath, err := filepath.Abs(config.ConfigFilePath)
		if err != nil {
			fmt.Printf("could not resolve relative config path: %v\n", err)
			return
		}
		config.ConfigFilePath = absCfgPath

		configFileReader, err := os.Open(absCfgPath)
		if err != nil {
			fmt.Printf("failed to open config file: %v", err)
			os.Exit(1)
		}

		configs := config.ReadConfigs(configFileReader)

		generate(configs)
	},
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove generated files",
	Run: func(cmd *cobra.Command, args []string) {
		absCfgPath, err := filepath.Abs(config.ConfigFilePath)
		if err != nil {
			fmt.Printf("could not resolve relative config path: %v\n", err)
			return
		}
		config.ConfigFilePath = absCfgPath

		configFileReader, err := os.Open(absCfgPath)
		if err != nil {
			fmt.Printf("failed to open config file: %v", err)
			os.Exit(1)
		}

		configs := config.ReadConfigs(configFileReader)

		clean(configs)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Can: v%s\n", config.SemVer)
	},
}

func main() {
	canCmd.PersistentFlags().StringVarP(
		&config.ConfigFilePath,
		"configFile",
		"c",
		"./can.yml",
		"config file (default is ./can.yml)")

	canCmd.PersistentFlags().BoolVarP(
		&config.Debug,
		"debug",
		"d",
		false,
		"Enable Debug logging")

	canCmd.PersistentFlags().BoolVarP(
		&config.Dryrun,
		"dry-run",
		"r",
		false,
		"Print actions instead of applying them to disk")

	canCmd.AddCommand(generateCmd)
	canCmd.AddCommand(cleanCmd)
	canCmd.AddCommand(versionCmd)

	if err := canCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func clean(configs <-chan []byte) {
	wg := sync.WaitGroup{}
	for configBytes := range configs {

		wg.Add(1)

		go func(configBytes []byte) {
			defer wg.Done()

			if len(configBytes) == 0 {
				if config.Debug {
					fmt.Println("Skipping empty config")
				}
				return
			}

			configReader := bytes.NewReader(configBytes)
			cfg := mustLoadConfig(configReader)

			if config.Dryrun {
				fmt.Printf("Would delete %s\n", cfg.GetOutputDir())
				return
			}

			err := os.RemoveAll(cfg.GetOutputDir())
			if err != nil {
				fmt.Println(err)
			}

		}(configBytes)
	}
	wg.Wait()
}

func generate(configs <-chan []byte) {
	wg := sync.WaitGroup{}
	for configBytes := range configs {

		wg.Add(1)

		go func(configBytes []byte) {
			defer wg.Done()

			if len(configBytes) == 0 {
				if config.Debug {
					fmt.Println("Skipping empty config")
				}
				return
			}

			configReader := bytes.NewReader(configBytes)
			cfg := mustLoadConfig(configReader)

			err := executeJobForConfig(cfg)
			if err != nil {
				fmt.Println(err)
			}
		}(configBytes)
	}
	wg.Wait()
}

func executeJobForConfig(cfg config.Data) error {
	fmt.Printf("Executing job: %s\n", cfg.Name)

	if config.Debug {
		fmt.Printf("Reading API specification from \"%s\"\n", cfg.GetOpenAPIFilepath())
	}
	apiSpec := mustLoadOpenApiFile(cfg.GetOpenAPIFilepath())

	engine := render.NewEngine(cfg)

	// Setup appropriate renderer via the `strategy` design pattern
	err := setStrategy(&engine, cfg.Template.Strategy)
	if err != nil {
		return err
	}

	apiSpec.SetMetadata(tree.Metadata{
		"package": cfg.Template.BasePackageName,
	})

	if _, err := tree.Traverse(apiSpec, engine.Render); err != nil {
		return err
	}
	return nil
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
func setStrategy(e *render.Engine, strategy string) error {
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

func mustLoadConfig(reader io.Reader) config.Data {
	cfg := config.Data{}
	err := cfg.Load(reader)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cfg
}
