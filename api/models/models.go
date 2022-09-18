// GENERATED CODE. DO NOT EDIT

package models

type UserDeleteRequest struct {
	Name string
	Body UserDeleteRequestBody
}

func (r UserDeleteRequest)IsValid() error {
	return nil
}

type UserDeleteRequestBody struct {
}

type UserDeleteResponse interface {
	mustImplementUserDeleteResponse()
}


type UserDelete204Response struct {
}
func (UserDelete204Response) mustImplementUserDeleteResponse(){}

type UserDelete400Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (UserDelete400Response) mustImplementUserDeleteResponse(){}

type UserDelete500Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (UserDelete500Response) mustImplementUserDeleteResponse(){}

type UserGetRequest struct {
	Name string
	Body UserGetRequestBody
}

func (r UserGetRequest)IsValid() error {
	return nil
}

type UserGetRequestBody struct {
}

type UserGetResponse interface {
	mustImplementUserGetResponse()
}


type UserGet200Response struct {
	Enabled bool `json:"enabled"`
	Name string `json:"name"`
	Options []string `json:"options"`
	Password string `json:"password"`
}
func (UserGet200Response) mustImplementUserGetResponse(){}

type UserGet400Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (UserGet400Response) mustImplementUserGetResponse(){}

type UserGet500Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (UserGet500Response) mustImplementUserGetResponse(){}

type UserPatchRequest struct {
	Name string
	Body UserPatchRequestBody
}

func (r UserPatchRequest)IsValid() error {
	return nil
}

type UserPatchRequestBody struct {
	Enabled bool `json:"enabled"`
	Name string `json:"name"`
	Options []string `json:"options"`
	Password string `json:"password"`
}

type UserPatchResponse interface {
	mustImplementUserPatchResponse()
}


type UserPatch201Response struct {
}
func (UserPatch201Response) mustImplementUserPatchResponse(){}

type UserPatch400Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (UserPatch400Response) mustImplementUserPatchResponse(){}

type UserPatch500Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (UserPatch500Response) mustImplementUserPatchResponse(){}

type UserPostRequest struct {
	Name string
	Body UserPostRequestBody
}

func (r UserPostRequest)IsValid() error {
	return nil
}

type UserPostRequestBody struct {
	Enabled bool `json:"enabled"`
	Name string `json:"name"`
	Options []string `json:"options"`
	Password string `json:"password"`
}

type UserPostResponse interface {
	mustImplementUserPostResponse()
}


type UserPost204Response struct {
}
func (UserPost204Response) mustImplementUserPostResponse(){}

type UserPost400Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (UserPost400Response) mustImplementUserPostResponse(){}

type UserPost500Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (UserPost500Response) mustImplementUserPostResponse(){}

type ProjectDeleteRequest struct {
	Id string
	Body ProjectDeleteRequestBody
}

func (r ProjectDeleteRequest)IsValid() error {
	return nil
}

type ProjectDeleteRequestBody struct {
}

type ProjectDeleteResponse interface {
	mustImplementProjectDeleteResponse()
}


type ProjectDelete200Response struct {
}
func (ProjectDelete200Response) mustImplementProjectDeleteResponse(){}

type ProjectDelete400Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (ProjectDelete400Response) mustImplementProjectDeleteResponse(){}

type ProjectDelete500Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (ProjectDelete500Response) mustImplementProjectDeleteResponse(){}

type ProjectGetRequest struct {
	Id string
	Body ProjectGetRequestBody
}

func (r ProjectGetRequest)IsValid() error {
	return nil
}

type ProjectGetRequestBody struct {
}

type ProjectGetResponse interface {
	mustImplementProjectGetResponse()
}


type ProjectGet200Response struct {
}
func (ProjectGet200Response) mustImplementProjectGetResponse(){}

type ProjectGet400Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (ProjectGet400Response) mustImplementProjectGetResponse(){}

type ProjectGet500Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (ProjectGet500Response) mustImplementProjectGetResponse(){}

type ProjectPatchRequest struct {
	Id string
	Body ProjectPatchRequestBody
}

func (r ProjectPatchRequest)IsValid() error {
	return nil
}

type ProjectPatchRequestBody struct {
}

type ProjectPatchResponse interface {
	mustImplementProjectPatchResponse()
}


type ProjectPatch200Response struct {
}
func (ProjectPatch200Response) mustImplementProjectPatchResponse(){}

type ProjectPatch400Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (ProjectPatch400Response) mustImplementProjectPatchResponse(){}

type ProjectPatch500Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (ProjectPatch500Response) mustImplementProjectPatchResponse(){}

type ProjectPostRequest struct {
	Id string
	Body ProjectPostRequestBody
}

func (r ProjectPostRequest)IsValid() error {
	return nil
}

type ProjectPostRequestBody struct {
}

type ProjectPostResponse interface {
	mustImplementProjectPostResponse()
}


type ProjectPost200Response struct {
}
func (ProjectPost200Response) mustImplementProjectPostResponse(){}

type ProjectPost400Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (ProjectPost400Response) mustImplementProjectPostResponse(){}

type ProjectPost500Response struct {
	Code string `json:"code"`
	Error string `json:"error"`
}
func (ProjectPost500Response) mustImplementProjectPostResponse(){}

