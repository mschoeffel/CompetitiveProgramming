package X0169_MajorityElement

import (
	"strconv"
	"testing"
)

type InputCase struct {
	nums     []int
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{nums: []int{3, 2, 3}, expected: 3, comment: "Example1"},
		{nums: []int{2, 2, 1, 1, 1, 2, 2}, expected: 2, comment: "Example2"},
		{nums: []int{2}, expected: 2, comment: "SingleElement"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.nums, element.expected, index, t)
			})

		t.Run("Test Sort: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testAppSort(element.nums, element.expected, index, t)
			})
	}
}

func testApp(nums []int, expected int, index int, t *testing.T) {
	actual := majorityElement(nums)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}

func testAppSort(nums []int, expected int, index int, t *testing.T) {
	actual := majorityElementSort(nums)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}

func BenchmarkApp(b *testing.B) {
	benchmarkSets := []InputCase{
		{nums: []int{3, 2, 3}, expected: 3, comment: "Example1"},
		{nums: []int{2, 2, 1, 1, 1, 2, 2}, expected: 2, comment: "Example2"},
		{nums: []int{2}, expected: 2, comment: "SingleElement"},
	}

	for index, element := range benchmarkSets {
		b.Run("Benchmark: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkApp(element.nums, b) })
		b.Run("Benchmark Sort: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkAppSort(element.nums, b) })
	}
}

func benchmarkApp(nums []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		majorityElement(nums)
	}
}

func benchmarkAppSort(nums []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		majorityElementSort(nums)
	}
}
