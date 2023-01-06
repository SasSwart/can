// GENERATED CODE. DO NOT EDIT

package api

import "github.com/gin-gonic/gin"

type Server interface {
  PetsSwaggerPetstore
  PetsPetidSwaggerPetstore
}

func RegisterServer(e *gin.Engine, srv Server) {
  RegisterPetsSwaggerPetstorePath(e, srv)
  RegisterPetsPetidSwaggerPetstorePath(e, srv)
}

type UnimplementedServer struct {
  UnimplementedPetsSwaggerPetstore
  UnimplementedPetsPetidSwaggerPetstore
}

func (u UnimplementedServer) InvalidRequest(c *gin.Context, err error) {
	c.JSON(400, err.Error())
}