package main

import "fmt"

/*
AVLNode defines a node in the AVL tree structure

	Time Complexity of almost operations associated: O(log n)
	Auxiliary Space: O(n)
*/
type AVLNode struct {
	left  *AVLNode
	right *AVLNode
	// height counts nodes (not edges)
	height int // for balancing the tree whenever a node's height violate the invariant of the AVL tree
	key    int
}

type AVLTree struct {
	root *AVLNode
}

// NewAVLTree returns an empty AVL Tree
func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func (n *AVLNode) Height() int {
	if n == nil {
		return 0
	}
	return n.height
}

func newAVLNode(key int) *AVLNode {
	return &AVLNode{
		left:   nil,
		right:  nil,
		key:    key,
		height: 1,
	}
}

func rightRotate(y *AVLNode) *AVLNode {
	// extract the current node's left child
	x := y.left
	// extract the right child of the current node's left child
	leftRightChild := x.right
	// the right child of the current node's left child is now the current node
	x.right = y
	// the current node's left child is now the right child of itself left child
	y.left = leftRightChild

	// it's important to update y's height (our input) first, because y now becomes a child of x
	y.updateHeight()
	// then update x's height
	x.updateHeight()

	// return the left child of the rotated node
	return x
}

func leftRotate(x *AVLNode) *AVLNode {
	// extract the current node's right child
	y := x.right
	// extract the left child of the current node's right child
	rightLeftChild := y.left
	// the left child of the current node's right child is now the current node
	y.left = x
	// the current node's right child is now the left child of itself right child
	x.right = rightLeftChild

	// it's important to update x's height (our input) first, because x now becomes a child of y
	x.updateHeight()
	// then update y's height
	y.updateHeight()

	// return the right child of the rotated node
	return y
}

// rotateForInsertion recursively perform necessary rotations to keep the current AVL Tree balance after an insertion.
func (n *AVLNode) rotateForInsertion(key int) *AVLNode {
	// after update the height of current node, get the current balance factor of the current node
	balFtr := n.getBalanceFactor()

	if balFtr > 1 { // the node is unbalance because of the left child (because `balance = left-right` so left's height is higher than right's height)
		if key < n.left.key { // case 1: the left child skewed on the left
			/*
								[30]
						   		/
							[27]				---> "TOTALLY LEFT SKEWED"
					    	/
					 	[24]

				- we rotate the current node to the right so this sub-tree becomes like this

							[27]
					       /	\				---> "BALANCE"
					 	[24]	[30]
			*/
			return rightRotate(n)
		} else if key > n.left.key { // case 2: the left child skewed on the right
			/*
								[30]
					   			/
							[24]				---> "LEFT_RIGHT SKEWED"
								\
				 				[27]
			*/
			n.left = leftRotate(n.left) // we need to rotate the left child of the current node (which is [24] in this ex) to the left so this sub-tree become like our 1st case
			/*
								[30]
					   			/
							[27]				---> "TOTALLY LEFT SKEWED"
				    		/
				 		[24]
			*/
			// now it's become our 1st case, let's do the same here
			return rightRotate(n)
		}
	} else if balFtr < -1 { // the node is unbalance because of the right child (because `balance = left-right` so right's height is higher than left's height)
		if key > n.right.key { // case 3: the right child skewed on the right
			/*
						[30]
						   	\
							[37]				---> "TOTALLY RIGHT SKEWED"
					    		\
					 			[39]

				- we rotate the current node to the left so this sub-tree becomes like this

							[37]
					       /	\				---> "BALANCE"
					 	[30]	[39]
			*/
			return leftRotate(n)
		} else if key < n.right.key { // case 4: the right child skewed on the left
			/*
							[30]
					   			\
								[39]			---> "RIGHT_LEFT SKEWED"
				    			/
				 			[37]
			*/
			n.right = rightRotate(n.right) // we need to rotate the right child of the current node (which is [39] in this ex) to the right so this sub-tree become like our 3rd case
			/*
						[30]
					   		\
							[37]				---> "TOTALLY RIGHT SKEWED"
				    			\
				 				[39]
			*/
			// now it's become our 3rd case, let's do the same here
			return leftRotate(n)
		}
	}

	return n
}

