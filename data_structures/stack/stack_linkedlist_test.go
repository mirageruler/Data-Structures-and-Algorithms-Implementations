package stack

import "testing"

func TestPush_StackLinkedList(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		manipulate func(s *StackLinkedList, val interface{}) interface{}
		want       interface{}
	}{
		{
			name: "Case return -1 due to nil stack",
			manipulate: func(s *StackLinkedList, v interface{}) interface{} {
				s = nil // make the stack nil
				return s.Push(v)
			},
			want: -1,
		},
		{
			name: "Case success when stack is empty",
			manipulate: func(s *StackLinkedList, v interface{}) interface{} {
				return s.Push(v)
			},
			want: 10,
		},
		{
			name: "Case success when stack is not empty",
			manipulate: func(s *StackLinkedList, v interface{}) interface{} {
				s.Push(5)
				return s.Push(v)
			},
			want: 10,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := NewStackLinkedList()
			got := c.manipulate(s, 10)
			if got != c.want {
				t.Errorf("got: %#v, want: %#v", got, c.want)
			}
		})
	}
}

func TestPop_StackLinkedList(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		manipulate func(s *StackLinkedList) interface{}
		want       interface{}
	}{
		{
			name: "Case return -1 due to nil stack",
			manipulate: func(s *StackLinkedList) interface{} {
				s = nil // make the stack nil
				return s.Pop()
			},
			want: -1,
		},
		{
			name: "Case success when stack have 1 elem",
			manipulate: func(s *StackLinkedList) interface{} {
				s.Push(10)
				return s.Pop()
			},
			want: 10,
		},
		{
			name: "Case success when stack have multiple elems",
			manipulate: func(s *StackLinkedList) interface{} {
				s.Push(5)
				s.Push(10)
				return s.Pop()
			},
			want: 10,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := NewStackLinkedList()
			got := c.manipulate(s)
			if got != c.want {
				t.Errorf("got: %#v, want: %#v", got, c.want)
			}
		})
	}
}

func TestPeak_StackLinkedList(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		manipulate func(s *StackLinkedList) interface{}
		want       interface{}
	}{
		{
			name: "Case return -1 due to nil stack",
			manipulate: func(s *StackLinkedList) interface{} {
				s = nil // make the stack nil
				return s.Peak()
			},
			want: -1,
		},
		{
			name: "Case success and return the pushed value",
			manipulate: func(s *StackLinkedList) interface{} {
				s.Push(5)
				s.Push(10)
				return s.Peak()
			},
			want: 10,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := NewStackLinkedList()
			got := c.manipulate(s)
			if got != c.want {
				t.Errorf("got: %#v, want: %#v", got, c.want)
			}
		})
	}
}

func TestIsEmpty_StackLinkedList(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		manipulate func(s *StackLinkedList) interface{}
		want       interface{}
	}{
		{
			name: "Case return true due to nil stack",
			manipulate: func(s *StackLinkedList) interface{} {
				s = nil // make the stack nil
				return s.IsEmpty()
			},
			want: true,
		},
		{
			name: "Case return false due to not empty stack",
			manipulate: func(s *StackLinkedList) interface{} {
				s.Push(5)
				s.Push(10)
				return s.IsEmpty()
			},
			want: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := NewStackLinkedList()
			got := c.manipulate(s)
			if got != c.want {
				t.Errorf("got: %#v, want: %#v", got, c.want)
			}
		})
	}
}
