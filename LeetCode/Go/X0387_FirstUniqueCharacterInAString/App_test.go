package X0387_FirstUniqueCharacterInAString

import (
	"strconv"
	"testing"
)

type InputCase struct {
	s        string
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{s: "leetcode", expected: 0, comment: "Example1"},
		{s: "loveleetcode", expected: 2, comment: "Example2"},
		{s: "aabb", expected: -1, comment: "Example3"},
		{s: "aaabbb", expected: -1, comment: "Triplets"},
		{s: "a", expected: 0, comment: "SingleChar"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.s, element.expected, index, t)
			})
	}
}

func testApp(s string, expected int, index int, t *testing.T) {
	actual := firstUniqChar(s)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
