package X0415_AddStrings

import (
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	num1     string
	num2     string
	expected string
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{num1: "11", num2: "123", expected: "134", comment: "Example1"},
		{num1: "456", num2: "77", expected: "533", comment: "Example2"},
		{num1: "0", num2: "0", expected: "0", comment: "Example3"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.num1, element.num2, element.expected, index, t)
			})
	}
}

func testApp(num1 string, num2 string, expected string, index int, t *testing.T) {
	actual := addStrings(num1, num2)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + expected +
			" Actual: " + actual)
	}
}
