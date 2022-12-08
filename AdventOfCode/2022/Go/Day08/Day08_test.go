package Day08

import (
	"AdventOfCode-2022/Utils"
	"testing"
)

func TestDay08Part1Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\test.txt")
	expected := 21
	actual := Day8Part1(data)
	if actual != expected {
		t.Errorf("Test Day 08 Part 1 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay08Part1Test2(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\map.txt")
	expected := 1700
	actual := Day8Part1(data)
	if actual != expected {
		t.Errorf("Test Day 08 Part 1 Test 2: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay08Part2Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\test.txt")
	expected := 8
	actual := Day8Part2(data)
	if actual != expected {
		t.Errorf("Test Day 08 Part 2 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay08Part2Test2(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\map.txt")
	expected := 470596
	actual := Day8Part2(data)
	if actual != expected {
		t.Errorf("Test Day 08 Part 2 Test 2: Expected: %v Actual: %v", expected, actual)
	}
}
