package X0495_TeemoAttacking

import (
	"strconv"
	"testing"
)

type InputCase struct {
	timeSeries []int
	duration   int
	expected   int
	comment    string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{timeSeries: []int{1, 4}, duration: 2, expected: 4, comment: "Example1"},
		{timeSeries: []int{1, 2}, duration: 2, expected: 3, comment: "Example2"},
		{timeSeries: []int{1}, duration: 5, expected: 5, comment: "SingleSeries"},
		{timeSeries: []int{1, 2, 3, 4, 5}, duration: 5, expected: 9, comment: "MaxOverlap"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.timeSeries, element.duration, element.expected, index, t)
			})
	}
}

func testApp(timeSeries []int, duration int, expected int, index int, t *testing.T) {
	actual := findPoisonedDuration(timeSeries, duration)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
