package X0476_NumberComplement

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
		{num: 5, expected: 2, comment: "Example1"},
		{num: 1, expected: 0, comment: "Example2"},
		{num: 0, expected: 1, comment: "Complement 0"},
		{num: 3, expected: 0, comment: "Complement multiple 1"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.num, element.expected, index, t)
			})
		t.Run("Test alternative: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testAppAlternative(element.num, element.expected, index, t)
			})
	}
}

func testApp(num int, expected int, index int, t *testing.T) {
	actual := findComplement(num)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}

func testAppAlternative(num int, expected int, index int, t *testing.T) {
	actual := findComplement(num)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}

func BenchmarkApp(b *testing.B) {
	benchmarkSets := []InputCase{
		{num: 5, expected: 2, comment: "Example1"},
		{num: 1, expected: 0, comment: "Example2"},
		{num: 0, expected: 1, comment: "Complement 0"},
		{num: 3, expected: 0, comment: "Complement multiple 1"},
		{num: 1024, expected: 1023, comment: "Big number one less"},
		{num: 1023, expected: 0, comment: "Big number 0"},
	}

	for index, element := range benchmarkSets {
		b.Run("Benchmark: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkApp(element.num, b) })
		b.Run("Benchmark alternative: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkAppAlternative(element.num, b) })
	}
}

func benchmarkApp(num int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		findComplement(num)
	}
}

func benchmarkAppAlternative(num int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		findComplementAlternative(num)
	}
}
