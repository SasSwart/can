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

type Engine struct {
	Renderer Renderer
}

func (e Engine) New(renderer Renderer) *Engine {
	return &Engine{Renderer: renderer}
}

type Renderer interface {
	SanitiseName(string) string

	SanitiseType(n *tree.Node) string

	GetOutputFile(n tree.Node) string

	SetRenderer(n *tree.Node) error
	GetRenderer(n *tree.Node) Renderer
}

// Render is the main parsing and rendering steps within the render library
func Render(config config.Config, data tree.NodeTraverser, templateFile string) ([]byte, error) {
	var absoluteTemplateDirectory string
	switch true {
	case filepath.IsAbs(config.Generator.TemplateDirectory):
		absoluteTemplateDirectory = filepath.Join(
			config.Generator.TemplateDirectory,
			config.Generator.TemplateName,
		)
	case filepath.IsAbs(config.ConfigFilePath):
		absoluteTemplateDirectory = filepath.Join(
			filepath.Dir(config.ConfigFilePath),
			config.Generator.TemplateDirectory,
			config.Generator.TemplateName,
		)
	default:
		absoluteTemplateDirectory = filepath.Join(
			config.WorkingDirectory,
			filepath.Dir(config.ConfigFilePath),
			config.Generator.TemplateDirectory,
			config.Generator.TemplateName,
		)
	}

	var absoluteOutputFile string
	switch true {
	case filepath.IsAbs(config.OutputPath):
		absoluteOutputFile = config.OutputPath
	case filepath.IsAbs(config.ConfigFilePath):
		absoluteOutputFile = filepath.Join(
			filepath.Dir(config.ConfigFilePath),
			config.OutputPath,
		)
	default:
		absoluteOutputFile = filepath.Join(
			config.WorkingDirectory,
			filepath.Dir(config.ConfigFilePath),
			config.OutputPath,
		)
	}
	absoluteOutputFile = filepath.Join(absoluteOutputFile, data.GetOutputFile())

	buff := bytes.NewBuffer([]byte{})

	templater := template.New(templateFile)

	templater.Funcs(templateFuncMap)

	parsedTemplate, err := templater.ParseGlob(fmt.Sprintf("%s/*.tmpl", absoluteTemplateDirectory))
	if err != nil {
		return nil, err
	}

	err = parsedTemplate.Execute(buff, data)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Rendering %s using %s\n", data.GetOutputFile(), templateFile)

	outputDirectory := filepath.Dir(absoluteOutputFile)
	if _, err := os.Stat(outputDirectory); errors.Is(err, os.ErrNotExist) {
		_ = os.MkdirAll(outputDirectory, 0755)
	}

	err = os.WriteFile(absoluteOutputFile, buff.Bytes(), 0644)
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
