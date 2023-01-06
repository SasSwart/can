package test

import "testing"

func TestFindRootFolder(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "/Users/alex/code/gin-in-a-can/render/docs/openapi.yml", expected: "/Users/alex/code/gin-in-a-can"},
		{input: "/Users/alex/code/gin-in-a-can/render", expected: "/Users/alex/code/gin-in-a-can"},
		{input: "github.com/sasswart/gin-in-a-can", expected: "github.com/sasswart/gin-in-a-can"},
	}

	for _, testCase := range tests {
		got := findRootFolder(testCase.input)
		if got != testCase.expected {
			t.Errorf("got %s, expected %s\n", got, testCase.expected)
		}
	}
}
