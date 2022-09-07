package X0283_MoveZeroes

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
		{nums: []int{0, 1, 0, 3, 12}, expected: []int{1, 3, 12, 0, 0}, comment: "Example1"},
		{nums: []int{0}, expected: []int{0}, comment: "Example2"},
		{nums: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}, comment: "NoZeroes"},
		{nums: []int{1, 2, 0, 0, 5}, expected: []int{1, 2, 5, 0, 0}, comment: "MultipleZeroesInRow"},
		{nums: []int{1}, expected: []int{1}, comment: "SingleNumber"},
		{nums: []int{0, 0, 0, 0, 0}, expected: []int{0, 0, 0, 0, 0}, comment: "OnlyZeroes"},
		{nums: []int{0, 0, 4, 0, 5}, expected: []int{4, 5, 0, 0, 0}, comment: "MultipleLeadingZeroesInRow"},
		{nums: []int{2, 0, 4, 0, 0}, expected: []int{2, 4, 0, 0, 0}, comment: "MultipleEndingZeroesInRow"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.nums, element.expected, index, t)
			})
	}
}

func testApp(nums []int, expected []int, index int, t *testing.T) {
	moveZeroes(nums)
	if !reflect.DeepEqual(nums, expected) {
		actualString, _ := json.Marshal(nums)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
