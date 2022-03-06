package linkedlist

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_AddAtBeg_Doubly(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		manipulate func(l *Doubly) *Doubly
		want       []interface{}
	}{
		{
			name: "Case add 2 elems at the beginning successfully",
			manipulate: func(l *Doubly) *Doubly {
				l.AddAtBeg(1)
				l.AddAtBeg(2)
				return l
			},
			want: []interface{}{2, 1},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			list := NewDoubly()
			list = c.manipulate(list)
			got := []interface{}{}
			current := list.Head
			for current != nil {
				got = append(got, current.Val)
				current = current.Next
			}
			require.Equal(t, got, c.want)
			require.Equal(t, len(got), list.Count())
		})
	}
}

func Test_AddAtEnd_Doubly(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		manipulate func(l *Doubly) *Doubly
		want       []interface{}
	}{
		{
			name: "Case add 2 elems at the end successfully",
			manipulate: func(l *Doubly) *Doubly {
				l.AddAtEnd(1)
				l.AddAtEnd(2)
				l.AddAtEnd(3)
				return l
			},
			want: []interface{}{1, 2, 3},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			list := NewDoubly()
			list = c.manipulate(list)
			got := []interface{}{}
			current := list.Head
			for current != nil {
				got = append(got, current.Val)
				current = current.Next
			}
			require.Equal(t, got, c.want)
			require.Equal(t, len(got), list.Count())
		})
	}
}

func Test_DelAtBeg_Doubly(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		manipulate func(l *Doubly) interface{}
		want       interface{}
	}{
		{
			name: "Case del the first elem successfully",
			manipulate: func(l *Doubly) interface{} {
				l.DelAtBeg()
				return l
			},
			want: []interface{}{2, 3},
		},
		{
			name: "Case return -1 due to nil list",
			manipulate: func(l *Doubly) interface{} {
				l.Head = nil
				return l.DelAtBeg()
			},
			want: -1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			list := NewDoubly()
			list.AddAtEnd(1)
			list.AddAtEnd(2)
			list.AddAtEnd(3)
			result := c.manipulate(list)
			switch result.(type) {
			case *Doubly:
				got := []interface{}{}
				current := list.Head
				for current != nil {
					got = append(got, current.Val)
					current = current.Next
				}
				require.Equal(t, got, c.want)
				require.Equal(t, len(got), list.Count())
			default:
				require.Equal(t, result, c.want)
			}

		})
	}
}

func Test_DelAtEnd_Doubly(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		manipulate func(l *Doubly) interface{}
		want       interface{}
	}{
		{
			name: "Case del the last elem successfully",
			manipulate: func(l *Doubly) interface{} {
				l.DelAtEnd()
				l.DelAtEnd()
				l.DelAtEnd()
				return l
			},
			want: []interface{}{},
		},
		{
			name: "Case return -1 due to nil list",
			manipulate: func(l *Doubly) interface{} {
				l.Head = nil
				return l.DelAtEnd()
			},
			want: -1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			list := NewDoubly()
			list.AddAtEnd(1)
			list.AddAtEnd(2)
			list.AddAtEnd(3)
			result := c.manipulate(list)
			switch result.(type) {
			case *Doubly:
				got := []interface{}{}
				current := list.Head
				for current != nil {
					got = append(got, current.Val)
					current = current.Next
				}
				require.Equal(t, got, c.want)
				require.Equal(t, len(got), list.Count())
			default:
				require.Equal(t, result, c.want)
			}

		})
	}
}

func Test_Reverse_Doubly(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		manipulate func(l *Doubly) *Doubly
		want       []interface{}
	}{
		{
			name: "Case reverse successfully",
			manipulate: func(l *Doubly) *Doubly {
				l.Reverse()
				return l
			},
			want: []interface{}{2, 4, 3, 5, 1},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			list := NewDoubly()
			list.AddAtEnd(1)
			list.AddAtEnd(5)
			list.AddAtEnd(3)
			list.AddAtEnd(4)
			list.AddAtEnd(2)
			list = c.manipulate(list)
			got := []interface{}{}
			current := list.Head
			for current != nil {
				got = append(got, current.Val)
				current = current.Next
			}
			require.Equal(t, got, c.want)
			require.Equal(t, len(got), list.Count())
		})
	}
}

func Test_ReverseSubList_Doubly(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		manipulate func(l *Doubly) interface{}
		want       interface{}
	}{
		{
			name: "Case reverse successfully",
			manipulate: func(l *Doubly) interface{} {
				l.ReverseSublist(2, 5)
				return l
			},
			want: []interface{}{1, 2, 4, 3, 5, 9, 6},
		},
		{
			name: "Case reverse fail due to left > right",
			manipulate: func(l *Doubly) interface{} {
				return l.ReverseSublist(5, 2)
			},
			want: errors.New("left boundary must smaller than right"),
		},
		{
			name: "Case reverse fail due to left < 0",
			manipulate: func(l *Doubly) interface{} {
				return l.ReverseSublist(-1, 5)
			},
			want: errors.New("left boundary starts from the first node"),
		},
		{
			name: "Case reverse fail due to right > list.length",
			manipulate: func(l *Doubly) interface{} {
				return l.ReverseSublist(3, 9)
			},
			want: errors.New("right boundary cannot be greater than the length of the linked list"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			list := NewDoubly()
			list.AddAtEnd(1)
			list.AddAtEnd(5)
			list.AddAtEnd(3)
			list.AddAtEnd(4)
			list.AddAtEnd(2)
			list.AddAtEnd(9)
			list.AddAtEnd(6)
			result := c.manipulate(list)
			switch result.(type) {
			case *Doubly:
				got := []interface{}{}
				current := list.Head
				for current != nil {
					got = append(got, current.Val)
					current = current.Next
				}
				require.Equal(t, got, c.want)
				require.Equal(t, len(got), list.Count())
			default:
				require.Equal(t, result, c.want)
			}
		})
	}
}
