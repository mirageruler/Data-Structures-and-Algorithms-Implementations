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

	/*
					  9
			   4             20
		   1       6     15      170

	*/

	//bfs := bst.BFS_Iterative()
	bfs := BFS_Recursive([]*Node{bst.Root}, []int{})
	fmt.Println("BFS: ", bfs)

	bs, _ := json.Marshal(bst)
	fmt.Println("BST: ", string(bs))

	t.Error()
}
