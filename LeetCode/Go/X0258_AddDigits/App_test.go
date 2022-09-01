package X0258_AddDigits

import (
	"strconv"
	"testing"
)

type InputCase struct {
	num      int
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{num: 38, expected: 2, comment: "Example1"},
		{num: 0, expected: 0, comment: "Example2"},
		{num: 9, expected: 9, comment: "SingleOne"},
		{num: 99, expected: 9, comment: "MoreComplex"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.num, element.expected, index, t)
			})

		t.Run("Test Math: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testAppMath(element.num, element.expected, index, t)
			})
	}
}

func testApp(num int, expected int, index int, t *testing.T) {
	result := addDigits(num)
	if result != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(result))
	}
}

func testAppMath(num int, expected int, index int, t *testing.T) {
	result := addDigitsMath(num)
	if result != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(result))
	}
}

func BenchmarkApp(b *testing.B) {
	benchmarkSets := []InputCase{
		{num: 38, expected: 2, comment: "Example1"},
		{num: 0, expected: 0, comment: "Example2"},
		{num: 9, expected: 9, comment: "SingleOne"},
		{num: 99, expected: 9, comment: "MoreComplex"},
	}

	for index, element := range benchmarkSets {
		b.Run("Benchmark: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkApp(element.num, b) })
		b.Run("Benchmark Math: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkAppMath(element.num, b) })
	}
}

func benchmarkApp(num int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		addDigits(num)
	}
}

func benchmarkAppMath(num int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		addDigitsMath(num)
	}
}
