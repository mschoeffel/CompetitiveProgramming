package X0441_ArrangingCoins

import (
	"strconv"
	"testing"
)

type InputCase struct {
	n        int
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{n: 5, expected: 2, comment: "Example1"},
		{n: 8, expected: 3, comment: "Example2"},
		{n: 1, expected: 1, comment: "One"},
		{n: 2, expected: 1, comment: "Two"},
		{n: 10, expected: 4, comment: "EvenStair"},
		{n: 11, expected: 4, comment: "EvenStairPlusOne"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.n, element.expected, index, t)
			})
	}
}

func testApp(n int, expected int, index int, t *testing.T) {
	actual := arrangeCoins(n)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
