// GENERATED CODE. DO NOT EDIT

package api

import "github.com/gin-gonic/gin"

type Server interface {
	Pets
	PetsPetid
}

func RegisterServer(e *gin.Engine, srv Server) {
	RegisterPetsPath(e, srv)
	RegisterPetsPetidPath(e, srv)
}

type UnimplementedServer struct {
	UnimplementedPets
	UnimplementedPetsPetid
}

// GENERATED INTERFACE. DO NOT EDIT

type Pets interface {
	PetsGet(*gin.Context, *PetsGetParameters, *PetsGetRequestBody) PetsGetResponse
	PetsPost(*gin.Context, *PetsPostParameters, *PetsPostRequestBody) PetsPostResponse
}

type UnimplementedPets struct{}

func (u UnimplementedPets) PetsGet(*gin.Context, *PetsGetParameters, *PetsGetRequestBody) PetsGetResponse {
	return PetsGet405Response{}
}
func (u UnimplementedPets) PetsPost(*gin.Context, *PetsPostParameters, *PetsPostRequestBody) PetsPostResponse {
	return PetsPost405Response{}
}

func RegisterPetsPath(e *gin.Engine, srv Pets) {
	e.GET("/path", func(c *gin.Context) {
		params := &PetsGetParameters{}
		body := &PetsGetRequestBody{}
		response := srv.PetsGet(c, params, body)
		c.JSON(response.GetStatus(), response)
	})
	e.POST("/path", func(c *gin.Context) {
		params := &PetsPostParameters{}
		body := &PetsPostRequestBody{}
		response := srv.PetsPost(c, params, body)
		c.JSON(response.GetStatus(), response)
	})
}

type PetsGetResponse interface {
	GetStatus() int
}

type PetsGet200Response struct {
	PetsGet200responsemodel
}

func (PetsGet200Response) GetStatus() int {
	return 200
}

type PetsGetdefaultResponse struct {
	PetsGetdefaultresponsemodel
}

func (PetsGetdefaultResponse) GetStatus() int {
	return 405
}

type PetsGet405Response struct{}

func (PetsGet405Response) GetStatus() int {
	return 405
}

type PetsGetParameters struct{}
type PetsGetRequestBody struct{}
type PetsPostResponse interface {
	GetStatus() int
}

type PetsPostdefaultResponse struct {
	PetsPostdefaultresponsemodel
}

func (PetsPostdefaultResponse) GetStatus() int {
	return 405
}

type PetsPost405Response struct{}

func (PetsPost405Response) GetStatus() int {
	return 405
}

type PetsPostParameters struct{}
type PetsPostRequestBody struct{}

// GENERATED MODEL. DO NOT EDIT

type PetsGet200responsemodel []PetsGet200responsemodelitem

// GENERATED MODEL. DO NOT EDIT

type PetsGet200responsemodelitem struct {
	id   PetsGet200responsemodelitemid
	name PetsGet200responsemodelitemname
	tag  PetsGet200responsemodelitemtag
}

// GENERATED MODEL. DO NOT EDIT

type PetsGet200responsemodelitemname string

// GENERATED MODEL. DO NOT EDIT

type PetsGet200responsemodelitemtag string

// GENERATED MODEL. DO NOT EDIT

type PetsGet200responsemodelitemid int

// GENERATED MODEL. DO NOT EDIT

type PetsGetdefaultresponsemodel struct {
	code    PetsGetdefaultresponsemodelcode
	message PetsGetdefaultresponsemodelmessage
}

// GENERATED MODEL. DO NOT EDIT

type PetsGetdefaultresponsemodelcode int

// GENERATED MODEL. DO NOT EDIT

type PetsGetdefaultresponsemodelmessage string

// GENERATED MODEL. DO NOT EDIT

type PetsPostdefaultresponsemodel struct {
	code    PetsPostdefaultresponsemodelcode
	message PetsPostdefaultresponsemodelmessage
}

// GENERATED MODEL. DO NOT EDIT

type PetsPostdefaultresponsemodelmessage string

// GENERATED MODEL. DO NOT EDIT

type PetsPostdefaultresponsemodelcode int

// GENERATED INTERFACE. DO NOT EDIT

type PetsPetid interface {
	PetsPetidget(*gin.Context, *PetsPetidgetParameters, *PetsPetidgetRequestBody) PetsPetidgetResponse
}

type UnimplementedPetsPetid struct{}

func (u UnimplementedPetsPetid) PetsPetidget(*gin.Context, *PetsPetidgetParameters, *PetsPetidgetRequestBody) PetsPetidgetResponse {
	return PetsPetidget405Response{}
}

func RegisterPetsPetidPath(e *gin.Engine, srv PetsPetid) {
	e.GET("/path", func(c *gin.Context) {
		params := &PetsPetidgetParameters{}
		body := &PetsPetidgetRequestBody{}
		response := srv.PetsPetidget(c, params, body)
		c.JSON(response.GetStatus(), response)
	})
}

type PetsPetidgetResponse interface {
	GetStatus() int
}

type PetsPetidget200Response struct {
	Petspetidget200responsemodel
}

func (PetsPetidget200Response) GetStatus() int {
	return 200
}

type PetsPetidgetdefaultResponse struct {
	Petspetidgetdefaultresponsemodel
}

func (PetsPetidgetdefaultResponse) GetStatus() int {
	return 405
}

type PetsPetidget405Response struct{}

func (PetsPetidget405Response) GetStatus() int {
	return 405
}

type PetsPetidgetParameters struct{}
type PetsPetidgetRequestBody struct{}

// GENERATED MODEL. DO NOT EDIT

type Petspetidget200responsemodel struct {
	id   Petspetidget200responsemodelid
	name Petspetidget200responsemodelname
	tag  Petspetidget200responsemodeltag
}

// GENERATED MODEL. DO NOT EDIT

type Petspetidget200responsemodelid int

// GENERATED MODEL. DO NOT EDIT

type Petspetidget200responsemodelname string

// GENERATED MODEL. DO NOT EDIT

type Petspetidget200responsemodeltag string

// GENERATED MODEL. DO NOT EDIT

type Petspetidgetdefaultresponsemodel struct {
	code    Petspetidgetdefaultresponsemodelcode
	message Petspetidgetdefaultresponsemodelmessage
}

// GENERATED MODEL. DO NOT EDIT

type Petspetidgetdefaultresponsemodelcode int

// GENERATED MODEL. DO NOT EDIT

type Petspetidgetdefaultresponsemodelmessage string
