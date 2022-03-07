package stack

// A Stack is a linear data structure which follows a particular order in which the operations are performed.
// The order is: Last in, first out (LIFO) or first in, last out (FILO).

// Node defines a single node as a building block of StackLinkedList.
// Note that this Node is for building a singly linked list.
type Node struct {
	Val  interface{}
	Next *Node
}

// NewNode returns the Node contains val for `Val` and a nil pointer for `Next`.
func NewNode(val interface{}) *Node {
	return &Node{Val: val}
}

// StackLinkedList defines the stack built from a singly linked list.
type StackLinkedList struct {
	Top *Node
}

// NewStackLinkedList returns the StackLinkedList with field to be filled with zeroed value.
func NewStackLinkedList() *StackLinkedList {
	return &StackLinkedList{}
}

// Push pushes val into the top of the stack and return it, -1 is returned if that stack is invalid.
func (s *StackLinkedList) Push(val interface{}) interface{} {
	if s == nil {
		return -1
	}

	n := NewNode(val)

	if s.Top == nil {
		s.Top = n
		return val
	}

	n.Next = s.Top
	s.Top = n

	return val
}

// Pop pops the top most element of the stack and return it, -1 is returned if that stack is empty or invalid.
func (s *StackLinkedList) Pop() interface{} {
	if s == nil || s.Top == nil {
		return -1
	}

	retVal := s.Top.Val

	if s.Top.Next == nil {
		s.Top = nil
		return retVal
	}

	s.Top = s.Top.Next

	return retVal
}

// Peak returns the top most element of the stack, -1 is returned if that stack is empty or invalid.
func (s *StackLinkedList) Peak() interface{} {
	if s == nil || s.Top == nil {
		return -1
	}

	return s.Top.Val
}

// IsEmpty tells whether the stack is empty or not, -1 is returned if that stack is invalid.
func (s *StackLinkedList) IsEmpty() interface{} {
	if s == nil {
		return true
	}

	return s.Top == nil
}
