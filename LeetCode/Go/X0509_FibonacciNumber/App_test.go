package X0509_FibonacciNumber

import (
	"strconv"
	"testing"
)

type InputCase struct {
	num      int
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{num: 2, expected: 1, comment: "Example1"},
		{num: 3, expected: 2, comment: "Example2"},
		{num: 4, expected: 3, comment: "Example3"},
		{num: 0, expected: 0, comment: "Zero"},
		{num: 1, expected: 1, comment: "One"},
		{num: 12, expected: 144, comment: "Bigger"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.num, element.expected, index, t)
			})
	}
}

func testApp(num int, expected int, index int, t *testing.T) {
	actual := fib(num)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
