package X0461_HammingDistance

import (
	"strconv"
	"testing"
)

type InputCase struct {
	x        int
	y        int
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{x: 1, y: 4, expected: 2, comment: "Example1"},
		{x: 3, y: 1, expected: 1, comment: "Example2"},
		{x: 4, y: 1, expected: 2, comment: "OrderIsIrrelevant"},
		{x: 0, y: 0, expected: 0, comment: "Zero"},
		{x: 4, y: 14, expected: 2, comment: "BiggerDifference"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.x, element.y, element.expected, index, t)
			})
	}
}

func testApp(x int, y int, expected int, index int, t *testing.T) {
	actual := hammingDistance(x, y)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
