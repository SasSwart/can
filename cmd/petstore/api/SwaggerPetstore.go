// GENERATED CODE. DO NOT EDIT

package api

import "github.com/gin-gonic/gin"

type Server interface {
  SwaggerPetstorePets
  SwaggerPetstorePetsPetid
}

func RegisterServer(e *gin.Engine, srv Server) {
  RegisterSwaggerPetstorePetsPath(e, srv)
  RegisterSwaggerPetstorePetsPetidPath(e, srv)
}

type UnimplementedServer struct {
  UnimplementedSwaggerPetstorePets
  UnimplementedSwaggerPetstorePetsPetid
}

func (u UnimplementedServer) InvalidRequest(c *gin.Context, err error) {
	c.JSON(400, err.Error())
}