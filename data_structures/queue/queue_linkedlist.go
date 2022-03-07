package queue

// A Queue is a linear data structure which follows a particular order in which the operations are performed.
// The order is: First in, first out (FIFO).

// Node defines a single node as a building block of QueueLinkedList.
// Note that this Node is for building a singly linked list.
type Node struct {
	Val  interface{}
	Next *Node
}

// NewNode returns the Node contains val for `Val` and a nil pointer for `Next`.
func NewNode(val interface{}) *Node {
	return &Node{Val: val}
}

// QueueLinkedList defines the queue built from a singly linked list.
type QueueLinkedList struct {
	Head *Node
}

// NewQueueLinkedList returns the QueueLinkedList with field to be filled with zeroed value.
func NewQueueLinkedList() *QueueLinkedList {
	return &QueueLinkedList{}
}

// Enqueue adds val to the end of the queue and return it, -1 is returned if that queue is invalid.
func (q *QueueLinkedList) Enqueue(val interface{}) interface{} {
	if q == nil {
		return -1
	}

	n := NewNode(val)

	if q.Head == nil {
		q.Head = n
		return val
	}

	q.Head.Next = n

	return val
}

// Dequeue removes the first element of the queue and return it, -1 is return it that queue is empty or invalid.
func (q *QueueLinkedList) Dequeue() interface{} {
	if q == nil || q.Head == nil {
		return -1
	}

	retVal := q.Head.Val

	if q.Head.Next == nil {
		q.Head = nil
		return retVal
	}

	q.Head = q.Head.Next

	return retVal
}

// Peak returns the first element of the queue, -1 is returned if that queue is empty or invalid.
func (q *QueueLinkedList) Peak() interface{} {
	if q == nil || q.Head == nil {
		return -1
	}

	return q.Head.Val
}

// IsEmpty tells whether the queue is empty or not, -1 is returned if that queue is invalid.
func (q *QueueLinkedList) IsEmpty() bool {
	if q == nil {
		return true
	}

	return q.Head == nil
}
