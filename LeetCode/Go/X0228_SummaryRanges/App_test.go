package X0228_SummaryRanges

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	nums     []int
	expected []string
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{nums: []int{0, 1, 2, 4, 5, 7}, expected: []string{"0->2", "4->5", "7"}, comment: "Example1"},
		{nums: []int{0, 2, 3, 4, 6, 8, 9}, expected: []string{"0", "2->4", "6", "8->9"}, comment: "Example2"},
		{nums: []int{}, expected: []string{}, comment: "EmptyInput"},
		{nums: []int{1}, expected: []string{"1"}, comment: "SingleNumberInput"},
		{nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, expected: []string{"0->9"}, comment: "SingleRangeInput"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.nums, element.expected, index, t)
			})
	}
}

func testApp(root []int, expected []string, index int, t *testing.T) {
	result := summaryRanges(root)
	if !reflect.DeepEqual(result, expected) {
		actualString, _ := json.Marshal(result)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
