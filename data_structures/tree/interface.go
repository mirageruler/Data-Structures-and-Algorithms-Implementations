package main

type BinaryNode interface {
	Key() int
	Val() interface{}
	LeftChild() BinaryNode
	RightChild() BinaryNode
}

// BST_Node implements BinaryNode
type BST_Node struct {
	key         int
	val         interface{}
	right, left *BST_Node
}

func (n BST_Node) Key() interface{} {
	return n.key
}

func (n BST_Node) LeftChild() *BST_Node {
	return n.left
}

func (n BST_Node) RightChild() *BST_Node {
	return n.right
}

func (n BST_Node) Val() interface{} {
	return n.val
}
