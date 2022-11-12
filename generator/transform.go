package generator

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// funcName splits up the path, removes illegal characters (for go generation), and formats it to camelCase for use in function names
func funcName(pathName string) string {
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

func ginPathName(pathName string) string {
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
