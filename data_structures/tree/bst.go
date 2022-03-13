package tree

// BST defines a binary search tree structure
type BST struct {
	Root *Node
}

// NewBST returns an empty BST
func NewBST() *BST {
	return &BST{}
}

// Insert inserts val into the current structure of the BST and return val, -1 is returned if that BST is invalid.
func (t *BST) Insert(val int) interface{} {
	if t == nil {
		return -1
	}

	n := NewNode(val)
	if t.Root == nil {
		t.Root = n
		return val
	}

	cur := t.Root
	for {
		if n.Val < cur.Val {
			// Go to the left of current Node
			if cur.Left == nil {
				// if the left child of the current node is not exist, we place the new node here
				cur.Left = n
				return val
			}
			cur = cur.Left
			continue
		} else if n.Val > cur.Val {
			// Go to the right of current Node
			if cur.Right == nil {
				// if the right child of the current node is not exist, we place the new node here
				cur.Right = n
				return val
			}
			cur = cur.Right
			continue
		}
	}

}

/*
						  9
				   4             20
			   1       6     15      170

	BFS : [9, 4, 20, 1, 6, 15, 170]
	DFS : [9, 4, 1, 6, 20, 15, 170]
*/

func (t *BST) BFS_Iterative() interface{} {
	cur := t.Root
	result := []int{}
	queue := []*Node{}

	queue = append(queue, cur)

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		cur = next
		result = append(result, cur.Val)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}

	return result
}

func BFS_Recursive(queue []*Node, list []int) interface{} {
	if len(queue) == 0 {
		return list
	}

	item := queue[0]
	queue = queue[1:]
	cur := item
	list = append(list, cur.Val)
	if cur.Left != nil {
		queue = append(queue, cur.Left)
	}
	if cur.Right != nil {
		queue = append(queue, cur.Right)
	}

	return BFS_Recursive(queue, list)
}

func (t *BST) DFS() interface{} {

	return nil
}
