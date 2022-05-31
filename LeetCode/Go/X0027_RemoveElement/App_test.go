package X0027_RemoveElement

import (
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	nums          []int
	val           int
	expected      int
	expectedArray []int
	comment       string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{nums: []int{3, 2, 2, 3}, val: 3, expected: 2, expectedArray: []int{2, 2, 2, 3}, comment: "Example1"},
		{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, expected: 5, expectedArray: []int{0, 1, 3, 0, 4, 0, 4, 2}, comment: "Example2"},
		{nums: []int{0}, val: 0, expected: 0, expectedArray: []int{0}, comment: "RemoveSingleElement"},
		{nums: []int{1, 2, 3, 4, 5}, val: 6, expected: 5, expectedArray: []int{1, 2, 3, 4, 5}, comment: "RemoveNone"},
		{nums: []int{}, val: 0, expected: 0, expectedArray: []int{}, comment: "EmptyNums"},
		{nums: []int{0, 0, 0, 0}, val: 0, expected: 0, expectedArray: []int{0, 0, 0, 0}, comment: "OnlyOneNum"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.nums, element.val, element.expected, element.expectedArray, index, t)
			})
	}
}

func testApp(nums []int, val int, expected int, expectedArray []int, index int, t *testing.T) {
	actual := removeElement(nums, val)
	if !reflect.DeepEqual(actual, expected) || !reflect.DeepEqual(nums, expectedArray) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
