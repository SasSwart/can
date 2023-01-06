package api

import "github.com/gin-gonic/gin"

// GENERATED INTERFACE. DO NOT EDIT

type PetsPetidSwaggerPetstore interface {
	GetPetsPetidSwaggerPetstore(*gin.Context, *GetPetsPetidSwaggerPetstoreParameters, *RequestBody) GetPetsPetidSwaggerPetstoreResponse
	InvalidRequest(*gin.Context, error)
}

type UnimplementedPetsPetidSwaggerPetstore struct{}

func (u UnimplementedPetsPetidSwaggerPetstore) GetPetsPetidSwaggerPetstore(*gin.Context, *GetPetsPetidSwaggerPetstoreParameters, *RequestBody) GetPetsPetidSwaggerPetstoreResponse {
	return GetPetsPetidSwaggerPetstore405Response{}
}
func (u UnimplementedPetsPetidSwaggerPetstore) InvalidRequest(c *gin.Context, err error) {
	c.JSON(400, err.Error())
}

func RegisterPetsPetidSwaggerPetstorePath(e *gin.Engine, srv PetsPetidSwaggerPetstore) {

	e.GET("/pets/{petId}", func(c *gin.Context) {
		params := &GetPetsPetidSwaggerPetstoreParameters{}
		body := &RequestBody{}
		err := c.ShouldBindJSON(body)
		if err != nil {
			srv.InvalidRequest(c, err)
		}
		response := srv.GetPetsPetidSwaggerPetstore(c, params, body)
		c.JSON(response.GetStatus(), response)
	})
}

type GetPetsPetidSwaggerPetstoreResponse interface {
	GetStatus() int
}

type GetPetsPetidSwaggerPetstore200Response GetPetsPetidSwaggerPetstore200ResponseModel

func (GetPetsPetidSwaggerPetstore200Response) GetStatus() int {
	return 200
}

type GetPetsPetidSwaggerPetstoreDefaultResponse GetPetsPetidSwaggerPetstoreDefaultResponseModel

func (GetPetsPetidSwaggerPetstoreDefaultResponse) GetStatus() int {
	return 200
}

type GetPetsPetidSwaggerPetstore405Response struct{}

func (GetPetsPetidSwaggerPetstore405Response) GetStatus() int {
	return 405
}
