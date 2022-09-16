// GENERATED CODE. DO NOT EDIT

package models

type UserDeleteRequest struct {
	NAME string
}

func (req UserDeleteRequest)Valid() error {
	return nil
}



type UserDeleteResponse interface {
    mustImplementUserDeleteResponse()
}
type UserDelete204Response struct {}
func (UserDelete204Response)mustImplementUserDeleteResponse() {}

type UserDelete400Response struct {}
func (UserDelete400Response)mustImplementUserDeleteResponse() {}

type UserDelete500Response struct {}
func (UserDelete500Response)mustImplementUserDeleteResponse() {}

type UserGetRequest struct {
	NAME string
}

func (req UserGetRequest)Valid() error {
	return nil
}



type UserGetResponse interface {
    mustImplementUserGetResponse()
}
type UserGet200Response struct {}
func (UserGet200Response)mustImplementUserGetResponse() {}

type UserGet400Response struct {}
func (UserGet400Response)mustImplementUserGetResponse() {}

type UserGet500Response struct {}
func (UserGet500Response)mustImplementUserGetResponse() {}

type UserPatchRequest struct {
	NAME string
}

func (req UserPatchRequest)Valid() error {
	return nil
}



type UserPatchResponse interface {
    mustImplementUserPatchResponse()
}
type UserPatch204Response struct {}
func (UserPatch204Response)mustImplementUserPatchResponse() {}

type UserPatch400Response struct {}
func (UserPatch400Response)mustImplementUserPatchResponse() {}

type UserPatch500Response struct {}
func (UserPatch500Response)mustImplementUserPatchResponse() {}

type UserPostRequest struct {
	NAME string
}

func (req UserPostRequest)Valid() error {
	return nil
}



type UserPostResponse interface {
    mustImplementUserPostResponse()
}
type UserPost204Response struct {}
func (UserPost204Response)mustImplementUserPostResponse() {}

type UserPost400Response struct {}
func (UserPost400Response)mustImplementUserPostResponse() {}

type UserPost500Response struct {}
func (UserPost500Response)mustImplementUserPostResponse() {}

type ProjectPostRequest struct {
	NAME string
}

func (req ProjectPostRequest)Valid() error {
	return nil
}



type ProjectPostResponse interface {
    mustImplementProjectPostResponse()
}
type ProjectPost200Response struct {}
func (ProjectPost200Response)mustImplementProjectPostResponse() {}

type ProjectPost400Response struct {}
func (ProjectPost400Response)mustImplementProjectPostResponse() {}

type ProjectPost500Response struct {}
func (ProjectPost500Response)mustImplementProjectPostResponse() {}

type ProjectDeleteRequest struct {
	NAME string
}

func (req ProjectDeleteRequest)Valid() error {
	return nil
}



type ProjectDeleteResponse interface {
    mustImplementProjectDeleteResponse()
}
type ProjectDelete200Response struct {}
func (ProjectDelete200Response)mustImplementProjectDeleteResponse() {}

type ProjectDelete400Response struct {}
func (ProjectDelete400Response)mustImplementProjectDeleteResponse() {}

type ProjectDelete500Response struct {}
func (ProjectDelete500Response)mustImplementProjectDeleteResponse() {}

type ProjectGetRequest struct {
	NAME string
}

func (req ProjectGetRequest)Valid() error {
	return nil
}



type ProjectGetResponse interface {
    mustImplementProjectGetResponse()
}
type ProjectGet200Response struct {}
func (ProjectGet200Response)mustImplementProjectGetResponse() {}

type ProjectGet400Response struct {}
func (ProjectGet400Response)mustImplementProjectGetResponse() {}

type ProjectGet500Response struct {}
func (ProjectGet500Response)mustImplementProjectGetResponse() {}

type ProjectPatchRequest struct {
	NAME string
}

func (req ProjectPatchRequest)Valid() error {
	return nil
}



type ProjectPatchResponse interface {
    mustImplementProjectPatchResponse()
}
type ProjectPatch200Response struct {}
func (ProjectPatch200Response)mustImplementProjectPatchResponse() {}

type ProjectPatch400Response struct {}
func (ProjectPatch400Response)mustImplementProjectPatchResponse() {}

type ProjectPatch500Response struct {}
func (ProjectPatch500Response)mustImplementProjectPatchResponse() {}




