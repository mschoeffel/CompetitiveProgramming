package X0290_WordPattern

import (
	"strconv"
	"testing"
)

type InputCase struct {
	pattern  string
	s        string
	expected bool
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{pattern: "abba", s: "dog cat cat dog", expected: true, comment: "Example1"},
		{pattern: "abba", s: "dog cat cat fish", expected: false, comment: "Example2"},
		{pattern: "aaaa", s: "dog cat cat dog", expected: false, comment: "Example3"},
		{pattern: "a", s: "dog", expected: true, comment: "SingleWord"},
		{pattern: "a", s: "dog cat", expected: false, comment: "DifferentLengthString"},
		{pattern: "ab", s: "dog", expected: false, comment: "DifferentLengthPattern"},
		{pattern: "abab", s: "dog cat dogg cat", expected: false, comment: "OneCharDifferent"},
		{pattern: "abab", s: "dog cat dogg cat", expected: false, comment: "OneCharDifferent"},
		{pattern: "abba", s: "dog dog dog dog", expected: false, comment: "PatternStringDuplication"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.pattern, element.s, element.expected, index, t)
			})
	}
}

func testApp(pattern string, s string, expected bool, index int, t *testing.T) {
	actual := wordPattern(pattern, s)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(actual))
	}
}
