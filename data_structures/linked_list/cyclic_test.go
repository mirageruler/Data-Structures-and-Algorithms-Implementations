package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func fillListForTesting(cl *Cyclic, n int) {
	for i := 1; i <= n; i++ {
		cl.Add(i)
	}
}

func Test_Add_Cyclic(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		manipulate func(l *Cyclic) *Cyclic
		want       []interface{}
	}{
		{
			name: "Case add 2 elems successfully",
			manipulate: func(cl *Cyclic) *Cyclic {
				cl.Add(1)
				cl.Add(2)
				return cl
			},
			want: []interface{}{1, 2},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cl := NewCyclic()
			cl = c.manipulate(cl)
			got := []interface{}{}
			current := cl.Head
			for i := 0; i < cl.Size; i++ {
				got = append(got, current.Val)
				current = current.Next
			}
			require.Equal(t, got, c.want)
			require.Equal(t, len(got), cl.Size)
		})
	}
}

func TestWalk_Cyclic(t *testing.T) {
	t.Parallel()
	cl := NewCyclic()
	fillListForTesting(cl, 3)

	want := 3
	got := cl.Walk()
	fmt.Printf("got: %#v\n", got)

	if got.Val != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestRotate_Cyclic(t *testing.T) {
	t.Parallel()
	cases := []struct {
		param int
		want  int
	}{
		{1, 2},
		{2, 3},
		{3, 1},
		{-1, 3},
		{-4, 3},
		{-6, 1},
	}

	for idx, c := range cases {
		cl := NewCyclic()
		fillListForTesting(cl, 3)
		cl.Rotate(c.param)
		got := cl.Head.Val
		if got != c.want {
			t.Errorf("got %v, want: %v for test idx = %d", got, c.want, idx)
		}
	}
}

func TestDelete_Cyclic(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		manipulate func(l *Cyclic) bool
		want       interface{}
	}{
		{
			name: "Case del fail due to nil list",
			manipulate: func(cl *Cyclic) bool {
				return cl.Delete()
			},
			want: false,
		},
		{
			name: "Case del success list with 1 elem",
			manipulate: func(cl *Cyclic) bool {
				cl.Add(1)
				return cl.Delete()
			},
			want: true,
		},
		{
			name: "Case del success list with multiple elems",
			manipulate: func(cl *Cyclic) bool {
				cl.Add(1)
				cl.Add(2)
				return cl.Delete()
			},
			want: true,
		},
	}

	for _, c := range cases {
		cl := NewCyclic()
		require.Equal(t, c.manipulate(cl), c.want)
	}
}

func TestDestroy_Cyclic(t *testing.T) {
	t.Parallel()
	cl := NewCyclic()
	fillListForTesting(cl, 100)

	cl.Destroy()

	require.Equal(t, cl.Size, 0)
}

func TestJosephusProblem_Cyclic(t *testing.T) {
	t.Parallel()
	cases := []struct {
		steps  int
		winner int
		count  int
	}{
		{5, 4, 8},
		{3, 8, 8},
		{8, 5, 8},
		{2, 14, 14},
		{13, 56, 58},
	}

	for _, c := range cases {
		cl := NewCyclic()
		fillListForTesting(cl, c.count)
		got := cl.JosephusProblem(c.steps)
		if got != c.winner {
			t.Errorf("got: %v, want: %v", got, c.winner)
		}
	}
}
