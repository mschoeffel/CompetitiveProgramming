package X0066_PlusOne

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	digits   []int
	expected []int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{digits: []int{1, 2, 3}, expected: []int{1, 2, 4}, comment: "Example1"},
		{digits: []int{4, 3, 2, 1}, expected: []int{4, 3, 2, 2}, comment: "Example2"},
		{digits: []int{9}, expected: []int{1, 0}, comment: "Example3"},
		{digits: []int{9, 9, 9, 9}, expected: []int{1, 0, 0, 0, 0}, comment: "OnlyOverflow"},
		{digits: []int{6, 7, 8, 9}, expected: []int{6, 7, 9, 0}, comment: "OneOverflow"},
		{digits: []int{0}, expected: []int{1}, comment: "SingleNullDigit"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.digits, element.expected, index, t)
			})
	}
}

func testApp(digits []int, expected []int, index int, t *testing.T) {
	actual := plusOne(digits)
	if !reflect.DeepEqual(actual, expected) {
		actualString, _ := json.Marshal(actual)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
