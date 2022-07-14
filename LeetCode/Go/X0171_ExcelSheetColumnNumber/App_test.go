package X0171_ExcelSheetColumnNumber

import (
	"strconv"
	"testing"
)

type InputCase struct {
	column   string
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{column: "A", expected: 1, comment: "Example1"},
		{column: "AB", expected: 28, comment: "Example2"},
		{column: "ZY", expected: 701, comment: "Example3"},
		{column: "ABA", expected: 729, comment: "ThreeRunes"},
		{column: "Z", expected: 26, comment: "Z"},
		{column: "AZ", expected: 52, comment: "LastZ"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.column, element.expected, index, t)
			})
	}
}

func testApp(column string, expected int, index int, t *testing.T) {
	actual := titleToNumber(column)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
