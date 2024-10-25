package X1233_RemoveSubFoldersFromTheFilesystem

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	root     []string
	expected []string
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{root: []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}, expected: []string{"/a", "/c/d", "/c/f"}, comment: "Example1"},
		{root: []string{"/a", "/a/b/c", "/a/b/d"}, expected: []string{"/a"}, comment: "Example2"},
		{root: []string{"/a/b/c", "/a/b/ca", "/a/b/d"}, expected: []string{"/a/b/c", "/a/b/ca", "/a/b/d"}, comment: "Example3"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.root, element.expected, index, t)
			})
	}
}

func testApp(root []string, expected []string, index int, t *testing.T) {
	actual := removeSubfolders(root)
	if !reflect.DeepEqual(actual, expected) {
		actualStr, _ := json.Marshal(actual)
		expectedStr, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedStr) +
			" Actual: " + string(actualStr))
	}
}
