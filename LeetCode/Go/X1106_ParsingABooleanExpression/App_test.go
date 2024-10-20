package X1106_ParsingABooleanExpression

import (
	"strconv"
	"testing"
)

type InputCase struct {
	expression string
	expected   bool
	comment    string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{expression: "&(|(f))", expected: false, comment: "Example1"},
		{expression: "|(f,f,f,t)", expected: true, comment: "Example2"},
		{expression: "!(&(f,t))", expected: true, comment: "Example3"},
		{expression: "f", expected: false, comment: "False"},
		{expression: "t", expected: true, comment: "True"},
		{expression: "!(t)", expected: false, comment: "InverseTrue"},
		{expression: "!(f)", expected: true, comment: "InverseFalse"},
		{expression: "&(f, f, f)", expected: false, comment: "AndFalseAll"},
		{expression: "&(t, f, t)", expected: false, comment: "AndFalseOne"},
		{expression: "&(t)", expected: true, comment: "AndTrue"},
		{expression: "|(f, f, f)", expected: false, comment: "OrFalse"},
		{expression: "|(f, t, f)", expected: true, comment: "OrTrue"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.expression, element.expected, index, t)
			})
	}
}

func testApp(expression string, expected bool, index int, t *testing.T) {
	actual := parseBoolExpr(expression)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(actual))
	}
}
