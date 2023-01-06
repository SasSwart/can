package api

import "github.com/gin-gonic/gin"

// GENERATED INTERFACE. DO NOT EDIT

type SwaggerPetstorePetsPetid interface {
  SwaggerPetstorePetsPetidGet(*gin.Context, *SwaggerPetstorePetsPetidGetParameters, *SwaggerPetstorePetsPetidGetRequestBody) SwaggerPetstorePetsPetidGetResponse
	InvalidRequest(*gin.Context, error)
}

type UnimplementedSwaggerPetstorePetsPetid struct {}
func (u UnimplementedSwaggerPetstorePetsPetid) SwaggerPetstorePetsPetidGet(*gin.Context, *SwaggerPetstorePetsPetidGetParameters, *SwaggerPetstorePetsPetidGetRequestBody) SwaggerPetstorePetsPetidGetResponse {
	return SwaggerPetstorePetsPetidGet405Response{}
}
func (u UnimplementedSwaggerPetstorePetsPetid) InvalidRequest(c *gin.Context, err error) {
	c.JSON(400, err.Error())
}

func RegisterSwaggerPetstorePetsPetidPath(e *gin.Engine, srv SwaggerPetstorePetsPetid) {
	
  e.GET("/pets/{petId}", func(c *gin.Context) {
  	params := &SwaggerPetstorePetsPetidGetParameters{}
  	body := &SwaggerPetstorePetsPetidGetRequestBody{}
		err := c.ShouldBindJSON(body)
		if err != nil {
			srv.InvalidRequest(c, err)
		}
    response := srv.SwaggerPetstorePetsPetidGet(c, params, body)
    c.JSON(response.GetStatus(), response)
  })
}
type SwaggerPetstorePetsPetidGetResponse interface {
	GetStatus() int
}

type SwaggerPetstorePetsPetidGet200Response SwaggerPetstorePetsPetidGet200ResponseModel

func (SwaggerPetstorePetsPetidGet200Response) GetStatus() int {
	return 200
}

type SwaggerPetstorePetsPetidGet500Response SwaggerPetstorePetsPetidGet500ResponseModel

func (SwaggerPetstorePetsPetidGet500Response) GetStatus() int {
	return 500
}
type SwaggerPetstorePetsPetidGet405Response struct {}
func (SwaggerPetstorePetsPetidGet405Response) GetStatus() int {
	return 405
}
