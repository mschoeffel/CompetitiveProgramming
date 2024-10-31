package X2463_MinimumTotalDistanceTraveled

import (
	"strconv"
	"testing"
)

type InputCase struct {
	robot    []int
	factory  [][]int
	expected int64
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{robot: []int{0, 4, 6}, factory: [][]int{{2, 2}, {6, 2}}, expected: 4, comment: "Example1"},
		{robot: []int{1, -1}, factory: [][]int{{-2, 1}, {2, 1}}, expected: 2, comment: "Example2"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.robot, element.factory, element.expected, index, t)
			})
	}
}

func testApp(robot []int, factory [][]int, expected int64, index int, t *testing.T) {
	actual := minimumTotalDistance(robot, factory)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(int(expected)) +
			" Actual: " + strconv.Itoa(int(actual)))
	}
}
