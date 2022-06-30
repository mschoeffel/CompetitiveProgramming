package X0119_PascalsTriangle2

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	rows     int
	expected []int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{rows: 3, expected: []int{1, 3, 3, 1}, comment: "Example1"},
		{rows: 0, expected: []int{1}, comment: "Example2"},
		{rows: 1, expected: []int{1, 1}, comment: "Example3"},
		{rows: 4, expected: []int{1, 4, 6, 4, 1}, comment: "OneMore"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.rows, element.expected, index, t)
			})
	}
}

func testApp(rows int, expected []int, index int, t *testing.T) {
	result := getRow(rows)
	if !reflect.DeepEqual(result, expected) {
		actualString, _ := json.Marshal(result)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
