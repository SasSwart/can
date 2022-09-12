package main

import (
	"github.com/gin-gonic/gin"

	"github.gom/sasswart/gin-in-a-can/api"
)

type server struct {
	api.UnimplementedServer
}

// Compile time assertion that server must implement the api.Server interface
var _ api.Server = server{}

func main() {
	r := gin.Default()

	// Add your auth, logging and other middleware here

	// Pass the server to the generated framework to register the OpenAPI routes
	api.RegisterServer(r, &server{})

	r.Run()
}
