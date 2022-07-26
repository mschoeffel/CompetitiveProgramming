package X0217_ContainsDuplicate

import (
	"strconv"
	"testing"
)

type InputCase struct {
	input    []int
	expected bool
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{input: []int{1, 2, 3, 1}, expected: true, comment: "Example1"},
		{input: []int{1, 2, 3, 4}, expected: false, comment: "Example2"},
		{input: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, expected: true, comment: "Example3"},
		{input: []int{1}, expected: false, comment: "SingleElement"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.input, element.expected, index, t)
			})
	}
}

func testApp(input []int, expected bool, index int, t *testing.T) {
	result := containsDuplicate(input)
	if result != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(result))
	}
}
