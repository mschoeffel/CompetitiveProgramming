package X0412_FizzBuzz

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	num      int
	expected []string
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{num: 3, expected: []string{"1", "2", "Fizz"}, comment: "Example1"},
		{num: 5, expected: []string{"1", "2", "Fizz", "4", "Buzz"}, comment: "Example2"},
		{num: 15, expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}, comment: "Example3"},
		{num: 1, expected: []string{"1"}, comment: "One"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.num, element.expected, index, t)
			})
	}
}

func testApp(num int, expected []string, index int, t *testing.T) {
	actual := fizzBuzz(num)
	if !reflect.DeepEqual(actual, expected) {
		actualString, _ := json.Marshal(actual)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
