package X1277_CountSquareSubmatricesWithAllOnes

import (
	"strconv"
	"testing"
)

type InputCase struct {
	matrix   [][]int
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{matrix: [][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}}, expected: 15, comment: "Example1"},
		{matrix: [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}, expected: 7, comment: "Example2"},
		{matrix: [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}, expected: 0, comment: "OnlyZero"},
		{matrix: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, expected: 14, comment: "OnlyOnes"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.matrix, element.expected, index, t)
			})
	}
}

func testApp(matrix [][]int, expected int, index int, t *testing.T) {
	actual := countSquares(matrix)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
