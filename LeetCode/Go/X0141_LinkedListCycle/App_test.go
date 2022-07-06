package X0141_LinkedListCycle

import (
	"strconv"
	"testing"
)

type InputCase struct {
	head     *ListNode
	expected bool
	comment  string
}

func TestApp(t *testing.T) {
	testExample1First := &ListNode{Val: 3}
	testExample1Second := &ListNode{Val: 2}
	testExample1Third := &ListNode{Val: 0}
	testExample1Fourth := &ListNode{Val: -4, Next: testExample1Second}
	testExample1Third.Next = testExample1Fourth
	testExample1Second.Next = testExample1Third
	testExample1First.Next = testExample1Second

	testExample2First := &ListNode{Val: 1}
	testExample2Second := &ListNode{Val: 2, Next: testExample2First}
	testExample2First.Next = testExample2Second

	testExample4First := &ListNode{Val: 1}
	testExample4Second := &ListNode{Val: 2}
	testExample4Third := &ListNode{Val: 1}
	testExample4First.Next = testExample4Second
	testExample4Second.Next = testExample4Third

	testSets := []InputCase{
		{head: testExample1First, expected: true, comment: "Example1"},
		{head: testExample2First, expected: true, comment: "Example2"},
		{head: &ListNode{Val: 1}, expected: false, comment: "Example3"},
		{head: testExample4First, expected: false, comment: "MultipleNodesNoCycle"},
		{head: nil, expected: false, comment: "NullInput"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.head, element.expected, index, t)
			})
		t.Run("Test Fast Slow: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testAppFastSlow(element.head, element.expected, index, t)
			})
	}
}

func testApp(head *ListNode, expected bool, index int, t *testing.T) {
	result := hasCycle(head)
	if result != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(result))
	}
}

func testAppFastSlow(head *ListNode, expected bool, index int, t *testing.T) {
	result := hasCycleFastSlow(head)
	if result != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(result))
	}
}

func BenchmarkApp(b *testing.B) {
	testExample1First := &ListNode{Val: 3}
	testExample1Second := &ListNode{Val: 2}
	testExample1Third := &ListNode{Val: 0}
	testExample1Fourth := &ListNode{Val: -4, Next: testExample1Second}
	testExample1Third.Next = testExample1Fourth
	testExample1Second.Next = testExample1Third
	testExample1First.Next = testExample1Second

	testExample2First := &ListNode{Val: 1}
	testExample2Second := &ListNode{Val: 2, Next: testExample2First}
	testExample2First.Next = testExample2Second

	testExample4First := &ListNode{Val: 1}
	testExample4Second := &ListNode{Val: 2}
	testExample4Third := &ListNode{Val: 3}
	testExample4First.Next = testExample4Second
	testExample4Second.Next = testExample4Third

	benchmarkSets := []InputCase{
		{head: testExample1First, expected: true, comment: "Example1"},
		{head: testExample2First, expected: true, comment: "Example2"},
		{head: &ListNode{Val: 1}, expected: false, comment: "Example3"},
		{head: testExample4First, expected: false, comment: "MultipleNodesNoCycle"},
		{head: nil, expected: false, comment: "NullInput"},
	}

	for index, element := range benchmarkSets {
		b.Run("Benchmark: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkApp(element.head, b) })
		b.Run("Benchmark Fast Slow: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkAppFastSlow(element.head, b) })
	}
}

func benchmarkApp(head *ListNode, b *testing.B) {
	for i := 0; i < b.N; i++ {
		hasCycle(head)
	}
}

func benchmarkAppFastSlow(head *ListNode, b *testing.B) {
	for i := 0; i < b.N; i++ {
		hasCycleFastSlow(head)
	}
}
