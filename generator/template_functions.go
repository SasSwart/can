package generator

import (
	"fmt"
	"strings"
)

func pathNameToGinPathName(pathName string) string {
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
