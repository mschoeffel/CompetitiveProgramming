package X0401_BinaryWatch

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	turnedOn int
	expected []string
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{turnedOn: 1, expected: []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"}, comment: "Example1"},
		{turnedOn: 9, expected: nil, comment: "Example2"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.turnedOn, element.expected, index, t)
			})
	}
}

func testApp(turnedOn int, expected []string, index int, t *testing.T) {
	actual := readBinaryWatch(turnedOn)
	if !reflect.DeepEqual(actual, expected) {
		expectedString, _ := json.Marshal(expected)
		actualString, _ := json.Marshal(actual)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
