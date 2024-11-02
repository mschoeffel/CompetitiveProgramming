package X2490_CircularSentence

import (
	"strconv"
	"testing"
)

type InputCase struct {
	sentence string
	expected bool
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{sentence: "leetcode exercises sound delightful", expected: true, comment: "Example1"},
		{sentence: "eetcode", expected: true, comment: "Example2"},
		{sentence: "Leetcode is cool", expected: false, comment: "Example3"},
		{sentence: "Leetcode ExerciseS sound DelightfuL", expected: false, comment: "LowerUpperCase"},
		{sentence: "l", expected: true, comment: "SingleRune"},
		{sentence: "lL", expected: false, comment: "TwoRunesSame"},
		{sentence: "lT", expected: false, comment: "TwoRunesDifferent"},
		{sentence: "a a a a a a a a a a a a", expected: true, comment: "SingleRuneWords"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.sentence, element.expected, index, t)
			})
	}
}

func testApp(sentence string, expected bool, index int, t *testing.T) {
	actual := isCircularSentence(sentence)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(actual))
	}
}
