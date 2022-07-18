package X0190_ReverseBits

import (
	"strconv"
	"testing"
)

type InputCase struct {
	num      uint32
	expected uint32
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{num: 43261596, expected: 964176192, comment: "Example1"},
		{num: 4294967293, expected: 3221225471, comment: "Example2"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.num, element.expected, index, t)
			})
	}
}

func testApp(num uint32, expected uint32, index int, t *testing.T) {
	actual := reverseBits(num)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatUint(uint64(expected), 10) +
			" Actual: " + strconv.FormatUint(uint64(actual), 10))
	}
}
