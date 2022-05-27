package X0013_RomanToInteger

import (
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	roman    string
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{roman: "III", expected: 3, comment: "Example1"},
		{roman: "LVIII", expected: 58, comment: "Example2"},
		{roman: "MCMXCIV", expected: 1994, comment: "Example3"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) { testApp(element.roman, element.expected, index, t) })
	}
}

func testApp(roman string, expected int, index int, t *testing.T) {
	actual := romanToInt(roman)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
