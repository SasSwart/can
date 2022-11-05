package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sasswart/gin-in-a-can/cmd/test_api/api/controller"
)

func main() {
	engine := gin.Default()

	controller.RegisterServer(engine, server{})
}

type server struct {
	controller.UnimplementedServer
}
