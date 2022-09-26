package X0345_ReverseVowelsOfAString

import (
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	s        string
	expected string
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{s: "hello", expected: "holle", comment: "Example1"},
		{s: "leetcode", expected: "leotcede", comment: "Example2"},
		{s: "abbabbi", expected: "ibbabba", comment: "UnevenVowels"},
		{s: "bbbbb", expected: "bbbbb", comment: "NoVowels"},
		{s: "aaeeii", expected: "iieeaa", comment: "OnlyVowels"},
		{s: "a", expected: "a", comment: "OneLength"},
		{s: "aI", expected: "Ia", comment: "TwoVowelsLength"},
		{s: "Aa", expected: "aA", comment: "TwoUpperCase"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.s, element.expected, index, t)
			})
	}
}

func testApp(s string, expected string, index int, t *testing.T) {
	actual := reverseVowels(s)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + expected +
			" Actual: " + actual)
	}
}
