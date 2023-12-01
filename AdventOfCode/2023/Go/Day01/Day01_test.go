package Day01

import (
	"AdventOfCode-2023/Utils"
	"strconv"
	"testing"
)

type InputCase struct {
	document []string
	expected int
	comment  string
}

func TestAppCases(t *testing.T) {
	testSets := []InputCase{
		{document: []string{"7"}, expected: 77, comment: "SingleDigit"},
		{document: []string{"72"}, expected: 72, comment: "DoubleDigit"},
		{document: []string{"7a"}, expected: 77, comment: "FirstWithLettersAfterwards"},
		{document: []string{"a7"}, expected: 77, comment: "LastWithLettersBefore"},
		{document: []string{"2a7"}, expected: 27, comment: "FirstAndLastWithLettersBetween"},
		{document: []string{"a2a7a"}, expected: 27, comment: "FirstAndLastWithLettersAround"},
		{document: []string{"a2a4a7a"}, expected: 27, comment: "MultiDigits"},
		{document: []string{"a24a42a77a8"}, expected: 28, comment: "Complex"},
	}

	for index, element := range testSets {
		t.Run("Test Day 01 Part 1: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) { testDay01Part1(element.document, element.expected, index, t) })
	}
}

func testDay01Part1(document []string, expected int, index int, t *testing.T) {
	actual := Day1Part1(document)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: [" + strconv.Itoa(expected) +
			"] Actual: [" + strconv.Itoa(actual) + "]")
	}
}

func TestDay01Part1Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\document.txt")
	expected := 54968
	actual := Day1Part1(data)
	if actual != expected {
		t.Errorf("Test Day 01 Part 1 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}
