package X0191_NumberOf1Bits

import (
	"strconv"
	"testing"
)

type InputCase struct {
	num      uint32
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{num: 11, expected: 3, comment: "Example1"},
		{num: 128, expected: 1, comment: "Example2"},
		{num: 4294967293, expected: 31, comment: "Example3"},
		{num: 0, expected: 0, comment: "Zero"},
		{num: 1, expected: 1, comment: "One"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.num, element.expected, index, t)
			})
	}
}

func testApp(num uint32, expected int, index int, t *testing.T) {
	actual := hammingWeight(num)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatUint(uint64(expected), 10) +
			" Actual: " + strconv.FormatUint(uint64(actual), 10))
	}
}
