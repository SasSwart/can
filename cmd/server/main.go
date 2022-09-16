package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.gom/sasswart/gin-in-a-can/api/controller"
	"github.gom/sasswart/gin-in-a-can/api/models"
)

type server struct {
	// Ensures compliance with the controller.Server interface even if not all functions are implemented
	controller.UnimplementedServer
}

// Compile time assertion that server must implement the api.Server interface
var _ controller.Server = server{}

// 1. Copy controller functions from the /api/controller/unimplemented.go to this file
// 2. Change the receivers to (server)
// 3. Implement your business logic
// eg:

func (server) UserGet(requestModel *models.UserGetRequest) (models.UserGetResponse, error) {
	fmt.Println(requestModel)
	return models.UserGet200Response{
		Enabled:  true,
		Name:     requestModel.Name,
		Password: "a strong password",
	}, nil
}
func (server) UserPost(requestModel *models.UserPostRequest) (models.UserPostResponse, error) {
	fmt.Println(requestModel)
	return models.UserPost204Response{}, nil
}

func main() {
	r := gin.Default()

	// Add your auth, logging and other middleware here

	// Pass the server to the generated framework to register the OpenAPI routes
	controller.RegisterServer(r, &server{})

	r.Run()
}
