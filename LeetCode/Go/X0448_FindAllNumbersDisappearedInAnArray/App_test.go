package X0448_FindAllNumbersDisappearedInAnArray

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	nums     []int
	expected []int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{nums: []int{4, 3, 2, 7, 8, 2, 3, 1}, expected: []int{5, 6}, comment: "Example1"},
		{nums: []int{1, 1}, expected: []int{2}, comment: "Example2"},
		{nums: []int{1, 3, 5, 4, 2}, expected: []int{}, comment: "NoneMissing"},
		{nums: []int{1}, expected: []int{}, comment: "SingleElement"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.nums, element.expected, index, t)
			})
	}
}

func testApp(nums []int, expected []int, index int, t *testing.T) {
	actual := findDisappearedNumbers(nums)
	if !reflect.DeepEqual(actual, expected) {
		actualString, _ := json.Marshal(actual)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
