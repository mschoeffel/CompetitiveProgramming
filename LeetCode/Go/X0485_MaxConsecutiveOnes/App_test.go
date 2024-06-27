package X0485_MaxConsecutiveOnes

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
		{nums: []int{1, 1, 0, 1, 1, 1}, expected: 3, comment: "Example1"},
		{nums: []int{1, 0, 1, 1, 0, 1}, expected: 2, comment: "Example2"},
		{nums: []int{0, 0, 0, 0, 0, 0, 0, 0}, expected: 0, comment: "OnlyZeros"},
		{nums: []int{1, 1, 1, 1, 1, 1, 1}, expected: 7, comment: "OnlyOnes"},
		{nums: []int{1, 1, 0, 1, 0, 1, 1, 1}, expected: 3, comment: "MaxAtEnd"},
		{nums: []int{1, 1, 1, 0, 1, 1, 0, 1}, expected: 3, comment: "MaxAtStart"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.nums, element.expected, index, t)
			})
	}
}

func testApp(nums []int, expected int, index int, t *testing.T) {
	actual := findMaxConsecutiveOnes(nums)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
