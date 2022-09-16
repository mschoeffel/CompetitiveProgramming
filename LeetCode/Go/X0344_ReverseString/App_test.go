package X0344_ReverseString

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	s        []byte
	expected []byte
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{s: []byte{'h', 'e', 'l', 'l', 'o'}, expected: []byte{'o', 'l', 'l', 'e', 'h'}, comment: "Example1"},
		{s: []byte{'H', 'a', 'n', 'n', 'a', 'h'}, expected: []byte{'h', 'a', 'n', 'n', 'a', 'H'}, comment: "Example2"},
		{s: []byte{'H'}, expected: []byte{'H'}, comment: "SingleChar"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.s, element.expected, index, t)
			})
	}
}

func testApp(s []byte, expected []byte, index int, t *testing.T) {
	reverseString(s)
	if !reflect.DeepEqual(s, expected) {
		actualString, _ := json.Marshal(s)
		expectedString, _ := json.Marshal(s)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
