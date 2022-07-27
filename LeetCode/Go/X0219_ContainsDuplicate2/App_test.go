package X0219_ContainsDuplicate2

import (
	"strconv"
	"testing"
)

type InputCase struct {
	nums     []int
	k        int
	expected bool
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{nums: []int{1, 2, 3, 1}, k: 3, expected: true, comment: "Example1"},
		{nums: []int{1, 0, 1, 1}, k: 1, expected: true, comment: "Example2"},
		{nums: []int{1, 2, 3, 1, 2, 3}, k: 2, expected: false, comment: "Example3"},
		{nums: []int{1}, k: 3, expected: false, comment: "SingleElement"},
		{nums: []int{1, 1, 3, 1}, k: 0, expected: false, comment: "KEqualsZero"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.nums, element.k, element.expected, index, t)
			})
	}
}

func testApp(nums []int, k int, expected bool, index int, t *testing.T) {
	result := containsNearbyDuplicate(nums, k)
	if result != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(result))
	}
}
