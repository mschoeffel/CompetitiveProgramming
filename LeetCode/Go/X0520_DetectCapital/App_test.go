package X0520_DetectCapital

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
		{s: "USA", expected: true, comment: "Example1"},
		{s: "FlaG", expected: false, comment: "Example2"},
		{s: "leetcode", expected: true, comment: "AllLower"},
		{s: "Leetcode", expected: true, comment: "FirstUpper"},
		{s: "LEETCODE", expected: true, comment: "AllUpper"},
		{s: "FFFFFFFFFFFFFFFFFFFFf", expected: false, comment: "LastLower"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.s, element.expected, index, t)
			})
	}
}

func testApp(s string, expected bool, index int, t *testing.T) {
	actual := detectCapitalUse(s)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(actual))
	}
}
