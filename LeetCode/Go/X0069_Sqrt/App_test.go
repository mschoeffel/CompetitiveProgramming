package X0069_Sqrt

import (
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	x        int
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{x: 4, expected: 2, comment: "Example1"},
		{x: 8, expected: 2, comment: "Example2"},
		{x: 2566439, expected: 1602, comment: "BigNumber"},
		{x: 11, expected: 3, comment: "Prime"},
		{x: 0, expected: 0, comment: "Zero"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.x, element.expected, index, t)
			})
		t.Run("Test Sort Search: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testAppSortSearch(element.x, element.expected, index, t)
			})
	}
}

func testApp(x int, expected int, index int, t *testing.T) {
	actual := mySqrt(x)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}

func testAppSortSearch(x int, expected int, index int, t *testing.T) {
	actual := mySqrtSortSearch(x)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}

func BenchmarkApp(b *testing.B) {
	benchmarkSets := []InputCase{
		{x: 4, expected: 2, comment: "Example1"},
		{x: 8, expected: 2, comment: "Example2"},
		{x: 2566439, expected: 1602, comment: "BigNumber"},
		{x: 11, expected: 3, comment: "Prime"},
		{x: 0, expected: 0, comment: "Zero"},
	}

	for index, element := range benchmarkSets {
		b.Run("Benchmark: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkApp(element.x, b) })
		b.Run("Benchmark Sort Search: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkAppSortSearch(element.x, b) })
	}
}

func benchmarkApp(x int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		mySqrt(x)
	}
}

func benchmarkAppSortSearch(x int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		mySqrtSortSearch(x)
	}
}
