package Day02

import (
	"AdventOfCode-2024/Utils"
	"testing"
)

func TestDay02Part1Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\input.txt")
	expected := 314
	actual := Day2Part1(data)
	if actual != expected {
		t.Errorf("Test Day 02 Part 1 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay02Part1Testdata(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\testdata.txt")
	expected := 2
	actual := Day2Part1(data)
	if actual != expected {
		t.Errorf("Test Day 02 Part 1 Testdata: Expected: %v Actual: %v", expected, actual)
	}
}
