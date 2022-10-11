package X0374_GuessNumberHigherOrLower

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
		{num: 16, expected: 6, comment: "Example1"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.num, element.expected, index, t)
			})
	}
}

func testApp(num int, expected int, index int, t *testing.T) {
	actual := guessNumber(num)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
