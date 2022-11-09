package X0455_AssignCookies

import (
	"strconv"
	"testing"
)

type InputCase struct {
	g        []int
	s        []int
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{g: []int{1, 2, 3}, s: []int{1, 1}, expected: 1, comment: "Example1"},
		{g: []int{1, 2}, s: []int{1, 2, 3}, expected: 2, comment: "Example2"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.g, element.s, element.expected, index, t)
			})
	}
}

func testApp(g []int, s []int, expected int, index int, t *testing.T) {
	actual := findContentChildren(g, s)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
