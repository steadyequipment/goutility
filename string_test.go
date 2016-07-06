package utility

import (
	"fmt"
	"testing"

	"os"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestStringAppendWithJoin(t *testing.T) {
	result := StringAppendWithJoin("", "", "")
	if result != "" {
		t.Fatal("Should be \"\"")
	}

	result = StringAppendWithJoin("a", "", "")
	if result != "a" {
		t.Fatal("Should be \"a\"")
	}

	result = StringAppendWithJoin("", "b", "")
	if result != "" {
		t.Fatal("Should be \"\"")
	}

	result = StringAppendWithJoin("", "", "c")
	if result != "c" {
		t.Fatal("Should be \"c\"")
	}

	result = StringAppendWithJoin("a", "b", "")
	if result != "a" {
		t.Fatal("Should be \"a\"")
	}

	result = StringAppendWithJoin("a", "", "c")
	if result != "ac" {
		t.Fatal("Should be \"ac\"")
	}

	result = StringAppendWithJoin("", "b", "c")
	if result != "c" {
		t.Fatal("Should be \"c\"")
	}

	result = StringAppendWithJoin("a", "b", "c")
	if result != "abc" {
		t.Fatal("Should be \"abc\"")
	}
}

func TestPathStrippedOfQuery(t *testing.T) {
	path := "this is a path"
	result := PathStrippedOfQuery(path)
	if result != path {
		t.Fatal("Result (\"%s\") does not match \"%s\"", result, path)
	}

	pathExtra := path + "{lalalaala}"
	result = PathStrippedOfQuery(pathExtra)
	if result != path {
		t.Fatal("Result (\"%s\") does not match \"%s\"", result, path)
	}

	pathExtraExtra := pathExtra + "{asdf}"
	result = PathStrippedOfQuery(pathExtraExtra)
	if result != path {
		t.Fatal("Result (\"%s\") does not match \"%s\"", result, path)
	}
}
