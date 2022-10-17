package X0392_IsSubsequence

import (
	"strconv"
	"testing"
)

type InputCase struct {
	s1       string
	s2       string
	expected bool
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{s1: "abc", s2: "ahbgdc", expected: true, comment: "Example1"},
		{s1: "axc", s2: "ahbgdc", expected: false, comment: "Example2"},
		{s1: "", s2: "", expected: true, comment: "BothZero"},
		{s1: "", s2: "abcd", expected: true, comment: "FirstZero"},
		{s1: "abcd", s2: "", expected: false, comment: "SecondZero"},
		{s1: "aabcbbc", s2: "ababbcbbbaabcbbc", expected: true, comment: "DoubleChars"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.s1, element.s2, element.expected, index, t)
			})
	}
}

func testApp(s1 string, s2 string, expected bool, index int, t *testing.T) {
	actual := isSubsequence(s1, s2)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(actual))
	}
}
