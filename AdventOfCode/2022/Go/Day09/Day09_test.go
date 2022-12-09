package Day09

import (
	"AdventOfCode-2022/Utils"
	"testing"
)

func TestDay09Part1Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\test.txt")
	expected := 13
	actual := Day9Part1(data)
	if actual != expected {
		t.Errorf("Test Day 09 Part 1 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay09Part1Test2(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\commands.txt")
	expected := 6384
	actual := Day9Part1(data)
	if actual != expected {
		t.Errorf("Test Day 09 Part 1 Test 2: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay09Part2Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\test.txt")
	expected := 1
	actual := Day9Part2(data)
	if actual != expected {
		t.Errorf("Test Day 09 Part 2 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay09Part2Test2(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\larger-test.txt")
	expected := 36
	actual := Day9Part2(data)
	if actual != expected {
		t.Errorf("Test Day 09 Part 2 Test 2: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay09Part2Test3(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\commands.txt")
	expected := 2734
	actual := Day9Part2(data)
	if actual != expected {
		t.Errorf("Test Day 09 Part 2 Test 3: Expected: %v Actual: %v", expected, actual)
	}
}
