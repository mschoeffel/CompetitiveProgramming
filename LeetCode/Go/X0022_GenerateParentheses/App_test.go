package X0022_GenerateParentheses

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	num      int
	expected []string
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{num: 3, expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, comment: "Example1"},
		{num: 1, expected: []string{"()"}, comment: "Example2"},
		{num: 2, expected: []string{"(())", "()()"}, comment: "2"},
		{num: 4, expected: []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}, comment: "4"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) { testApp(element.num, element.expected, index, t) })
	}
}

func testApp(num int, expected []string, index int, t *testing.T) {
	actual := generateParenthesis(num)
	if !reflect.DeepEqual(actual, expected) {
		actualString, _ := json.Marshal(actual)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
