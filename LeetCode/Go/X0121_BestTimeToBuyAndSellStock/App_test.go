package X0121_BestTimeToBuyAndSellStock

import (
	"strconv"
	"testing"
)

type InputCase struct {
	prices   []int
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{prices: []int{7, 1, 5, 3, 6, 4}, expected: 5, comment: "Example1"},
		{prices: []int{7, 6, 4, 3, 1}, expected: 0, comment: "Example2"},
		{prices: []int{}, expected: 0, comment: "EmptyArray"},
		{prices: []int{7, 2, 5, 3, 10, 1, 4, 9}, expected: 8, comment: "MoreComplex"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.prices, element.expected, index, t)
			})
	}
}

func testApp(prices []int, expected int, index int, t *testing.T) {
	result := maxProfit(prices)
	if result != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(result))
	}
}
