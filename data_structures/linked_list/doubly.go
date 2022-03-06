package linkedlist

import "errors"

// Doubly structure with just the Head node
type Doubly struct {
	Head *Node
}

// NewDoubly returns a pointer to Doubly struct with zeroed value
func NewDoubly() *Doubly {
	return &Doubly{}
}

// AddAtBeg adds a node to the beginning of the list
func (l *Doubly) AddAtBeg(val interface{}) {
	n := NewNode(val)
	n.Next = l.Head

	if l.Head != nil {
		l.Head.Prev = n
	}

	l.Head = n
}

// AddAtEnd adds a node at the end of the list
func (l *Doubly) AddAtEnd(val interface{}) {
	n := NewNode(val)
	if l.Head == nil {
		l.Head = n
		return
	}

	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = n
	n.Prev = cur
}

// DelAtBeg deletes the beginning node (at Head) of the list and return the deleted value, -1 is returned if the list is nil.
func (l *Doubly) DelAtBeg() interface{} {
	if l.Head == nil {
		return -1
	}

	cur := l.Head
	l.Head = cur.Next

	if l.Head != nil {
		l.Head.Prev = nil
	}

	return cur.Val
}

// DelAtEnd deletes the last node (at the tail) of the list and return the deleted value, -1 is returned if the list is nil.
func (l *Doubly) DelAtEnd() interface{} {
	// no item
	if l.Head == nil {
		return -1
	}

	// only one item
	if l.Head.Next == nil {
		return l.DelAtBeg()
	}

	// more than one, go to the SECOND last
	cur := l.Head
	for cur.Next.Next != nil {
		cur = cur.Next
	}

	retVal := cur.Next.Val
	cur.Next = nil

	return retVal
}

// Count returns the number of elements in the list.
func (l *Doubly) Count() interface{} {
	var ctr int

	for cur := l.Head; cur != nil; cur = cur.Next {
		ctr++
	}

	return ctr
}

// Reverse reverse the list
func (l *Doubly) Reverse() {
	var prev, next *Node
	cur := l.Head

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		cur.Prev = next
		prev = cur
		cur = next
	}

	l.Head = prev
}

// ReverseSublist reverse the list from the a-th (on the left side) node to the b-th(on the right side) node.
func (l *Doubly) ReverseSublist(left, right int) error {
	if err := l.checkRangeFromIndex(left, right); err != nil {
		return err
	}

	// Use this approach because it will perfectly solve the case left = 1
	tmpNode := NewNode(-1)
	tmpNode.Next = l.Head
	pre := tmpNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		//prev := cur.Prev
		cur.Next = next.Next // 12 -> 3
		next.Next.Prev = cur // 3  -> 12
		next.Next = pre.Next // 8  -> 12
		pre.Next.Prev = next // 12 -> 8
		next.Prev = pre      // 8  -> 5
		pre.Next = next      // 5  -> 8
	}

	l.Head = tmpNode.Next

	return nil
}

func (l *Doubly) checkRangeFromIndex(left, right int) error {
	if left > right {
		return errors.New("left boundary must smaller than right")
	} else if left < 1 {
		return errors.New("left boundary starts from the first node")
	} else if right > l.Count().(int) {
		return errors.New("right boundary cannot be greater than the length of the linked list")
	}

	return nil
}
