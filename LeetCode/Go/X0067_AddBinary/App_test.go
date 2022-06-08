package X0067_AddBinary

import (
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	a        string
	b        string
	expected string
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{a: "11", b: "1", expected: "100", comment: "Example1"},
		{a: "1010", b: "1011", expected: "10101", comment: "Example2"},
		{a: "11011", b: "1011", expected: "100110", comment: "MoreComplex"},
		{a: "11011", b: "0", expected: "11011", comment: "OneZero"},
		{a: "11111", b: "11111", expected: "111110", comment: "Overflow"},
		{a: "11111", b: "1", expected: "100000", comment: "OverflowSingle"},
		{a: "0", b: "0", expected: "0", comment: "OnlyZero"},
		{a: "1", b: "1", expected: "10", comment: "Simple"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.a, element.b, element.expected, index, t)
			})
	}
}

func testApp(a string, b string, expected string, index int, t *testing.T) {
	actual := addBinary(a, b)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + expected +
			" Actual: " + actual)
	}
}
