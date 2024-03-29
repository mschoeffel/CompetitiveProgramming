package X0232_ImplementQueueUsingStacks

type MyQueue struct {
	vals []int
}

func Main() {
	myQueue := Constructor()
	myQueue.Push(1)
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	rvals := make([]int, 0, len(q.vals))
	for len(q.vals) > 0 {
		rvals = append(rvals, q.Pop())
	}
	q.vals = q.vals[:0]
	q.vals = append(q.vals, x)
	for len(rvals) > 0 {
		q.vals = append(q.vals, rvals[len(rvals)-1])
		rvals = rvals[:len(rvals)-1]
	}
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	v := q.vals[len(q.vals)-1]
	q.vals = q.vals[:len(q.vals)-1]
	return v
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	return q.vals[len(q.vals)-1]
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return len(q.vals) == 0
}
