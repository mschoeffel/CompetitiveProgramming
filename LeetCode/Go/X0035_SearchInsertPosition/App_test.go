package X0035_SearchInsertPosition

import (
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	nums     []int
	target   int
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{nums: []int{1, 3, 5, 6}, target: 5, expected: 2, comment: "Example1"},
		{nums: []int{1, 3, 5, 6}, target: 2, expected: 1, comment: "Example2"},
		{nums: []int{1, 3, 5, 6}, target: 7, expected: 4, comment: "Example3"},
		{nums: []int{2, 3, 5, 6}, target: 1, expected: 0, comment: "FirstElement"},
		{nums: []int{}, target: 1, expected: 0, comment: "NumsEmpty"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.nums, element.target, element.expected, index, t)
			})
		t.Run("Test Build In: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testAppBuildIn(element.nums, element.target, element.expected, index, t)
			})
	}
}

func testApp(nums []int, target int, expected int, index int, t *testing.T) {
	actual := searchInsert(nums, target)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}

func testAppBuildIn(nums []int, target int, expected int, index int, t *testing.T) {
	actual := searchInsertBuildIn(nums, target)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}

func BenchmarkApp(b *testing.B) {
	benchmarkSets := []InputCase{
		{nums: []int{1, 3, 5, 6}, target: 5, comment: "Example1"},
		{nums: []int{1, 3, 5, 6}, target: 2, comment: "Example2"},
		{nums: []int{1, 3, 5, 6}, target: 7, comment: "Example3"},
		{nums: []int{2, 3, 5, 6}, target: 1, comment: "FirstElement"},
		{nums: []int{}, target: 1, comment: "NumsEmpty"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, target: 12, comment: "LongSampleIncluded"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, target: 30, comment: "LongSampleExcluded"},
	}

	for index, element := range benchmarkSets {
		b.Run("Benchmark: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkApp(element.nums, element.target, b) })
		b.Run("Benchmark BuildIn: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkAppBuildIn(element.nums, element.target, b) })
	}
}

func benchmarkApp(nums []int, target int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchInsert(nums, target)
	}
}

func benchmarkAppBuildIn(nums []int, target int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchInsertBuildIn(nums, target)
	}
}
