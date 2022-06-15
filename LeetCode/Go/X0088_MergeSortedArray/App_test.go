package X0088_MergeSortedArray

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	nums1    []int
	m        int
	nums2    []int
	n        int
	expected []int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3, expected: []int{1, 2, 2, 3, 5, 6}, comment: "Example1"},
		{nums1: []int{1}, m: 1, nums2: []int{}, n: 0, expected: []int{1}, comment: "Example2"},
		{nums1: []int{0}, m: 0, nums2: []int{1}, n: 1, expected: []int{1}, comment: "Example3"},
		{nums1: []int{-1000000000, 0, 0, 3, 3, 3, 0, 0, 0}, m: 6, nums2: []int{1, 2, 2}, n: 3, expected: []int{-1000000000, 0, 0, 1, 2, 2, 3, 3, 3}, comment: "Negative"},
		{nums1: []int{1, 1, 3, 4, 4, 5, 6, 6, 7, 0, 0, 0, 0, 0, 0, 0}, m: 9, nums2: []int{1, 2, 3, 4, 5, 5, 6}, n: 7, expected: []int{1, 1, 1, 2, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7}, comment: "MoreComplicated"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				nums1Temp := make([]int, len(element.nums1))
				copy(nums1Temp, element.nums1)
				testApp(nums1Temp, element.m, element.nums2, element.n, element.expected, index, t)
			})

		t.Run("Test Optimized: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				nums1Temp := make([]int, len(element.nums1))
				copy(nums1Temp, element.nums1)
				testAppOptimized(nums1Temp, element.m, element.nums2, element.n, element.expected, index, t)
			})

		t.Run("Test BuiltIn: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				nums1Temp := make([]int, len(element.nums1))
				copy(nums1Temp, element.nums1)
				testAppBuiltIn(nums1Temp, element.m, element.nums2, element.n, element.expected, index, t)
			})
	}
}

func testApp(nums1 []int, m int, nums2 []int, n int, expected []int, index int, t *testing.T) {
	merge(nums1, m, nums2, n)
	if !reflect.DeepEqual(nums1, expected) {
		actualString, _ := json.Marshal(nums1)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}

func testAppOptimized(nums1 []int, m int, nums2 []int, n int, expected []int, index int, t *testing.T) {
	mergeOptimized(nums1, m, nums2, n)
	if !reflect.DeepEqual(nums1, expected) {
		actualString, _ := json.Marshal(nums1)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}

func testAppBuiltIn(nums1 []int, m int, nums2 []int, n int, expected []int, index int, t *testing.T) {
	mergeBuiltIn(nums1, m, nums2, n)
	if !reflect.DeepEqual(nums1, expected) {
		actualString, _ := json.Marshal(nums1)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}

func BenchmarkApp(b *testing.B) {
	benchmarkSets := []InputCase{
		{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3, expected: []int{1, 2, 2, 3, 5, 6}, comment: "Example1"},
		{nums1: []int{1}, m: 1, nums2: []int{}, n: 0, expected: []int{1}, comment: "Example2"},
		{nums1: []int{0}, m: 0, nums2: []int{1}, n: 1, expected: []int{1}, comment: "Example3"},
		{nums1: []int{-1000000000, 0, 0, 3, 3, 3, 0, 0, 0}, m: 6, nums2: []int{1, 2, 2}, n: 3, expected: []int{-1000000000, 0, 0, 1, 2, 2, 3, 3, 3}, comment: "Negative"},
		{nums1: []int{1, 1, 3, 4, 4, 5, 6, 6, 7, 0, 0, 0, 0, 0, 0, 0}, m: 9, nums2: []int{1, 2, 3, 4, 5, 5, 6}, n: 7, expected: []int{1, 1, 1, 2, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7}, comment: "MoreComplicated"},
	}

	for index, element := range benchmarkSets {
		b.Run("Benchmark: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkApp(element.nums1, element.m, element.nums2, element.n, b) })
		b.Run("Benchmark Optimized: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkAppOptimized(element.nums1, element.m, element.nums2, element.n, b) })
		b.Run("Benchmark BuiltIn: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkAppBuiltIn(element.nums1, element.m, element.nums2, element.n, b) })
	}
}

func benchmarkApp(nums1 []int, m int, nums2 []int, n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums1Temp := make([]int, len(nums1))
		copy(nums1Temp, nums1)
		merge(nums1Temp, m, nums2, n)
	}
}

func benchmarkAppOptimized(nums1 []int, m int, nums2 []int, n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums1Temp := make([]int, len(nums1))
		copy(nums1Temp, nums1)
		mergeOptimized(nums1Temp, m, nums2, n)
	}
}

func benchmarkAppBuiltIn(nums1 []int, m int, nums2 []int, n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums1Temp := make([]int, len(nums1))
		copy(nums1Temp, nums1)
		mergeBuiltIn(nums1Temp, m, nums2, n)
	}
}
