package X0389_FindTheDifference

import (
	"strconv"
	"testing"
)

type InputCase struct {
	s1       string
	s2       string
	expected byte
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{s1: "abcd", s2: "abcde", expected: byte('e'), comment: "Example1"},
		{s1: "", s2: "y", expected: byte('y'), comment: "Example2"},
		{s1: "abbcdccaaf", s2: "acbcfccabad", expected: byte('c'), comment: "MoreComplex"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.s1, element.s2, element.expected, index, t)
			})
	}
}

func testApp(s1 string, s2 string, expected byte, index int, t *testing.T) {
	actual := findTheDifference(s1, s2)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expected) +
			" Actual: " + string(actual))
	}
}
