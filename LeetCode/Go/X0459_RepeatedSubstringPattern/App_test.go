package X0459_RepeatedSubstringPattern

import (
	"strconv"
	"testing"
)

type InputCase struct {
	s        string
	expected bool
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{s: "abab", expected: true, comment: "Example1"},
		{s: "aba", expected: false, comment: "Example2"},
		{s: "abcabcabcabc", expected: true, comment: "Example3"},
		{s: "a", expected: false, comment: "SingleChar"},
		{s: "aabb", expected: false, comment: "MultipleTrap"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.s, element.expected, index, t)
			})
	}
}

func testApp(s string, expected bool, index int, t *testing.T) {
	actual := repeatedSubstringPattern(s)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(actual))
	}
}
