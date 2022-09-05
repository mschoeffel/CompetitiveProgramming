package X0268_MissingNumber

import (
	"strconv"
	"testing"
)

type InputCase struct {
	num      []int
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{num: []int{3, 0, 1}, expected: 2, comment: "Example1"},
		{num: []int{0, 1}, expected: 2, comment: "Example2"},
		{num: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, expected: 8, comment: "Example3"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.num, element.expected, index, t)
			})
	}
}

func testApp(num []int, expected int, index int, t *testing.T) {
	result := missingNumber(num)
	if result != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(result))
	}
}
