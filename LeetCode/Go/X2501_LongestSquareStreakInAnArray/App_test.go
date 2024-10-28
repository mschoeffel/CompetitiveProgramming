package X2501_LongestSquareStreakInAnArray

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
		{nums: []int{4, 3, 6, 16, 8, 2}, expected: 3, comment: "Example1"},
		{nums: []int{2, 3, 5, 6, 7}, expected: -1, comment: "Example2"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.nums, element.expected, index, t)
			})
	}
}

func testApp(nums []int, expected int, index int, t *testing.T) {
	actual := longestSquareStreak(nums)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
