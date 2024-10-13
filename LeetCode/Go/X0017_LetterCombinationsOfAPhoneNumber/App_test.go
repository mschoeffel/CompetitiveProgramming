package X0017_LetterCombinationsOfAPhoneNumber

import (
	"encoding/json"
	"strconv"
	"testing"
)

type InputCase struct {
	num      string
	expected []string
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{num: "23", expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, comment: "Example1"},
		{num: "", expected: []string{}, comment: "Example2"},
		{num: "2", expected: []string{"a", "b", "c"}, comment: "Example3"},
		{num: "3", expected: []string{"d", "e", "f"}, comment: "3"},
		{num: "4", expected: []string{"g", "h", "i"}, comment: "4"},
		{num: "5", expected: []string{"j", "k", "l"}, comment: "5"},
		{num: "6", expected: []string{"m", "n", "o"}, comment: "6"},
		{num: "7", expected: []string{"p", "q", "r", "s"}, comment: "7"},
		{num: "8", expected: []string{"t", "u", "v"}, comment: "8"},
		{num: "9", expected: []string{"w", "x", "y", "z"}, comment: "9"},
		{num: "234", expected: []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}, comment: "ThreeDigits"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.num, element.expected, index, t)
			})
	}
}

func testApp(num string, expected []string, index int, t *testing.T) {
	actual := letterCombinations(num)
	if !containsAllExpected(actual, expected) {
		expectedJson, _ := json.Marshal(expected)
		actualJson, _ := json.Marshal(actual)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedJson) +
			" Actual: " + string(actualJson))
	}
}

func containsAllExpected(actual []string, expected []string) bool {
	for _, expectedElement := range expected {
		contains := false
		for _, actualElement := range actual {
			if expectedElement == actualElement {
				contains = true
			}
		}
		if !contains {
			return false
		}
	}
	return true
}
