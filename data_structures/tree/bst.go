package tree

import "math"

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

func (t *BST) LookUp(val int) interface{} {
	if t.Root == nil {
		return false
	}

	cur := t.Root
	for cur != nil {
		if val < cur.Val {
			cur = cur.Left
		} else if val > cur.Val {
			cur = cur.Right
		} else if val == cur.Val {
			return cur.Val
		}
	}

	return -1
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

func DFSInOrder(root *Node) *[]int {
	return traverseInOrder(root, &[]int{})
}

func DFSPreOrder(root *Node) *[]int {
	return traversePreOrder(root, &[]int{})
}

func DFSPostOrder(root *Node) *[]int {
	return traversePostOrder(root, &[]int{})
}

func traverseInOrder(cur *Node, result *[]int) *[]int {
	if cur.Left != nil {
		traverseInOrder(cur.Left, result)
	}

	*result = append(*result, cur.Val)
	if cur.Right != nil {
		traverseInOrder(cur.Right, result)
	}
	return result
}

func traversePreOrder(cur *Node, result *[]int) *[]int {
	*result = append(*result, cur.Val)
	if cur.Left != nil {
		traversePreOrder(cur.Left, result)
	}

	if cur.Right != nil {
		traversePreOrder(cur.Right, result)
	}
	return result

}

func traversePostOrder(cur *Node, result *[]int) *[]int {
	if cur.Left != nil {
		traversePostOrder(cur.Left, result)
	}

	if cur.Right != nil {
		traversePostOrder(cur.Right, result)
	}
	*result = append(*result, cur.Val)
	return result
}

// func AverageOfSubArrayOfSizeK(a []int, k int) []float64 {
// 	l := len(a)

// 	result := make([]float64, l-k+1)

// 	for i := 0; i < l-k+1; i++ {
// 		sum := 0
// 		for j := i; j < i+k; j++ {
// 			sum += a[j]
// 		}
// 		result[i] = float64(sum) / float64(k)
// 	}

// 	return result
// }

func AverageOfSubArrayOfSizeK(a []int, k int) []float64 {
	l := len(a)
	result := make([]float64, l-k+1)
	windowStart := 0
	windowSum := 0

	for windowEnd := 0; windowEnd < k; windowEnd++ {
		windowSum += a[windowEnd]
	}

	for windowEnd := k; windowEnd < l; windowEnd++ {
		result[windowStart] = float64(windowSum) / float64(k)
		windowSum += a[windowEnd]
		windowSum -= a[windowStart]
		windowStart++
		result[windowStart] = float64(windowSum) / float64(k)
	}

	return result
}

// [4, 2, 1, 7, 8, 1, 2, 8, 1, 0]
func MaxSumSubArray(a []int, k int) int {
	currentRunningSum := 0
	maxValue := 0

	for i := 0; i < len(a); i++ {
		currentRunningSum += a[i]
		if i >= k-1 {
			if maxValue < currentRunningSum {
				maxValue = currentRunningSum
			}
			currentRunningSum -= a[i-(k-1)]
		}
	}

	return maxValue
}

func SmallestSubarrayGivenSum(target int, input []int) int {
	minWindowSize := math.Inf(1)
	currentWindowSum := 0
	windowStart := 0

	for windowEnd := 0; windowEnd < len(input); windowEnd++ {
		currentWindowSum += input[windowEnd]
		for currentWindowSum >= target {
			if minWindowSize > float64(windowEnd-windowStart+1) {
				minWindowSize = float64(windowEnd - windowStart + 1)
			}
			currentWindowSum -= input[windowStart]
			windowStart++
		}
	}

	if minWindowSize == math.Inf(1) {
		return -1
	}

	return int(minWindowSize)
}
