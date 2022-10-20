package X0405_ConvertANumberToHexadecimal

import (
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	num      int
	expected string
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{num: 26, expected: "1a", comment: "Example1"},
		{num: -1, expected: "ffffffff", comment: "Example2"},
		{num: 0, expected: "0", comment: "Zero"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.num, element.expected, index, t)
			})
	}
}

func testApp(num int, expected string, index int, t *testing.T) {
	actual := toHex(num)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + expected +
			" Actual: " + actual)
	}
}
