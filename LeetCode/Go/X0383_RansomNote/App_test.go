package X0383_RansomNote

import (
	"strconv"
	"testing"
)

type InputCase struct {
	note     string
	mag      string
	expected bool
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{note: "a", mag: "b", expected: false, comment: "Example1"},
		{note: "aa", mag: "ab", expected: false, comment: "Example2"},
		{note: "aa", mag: "aab", expected: true, comment: "Example3"},
		{note: "a", mag: "a", expected: true, comment: "SingleSameChar"},
		{note: "abaddbcf", mag: "ffbabadcd", expected: true, comment: "MoreComplexHappyCase"},
		{note: "abaddbcf", mag: "ffbabdcd", expected: false, comment: "MoreComplexBadCase"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.note, element.mag, element.expected, index, t)
			})
	}
}

func testApp(note string, mag string, expected bool, index int, t *testing.T) {
	actual := canConstruct(note, mag)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(actual))
	}
}
