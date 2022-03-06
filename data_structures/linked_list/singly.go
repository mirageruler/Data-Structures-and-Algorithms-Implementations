package linkedlist

import (
	"errors"
	"fmt"
	"io"
)

type Singly struct {
	length int

	// Note that Node here holds both Next and Prev Node
	// However only the Next Node is used in Singly methods.
	Head *Node
}

// NewSingly returns a new instance of a linked list.
func NewSingly() *Singly {
	return &Singly{}
}

// AddAtBeg adds a new singly Node with given value at the beginning of the list.
func (l *Singly) AddAtBeg(val interface{}) {
	n := NewNode(val)
	n.Next = l.Head

	l.Head = n
	l.length++
}

// AddAtEnd adds a new singly Node with given value at the end of the list.
func (l *Singly) AddAtEnd(val interface{}) {
	n := NewNode(val)

	if l.Head == nil {
		l.Head = n
		l.length++
		return
	}

	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = n
	l.length++
}

// DelAtBeg deletes the singly Node at the head(beginning) of the list and returns its value. -1 is returned if the list is empty.
func (l *Singly) DelAtBeg() interface{} {
	if l.Head == nil {
		return -1
	}

	cur := l.Head
	l.Head = cur.Next
	l.length--

	return cur.Val
}

// DelAtEnd deletes the singly Node at the tail(end) of the list and returns its value. -1 is returned if the list is empty
func (l *Singly) DelAtEnd() interface{} {
	if l.Head == nil {
		return -1
	}

	cur := l.Head
	for ; cur.Next.Next != nil; cur = cur.Next {
	}

	cur.Next = nil
	l.length--

	return cur.Val
}

// Count returns the current size of the list.
func (l *Singly) Count() int {
	return l.length
}

// Display prints to io.writer the elements of the list.
func (l *Singly) Display(r io.Writer) (int, error) {
	var bw int
	for cur := l.Head; cur != nil; cur = cur.Next {
		n, err := fmt.Fprint(r, cur.Val)
		if err != nil {
			return n, err
		}
		if cur.Next != nil {
			n, err := fmt.Fprint(r, " -> ")
			if err != nil {
				return n, err
			}
			bw += n
		}
		bw += n
	}

	return bw, nil
}

// Reverse reverses the list.
func (l *Singly) Reverse() {
	var prev, Next *Node
	cur := l.Head

	// Changing the pointer destination of each node reversely
	for cur != nil {
		Next = cur.Next
		cur.Next = prev
		prev = cur
		cur = Next
		// Below is alternative, but harder to read and need to deeply understand Golang multiple assignments that
		// the assignments order is not important.
		// Although the running process manipulate these variables slightly different but the end result is the same.
		// cur.Next, prev, cur = prev, cur, cur.Next
	}

	// The reversed list IS STORED in prev, so we need to assign it to the original list's Head!
	l.Head = prev
}

// explanations for Reverse():
// -	original: 5 -> 12 -> 8 -> 3 -> 6 -> 7 -> 2 -> nil
// -	after 1st iteration of the loop: nil <- 5 [] 12 -> 8 -> 3 -> 6 -> 7 -> 2 -> nil
// -	after 2nd iteration of the loop: nil <- 5 <- 12 [] 8 -> 3 -> 6 -> 7 -> 2 -> nil
// -	after 3rd iteration of the loop: nil <- 5 <- 12 <- 8 [] 3 -> 6 -> 7 -> 2 -> nil
// - 	So on...

// ReverseSublist reverses the list from the a-th(on the left side) node to the b-th(on the right side) node.
func (l *Singly) ReverseSublist(left, right int) error {
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
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	l.Head = tmpNode.Next

	return nil
}

/* explanations for ReverseSublist(left, right int):
 - 	Let's say ReverseSublist(2, 6)
 -	original: 5 -> 12 -> 8 -> 3 -> 6 -> 7  -> 2 -> nil
 - 	expect:   5 -> 7  -> 6 -> 3 -> 8 -> 12 -> 2 -> nil
 - 	tmpNode{val: -1, Next: Node{val:5, Node{val:12, Node{....}}}}
 -	pre = tmpNode, so we have some thing like this -1 -> 5 -> 12 -> 8 -> 3 -> 6 -> 07 -> 2 -> nil
 -	run the 1st loop to traverse pre, we got pre = Node{val: 5, Node{val: 12, Node{...}}}
 -	cur := pre.Next ==> cur = Node{val: 12, Node{val: 8, Node{...}}}
 -	run the 2nd loop to traverse cur and pre:
 		- during the 1st iteration of the 2nd loop: we make 12 -> 3, we make 8 -> 12, we make 5 -> 8. So in the end we got 5 -> 8 -> 12 -> 3  -> 6  -> 7 ->  2 -> nil
		- during the 2nd iteration of the 2nd loop:
		- .......

		l: 5 -> 12 -> 8 -> 3 -> 6 -> 7 -> 2 -> nil
		Analysing the 1st iteration of the 1st loop: (cur = 12 & cur.Next = 8, pre = 5 & pre.Next = 12)
		next := cur.Next 			==> next = 8
 		cur.Next = next.Next = 3 	==> Make 12 -> 3
 		next.Next = pre.Next = 12	==> Make 8  -> 12
		pre.Next = next = 8			==> Make 5  -> 8
		So we got:  5 -> 8 -> 12 -> 3 -> ..... (not manipulated)

		l: 5 -> 8 -> 12 -> 3 -> 6 -> 7 -> 2 -> nil
		Analysing the 2nd iteration of the 2nd loop: (cur = 12 & cur.Next = 3, pre = 5 & pre.Next = 8)
		next := cur.Next 			==> next = 3
		cur.Next = next.Next = 6 	==> Make 12 -> 6
		next.Next = pre.Next = 8	==> Make 3  -> 8
		pre.Next = next = 3			==> Make 5  -> 3
		So we got: 5 -> 3 -> 8 -> 12 -> 6 -> 7 -> 2 -> nil

		l: 5 -> 3 -> 8 -> 12 -> 6 -> 7 -> 2 -> nil
		Analysing the 3rd iteration of the 2nd loop: (cur = 12 & cur.Next = 6, per = 5 & pre.Next = 3)
		next := cur.Next 			==> next = 6
		cur.Next = next.Next = 7	==> Make 12 -> 7
		next.Next = pre.Next = 3 	==> Make 6  -> 3
		pre.Next = next = 6 		==> Make 5  -> 6
		So we got: 5 -> 6 - > 3 -> 8 -> 12 -> 7 -> 2 -> nil

*/

func (l *Singly) checkRangeFromIndex(left, right int) error {
	if left > right {
		return errors.New("left boundary must smaller than right")
	} else if left < 1 {
		return errors.New("left boundary starts from the first node")
	} else if right > l.length {
		return errors.New("right boundary cannot be greater than the length of the linked list")
	}

	return nil
}
