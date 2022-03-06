package linkedlist

import "fmt"

// Cyclic which cycles the linked list in this implementation
type Cyclic struct {
	Size int
	Head *Node
}

// NewCyclic returns a new cyclic with zeroed values for each field.
func NewCyclic() *Cyclic {
	return &Cyclic{}
}

// Inserting the first node is a special case. It will point to itself.
// For other cases, the node will be added to the end of the list (determined by cl.Head.Prev).
// Complexity O(1)
func (l *Cyclic) Add(val interface{}) {
	n := NewNode(val)
	l.Size++

	if l.Head == nil {
		n.Prev = n
		n.Next = n
		l.Head = n
	} else {
		// Node: The order of assigning pointers here is very IMPORTANT!
		n.Prev = l.Head.Prev // new node.Prev             -> currently last node
		n.Next = l.Head      // new node.Next             -> currently first node
		l.Head.Prev.Next = n // currently last node.Next  -> new node
		l.Head.Prev = n      // currently first node.Prev -> new node
	}
}

// Rotate rotates the list p steps.
// This method is interesting for optimization.
// For first optimaztion we must decrease p value
// so that it only ranges from 0 to cl.Size()-1.
// For this we need to use the operation of division modulo,
// be careful if p < 0, if that the case then make it positive.
// This can be done without violating the meaning of p by adding to it
// a multiple of cl.Size(). Now we can decrease p modulo to cl.Size() to
// rotate the list by the minimum number of steps.
// We leverage the fact that moving forward in a circle by p steps
// is the same as moving backward cl.Size()-p steps.
// Therefore, if p > cl.Size()/2, we can just rotate the list by cl.Size()-p steps back.
// Complexity O(n)
func (l *Cyclic) Rotate(p int) {
	if p < 0 {
		// Make p positive without changing the meaning of it as steps to rotate in the cycle
		multiple := l.Size - 1 - p/l.Size
		p += multiple * l.Size
	}

	p %= l.Size // optimizing(decrease) p without changing it's meaning for steps to rotate

	if p > l.Size/2 {
		for i := 0; i < l.Size-p; i++ {
			l.Head = l.Head.Prev
		}
	} else if p == 0 {
		return
	} else {
		for i := 0; i < p; i++ {
			l.Head = l.Head.Next
		}
	}
}

// Delete deletes the current item
func (l *Cyclic) Delete() bool {
	var prevItem, curItem, nextItem *Node

	if l.Size == 0 {
		return false
	}

	curItem = l.Head
	prevItem = curItem.Prev
	nextItem = curItem.Next

	if l.Size == 1 {
		l.Head = nil
	} else {
		l.Head = nextItem
		nextItem.Prev = prevItem
		prevItem.Next = nextItem
	}

	l.Size--

	return true
}

// Destroy destroys all items in the list
func (l *Cyclic) Destroy() {
	for l.Delete() {
		continue
	}
}

// Walk prints to stdout the list body
func (l *Cyclic) Walk() *Node {
	for i := 0; i < l.Size; i++ {
		fmt.Printf("%v \n", l.Head.Val)

		if i != l.Size-1 {
			l.Head = l.Head.Next
		}
	}

	return l.Head
}

// JosephusProblem solves the below problem described in wiki
// https://en.wikipedia.org/wiki/Josephus_problem
func (l *Cyclic) JosephusProblem(k int) int {
	for l.Size > 1 {
		l.Rotate(k)
		l.Delete()
		l.Rotate(-1)
	}

	retVal := l.Head.Val.(int)
	l.Destroy()

	return retVal
}
