package X0492_ConstructTheRectangle

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	area     int
	expected []int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{area: 4, expected: []int{2, 2}, comment: "Example1"},
		{area: 37, expected: []int{37, 1}, comment: "Example2"},
		{area: 122122, expected: []int{427, 286}, comment: "Example3"},
		{area: 1, expected: []int{1, 0}, comment: "One"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.area, element.expected, index, t)
			})
	}
}

func testApp(area int, expected []int, index int, t *testing.T) {
	actual := constructRectangle(area)
	if !reflect.DeepEqual(actual, expected) {
		expectedString, _ := json.Marshal(expected)
		actualString, _ := json.Marshal(actual)

		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
