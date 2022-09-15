package X0342_PowerOfFour

import (
	"strconv"
	"testing"
)

type InputCase struct {
	num      int
	expected bool
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{num: 16, expected: true, comment: "Example1"},
		{num: 5, expected: false, comment: "Example2"},
		{num: 1, expected: true, comment: "Example3"},
		{num: 3, expected: false, comment: "Three"},
		{num: 4, expected: true, comment: "Four"},
		{num: -1, expected: false, comment: "Negative"},
		{num: -16777216, expected: false, comment: "NegativeLong"},
		{num: 16777216, expected: true, comment: "PositiveLongHappyCase"},
		{num: 16777226, expected: false, comment: "PositiveLongBadCase"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.num, element.expected, index, t)
			})

		t.Run("Test BitPattern: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testAppBitPattern(element.num, element.expected, index, t)
			})
	}
}

func testApp(num int, expected bool, index int, t *testing.T) {
	actual := isPowerOfFour(num)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(actual))
	}
}

func testAppBitPattern(num int, expected bool, index int, t *testing.T) {
	actual := isPowerOfFourBitPattern(num)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(actual))
	}
}

func BenchmarkApp(b *testing.B) {
	benchmarkSets := []InputCase{
		{num: 16, expected: true, comment: "Example1"},
		{num: 5, expected: false, comment: "Example2"},
		{num: 1, expected: true, comment: "Example3"},
		{num: 3, expected: false, comment: "Three"},
		{num: 4, expected: true, comment: "Four"},
		{num: -1, expected: false, comment: "Negative"},
		{num: -16777216, expected: false, comment: "NegativeLong"},
		{num: 16777216, expected: false, comment: "PositiveLongHappyCase"},
		{num: 16777226, expected: false, comment: "PositiveLongBadCase"},
	}

	for index, element := range benchmarkSets {
		b.Run("Benchmark: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkApp(element.num, b) })
		b.Run("Benchmark BitPattern: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkAppMath(element.num, b) })
	}
}

func benchmarkApp(num int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPowerOfFour(num)
	}
}

func benchmarkAppMath(num int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPowerOfFourBitPattern(num)
	}
}
