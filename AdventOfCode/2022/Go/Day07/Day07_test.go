package Day07

import (
	"AdventOfCode-2022/Utils"
	"testing"
)

func TestDay07Part1Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\test.txt")
	expected := 95437
	actual := Day7Part1(data)
	if actual != expected {
		t.Errorf("Test Day 07 Part 1 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay07Part1Test2(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\commands.txt")
	expected := 1307902
	actual := Day7Part1(data)
	if actual != expected {
		t.Errorf("Test Day 07 Part 1 Test 2: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay07Part2Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\test.txt")
	expected := 24933642
	actual := Day7Part2(data)
	if actual != expected {
		t.Errorf("Test Day 07 Part 2 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay07Part2Test2(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\commands.txt")
	expected := 7068748
	actual := Day7Part2(data)
	if actual != expected {
		t.Errorf("Test Day 07 Part 2 Test 2: Expected: %v Actual: %v", expected, actual)
	}
}
