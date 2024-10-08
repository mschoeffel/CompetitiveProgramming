package X0003_LongestSubstringWithoutRepeatingCharacters

import (
	"strconv"
	"testing"
)

type InputCase struct {
	s        string
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{s: "abcabcbb", expected: 3, comment: "Example1"},
		{s: "bbbbb", expected: 1, comment: "Example2"},
		{s: "pwwkew", expected: 3, comment: "Example3"},
		{s: "", expected: 0, comment: "EmptyString"},
		{s: "abababcd", expected: 4, comment: "RepeatedUntilEnd"},
		{s: "abcdabab", expected: 4, comment: "RepeatedAfterStart"},
		{s: "abababcdabab", expected: 4, comment: "RepeatedAtStartAndEnd"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.s, element.expected, index, t)
			})
	}
}

func testApp(s string, expected int, index int, t *testing.T) {
	actual := lengthOfLongestSubstring(s)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
