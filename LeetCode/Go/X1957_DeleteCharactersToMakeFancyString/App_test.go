package X1957_DeleteCharactersToMakeFancyString

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
		{s: "leeetcode", expected: "leetcode", comment: "Example1"},
		{s: "aaabaaaa", expected: "aabaa", comment: "Example2"},
		{s: "aab", expected: "aab", comment: "Example3"},
		{s: "aaaaaaa", expected: "aa", comment: "OnlyAMultipleTimes"},
		{s: "aaaabaabbaaaa", expected: "aabaabbaa", comment: "Inners"},
		{s: "a", expected: "a", comment: "SingleChar"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.s, element.expected, index, t)
			})
	}
}

func testApp(s string, expected string, index int, t *testing.T) {
	actual := makeFancyString(s)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + expected +
			" Actual: " + actual)
	}
}
