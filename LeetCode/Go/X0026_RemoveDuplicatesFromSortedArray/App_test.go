package X0026_RemoveDuplicatesFromSortedArray

import (
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	nums          []int
	expected      int
	expectedArray []int
	comment       string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{nums: []int{1, 1, 2}, expected: 2, expectedArray: []int{1, 2, 2}, comment: "Example1"},
		{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, expected: 5, expectedArray: []int{0, 1, 2, 3, 4, 2, 2, 3, 3, 4}, comment: "Example2"},
		{nums: []int{0}, expected: 1, expectedArray: []int{0}, comment: "SingleElementInNums"},
		{nums: []int{}, expected: 0, expectedArray: []int{}, comment: "EmptyNums"},
		{nums: []int{0, 0, 0, 0}, expected: 1, expectedArray: []int{0, 0, 0, 0}, comment: "OnlyOneNum"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) { testApp(element.nums, element.expected, element.expectedArray, index, t) })
	}
}

func testApp(nums []int, expected int, expectedArray []int, index int, t *testing.T) {
	actual := removeDuplicates(nums)
	if !reflect.DeepEqual(actual, expected) || !reflect.DeepEqual(nums, expectedArray) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
