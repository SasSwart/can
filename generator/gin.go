package generator

import (
	"bytes"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"text/template"
)

func Generate(config TemplateConfig, templateFile string) ([]byte, error) {
	buff := bytes.NewBuffer([]byte{})

	templater := template.New(templateFile)

	templater.Funcs(templateFuncMap)

	parsedTemplate, err := templater.ParseGlob("templates/*.tmpl")
	if err != nil {
		return nil, err
	}

	err = parsedTemplate.Execute(buff, config)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

var templateFuncMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
	"ToTitle": toTitle,
}

func toTitle(s string) string {
	caser := cases.Title(language.English)
	return caser.String(s)
}
