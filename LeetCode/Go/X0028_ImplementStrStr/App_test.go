package X0028_ImplementStrStr

import (
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	haystack string
	needle   string
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{haystack: "hello", needle: "ll", expected: 2, comment: "Example1"},
		{haystack: "aaaaa", needle: "bba", expected: -1, comment: "Example2"},
		{haystack: "hello", needle: "lo", expected: 3, comment: "NeedleAtSecondAppearance"},
		{haystack: "hello", needle: "oll", expected: -1, comment: "EndOverflow"},
		{haystack: "hello", needle: "", expected: 0, comment: "NeedleEmpty"},
		{haystack: "l", needle: "ll", expected: -1, comment: "NeedleLongerThanHaystack"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.haystack, element.needle, element.expected, index, t)
			})
	}
}

func testApp(haystack string, needle string, expected int, index int, t *testing.T) {
	actual := strStr(haystack, needle)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
