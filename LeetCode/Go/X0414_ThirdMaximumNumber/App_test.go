package X0414_ThirdMaximumNumber

import (
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
		{nums: []int{3, 2, 1}, expected: 1, comment: "Example1"},
		{nums: []int{1, 2}, expected: 2, comment: "Example2"},
		{nums: []int{2, 2, 3, 1}, expected: 1, comment: "Example3"},
		{nums: []int{2, 2, 1, 2, 1}, expected: 2, comment: "OnlyTwoButMultiple"},
		{nums: []int{-4, 0, 5, 2, 5, 6, 4, 7, 6}, expected: 5, comment: "MoreComplex"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.nums, element.expected, index, t)
			})
	}
}

func testApp(nums []int, expected int, index int, t *testing.T) {
	actual := thirdMax(nums)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
