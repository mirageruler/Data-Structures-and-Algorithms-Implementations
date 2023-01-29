package main

// Node defines a node in any tree structure
// Note that we only support integer value for the Node's value so we can implement the tree structures with ease.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// NewNode returns a Node contains val as `Val` and nil pointer for `Left` and `Right`
func NewNode(val int) *Node {
	return &Node{Val: val}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
