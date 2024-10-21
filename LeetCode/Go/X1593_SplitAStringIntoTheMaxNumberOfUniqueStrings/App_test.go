package X1593_SplitAStringIntoTheMaxNumberOfUniqueStrings

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
		{s: "ababccc", expected: 5, comment: "Example1"},
		{s: "aba", expected: 2, comment: "Example2"},
		{s: "aa", expected: 1, comment: "Example3"},
		{s: "a", expected: 1, comment: "SingleChar"},
		{s: "aaaa", expected: 2, comment: "OneCharRepeated"},
		{s: "aabcbaa", expected: 5, comment: "Palindrome"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.s, element.expected, index, t)
			})
	}
}

func testApp(s string, expected int, index int, t *testing.T) {
	actual := maxUniqueSplit(s)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
