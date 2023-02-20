package render

import (
	"github.com/sasswart/gin-in-a-can/tree"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"text/template"
)

type Config struct {
	ModuleName        string
	BasePackageName   string
	TemplateDirectory string
	TemplateName      string
	OutputPath        string
	WorkingDirectory  string
	ConfigFilePath    string
}
type Renderer interface {
	SanitiseName(string) string
	SanitiseType(n tree.NodeTraverser) string

	GetOutputFile(n tree.NodeTraverser) string
	//getName(n tree.NodeTraverser) string
}

// misc package functions
var templateFuncMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
	"ToTitle": toTitleCase,
}

func toTitleCase(s string) string {
	caser := cases.Title(language.English)
	return caser.String(s)
}
