package test

import (
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/media"
	"github.com/sasswart/gin-in-a-can/openapi/operation"
	"github.com/sasswart/gin-in-a-can/openapi/path"
	"github.com/sasswart/gin-in-a-can/openapi/request"
	"github.com/sasswart/gin-in-a-can/openapi/response"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/tree"
	"net/http"
)

var md = tree.Metadata{"package": "testPackage", "some": "metadata"}

func OpenAPITree() *openapi.OpenAPI {
	root := openapi.OpenAPI{
		Node: tree.Node{Name: "root"},
	}
	root.SetMetadata(md)
	p := path.Item{
		Node: tree.Node{Name: "Path Item"},
	}
	root.SetChild("/endpoint", &p)
	p.SetParent(&root)

	post := operation.Operation{
		Node: tree.Node{Name: "post"},
	}
	p.SetChild(http.MethodPost, &post)
	post.SetParent(&p)

	postRequestBody := request.Body{
		Node: tree.Node{Name: "postRequestBody"},
	}
	post.SetChild(request.BodyKey, &postRequestBody)
	postRequestBody.SetParent(&post)

	postMt := media.Type{
		Node: tree.Node{
			Name: "postMt",
		},
	}
	postRequestBody.SetChild(media.JSONKey, &postMt)
	postMt.SetParent(&postRequestBody)

	postSchema := schema.Schema{Node: tree.Node{
		Name: "postSchema",
	}}
	postMt.SetChild(schema.PropertyKey, &postSchema)
	postSchema.SetParent(&postMt)

	// ------------------//

	get := operation.Operation{
		Node: tree.Node{Name: "operation"},
	}
	p.SetChild(http.MethodGet, &get)
	get.SetParent(&p)

	getRequestBody := request.Body{
		Node: tree.Node{Name: "getRequestBody"},
	}
	get.SetChild(request.BodyKey, &getRequestBody)
	getRequestBody.SetParent(&get)

	getMt := media.Type{
		Node: tree.Node{
			Name: "getMt",
		},
	}
	getRequestBody.SetChild(media.JSONKey, &getMt)
	getMt.SetParent(&getRequestBody)

	getSchema := schema.Schema{Node: tree.Node{
		Name: "getSchema",
	}}
	getMt.SetChild(schema.PropertyKey, &getSchema)
	getSchema.SetParent(&getMt)

	getResponse200 := response.Response{
		Content: map[string]media.Type{},
		Node: tree.Node{
			Name: "getResponse200",
		},
	}
	getResponse204 := response.Response{
		Content: map[string]media.Type{},
		Node: tree.Node{
			Name: "getResponse204",
		},
	}
	getResponse200.SetParent(&get)
	getResponse204.SetParent(&get)
	get.SetChild("200", &getResponse200)
	get.SetChild("204", &getResponse204)

	get200Mt := media.Type{
		Node: tree.Node{
			Name: "get200Mt",
		},
		Schema: &schema.Schema{
			Properties:  schema.Properties{},
			Items:       &schema.Schema{},
			Description: "turn something on or off",
			Type:        "boolean",
		},
	}
	get204Mt := media.Type{
		Node: tree.Node{
			Name: "get204Mt",
		},
		Schema: &schema.Schema{
			Properties: schema.Properties{},
			Items:      &schema.Schema{},
		},
	}
	getResponse200.SetChild(media.JSONKey, &get200Mt)
	getResponse204.SetChild(media.JSONKey, &get204Mt)

	// sanity assertions
	notNil := Dig(&root, "/endpoint", http.MethodGet, "204", media.JSONKey, schema.PropertyKey)
	if notNil == nil {
		panic("BORKED")
	}

	return &root
}
