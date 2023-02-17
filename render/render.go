package render

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/sasswart/gin-in-a-can/config"
	"github.com/sasswart/gin-in-a-can/tree"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Config struct {
	ModuleName        string
	BasePackageName   string
	TemplateDirectory string
	TemplateName      string
}
type EngineInterface interface {
	New(renderer Renderer, config Config) *Engine
	GetRenderer() Renderer
	Render(config config.Config, data tree.NodeTraverser, templateFile string) ([]byte, error)

	// TODO Not needed after New() constructor introduced
	//SetRenderer(renderer Renderer)
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

type Renderer interface {
	SanitiseName(string) string
	SanitiseType(n tree.NodeTraverser) string

	GetOutputFile(n tree.NodeTraverser) string
}

// Render contains the parsing and rendering steps
func (e Engine) Render(config config.Config, data tree.NodeTraverser, templateFile string) ([]byte, error) {
	var templateDirAbs string
	switch true {
	case filepath.IsAbs(e.config.TemplateDirectory):
		templateDirAbs = filepath.Join(
			e.config.TemplateDirectory,
			e.config.TemplateName,
		)
		// TODO Render shouldn't have to know about the config file path. This seems hacky
	case filepath.IsAbs(config.ConfigFilePath):
		templateDirAbs = filepath.Join(
			filepath.Dir(config.ConfigFilePath),
			e.config.TemplateDirectory,
			e.config.TemplateName,
		)
	default:
		templateDirAbs = filepath.Join(
			config.WorkingDirectory,
			filepath.Dir(config.ConfigFilePath),
			e.config.TemplateDirectory,
			e.config.TemplateName,
		)
	}

	var outputFileAbs string
	switch true {
	case filepath.IsAbs(config.OutputPath):
		outputFileAbs = config.OutputPath
	case filepath.IsAbs(config.ConfigFilePath):
		outputFileAbs = filepath.Join(
			filepath.Dir(config.ConfigFilePath),
			config.OutputPath,
		)
	default:
		outputFileAbs = filepath.Join(
			config.WorkingDirectory,
			filepath.Dir(config.ConfigFilePath),
			config.OutputPath,
		)
	}
	outputFileAbs = filepath.Join(outputFileAbs, data.GetOutputFile())

	buff := bytes.NewBuffer([]byte{})

	templater := template.New(templateFile)

	templater.Funcs(templateFuncMap)

	parsedTemplate, err := templater.ParseGlob(fmt.Sprintf("%s/*.tmpl", templateDirAbs))
	if err != nil {
		return nil, err
	}

	err = parsedTemplate.Execute(buff, data)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Rendering %s using %s\n", data.GetOutputFile(), templateFile)

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

var templateFuncMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
	"ToTitle": toTitleCase,
}

func toTitleCase(s string) string {
	caser := cases.Title(language.English)
	return caser.String(s)
}
