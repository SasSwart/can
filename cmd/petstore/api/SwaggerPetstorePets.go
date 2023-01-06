package api

import "github.com/gin-gonic/gin"

// GENERATED INTERFACE. DO NOT EDIT

type PetsSwaggerPetstore interface {
  GetPetsSwaggerPetstore(*gin.Context, *GetPetsSwaggerPetstoreParameters, *GetPetsSwaggerPetstoreRequestBody) GetPetsSwaggerPetstoreResponse
  PostPetsSwaggerPetstore(*gin.Context, *PostPetsSwaggerPetstoreParameters, *PostPetsSwaggerPetstoreRequestBody) PostPetsSwaggerPetstoreResponse
	InvalidRequest(*gin.Context, error)
}

type UnimplementedPetsSwaggerPetstore struct {}
func (u UnimplementedPetsSwaggerPetstore) GetPetsSwaggerPetstore(*gin.Context, *GetPetsSwaggerPetstoreParameters, *GetPetsSwaggerPetstoreRequestBody) GetPetsSwaggerPetstoreResponse {
	return GetPetsSwaggerPetstore405Response{}
}
func (u UnimplementedPetsSwaggerPetstore) PostPetsSwaggerPetstore(*gin.Context, *PostPetsSwaggerPetstoreParameters, *PostPetsSwaggerPetstoreRequestBody) PostPetsSwaggerPetstoreResponse {
	return PostPetsSwaggerPetstore405Response{}
}
func (u UnimplementedPetsSwaggerPetstore) InvalidRequest(c *gin.Context, err error) {
	c.JSON(400, err.Error())
}

func RegisterPetsSwaggerPetstorePath(e *gin.Engine, srv PetsSwaggerPetstore) {
	
  e.GET("/pets", func(c *gin.Context) {
  	params := &GetPetsSwaggerPetstoreParameters{}
  	body := &GetPetsSwaggerPetstoreRequestBody{}
		err := c.ShouldBindJSON(body)
		if err != nil {
			srv.InvalidRequest(c, err)
		}
    response := srv.GetPetsSwaggerPetstore(c, params, body)
    c.JSON(response.GetStatus(), response)
  })
	
  e.POST("/pets", func(c *gin.Context) {
  	params := &PostPetsSwaggerPetstoreParameters{}
  	body := &PostPetsSwaggerPetstoreRequestBody{}
		err := c.ShouldBindJSON(body)
		if err != nil {
			srv.InvalidRequest(c, err)
		}
    response := srv.PostPetsSwaggerPetstore(c, params, body)
    c.JSON(response.GetStatus(), response)
  })
}
type GetPetsSwaggerPetstoreResponse interface {
	GetStatus() int
}

type GetPetsSwaggerPetstore200Response GetPetsSwaggerPetstore200ResponseModel

func (GetPetsSwaggerPetstore200Response) GetStatus() int {
	return 200
}

type GetPetsSwaggerPetstore500Response GetPetsSwaggerPetstore500ResponseModel

func (GetPetsSwaggerPetstore500Response) GetStatus() int {
	return 500
}
type GetPetsSwaggerPetstore405Response struct {}
func (GetPetsSwaggerPetstore405Response) GetStatus() int {
	return 405
}
type PostPetsSwaggerPetstoreResponse interface {
	GetStatus() int
}


type PostPetsSwaggerPetstore500Response PostPetsSwaggerPetstore500ResponseModel

func (PostPetsSwaggerPetstore500Response) GetStatus() int {
	return 500
}
type PostPetsSwaggerPetstore405Response struct {}
func (PostPetsSwaggerPetstore405Response) GetStatus() int {
	return 405
}
