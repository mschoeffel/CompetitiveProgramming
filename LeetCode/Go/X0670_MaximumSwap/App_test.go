package X0670_MaximumSwap

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
		{num: 2736, expected: 7236, comment: "Example1"},
		{num: 9973, expected: 9973, comment: "Example2"},
		{num: 98368, expected: 98863, comment: "Example3"},
		{num: 8888, expected: 8888, comment: "OnlyBiggest"},
		{num: 1234, expected: 4231, comment: "Ascending"},
		{num: 4321, expected: 4321, comment: "Descending"},
		{num: 4324, expected: 4423, comment: "BiggestAtStartAndEnd"},
		{num: 434224, expected: 444223, comment: "SwapMiddle"},
		{num: 1, expected: 1, comment: "SingleDigit"},
		{num: 99998368, expected: 99998863, comment: "MultipleLeading"},
		{num: 99887763262, expected: 99887766232, comment: "SteppingDown"},
		{num: 99298368, expected: 99928368, comment: "Combination"},
		{num: 1993, expected: 9913, comment: "LeadingLow"},
		{num: 100, expected: 100, comment: "Hundred"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.num, element.expected, index, t)
			})
	}
}

func testApp(num int, expected int, index int, t *testing.T) {
	actual := maximumSwap(num)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
