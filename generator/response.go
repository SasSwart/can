package generator

import "github.com/sasswart/gin-in-a-can/model"

type Response struct {
	Name string
	model.Model
	StatusCode string
}