// rotateForDeletion recursively perform necessary rotations to keep the current AVL Tree balance after a deletion.
func (n *AVLNode) rotateForDeletion() *AVLNode {
	// after update the height of current node, get the current balance factor of the current node
	balFtr := n.getBalanceFactor()

	if balFtr > 1 { // the node is unbalance because of the left child (because `balance = left-right` so left's height is higher than right's height)
		// Check by getBalanceFactor from the current node's left child.
		// Not like the `InsertNode` method that we can compare the entered values, because the node we are looking for is already deleted
		if n.left.getBalanceFactor() >= 0 { // case 1: the left child skewed on the left
			/*
								[30]
						   		/
							[27]				---> "TOTALLY LEFT SKEWED"
					    	/
					 	[24]

				- we rotate the current node to the right so this sub-tree becomes like this

							[27]
					       /	\				---> "BALANCE"
					 	[24]	[30]
			*/
			return rightRotate(n)
		} else { // case 2: the left child skewed on the right
			/*
								[30]
					   			/
							[24]				---> "LEFT_RIGHT SKEWED"
								\
				 				[27]
			*/
			n.left = leftRotate(n.left) // we need to rotate the left child of the current node (which is [24] in this ex) to the left so this sub-tree become like our 1st case
			/*
								[30]
					   			/
							[27]				---> "TOTALLY LEFT SKEWED"
				    		/
				 		[24]
			*/
			// now it's become our 1st case, let's do the same here
			return rightRotate(n)
		}
	}

	if balFtr < -1 { // the node is unbalance because of the right child (because `balance = left-right` so right's height is higher than left's height)
		// Check by getBalanceFactor from the current node's right child.
		// Not like the `InsertNode` method that we can compare the entered values, because the node we are looking for is already deleted
		if n.right.getBalanceFactor() <= 0 { // case 3: the right child skewed on the right
			/*
						[30]
						   	\
							[37]				---> "TOTALLY RIGHT SKEWED"
					    		\
					 			[39]

				- we rotate the current node to the left so this sub-tree becomes like this

							[37]
					       /	\				---> "BALANCE"
					 	[30]	[39]
			*/
			return leftRotate(n)
		} else { // case 4: the right child skewed on the left
			/*
							[30]
					   			\
								[39]			---> "RIGHT_LEFT SKEWED"
				    			/
				 			[37]
			*/
			n.right = rightRotate(n.right) // we need to rotate the right child of the current node (which is [39] in this ex) to the right so this sub-tree become like our 3rd case
			/*
						[30]
					   		\
							[37]				---> "TOTALLY RIGHT SKEWED"
				    			\
				 				[39]
			*/
			// now it's become our 3rd case, let's do the same here
			return leftRotate(n)
		}
	}

	return n
}

// getBalanceFactor calculate the differrence between the left & right child of the input `n` *AVLNode by using the formula `balF = n.left - n.right`
func (n *AVLNode) getBalanceFactor() int {
	if n == nil {
		return 0
	}
	return n.left.Height() - n.right.Height()
}

// updateHeight update height of the current node
func (n *AVLNode) updateHeight() {
	// compare the maximum height of the children + its own height
	n.height = max(n.left.Height(), n.right.Height()) + 1
}

