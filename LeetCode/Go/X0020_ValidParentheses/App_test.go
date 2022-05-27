package X0020_ValidParentheses

import (
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	parentheses string
	expected    bool
	comment     string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{parentheses: "()", expected: true, comment: "Example1"},
		{parentheses: "()[]{}", expected: true, comment: "Example2"},
		{parentheses: "(]", expected: false, comment: "Example3"},
		{parentheses: "({()[{}]})([]{})", expected: true, comment: "NestedValid"},
		{parentheses: "({()[{]}]})()", expected: false, comment: "NestedInvalid"},
		{parentheses: "(", expected: false, comment: "OneOpenLeft"},
		{parentheses: "}", expected: false, comment: "OneCloseLeft"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) { testApp(element.parentheses, element.expected, index, t) })
	}
}

func testApp(parentheses string, expected bool, index int, t *testing.T) {
	actual := isValid(parentheses)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(actual))
	}
}
