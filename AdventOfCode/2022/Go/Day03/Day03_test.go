package Day03

import (
	"AdventOfCode-2022/Utils"
	"testing"
)

func TestDay03Part1Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\contents.txt")
	expected := 7990
	actual := Day3Part1(data)
	if actual != expected {
		t.Errorf("Test Day 03 Part 1 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay03Part2Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\contents.txt")
	expected := 2602
	actual := Day3Part2(data)
	if actual != expected {
		t.Errorf("Test Day 03 Part 2 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}
