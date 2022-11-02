package X0434_NumberOfSegmentsInAString

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
		{s: "Hello, my name is John", expected: 5, comment: "Example1"},
		{s: "Hello", expected: 1, comment: "Example2"},
		{s: "", expected: 0, comment: "Empty"},
		{s: "     ", expected: 0, comment: "OnlySpaces"},
		{s: "TEST     abc", expected: 2, comment: "MultipleSpaces"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.s, element.expected, index, t)
			})
	}
}

func testApp(s string, expected int, index int, t *testing.T) {
	actual := countSegments(s)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
