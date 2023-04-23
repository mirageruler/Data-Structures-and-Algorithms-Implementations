package main

type fiboHeapNode struct {
	id          int
	distance    int
	degree      int
	child       *fiboHeapNode
	left, right *fiboHeapNode
	parent      *fiboHeapNode
	marked      bool
}

type FibonacciHeap struct {
	minNode *fiboHeapNode
	count   int
}

func (h *FibonacciHeap) Insert(id, distance int) *fiboHeapNode {
	newNode := &fiboHeapNode{
		distance: distance,
		id:       id,
		left:     nil,
		right:    nil,
	}
	newNode.left = newNode
	newNode.right = newNode

	if h.minNode == nil {
		h.minNode = newNode
	} else {
		h.mergeLists(h.minNode, newNode)
		if newNode.distance < h.minNode.distance {
			h.minNode = newNode
		}
	}
	h.count++
	return newNode
}

func (h *FibonacciHeap) FindMin() *fiboHeapNode {
	return h.minNode
}

func (h *FibonacciHeap) ExtractMin() *fiboHeapNode {
	minNode := h.FindMin()
	h.DeleteMin()
	return minNode
}

func (h *FibonacciHeap) DeleteMin() {
	if h.minNode == nil {
		return
	}

	// add children of the min node to the root list
	for child := h.minNode.child; child != nil; child = child.right {
		h.minNode.parent = nil
		h.mergeLists(h.minNode, child)
	}

	// remove the min node from the root list
	h.removeNode(h.minNode)

	// consolidate trees with the same degree
	degreeMap := make(map[int]*fiboHeapNode)
	for x := h.minNode; x != nil; x = x.right {
		for degreeMap[x.degree] != nil {
			y := degreeMap[x.degree]
			if x.distance > y.distance {
				x, y = y, x
			}
			h.link(y, x)
			degreeMap[x.degree-1] = nil
		}
		degreeMap[x.degree] = x
	}

	// find new minimum node
	h.minNode = nil
	for _, node := range degreeMap {
		if node == nil {
			continue
		}
		if h.minNode == nil || node.distance < h.minNode.distance {
			h.minNode = node
		}
	}

	h.count--
}

func (h *FibonacciHeap) DecreaseKey(node *fiboHeapNode, newDistance int) {
	if newDistance > node.distance {
		return
	}
	node.distance = newDistance
	parent := node.parent
	if parent != nil && node.distance < parent.distance {
		h.cut(node, parent)
		h.cascadingCut(parent)
	}
	if node.distance < h.minNode.distance {
		h.minNode = node
	}
}

func (h *FibonacciHeap) Len() int {
	return h.count
}

func (h *FibonacciHeap) mergeLists(a, b *fiboHeapNode) {
	a.right.left = b
	b.right.left = a
	a.right, b.right = b, a.right
}

func (h *FibonacciHeap) removeNode(node *fiboHeapNode) {
	node.left.right = node.right
	node.right.left = node.left
}

func (h *FibonacciHeap) link(child, parent *fiboHeapNode) {
	h.removeNode(child)
	child.parent = parent
	child.marked = false
	if parent.child == nil {
		parent.child = child
		child.left, child.right = child, child
	} else {
		h.mergeLists(parent.child, child)
	}
	parent.degree++
}

func (h *FibonacciHeap) cut(child, parent *fiboHeapNode) {
	h.removeNode(child)
	parent.degree--
	if parent.child == child {
		parent.child = child.right
	}
	child.parent = nil
	child.marked = false
	h.mergeLists(h.minNode, child)
}

func (h *FibonacciHeap) cascadingCut(node *fiboHeapNode) {
	parent := node.parent
	if parent != nil {
		if !node.marked {
			node.marked = true
		} else {
			h.cut(node, parent)
			h.cascadingCut(parent)
		}
	}
}
