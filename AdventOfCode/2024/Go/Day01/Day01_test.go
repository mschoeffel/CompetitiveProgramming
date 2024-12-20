package Day01

import (
	"AdventOfCode-2024/Utils"
	"testing"
)

func TestDay01Part1Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\input.txt")
	expected := 1666427
	actual := Day1Part1(data)
	if actual != expected {
		t.Errorf("Test Day 01 Part 1 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay01Part2Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\input.txt")
	expected := 24316233
	actual := Day1Part2(data)
	if actual != expected {
		t.Errorf("Test Day 01 Part 2 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}
