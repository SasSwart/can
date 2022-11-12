package generator

import (
	"regexp"
	"testing"
)

func TestFuncName(t *testing.T) {
	//TODO: add more test data
	testStrings := []string{
		"",
		"/",
		"/path",
		"/path-name",
		"/path/name",
		"/path_name",
		"/path:name",
		"/path{name}",
		"/path.name",
	}
	for _, testString := range testStrings {
		res := funcName(testString)

		// test Go Compatibility
		// TODO: format this regex to go function name spec
		isGoCompat := regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(res)
		if !isGoCompat {
			t.Logf("testString %v is not go compatible", res)
			t.Fail()
		}

		//test empty string
		if testString == "" && res != "" {
			t.Log("testString was empty but result of function call was not")
			t.Fail()
		}

	}
}
