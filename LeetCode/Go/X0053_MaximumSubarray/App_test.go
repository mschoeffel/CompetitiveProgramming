package X0053_MaximumSubarray

import (
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	nums     []int
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, expected: 6, comment: "Example1"},
		{nums: []int{1}, expected: 1, comment: "Example2"},
		{nums: []int{5, 4, -1, 7, 8}, expected: 23, comment: "Example3"},
		{nums: []int{0}, expected: 0, comment: "NumsEmpty"},
		{nums: []int{-2, -3, -1, -4}, expected: -1, comment: "OnlyNegativeNums"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.nums, element.expected, index, t)
			})
	}
}

func testApp(nums []int, expected int, index int, t *testing.T) {
	actual := maxSubArray(nums)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
