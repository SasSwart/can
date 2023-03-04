package render_test

import (
	"github.com/sasswart/gin-in-a-can/config"
	"github.com/sasswart/gin-in-a-can/openapi"
	"github.com/sasswart/gin-in-a-can/openapi/media"
	"github.com/sasswart/gin-in-a-can/openapi/operation"
	"github.com/sasswart/gin-in-a-can/openapi/path"
	"github.com/sasswart/gin-in-a-can/openapi/request"
	"github.com/sasswart/gin-in-a-can/openapi/response"
	"github.com/sasswart/gin-in-a-can/openapi/schema"
	"github.com/sasswart/gin-in-a-can/render"
	golang "github.com/sasswart/gin-in-a-can/render/go"
	"github.com/sasswart/gin-in-a-can/test"
	"github.com/sasswart/gin-in-a-can/tree"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/http"
	"os"
	"strings"
	"testing"
	"text/template"
)

var md = tree.Metadata{"package": "testPackage", "some": "metadata"}

func buildTestSpec() *openapi.OpenAPI {
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
	notNil := test.Dig(&root, "/endpoint", http.MethodGet, "204", media.JSONKey, schema.PropertyKey)
	if notNil == nil {
		panic("BORKED")
	}

	return &root
}

func Test_Render_Render(t *testing.T) {
	tempFolder, _ := os.MkdirTemp(os.TempDir(), "CanTestArtifacts")
	defer os.RemoveAll(tempFolder)

	cfg := newTestConfig()
	err := cfg.Load()
	if err != nil {
		t.Errorf(err.Error())
	}
	cfg.OutputPath = tempFolder
	e := render.Engine{}.With(&golang.Renderer{Base: &render.Base{}}, cfg)
	r := *e.GetRenderer()
	r.SetTemplateFuncMap(&template.FuncMap{
		"ToUpper": strings.ToUpper,
		"ToTitle": func(s string) string {
			caser := cases.Title(language.English)
			return caser.String(s)
		},
		"SanitiseName": r.SanitiseName,
		"SanitiseType": r.SanitiseType,
	})
	if r.GetTemplateFuncMap() == nil {
		t.Errorf("TemplateFuncMap should NOT be nil")
	}
	_, err = tree.Traverse(buildTestSpec(), e.BuildRenderNode())
	if err != nil {
		t.Errorf(err.Error())
	}
	// TODO What dictates whether or not the schema type is rendered to file?
	// 	We're not currently rendering any 204 response schemas
}
func newTestConfig() config.Data {
	config.ConfigFilePath = "../config/config_test.yml"
	config.Debug = true
	return config.Data{
		Template: config.Template{
			Name: "go-gin",
		},
		TemplatesDir: "../templates",
		OpenAPIFile:  "../openapi/test/fixtures/validation_no_refs.yaml",
		OutputPath:   ".",
	}
}
