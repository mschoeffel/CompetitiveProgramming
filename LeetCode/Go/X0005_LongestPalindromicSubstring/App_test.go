package X0005_LongestPalindromicSubstring

import (
	"strconv"
	"testing"
)

type InputCase struct {
	s        string
	expected string
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{s: "babad", expected: "aba", comment: "Example1"},
		{s: "cbbd", expected: "bb", comment: "Example2"},
		{s: "abc", expected: "c", comment: "SinglePalindrome"},
		{s: "a", expected: "a", comment: "OneLetter"},
		{s: "abbc", expected: "bb", comment: "TwoLetters"},
		{s: "abbbc", expected: "bbb", comment: "ThreeLetters"},
		{s: "abbbbc", expected: "bbbb", comment: "FourLetters"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.s, element.expected, index, t)
			})
	}
}

func testApp(s string, expected string, index int, t *testing.T) {
	actual := longestPalindrome(s)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + expected +
			" Actual: " + actual)
	}
}
