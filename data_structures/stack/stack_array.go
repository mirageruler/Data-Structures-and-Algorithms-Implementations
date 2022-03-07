package stack

// A Stack is a linear data structure which follows a particular order in which the operations are performed.
// The order is: Last in, first out (LIFO) or first in, last out (FILO).

// StackArray defines a stack built from the underlying Go slice.
type StackArray struct {
	Data []interface{}
}

// NewStackArray returns the StackArray with an empty underlying Go slice.
func NewStackArray() *StackArray {
	return &StackArray{}
}

// Push pushes val into the top of the stack and return it, -1 is returned if that stack is invalid.
func (s *StackArray) Push(val interface{}) interface{} {
	if s == nil {
		return -1
	}

	s.Data = append([]interface{}{val}, s.Data...)
	return val
}

// Pop pops the top most element of the stack and return it, -1 is returned if that stack is empty or invalid.
func (s *StackArray) Pop() interface{} {
	if s == nil || len(s.Data) == 0 {
		return -1
	}

	retVal := s.Data[0]
	s.Data = s.Data[1:]

	return retVal
}

// Peak returns the top most element of the stack, -1 is returned if that stack is empty or invalid.
func (s *StackArray) Peak() interface{} {
	if s == nil || len(s.Data) == 0 {
		return -1
	}

	return s.Data[0]
}

// IsEmpty tells whether the stack is empty or not, -1 is returned if that stack is invalid.
func (s *StackArray) IsEmpty() bool {
	if s == nil {
		return true
	}

	return len(s.Data) == 0
}
