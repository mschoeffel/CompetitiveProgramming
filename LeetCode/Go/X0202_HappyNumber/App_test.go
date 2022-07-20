package X0202_HappyNumber

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
		{num: 19, expected: true, comment: "Example1"},
		{num: 2, expected: false, comment: "Example2"},
		{num: 1, expected: true, comment: "One"},
		{num: 11, expected: false, comment: "DoubleOne"},
		{num: 7, expected: true, comment: "Seven"},
		{num: 1111111, expected: true, comment: "SevenWithOnes"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.num, element.expected, index, t)
			})
	}
}

func testApp(num int, expected bool, index int, t *testing.T) {
	actual := isHappy(num)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(actual))
	}
}
