package X0136_SingleNumber

import (
	"strconv"
	"testing"
)

type InputCase struct {
	input    []int
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{input: []int{2, 2, 1}, expected: 1, comment: "Example1"},
		{input: []int{4, 1, 2, 1, 2}, expected: 4, comment: "Example2"},
		{input: []int{1}, expected: 1, comment: "Example3"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.input, element.expected, index, t)
			})

		t.Run("Test XOr: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testAppXOr(element.input, element.expected, index, t)
			})
	}
}

func testApp(input []int, expected int, index int, t *testing.T) {
	result := singleNumber(input)
	if result != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(result))
	}
}

func testAppXOr(input []int, expected int, index int, t *testing.T) {
	result := singleNumberXOr(input)
	if result != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(result))
	}
}

func BenchmarkApp(b *testing.B) {
	benchmarkSets := []InputCase{
		{input: []int{2, 2, 1}, expected: 1, comment: "Example1"},
		{input: []int{4, 1, 2, 1, 2}, expected: 4, comment: "Example2"},
		{input: []int{1}, expected: 1, comment: "Example3"},
	}

	for index, element := range benchmarkSets {
		b.Run("Benchmark: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkApp(element.input, b) })
		b.Run("Benchmark XOr: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkAppXOr(element.input, b) })
	}
}

func benchmarkApp(input []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		singleNumber(input)
	}
}

func benchmarkAppXOr(input []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		singleNumberXOr(input)
	}
}
