package X0263_UglyNumber

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
		{num: 6, expected: true, comment: "Example1"},
		{num: 1, expected: true, comment: "Example2"},
		{num: 14, expected: false, comment: "Example3"},
		{num: 0, expected: false, comment: "Zero"},
		{num: -1, expected: false, comment: "Negative"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.num, element.expected, index, t)
			})
	}
}

func testApp(num int, expected bool, index int, t *testing.T) {
	result := isUgly(num)
	if result != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(result))
	}
}
