package openapi

type Components struct {
	Schemas       map[string]Schema
	Responses     map[string]Response
	RequestBodies map[string]RequestBody
}
