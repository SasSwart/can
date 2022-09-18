package generator

import "github.gom/sasswart/gin-in-a-can/openapi"

type Responses struct {
	Interface string
	Responses map[string]Response
}

type Response struct {
	Name   string
	Schema openapi.Schema
}
