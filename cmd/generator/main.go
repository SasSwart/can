package main

import (
	"fmt"
	"os"

	"github.gom/sasswart/gin-in-a-can/generator"
	"github.gom/sasswart/gin-in-a-can/openapi"
)

func main() {
	apiSpec, err := openapi.LoadOpenAPI("openapi.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Lex
	generator.Generate(apiSpec)
	// Parse
	// Generate
}
