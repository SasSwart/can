package generator

import (
	"bytes"
	"strings"
	"text/template"
)

func GenerateController(config TemplateConfig) []byte {
	buff := bytes.NewBuffer([]byte{})

	templater := template.New("controller.tmpl")

	templater.Funcs(templateFuncMap)

	parsedTemplate, err := templater.ParseGlob("templates/*.tmpl")
	if err != nil {
		panic(err)
	}

	err = parsedTemplate.Execute(buff, config)
	if err != nil {
		panic(err)
	}

	return buff.Bytes()
}

func GenerateModels(config TemplateConfig) []byte {
	buff := bytes.NewBuffer([]byte{})

	templater := template.New("models.tmpl")

	templater.Funcs(templateFuncMap)

	parsedTemplate, err := templater.ParseGlob("templates/*.tmpl")
	if err != nil {
		panic(err)
	}

	err = parsedTemplate.Execute(buff, config)
	if err != nil {
		panic(err)
	}

	return buff.Bytes()
}

func GenerateUnimplementedServer(config TemplateConfig) []byte {
	buff := bytes.NewBuffer([]byte{})

	templater := template.New("unimplemented_server.tmpl")

	templater.Funcs(templateFuncMap)

	parsedTemplate, err := templater.ParseGlob("templates/*.tmpl")
	if err != nil {
		panic(err)
	}

	err = parsedTemplate.Execute(buff, config)
	if err != nil {
		panic(err)
	}

	return buff.Bytes()
}

var templateFuncMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
	"ToTitle": strings.ToTitle,
}
