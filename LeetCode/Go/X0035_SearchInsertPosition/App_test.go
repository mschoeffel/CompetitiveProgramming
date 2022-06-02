package X0035_SearchInsertPosition

import (
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	nums     []int
	target   int
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{nums: []int{1, 3, 5, 6}, target: 5, expected: 2, comment: "Example1"},
		{nums: []int{1, 3, 5, 6}, target: 2, expected: 1, comment: "Example2"},
		{nums: []int{1, 3, 5, 6}, target: 7, expected: 4, comment: "Example3"},
		{nums: []int{2, 3, 5, 6}, target: 1, expected: 0, comment: "FirstElement"},
		{nums: []int{}, target: 1, expected: 0, comment: "NumsEmpty"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.nums, element.target, element.expected, index, t)
			})
		t.Run("Test Build In: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testAppBuildIn(element.nums, element.target, element.expected, index, t)
			})
	}
}

func testApp(nums []int, target int, expected int, index int, t *testing.T) {
	actual := searchInsert(nums, target)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}

func testAppBuildIn(nums []int, target int, expected int, index int, t *testing.T) {
	actual := searchInsertBuildIn(nums, target)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
