package render

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/sasswart/gin-in-a-can/tree"
	"os"
	"path/filepath"
	"text/template"
)

type EngineInterface interface {
	New(renderer Renderer, config Config) *Engine
	GetRenderer() Renderer
	Render(data tree.NodeTraverser, templateFile string) ([]byte, error)
}
type Engine struct {
	renderer Renderer
	config   Config
}

var _ EngineInterface = Engine{}

func (e Engine) New(renderer Renderer, config Config) *Engine {
	return &Engine{renderer: renderer, config: config}
}

func (e Engine) GetRenderer() Renderer {
	return e.renderer
}

// Render contains the parsing and rendering steps
func (e Engine) Render(node tree.NodeTraverser, templateFile string) ([]byte, error) {
	var templateDirAbs string
	switch true {
	case filepath.IsAbs(e.config.TemplateDirectory):
		templateDirAbs = filepath.Join(
			e.config.TemplateDirectory,
			e.config.TemplateName,
		)
		// TODO Render shouldn't have to know about the ConfigFilePath file path. This seems hacky
	case filepath.IsAbs(e.config.ConfigFilePath):
		templateDirAbs = filepath.Join(
			filepath.Dir(e.config.ConfigFilePath),
			e.config.TemplateDirectory,
			e.config.TemplateName,
		)
	default:
		templateDirAbs = filepath.Join(
			e.config.ConfigFilePath,
			filepath.Dir(e.config.ConfigFilePath),
			e.config.TemplateDirectory,
			e.config.TemplateName,
		)
	}

	var outputFileAbs string
	switch true {
	case filepath.IsAbs(e.config.ConfigFilePath):
		outputFileAbs = e.config.ConfigFilePath
	case filepath.IsAbs(e.config.ConfigFilePath):
		outputFileAbs = filepath.Join(
			filepath.Dir(e.config.ConfigFilePath),
			e.config.OutputPath,
		)
	default:
		outputFileAbs = filepath.Join(
			e.config.WorkingDirectory,
			filepath.Dir(e.config.ConfigFilePath),
			e.config.OutputPath,
		)
	}
	outputFileAbs = filepath.Join(outputFileAbs, e.renderer.GetOutputFile(node))

	buff := bytes.NewBuffer([]byte{})

	templater := template.New(templateFile)

	templater.Funcs(templateFuncMap)

	parsedTemplate, err := templater.ParseGlob(fmt.Sprintf("%s/*.tmpl", templateDirAbs))
	if err != nil {
		return nil, err
	}

	err = parsedTemplate.Execute(buff, node)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Rendering %s using %s\n", e.renderer.GetOutputFile(node), templateFile)

	outputDirAbs := filepath.Dir(outputFileAbs)
	if _, err := os.Stat(outputDirAbs); errors.Is(err, os.ErrNotExist) {
		_ = os.MkdirAll(outputDirAbs, 0755)
	}

	err = os.WriteFile(outputFileAbs, buff.Bytes(), 0644)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}
