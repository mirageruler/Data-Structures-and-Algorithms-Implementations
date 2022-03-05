package linkedlist

// Node represent the linkedlist node.
// This Node is shared across different implementation of linked list (singly, doubly, cyclic)
type Node struct {
	Val  interface{}
	Next *Node
	Prev *Node
}

// NewNode creates a new node with `val` argument.
func NewNode(val interface{}) *Node {
	return &Node{val, nil, nil}
}
