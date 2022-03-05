package linkedlist

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_AddAtBeg(t *testing.T) {
	cases := []struct {
		name       string
		manipulate func(l *Singly) *Singly
		want       []interface{}
	}{
		{
			name: "Case add 2 elems at the beginning successfully",
			manipulate: func(l *Singly) *Singly {
				l.AddAtBeg(1)
				l.AddAtBeg(2)
				return l
			},
			want: []interface{}{2, 1},
		},
	}

	for _, c := range cases {
		t.Parallel()
		t.Run(c.name, func(t *testing.T) {
			list := NewSingly()
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

func Test_AddAtEnd(t *testing.T) {
	cases := []struct {
		name       string
		manipulate func(l *Singly) *Singly
		want       []interface{}
	}{
		{
			name: "Case add 2 elems at the end successfully",
			manipulate: func(l *Singly) *Singly {
				l.AddAtEnd(1)
				l.AddAtEnd(2)
				l.AddAtEnd(3)
				return l
			},
			want: []interface{}{1, 2, 3},
		},
	}

	for _, c := range cases {
		t.Parallel()
		t.Run(c.name, func(t *testing.T) {
			list := NewSingly()
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

func Test_DelAtBeg(t *testing.T) {
	cases := []struct {
		name       string
		manipulate func(l *Singly) interface{}
		want       interface{}
	}{
		{
			name: "Case del the first elem successfully",
			manipulate: func(l *Singly) interface{} {
				l.DelAtBeg()
				return l
			},
			want: []interface{}{2, 3},
		},
		{
			name: "Case return -1 due to nil list",
			manipulate: func(l *Singly) interface{} {
				l.Head = nil
				return l.DelAtBeg()
			},
			want: -1,
		},
	}

	for _, c := range cases {
		//t.Parallel()
		t.Run(c.name, func(t *testing.T) {
			list := NewSingly()
			list.AddAtEnd(1)
			list.AddAtEnd(2)
			list.AddAtEnd(3)
			result := c.manipulate(list)
			switch result.(type) {
			case *Singly:
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

func Test_DelAtEnd(t *testing.T) {
	cases := []struct {
		name       string
		manipulate func(l *Singly) interface{}
		want       interface{}
	}{
		{
			name: "Case del the last elem successfully",
			manipulate: func(l *Singly) interface{} {
				l.DelAtEnd()
				return l
			},
			want: []interface{}{1, 2},
		},
		{
			name: "Case return -1 due to nil list",
			manipulate: func(l *Singly) interface{} {
				l.Head = nil
				return l.DelAtEnd()
			},
			want: -1,
		},
	}

	for _, c := range cases {
		//t.Parallel()
		t.Run(c.name, func(t *testing.T) {
			list := NewSingly()
			list.AddAtEnd(1)
			list.AddAtEnd(2)
			list.AddAtEnd(3)
			result := c.manipulate(list)
			switch result.(type) {
			case *Singly:
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

func Test_Display(t *testing.T) {
	list := NewSingly()
	list.AddAtEnd(1)
	list.AddAtEnd(2)
	list.AddAtEnd(3)

	buffer := bytes.Buffer{}

	_, err := list.Display(&buffer)
	got := buffer.String()
	want := "1 -> 2 -> 3"
	require.Nil(t, err)
	require.Equal(t, got, want)
}

func Test_Reverse(t *testing.T) {
	cases := []struct {
		name       string
		manipulate func(l *Singly) *Singly
		want       []interface{}
	}{
		{
			name: "Case reverse successfully",
			manipulate: func(l *Singly) *Singly {
				l.Reverse()
				return l
			},
			want: []interface{}{2, 4, 3, 5, 1},
		},
	}

	for _, c := range cases {
		t.Parallel()
		t.Run(c.name, func(t *testing.T) {
			list := NewSingly()
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

func Test_ReverseSubList(t *testing.T) {
	cases := []struct {
		name       string
		manipulate func(l *Singly) interface{}
		want       interface{}
	}{
		{
			name: "Case reverse successfully",
			manipulate: func(l *Singly) interface{} {
				l.ReverseSublist(2, 5)
				return l
			},
			want: []interface{}{1, 2, 4, 3, 5, 9, 6},
		},
		{
			name: "Case reverse fail due to left > right",
			manipulate: func(l *Singly) interface{} {
				return l.ReverseSublist(5, 2)
			},
			want: errors.New("left boundary must smaller than right"),
		},
		{
			name: "Case reverse fail due to left < 0",
			manipulate: func(l *Singly) interface{} {
				return l.ReverseSublist(-1, 5)
			},
			want: errors.New("left boundary starts from the first node"),
		},
		{
			name: "Case reverse fail due to right > list.length",
			manipulate: func(l *Singly) interface{} {
				return l.ReverseSublist(3, 9)
			},
			want: errors.New("right boundary cannot be greater than the length of the linked list"),
		},
	}

	// 1 5 3 4 2 9 6 ==> 1 2 4 3 5 9 6
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			list := NewSingly()
			list.AddAtEnd(1)
			list.AddAtEnd(5)
			list.AddAtEnd(3)
			list.AddAtEnd(4)
			list.AddAtEnd(2)
			list.AddAtEnd(9)
			list.AddAtEnd(6)
			result := c.manipulate(list)
			switch result.(type) {
			case *Singly:
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
