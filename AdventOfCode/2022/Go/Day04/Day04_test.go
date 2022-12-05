package Day04

import (
	"AdventOfCode-2022/Utils"
	"testing"
)

func TestDay04Part1Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\sections.txt")
	expected := 524
	actual := Day4Part1(data)
	if actual != expected {
		t.Errorf("Test Day 04 Part 1 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay04Part2Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\sections.txt")
	expected := 798
	actual := Day4Part2(data)
	if actual != expected {
		t.Errorf("Test Day 04 Part 2 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}
