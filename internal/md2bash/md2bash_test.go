package md2bash

import (
	"reflect"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestExtractScriptsFromMde(t *testing.T) {

	text := "Example 0\n\n```\necho \"Hello world\"\n```\n\n"
	text += "Example 1\n\n```word\necho \"Hello world\"\n```\n\n"
	text += "Example 2\n\n```two words\necho \"Hello world\"\n```\n\n"
	text += "Example 3\n\n```a few unrelated words\necho \"Hello world\"\n```\n\n"
	text += "Example 4\n\n``` a-few_UnrelatEd w0rds\necho \"Hello world\"\n```\n\n"

	pup := [][]string{
		{"```\necho \"Hello world\"\n```\n", "echo \"Hello world\""},
		{"```word\necho \"Hello world\"\n```\n", "echo \"Hello world\""},
		{"```two words\necho \"Hello world\"\n```\n", "echo \"Hello world\""},
		{"```a few unrelated words\necho \"Hello world\"\n```\n", "echo \"Hello world\""},
		{"``` a-few_UnrelatEd w0rds\necho \"Hello world\"\n```\n", "echo \"Hello world\""},
	}

	matches := extractScriptsFromMd([]byte(text))

	if reflect.DeepEqual(pup, matches) == false {
		t.Fatalf("Arrays are not equal")
	}
}
