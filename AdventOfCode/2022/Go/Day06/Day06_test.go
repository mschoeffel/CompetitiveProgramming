package Day06

import (
	"AdventOfCode-2022/Utils"
	"testing"
)

func TestDay06Part1Test1(t *testing.T) {
	data := "bvwbjplbgvbhsrlpgdmjqwftvncz"
	expected := 5
	actual := Day6Part1(data)
	if actual != expected {
		t.Errorf("Test Day 06 Part 1 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay06Part1Test2(t *testing.T) {
	data := "nppdvjthqldpwncqszvftbrmjlhg"
	expected := 6
	actual := Day6Part1(data)
	if actual != expected {
		t.Errorf("Test Day 06 Part 1 Test 2: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay06Part1Test3(t *testing.T) {
	data := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	expected := 10
	actual := Day6Part1(data)
	if actual != expected {
		t.Errorf("Test Day 06 Part 1 Test 3: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay06Part1Test4(t *testing.T) {
	data := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
	expected := 11
	actual := Day6Part1(data)
	if actual != expected {
		t.Errorf("Test Day 06 Part 1 Test 4: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay06Part1Test5(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\datastream.txt")[0]
	expected := 1766
	actual := Day6Part1(data)
	if actual != expected {
		t.Errorf("Test Day 06 Part 1 Test 4: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay06Part2Test1(t *testing.T) {
	data := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	expected := 19
	actual := Day6Part2(data)
	if actual != expected {
		t.Errorf("Test Day 06 Part 2 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay06Part2Test2(t *testing.T) {
	data := "bvwbjplbgvbhsrlpgdmjqwftvncz"
	expected := 23
	actual := Day6Part2(data)
	if actual != expected {
		t.Errorf("Test Day 06 Part 2 Test 2: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay06Part2Test3(t *testing.T) {
	data := "nppdvjthqldpwncqszvftbrmjlhg"
	expected := 23
	actual := Day6Part2(data)
	if actual != expected {
		t.Errorf("Test Day 06 Part 2 Test 3: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay06Part2Test4(t *testing.T) {
	data := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	expected := 29
	actual := Day6Part2(data)
	if actual != expected {
		t.Errorf("Test Day 06 Part 2 Test 4: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay06Part2Test5(t *testing.T) {
	data := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
	expected := 26
	actual := Day6Part2(data)
	if actual != expected {
		t.Errorf("Test Day 06 Part 2 Test 5: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay06Part2Test6(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\datastream.txt")[0]
	expected := 2383
	actual := Day6Part2(data)
	if actual != expected {
		t.Errorf("Test Day 06 Part 1 Test 6: Expected: %v Actual: %v", expected, actual)
	}
}
