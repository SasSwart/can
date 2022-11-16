package sanitizer

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

var supportedLangs = []string{
	"go",
	"gin",
}

type Sanitizer struct {
	language string
}

func NewSanitizer(language string) (*Sanitizer, error) {
	for _, lang := range supportedLangs {
		if strings.ToLower(language) == lang {
			return &Sanitizer{
				language: language,
			}, nil
		}
	}
	return nil, fmt.Errorf("language not supported: %v", language)
}

// Sanitize removes colons and dashes from a string
func (s Sanitizer) Sanitize(str string) string {
	switch s.language {
	case "go", "gin":
		noColons := strings.ReplaceAll(str, ":", "_")
		noDashes := strings.ReplaceAll(noColons, "-", "_")
		return noDashes
	}
	return ""
}

// FuncName splits up the path, removes illegal characters for the generation of go code and formats it to camelCase for use in function names
func (s Sanitizer) FuncName(pathName string) string {
	switch s.language {
	case "go", "gin":
		caser := cases.Title(language.English)

		// Replace - with _ (- is not allowed in go func names)
		pathSegments := strings.Split(pathName, "-")
		nameSegments := make([]string, len(pathSegments))
		for i, segment := range pathSegments {
			nameSegments[i] = caser.String(segment)
		}
		pathName = strings.Join(nameSegments, "_")

		// Replace : with _ (- is not allowed in go func names)
		pathSegments = strings.Split(pathName, ":")
		nameSegments = make([]string, len(pathSegments))
		for i, segment := range pathSegments {
			nameSegments[i] = caser.String(segment)
		}
		pathName = strings.Join(nameSegments, "_")

		// Convert from '/' delimited path to Camelcase func names
		pathSegments = strings.Split(pathName, "/")
		nameSegments = make([]string, len(pathSegments))
		for i, segment := range pathSegments {
			if len(segment) == 0 {
				continue
			}
			if segment[0] == '{' {
				continue
			}

			nameSegments[i] = caser.String(segment)
		}
		return strings.Join(nameSegments, "")
	}
	return ""
}

func (s Sanitizer) GinPathName(pathName string) string {
	switch s.language {
	case "gin":
		pathSegments := strings.Split(pathName, "/")
		for i, segment := range pathSegments {
			if len(segment) == 0 {
				continue
			}
			if segment[0] == '{' {
				pathSegments[i] = fmt.Sprintf(":%s", segment[1:len(segment)-1])
			}
		}
		return strings.Join(pathSegments, "/")
	}
	return ""
}
