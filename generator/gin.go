package generator

import (
	"bytes"
	"strings"
	"text/template"

	"github.gom/sasswart/gin-in-a-can/openapi"
)

func Generate(apiSpec *openapi.OpenAPI) []byte {
	buff := bytes.NewBuffer([]byte{})

	templater := template.New("server.tmpl")

	templater.Funcs(template.FuncMap{
		"ToUpper": strings.ToUpper,
	})

	parsedTemplate, err := templater.ParseGlob("templates/*.tmpl")
	if err != nil {
		panic(err)
	}

	err = parsedTemplate.Execute(buff, NewServerInterface(apiSpec))
	if err != nil {
		panic(err)
	}

	return buff.Bytes()
}
