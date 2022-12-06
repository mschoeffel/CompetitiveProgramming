package Day05

import (
	"AdventOfCode-2022/Utils"
	"testing"
)

func TestDay05Part1Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\crane.txt")
	expected := "TBVFVDZPN"
	actual := Day5Part1(data)
	if actual != expected {
		t.Errorf("Test Day 05 Part 1 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay05Part2Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\crane.txt")
	expected := "VLCWHTDSZ"
	actual := Day5Part2(data)
	if actual != expected {
		t.Errorf("Test Day 05 Part 2 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}
