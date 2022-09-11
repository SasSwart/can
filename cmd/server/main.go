package main

import (
	"github.com/gin-gonic/gin"

	"github.gom/sasswart/gin-in-a-can/api"
)

var _ api.Server = server{}

type server struct {
	api.UnimplementedServer
}

func main() {
	r := gin.Default()

	api.RegisterServer(r, &server{})

	r.Run()
}
