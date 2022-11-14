package sanitizer

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

// GoSanitize removes colons and dashes from a string
func GoSanitize(s string) string {
	noColons := strings.ReplaceAll(s, ":", "_")
	noDashes := strings.ReplaceAll(noColons, "-", "_")
	return noDashes
}

// GoFuncName splits up the path, removes illegal characters for the generation of go code and formats it to camelCase for use in function names
func GoFuncName(pathName string) string {
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

func GinPathName(pathName string) string {
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
