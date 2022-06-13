package X0070_ClimbingStairs

import (
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	stairs   int
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{stairs: 2, expected: 2, comment: "Example1"},
		{stairs: 3, expected: 3, comment: "Example2"},
		{stairs: 4, expected: 5, comment: "4Stairs"},
		{stairs: 1, expected: 1, comment: "1Stair"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.stairs, element.expected, index, t)
			})
	}
}

func testApp(stairs int, expected int, index int, t *testing.T) {
	actual := climbStairs(stairs)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
