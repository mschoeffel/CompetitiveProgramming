package X0083_RemoveDuplicatesFromSortedList

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	list     *ListNode
	expected *ListNode
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{list: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2}}, comment: "Example1"},
		{list: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}, comment: "Example2"},
		{list: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7}}}}}}}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 7}}}}}, comment: "MoreComplexOne"},
		{list: nil,
			expected: nil, comment: "ListNull"},
		{list: &ListNode{Val: 5},
			expected: &ListNode{Val: 5}, comment: "ListOnlyOneElement"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) { testApp(element.list, element.expected, index, t) })
		t.Run("Test In Place: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) { testAppInPlace(element.list, element.expected, index, t) })
	}
}

func testApp(list *ListNode, expected *ListNode, index int, t *testing.T) {
	actual := deleteDuplicates(list)
	if !reflect.DeepEqual(actual, expected) {
		actualString, _ := json.Marshal(actual)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}

func testAppInPlace(list *ListNode, expected *ListNode, index int, t *testing.T) {
	actual := deleteDuplicatesInPlace(list)
	if !reflect.DeepEqual(actual, expected) {
		actualString, _ := json.Marshal(actual)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}

func BenchmarkApp(b *testing.B) {
	benchmarkSets := []InputCase{
		{list: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2}}, comment: "Example1"},
		{list: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}, comment: "Example2"},
		{list: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7}}}}}}}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 7}}}}}, comment: "MoreComplexOne"},
		{list: nil,
			expected: nil, comment: "ListNull"},
		{list: &ListNode{Val: 5},
			expected: &ListNode{Val: 5}, comment: "ListOnlyOneElement"},
	}

	for index, element := range benchmarkSets {
		b.Run("Benchmark: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkApp(element.list, b) })
		b.Run("Benchmark In Place: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkAppInPlace(element.list, b) })
	}
}

func benchmarkApp(list *ListNode, b *testing.B) {
	for i := 0; i < b.N; i++ {
		deleteDuplicates(list)
	}
}

func benchmarkAppInPlace(list *ListNode, b *testing.B) {
	for i := 0; i < b.N; i++ {
		deleteDuplicatesInPlace(list)
	}
}
