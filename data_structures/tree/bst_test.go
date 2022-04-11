package tree

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestInsert_BST(t *testing.T) {
	t.Parallel()

	bst := NewBST()

	bst.Insert(9)
	bst.Insert(4)
	bst.Insert(6)
	bst.Insert(20)
	bst.Insert(170)
	bst.Insert(15)
	bst.Insert(1)

	bs, _ := json.Marshal(bst)
	fmt.Println("BST: ", string(bs))

	/*
					  9
			   4             20
		   1       6     15      170

	*/

	fmt.Println("looked:", bst.LookUp(170))

	bfs := BFS_Recursive([]*Node{bst.Root}, []int{})
	fmt.Println("BFS: ", bfs)

	dfsInOrder := DFSInOrder(bst.Root)
	fmt.Println("DFS - InOrder:", *dfsInOrder)

	dfsPostOrder := DFSPostOrder(bst.Root)
	fmt.Println("DFS - PostOrder:", *dfsPostOrder)

	dfsPreOrder := DFSPreOrder(bst.Root)
	fmt.Println("DFS - PreOrder:", *dfsPreOrder)

	t.Error()
}

func TestFoo(t *testing.T) {
	input := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}
	out := AverageOfSubArrayOfSizeK(input, 5)

	fmt.Println("OUT:", out)
	t.Error()
}

func TestBar(t *testing.T) {
	input := []int{4, 2, 2, 7, 8, 1, 2, 8, 10}
	out := SmallestSubarrayGivenSum(18, input)

	fmt.Println("OUT:", out)
	t.Error()
}
