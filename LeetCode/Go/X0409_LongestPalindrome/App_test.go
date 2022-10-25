package X0409_LongestPalindrome

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
		{s: "abccccdd", expected: 7, comment: "Example1"},
		{s: "a", expected: 1, comment: "Example2"},
		{s: "aAbbcCdedd", expected: 5, comment: "UpperLowerCase"},
		{s: "ccc", expected: 3, comment: "SingleThreeChars"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.s, element.expected, index, t)
			})
	}
}

func testApp(s string, expected int, index int, t *testing.T) {
	actual := longestPalindrome(s)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
