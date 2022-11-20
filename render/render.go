package render

import (
	"bytes"
	"fmt"
	"github.com/sasswart/gin-in-a-can/config"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"path/filepath"
	"strings"
	"text/template"
)

// Render is the main parsing and rendering steps within the render library
func Render(config config.Config, data any, templateFile string) ([]byte, error) {
	buff := bytes.NewBuffer([]byte{})

	templater := template.New(templateFile)

	templater.Funcs(templateFuncMap)

	absoluteTemplateDirectory := filepath.Join(
		config.WorkingDirectory,
		filepath.Dir(config.OpenAPI.OpenAPIFile),
		config.Generator.TemplateDirectory,
	)
	parsedTemplate, err := templater.ParseGlob(fmt.Sprintf("%s/*.tmpl", absoluteTemplateDirectory))
	if err != nil {
		return nil, err
	}

	err = parsedTemplate.Execute(buff, data)
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
