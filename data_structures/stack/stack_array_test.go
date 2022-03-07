package stack

import "testing"

func TestPush_StackArray(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		manipulate func(*StackArray, interface{}) interface{}
		want       interface{}
	}{
		{
			name: "Case return -1 due to nil stack",
			manipulate: func(s *StackArray, v interface{}) interface{} {
				s.Push(v)
				s = nil // make the stack nil
				return s.Push(v)
			},
			want: -1,
		},
		{
			name: "Case success and return the pushed value",
			manipulate: func(s *StackArray, v interface{}) interface{} {
				return s.Push(v)
			},
			want: 10,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := NewStackArray()
			got := c.manipulate(s, 10)
			if got != c.want {
				t.Errorf("got: %#v, want: %#v", got, c.want)
			}
		})
	}
}

func TestPop_StackArray(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		manipulate func(*StackArray) interface{}
		want       interface{}
	}{
		{
			name: "Case return -1 due to empty stack",
			manipulate: func(s *StackArray) interface{} {
				return s.Pop()
			},
			want: -1,
		},
		{
			name: "Case success and return the popped value",
			manipulate: func(s *StackArray) interface{} {
				s.Push(5)
				s.Push(10)
				return s.Pop()
			},
			want: 10,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := NewStackArray()
			got := c.manipulate(s)
			if got != c.want {
				t.Errorf("got: %#v, want: %#v", got, c.want)
			}
		})
	}
}

func TestPeak_StackArray(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		manipulate func(*StackArray) interface{}
		want       interface{}
	}{
		{
			name: "Case return -1 due to empty stack",
			manipulate: func(s *StackArray) interface{} {
				return s.Peak()
			},
			want: -1,
		},
		{
			name: "Case success and return the top most value of the stack",
			manipulate: func(s *StackArray) interface{} {
				s.Push(5)
				s.Push(10)
				return s.Peak()
			},
			want: 10,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := NewStackArray()
			got := c.manipulate(s)
			if got != c.want {
				t.Errorf("got: %#v, want: %#v", got, c.want)
			}
		})
	}
}

func TestIsEmpty_StackArray(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		manipulate func(*StackArray) interface{}
		want       interface{}
	}{
		{
			name: "Case return true due to nil stack",
			manipulate: func(s *StackArray) interface{} {
				s = nil // make the stack nil
				return s.IsEmpty()
			},
			want: true,
		},
		{
			name: "Case return false due to not empty stack",
			manipulate: func(s *StackArray) interface{} {
				s.Push(5)
				s.Push(10)
				return s.IsEmpty()
			},
			want: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := NewStackArray()
			got := c.manipulate(s)
			if got != c.want {
				t.Errorf("got: %#v, want: %#v", got, c.want)
			}
		})
	}
}
