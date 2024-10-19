package X1545_FindKthBitInNthBinaryString

import (
	"strconv"
	"testing"
)

type InputCase struct {
	n        int
	k        int
	expected byte
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{n: 3, k: 1, expected: '0', comment: "Example1"},
		{n: 4, k: 11, expected: '1', comment: "Example2"},
		{n: 1, k: 1, expected: '0', comment: "Ones"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.n, element.k, element.expected, index, t)
			})
	}
}

func testApp(n int, k int, expected byte, index int, t *testing.T) {
	actual := findKthBit(n, k)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expected) +
			" Actual: " + string(actual))
	}
}
