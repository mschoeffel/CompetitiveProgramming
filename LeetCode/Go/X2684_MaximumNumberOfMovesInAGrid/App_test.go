package X2684_MaximumNumberOfMovesInAGrid

import (
	"strconv"
	"testing"
)

type InputCase struct {
	grid     [][]int
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{grid: [][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}, expected: 3, comment: "Example1"},
		{grid: [][]int{{3, 2, 4}, {2, 1, 9}, {1, 1, 7}}, expected: 0, comment: "Example2"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.grid, element.expected, index, t)
			})
	}
}

func testApp(grid [][]int, expected int, index int, t *testing.T) {
	actual := maxMoves(grid)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
