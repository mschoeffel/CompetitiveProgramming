package X0168_ExcelSheetColumnTitle

import (
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	column   int
	expected string
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{column: 1, expected: "A", comment: "Example1"},
		{column: 28, expected: "AB", comment: "Example2"},
		{column: 701, expected: "ZY", comment: "Example3"},
		{column: 729, expected: "ABA", comment: "ThreeRunes"},
		{column: 26, expected: "Z", comment: "Z"},
		{column: 52, expected: "AZ", comment: "LastZ"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.column, element.expected, index, t)
			})
	}
}

func testApp(column int, expected string, index int, t *testing.T) {
	actual := convertToTitle(column)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + expected +
			" Actual: " + actual)
	}
}
