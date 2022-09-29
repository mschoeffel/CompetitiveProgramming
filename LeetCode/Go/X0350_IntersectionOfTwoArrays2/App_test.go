package X0350_IntersectionOfTwoArrays2

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
		{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}, expected: []int{2, 2}, comment: "Example1"},
		{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}, expected: []int{9, 4}, comment: "Example2"},
		{nums1: []int{1, 2, 3}, nums2: []int{4, 5, 6}, expected: nil, comment: "NoIntersection"},
		{nums1: []int{}, nums2: []int{4, 5, 6}, expected: nil, comment: "Num1Empty"},
		{nums1: []int{1, 2, 3}, nums2: []int{}, expected: nil, comment: "Num2Empty"},
		{nums1: []int{4, 9, 5, 9, 2, 1, 1, 4, 1}, nums2: []int{9, 4, 9, 8, 4, 1, 1, 4, 1}, expected: []int{9, 4, 9, 4, 1, 1, 1}, comment: "MorComplex"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.nums1, element.nums2, element.expected, index, t)
			})
	}
}

func testApp(nums1 []int, nums2 []int, expected []int, index int, t *testing.T) {
	actual := intersect(nums1, nums2)
	if !reflect.DeepEqual(actual, expected) {
		actualString, _ := json.Marshal(actual)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
