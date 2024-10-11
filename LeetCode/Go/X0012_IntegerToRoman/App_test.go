package X0012_IntegerToRoman

import (
	"strconv"
	"testing"
)

type InputCase struct {
	num      int
	expected string
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{num: 3749, expected: "MMMDCCXLIX", comment: "Example1"},
		{num: 58, expected: "LVIII", comment: "Example2"},
		{num: 1994, expected: "MCMXCIV", comment: "Example3"},
		{num: 1, expected: "I", comment: "SingleElement"},
		{num: 4, expected: "IV", comment: "Four"},
		{num: 9, expected: "IX", comment: "Nine"},
		{num: 40, expected: "XL", comment: "Fourty"},
		{num: 90, expected: "XC", comment: "Ninety"},
		{num: 400, expected: "CD", comment: "FourHundred"},
		{num: 900, expected: "CM", comment: "NineHundred"},
		{num: 2999, expected: "MMCMXCIX", comment: "Nines"},
		{num: 2444, expected: "MMCDXLIV", comment: "Fours"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.num, element.expected, index, t)
			})
	}
}

func testApp(num int, expected string, index int, t *testing.T) {
	actual := intToRoman(num)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + expected +
			" Actual: " + actual)
	}
}
