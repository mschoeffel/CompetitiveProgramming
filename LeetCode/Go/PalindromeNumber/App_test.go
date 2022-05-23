package PalindromeNumber

import (
	"reflect"
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
		{num: 121, expected: true, comment: "Example1"},
		{num: -121, expected: false, comment: "Example2"},
		{num: 10, expected: false, comment: "Example3"},
		{num: 1234567890987654321, expected: true, comment: "LongNumber"},
		{num: 22222, expected: true, comment: "AllSameDigits"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) { testApp(element.num, element.expected, index, t) })
		t.Run("Test ReverseNumber: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) { testAppReverse(element.num, element.expected, index, t) })
	}
}

func testApp(num int, expected bool, index int, t *testing.T) {
	actual := palindromeNumber(num)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(actual))
	}
}

func testAppReverse(num int, expected bool, index int, t *testing.T) {
	actual := palindromeNumberReverseNumber(num)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(actual))
	}
}

func BenchmarkApp(b *testing.B) {
	benchmarkSets := []InputCase{
		{num: 121, expected: true, comment: "Example1"},
		{num: -121, expected: false, comment: "Example2"},
		{num: 10, expected: false, comment: "Example3"},
		{num: 1234567890987654321, expected: true, comment: "LongNumber"},
		{num: 22222, expected: true, comment: "AllSameDigits"},
	}

	for index, element := range benchmarkSets {
		b.Run("Benchmark: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkApp(element.num, b) })
		b.Run("Benchmark ReverseNumber: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkAppReverse(element.num, b) })
	}
}

func benchmarkApp(num int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		palindromeNumber(num)
	}
}

func benchmarkAppReverse(num int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		palindromeNumberReverseNumber(num)
	}
}