// FindMinimum will return a pointer to an AVLNode that has the lowest key traversed from the the input `n`
func FindMinimum(n *AVLNode) *AVLNode {
	if n == nil {
		return n
	}
	cur := n
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

// InsertNode insert a new AVLNode with given `key` into the current AVLTree structure and will try to maintain the "balance" property of the tree after that which may or may not change the tree structure
func (t *AVLTree) InsertNode(key int) *AVLNode {
	if t.root == nil {
		t.root = newAVLNode(key)
		return t.root
	}

	return insertNode(t.root, key)
}

// insertNode is a recursive helper func for `InserNode`
func insertNode(n *AVLNode, key int) *AVLNode {
	// base case
	if n == nil {
		return newAVLNode(key)
	}

	if key < n.key {
		// key is less than the current node's key, recursively try to insert the node into the current left sub-tree
		n.left = insertNode(n.left, key)
	} else if key > n.key {
		// key is larger than the current node's key, recursively try to insert the node into the current right sub-tree
		n.right = insertNode(n.right, key)
	} else {
		// there's already a node with that given `key`, return and do nothing more.
		return n
	}

	// we've successfully added a new node with the given `key` into our current AVL Tree structure (it's either inserted as the left or the right child of a node in the tree because of these 2 recursive calls above),
	// when going all the way up, we will need to re-calculate the height of each node by using this formula
	n.updateHeight()

	return n.rotateForInsertion(key)
}

// DeleteNode find and delete an AVLNode with the specified `key` from the current AVLTree structure and will try to maintain the "balance" property of the tree after that which may or may not change the tree structure
func (t *AVLTree) DeleteNode(key int) *AVLNode {
	if t.root == nil {
		return t.root
	}

	return deleteNode(t.root, key)
}

// deleteNode is a recursive helper func for `DeleteNode`
func deleteNode(n *AVLNode, key int) *AVLNode {
	// the base case
	if n == nil {
		return n
	}

	if key < n.key {
		// key is less than the current node's key, recursively try to delete the node with the given 'key' from the current left sub-tree
		n.left = deleteNode(n.left, key)
	} else if key > n.key {
		// key is larger than the current node's key, recursively try to delete the node with the given 'key' from the current right sub-tree
		n.right = deleteNode(n.right, key)
	} else {
		// we finally found the node the delete from the tree
		if n.left == nil || n.right == nil { // if the node's left child OR right child is empty (both left & right empty will also fall into this case)
			caseLeftOrRightChildIsEmpty(n)
		} else { // if the node that is going to be deleted has both left and right child
			// find the minimum node in the right sub-tree of this node
			tmp := FindMinimum(n.right)
			// replace the current node's key with the node we've found right above
			n.key = tmp.key
			// recursively delete that minimum node from the right sub-tree of the current node (because it is already placed at the current node by the logic of the line above)
			n.right = deleteNode(n.right, tmp.key)
		}
	}
	// for the case when we've delete n
	if n == nil {
		return n
	}

	// we've successfully delete a node with the given `key` from our current AVL Tree structure
	// when going all the way up, we will need to re-calculate the height of each node by using this formula
	n.updateHeight()

	return n.rotateForDeletion()
}

// caseLeftOrRightChildIsEmpty handle deletion for a node in case it's left child or right child is empty or both of them are empty
func caseLeftOrRightChildIsEmpty(n *AVLNode) {
	tmp := n.left
	if tmp == nil {
		tmp = n.right
	}
	if tmp == nil {
		n = nil
	} else {
		*n = *tmp
	}
}

// Find find the *AVLNode with given `key`
func (t AVLTree) Find(key int) *AVLNode {
	if t.root == nil {
		return t.root
	}
	node := t.root

	return find(node, key)
}

// find is a recursive helper func for the `Find` method
func find(n *AVLNode, key int) *AVLNode {
	if n == nil {
		return nil
	}

	switch nKey := n.key; {
	case key == nKey:
		return n
	case key > nKey:
		return find(n.right, key)
	case key < nKey:
		return find(n.left, key)
	}

	return nil
}

// InOrder prints the tree elements into the screen in an `in-order` order
func (t AVLTree) InOrder() {
	var inorder func(n *AVLNode)
	inorder = func(n *AVLNode) {
		if n == nil {
			return
		}
		inorder(n.left)
		fmt.Printf("%d ", n.key)
		inorder(n.right)
	}
	inorder(t.root)
	fmt.Println()
}

// PreOrder prints the tree elements into the screen in a `pre-order` order
func (t AVLTree) PreOrder() {
	var preorder func(n *AVLNode)
	preorder = func(n *AVLNode) {
		if n == nil {
			return
		}
		fmt.Printf("%d ", n.key)
		preorder(n.left)
		preorder(n.right)
	}
	preorder(t.root)
	fmt.Println()
}

// PostOrder prints the tree elements into the screen in a `post-order` order
func (t AVLTree) PostOrder() {
	var postorder func(n *AVLNode)
	postorder = func(n *AVLNode) {
		if n == nil {
			return
		}
		postorder(n.left)
		postorder(n.right)
		fmt.Printf("%d ", n.key)
	}
	postorder(t.root)
	fmt.Println()
}

// Prints the AVL tree
func PrintTree(root *AVLNode, indent string, last bool) {
	if root != nil {
		fmt.Print(indent)
		if last {
			fmt.Print("R----")
			indent += "   "
		} else {
			fmt.Print("L----")
			indent += "|  "
		}
		fmt.Println(root.key)
		PrintTree(root.left, indent, false)
		PrintTree(root.right, indent, true)
	}
}

func (t AVLTree) MinNodesInCurrentHeight() int {
	var (
		res         int
		minNumOnAVL func(height, a, b int) int
	)

	minNumOnAVL = func(height, a, b int) int {
		// Base Conditions
		if height == 0 {
			return 1
		}
		if height == 1 {
			return b
		}

		// Tail Recursive Call
		return minNumOnAVL(height-1, b, a+b+1)
	}
	/*
		 Explanation:
		 - For height = 0, we can only have a single node in an AVL tree, i.e. n(0) = 1	(Our first base case)
		 - For height = 1, we can have a minimum of two nodes in an AVL tree, i.e. n(1) = 2	(Our second base case)
		 - Now for any height ‘h’, root will have two subtrees (left and right). Out of which one has to be of height h-1 and other of h-2. [root node excluded]
		==> use the formula base on the recurrence relation `n(h) = 1 + n(h-1) + n(h-2)`
	*/
	minNumOnAVL(t.root.height, 1, 2)
	return res
}

func main() {
	// Creating AVL tree and
	// inserting data in it
	avlTree := NewAVLTree()
	avlTree.InsertNode(33)
	avlTree.InsertNode(13)
	avlTree.InsertNode(53)
	avlTree.InsertNode(9)
	avlTree.InsertNode(21)
	avlTree.InsertNode(61)
	avlTree.InsertNode(8)
	avlTree.InsertNode(11)

	// Printing AVL Tree
	PrintTree(avlTree.root, "", true)

	// Deleting a node from AVL Tree
	avlTree.DeleteNode(13)
	fmt.Println("After deleting ")
	PrintTree(avlTree.root, "", true)

	fmt.Printf("FIND %d\n", avlTree.Find(9).key)
	fmt.Printf("Minimum number of nodes for height %d is %d \n", avlTree.root.height, avlTree.MinNodesInCurrentHeight())
}
