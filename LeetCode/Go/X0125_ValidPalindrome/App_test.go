package X0125_ValidPalindrome

import (
	"strconv"
	"testing"
)

type InputCase struct {
	input    string
	expected bool
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{input: "A man, a plan, a canal: Panama", expected: true, comment: "Example1"},
		{input: "race a car", expected: false, comment: "Example2"},
		{input: " ", expected: true, comment: "Example3"},
		{input: " ,-  ", expected: true, comment: "MultipleNonAlphaNumeric"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.input, element.expected, index, t)
			})
	}
}

func testApp(input string, expected bool, index int, t *testing.T) {
	result := isPalindrome(input)
	if result != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(result))
	}
}
