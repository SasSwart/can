package api

import "github.com/gin-gonic/gin"

// GENERATED INTERFACE. DO NOT EDIT

type PetsSwaggerPetstore interface {
	GetPetsSwaggerPetstore(*gin.Context, *GetPetsSwaggerPetstoreParameters, *RequestBody) GetPetsSwaggerPetstoreResponse
	PostPetsSwaggerPetstore(*gin.Context, *PostPetsSwaggerPetstoreParameters, *RequestBody) PostPetsSwaggerPetstoreResponse
	InvalidRequest(*gin.Context, error)
}

type UnimplementedPetsSwaggerPetstore struct{}

func (u UnimplementedPetsSwaggerPetstore) GetPetsSwaggerPetstore(*gin.Context, *GetPetsSwaggerPetstoreParameters, *RequestBody) GetPetsSwaggerPetstoreResponse {
	return GetPetsSwaggerPetstore405Response{}
}
func (u UnimplementedPetsSwaggerPetstore) PostPetsSwaggerPetstore(*gin.Context, *PostPetsSwaggerPetstoreParameters, *RequestBody) PostPetsSwaggerPetstoreResponse {
	return PostPetsSwaggerPetstore405Response{}
}
func (u UnimplementedPetsSwaggerPetstore) InvalidRequest(c *gin.Context, err error) {
	c.JSON(400, err.Error())
}

func RegisterPetsSwaggerPetstorePath(e *gin.Engine, srv PetsSwaggerPetstore) {

	e.GET("/pets", func(c *gin.Context) {
		params := &GetPetsSwaggerPetstoreParameters{}
		body := &RequestBody{}
		err := c.ShouldBindJSON(body)
		if err != nil {
			srv.InvalidRequest(c, err)
		}
		response := srv.GetPetsSwaggerPetstore(c, params, body)
		c.JSON(response.GetStatus(), response)
	})

	e.POST("/pets", func(c *gin.Context) {
		params := &PostPetsSwaggerPetstoreParameters{}
		body := &RequestBody{}
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

type GetPetsSwaggerPetstore200Response GetPetsSwaggerPetstoreDefaultResponse

func (GetPetsSwaggerPetstore200Response) GetStatus() int {
	return 200
}

type GetPetsSwaggerPetstoreDefaultResponse GetPetsSwaggerPetstoreDefaultResponseModel

func (GetPetsSwaggerPetstoreDefaultResponse) GetStatus() int {
	return 200
}

type GetPetsSwaggerPetstore405Response struct{}

func (GetPetsSwaggerPetstore405Response) GetStatus() int {
	return 405
}

type PostPetsSwaggerPetstoreResponse interface {
	GetStatus() int
}

type PostPetsSwaggerPetstoreDefaultResponse PostPetsSwaggerPetstoreDefaultResponseModel

func (PostPetsSwaggerPetstoreDefaultResponse) GetStatus() int {
	return 200
}

type PostPetsSwaggerPetstore405Response struct{}

func (PostPetsSwaggerPetstore405Response) GetStatus() int {
	return 405
}
