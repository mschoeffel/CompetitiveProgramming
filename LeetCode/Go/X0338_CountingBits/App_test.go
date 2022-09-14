package X0338_CountingBits

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	num      int
	expected []int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{num: 2, expected: []int{0, 1, 1}, comment: "Example1"},
		{num: 5, expected: []int{0, 1, 1, 2, 1, 2}, comment: "Example2"},
		{num: 1, expected: []int{0, 1}, comment: "Single"},
		{num: 0, expected: []int{0}, comment: "Zero"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.num, element.expected, index, t)
			})
	}
}

func testApp(num int, expected []int, index int, t *testing.T) {
	actual := countBits(num)
	if !reflect.DeepEqual(actual, expected) {
		actualString, _ := json.Marshal(actual)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
