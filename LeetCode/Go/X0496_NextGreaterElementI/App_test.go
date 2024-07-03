package X0496_NextGreaterElementI

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	nums1    []int
	nums2    []int
	expected []int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{nums1: []int{4, 1, 2}, nums2: []int{1, 3, 4, 2}, expected: []int{-1, 3, -1}, comment: "Example1"},
		{nums1: []int{2, 4}, nums2: []int{1, 2, 3, 4}, expected: []int{3, -1}, comment: "Example2"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.nums1, element.nums2, element.expected, index, t)
			})
	}
}

func testApp(nums1 []int, nums2 []int, expected []int, index int, t *testing.T) {
	actual := nextGreaterElement(nums1, nums2)
	if !reflect.DeepEqual(actual, expected) {
		actualString, _ := json.Marshal(actual)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
