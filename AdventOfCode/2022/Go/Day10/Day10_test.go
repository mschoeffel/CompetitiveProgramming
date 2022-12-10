package Day10

import (
	"AdventOfCode-2022/Utils"
	"testing"
)

func TestDay10Part1Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\test.txt")
	expected := 13140
	actual := Day10Part1(data)
	if actual != expected {
		t.Errorf("Test Day 10 Part 1 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay10Part1Test2(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\program.txt")
	expected := 12740
	actual := Day10Part1(data)
	if actual != expected {
		t.Errorf("Test Day 10 Part 1 Test 2: Expected: %v Actual: %v", expected, actual)
	}
}
