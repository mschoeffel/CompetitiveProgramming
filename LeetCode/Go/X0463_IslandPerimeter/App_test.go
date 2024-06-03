package X0463_IslandPerimeter

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
		{grid: [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}, expected: 16, comment: "Example1"},
		{grid: [][]int{{1}}, expected: 4, comment: "Example2"},
		{grid: [][]int{{1, 0}}, expected: 4, comment: "Example3"},
		{grid: [][]int{{0}}, expected: 0, comment: "OnlyOneZero"},
		{grid: [][]int{{0, 0}, {0, 0}}, expected: 0, comment: "OnlyGridZero"},
		{grid: [][]int{{1, 1}, {1, 1}}, expected: 8, comment: "OnlyGridOne"},
		{grid: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, expected: 16, comment: "Island"},
		{grid: [][]int{{1, 1, 1}}, expected: 8, comment: "SingleRow"},
		{grid: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, expected: 12, comment: "FullNine"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.grid, element.expected, index, t)
			})
	}
}

func testApp(grid [][]int, expected int, index int, t *testing.T) {
	actual := islandPerimeter(grid)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
