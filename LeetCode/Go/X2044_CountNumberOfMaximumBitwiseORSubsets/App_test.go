package X2044_CountNumberOfMaximumBitwiseORSubsets

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
		{num: []int{3, 1}, expected: 2, comment: "Example1"},
		{num: []int{2, 2, 2}, expected: 7, comment: "Example2"},
		{num: []int{3, 2, 1, 5}, expected: 6, comment: "Example3"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.num, element.expected, index, t)
			})
	}
}

func testApp(num []int, expected int, index int, t *testing.T) {
	actual := countMaxOrSubsets(num)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
