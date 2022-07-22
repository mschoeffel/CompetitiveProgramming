package X0205_IsomorphicStrings

import (
	"strconv"
	"testing"
)

type InputCase struct {
	s        string
	t        string
	expected bool
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{s: "egg", t: "add", expected: true, comment: "Example1"},
		{s: "foo", t: "bar", expected: false, comment: "Example2"},
		{s: "paper", t: "title", expected: true, comment: "Example3"},
		{s: "a", t: "b", expected: true, comment: "OnlyOneChar"},
		{s: "egg", t: "egg", expected: true, comment: "SameInput"},
		{s: "badc", t: "baba", expected: false, comment: "FailedCase"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.s, element.t, element.expected, index, t)
			})
	}
}

func testApp(a string, b string, expected bool, index int, t *testing.T) {
	actual := isIsomorphic(a, b)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(actual))
	}
}
