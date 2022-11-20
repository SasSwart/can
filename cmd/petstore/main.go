package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sasswart/gin-in-a-can/cmd/petstore/api"
)

func main() {
	engine := gin.Default()

	api.RegisterServer(engine, server{})

	engine.Run()
}

type server struct {
	api.UnimplementedServer
}
