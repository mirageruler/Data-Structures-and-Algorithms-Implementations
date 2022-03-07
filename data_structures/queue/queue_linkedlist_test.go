package queue

import "testing"

func TestEnqueue_QueueLinkedList(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		manipulate func(*QueueLinkedList, interface{}) interface{}
		want       interface{}
	}{
		{
			name: "Case return -1 due to nil queue",
			manipulate: func(q *QueueLinkedList, v interface{}) interface{} {
				q.Enqueue(v)
				q = nil // make the queue nil
				return q.Enqueue(v)
			},
			want: -1,
		},
		{
			name: "Case success when queue is empty",
			manipulate: func(q *QueueLinkedList, v interface{}) interface{} {
				return q.Enqueue(v)
			},
			want: 10,
		},
		{
			name: "Case success when queue is not empty",
			manipulate: func(q *QueueLinkedList, v interface{}) interface{} {
				q.Enqueue(5)
				return q.Enqueue(v)
			},
			want: 10,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			q := NewQueueLinkedList()
			got := c.manipulate(q, 10)
			if got != c.want {
				t.Errorf("got: %#v, want: %#v", got, c.want)
			}
		})
	}
}

func TestDequeue_QueueLinkedList(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		manipulate func(*QueueLinkedList) interface{}
		want       interface{}
	}{
		{
			name: "Case return -1 due to empty queue",
			manipulate: func(q *QueueLinkedList) interface{} {
				return q.Dequeue()
			},
			want: -1,
		},
		{
			name: "Case success when queue have 1 elem",
			manipulate: func(q *QueueLinkedList) interface{} {
				q.Enqueue(10)
				return q.Dequeue()
			},
			want: 10,
		},
		{
			name: "Case success when queue have multiple elems",
			manipulate: func(q *QueueLinkedList) interface{} {
				q.Enqueue(5)
				q.Enqueue(10)
				return q.Dequeue()
			},
			want: 5,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			q := NewQueueLinkedList()
			got := c.manipulate(q)
			if got != c.want {
				t.Errorf("got: %#v, want: %#v", got, c.want)
			}
		})
	}
}

func TestPeak_QueueLinkedList(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		manipulate func(*QueueLinkedList) interface{}
		want       interface{}
	}{
		{
			name: "Case return -1 due to empty queue",
			manipulate: func(q *QueueLinkedList) interface{} {
				return q.Peak()
			},
			want: -1,
		},
		{
			name: "Case success and return the first elem",
			manipulate: func(q *QueueLinkedList) interface{} {
				q.Enqueue(5)
				q.Enqueue(10)
				return q.Peak()
			},
			want: 5,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			q := NewQueueLinkedList()
			got := c.manipulate(q)
			if got != c.want {
				t.Errorf("got: %#v, want: %#v", got, c.want)
			}
		})
	}
}

func TestIsEmpty_QueueLinkedList(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		manipulate func(*QueueLinkedList) interface{}
		want       interface{}
	}{
		{
			name: "Case return true due to nil queue",
			manipulate: func(q *QueueLinkedList) interface{} {
				q.Enqueue(8)
				q = nil // make the queue nil
				return q.IsEmpty()
			},
			want: true,
		},
		{
			name: "Case return false due to not empty queue",
			manipulate: func(q *QueueLinkedList) interface{} {
				q.Enqueue(5)
				q.Enqueue(10)
				return q.IsEmpty()
			},
			want: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			q := NewQueueLinkedList()
			got := c.manipulate(q)
			if got != c.want {
				t.Errorf("got: %#v, want: %#v", got, c.want)
			}
		})
	}
}
