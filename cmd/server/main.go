package main

import (
	"github.com/gin-gonic/gin"
	"github.gom/sasswart/gin-in-a-can/api/controller"
)

type server struct {
	controller.UnimplementedServer
}

// Compile time assertion that server must implement the api.Server interface
var _ controller.Server = server{}

// 1. Copy controller functions from the /api/controller/unimplemented.go to this file
// 2. Change the receivers to (server)
// 3. Implement your business logic

func main() {
	r := gin.Default()

	// Add your auth, logging and other middleware here

	// Pass the server to the generated framework to register the OpenAPI routes
	controller.RegisterServer(r, &server{})

	r.Run()
}
