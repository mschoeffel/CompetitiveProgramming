package X0058_LengthOfLastWord

import (
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	words    string
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{words: "Hello World", expected: 5, comment: "Example1"},
		{words: "   fly me   to   the moon  ", expected: 4, comment: "Example2"},
		{words: "luffy is still joyboy", expected: 6, comment: "Example3"},
		{words: "joyboy", expected: 6, comment: "OnlyOneWord"},
		{words: "joyboy l", expected: 1, comment: "OnlyOneChar"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.words, element.expected, index, t)
			})
	}
}

func testApp(words string, expected int, index int, t *testing.T) {
	actual := lengthOfLastWord(words)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
