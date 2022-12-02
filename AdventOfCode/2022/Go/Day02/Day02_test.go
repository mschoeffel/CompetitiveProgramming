package Day02

import (
	"AdventOfCode-2022/Utils"
	"testing"
)

func TestDay02Part1Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\guide.txt")
	expected := 13675
	actual := Day2Part1(data)
	if actual != expected {
		t.Errorf("Test Day 02 Part 1 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay02Part2Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\guide.txt")
	expected := 14184
	actual := Day2Part2(data)
	if actual != expected {
		t.Errorf("Test Day 02 Part 2 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}
