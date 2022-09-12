package X0292_NimGame

import (
	"strconv"
	"testing"
)

type InputCase struct {
	num      int
	expected bool
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{num: 4, expected: false, comment: "Example1"},
		{num: 1, expected: true, comment: "Example2"},
		{num: 2, expected: true, comment: "Example3"},
		{num: 3, expected: true, comment: "Three"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.num, element.expected, index, t)
			})
	}
}

func testApp(num int, expected bool, index int, t *testing.T) {
	actual := canWinNim(num)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(actual))
	}
}
