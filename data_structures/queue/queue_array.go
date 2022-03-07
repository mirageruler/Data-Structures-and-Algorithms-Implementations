package queue

// A Queue is a linear data structure which follows a particular order in which the operations are performed.
// The order is: First in, first out (FIFO).

// QueueArray defines a queue built from the underlying Go slice
type QueueArray struct {
	Data []interface{}
}

// NewQueueArray returns the QueueArray with an empty underlying Go slice
func NewQueueArray() *QueueArray {
	return &QueueArray{}
}

// Enqueue adds val to the end of the queue and return it, -1 is returned if that queue is invalid.
func (q *QueueArray) Enqueue(val interface{}) interface{} {
	if q == nil {
		return -1
	}

	q.Data = append(q.Data, val)

	return val
}

// Dequeue removes the first element of the queue and return it, -1 is return it that queue is empty or invalid.
func (q *QueueArray) Dequeue() interface{} {
	if q == nil || len(q.Data) == 0 {
		return -1
	}

	retVal := q.Data[0]
	q.Data = q.Data[1:]

	return retVal
}

// FrontQueue returns the first(oldest) element of the queue, -1 is returned if that queue is empty or invalid.
func (q *QueueArray) FrontQueue() interface{} {
	if q == nil || len(q.Data) == 0 {
		return -1
	}

	return q.Data[0]
}

// BackQueue returns the last(newest) element of the queue, -1 is returned if that queue is empty or invalid.
func (q *QueueArray) BackQueue() interface{} {
	if q == nil || len(q.Data) == 0 {
		return -1
	}

	return q.Data[len(q.Data)-1]
}

// IsEmpty tells whether the queue is empty or not, -1 is returned if that queue is invalid.
func (q *QueueArray) IsEmpty() bool {
	if q == nil {
		return true
	}

	return len(q.Data) == 0
}
