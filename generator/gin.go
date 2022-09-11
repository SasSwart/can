package generator

import (
	"os"
	"strings"
	"text/template"

	"github.gom/sasswart/gin-in-a-can/openapi"
)

func Generate(apiSpec *openapi.OpenAPI) map[string][]byte {
	buffers := make(map[string][]byte)
	buffers["server.go"] = server(apiSpec)

	return buffers
}

func server(apiSpec *openapi.OpenAPI) []byte {
	templater := template.New("server.tmpl")

	templater.Funcs(template.FuncMap{
		"ToUpper": strings.ToUpper,
	})

	parsedTemplate, err := templater.ParseGlob("templates/*.tmpl")
	if err != nil {
		panic(err)
	}

	err = parsedTemplate.Execute(os.Stdout, NewServerInterface(apiSpec))
	if err != nil {
		panic(err)
	}

	return []byte{}
}
