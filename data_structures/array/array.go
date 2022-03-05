package array

import (
	"fmt"
)

// Array type define a value of type `Array`, each elements can be value of arbitrary types
type Array struct {
	Length int
	Data   map[int]interface{}
}

// checkValidIndex does check if the `Array` caller has any value with the given key(index).
func (a *Array) checkValidIndex(index int) bool {
	if _, ok := a.Data[index]; ok {
		return true
	}
	return false
}

// Get does get and return the element at the specified index of the `Array` caller.
func (a *Array) Get(index int) (interface{}, error) {
	// Validate the `Array` caller
	if a == nil || a.Data == nil {
		return a, fmt.Errorf("invalid array")
	}

	// Check if the specified index is valid/exists
	if a.checkValidIndex(index) {
		return a.Data[index], nil
	}

	// Case invalid/not-exists, print and return the error
	return nil, fmt.Errorf("invalid index")
}

// Pop does pop the specified element into the `Array` caller at the current last index.
func (a *Array) Pop() error {
	// Validate the `Array` caller
	if a == nil || a.Data == nil {
		return fmt.Errorf("array is nil! please check again")
	}

	// Handle popping
	{
		delete(a.Data, a.Length-1)
		a.Length--
	}

	return nil
}

// Push does push the specified element into the `Array` caller at the new last index.
func (a *Array) Push(item interface{}) error {
	// Validate the `Array` caller
	if a == nil {
		return fmt.Errorf("array is nil! please check again")
	}

	// If the `Array` caller is empty/nil then initialize one.
	if a.Data == nil {
		a.Data = map[int]interface{}{}
	}

	// Handle pushing
	a.Data[a.Length] = item
	a.Length++

	return nil
}

func (a *Array) Delete(index int) error {
	// Validate the `Array` caller
	if a == nil || a.Data == nil {
		return fmt.Errorf("array is nil! please check again")
	}

	// Check if the specified index is valid/exists
	if a.checkValidIndex(index) {
		// Case valid/exists, handle deleting the specified element
		a.shiftItems(index, "left")

		delete(a.Data, a.Length-1)
		a.Length--

		return nil

	}

	// Case invalid/not-exists, print and return the error
	return fmt.Errorf("invalid index")
}

func (a *Array) Insert(item interface{}, index int) error {
	// Validate the `Array` caller
	if a == nil {
		return fmt.Errorf("array is nil! please check again")
	}

	// If the `Array` caller is empty/nil then initialize one.
	if a.Data == nil || a.Length == 0 {
		return a.Push(item)
	}

	// Handle inserting
	if a.checkValidIndex(index) {
		a.Length++

		a.shiftItems(index, "right")
		a.Data[index] = item

		return nil
	}

	return fmt.Errorf("invalid index")
}

// shiftItems does shift elements to the previous index, from the specified index to the last index (BUT not including) of the `Array` caller.
func (a *Array) shiftItems(index int, direction string) {
	if direction == "left" {
		for i := index; i < a.Length-1; i++ {
			a.Data[i] = a.Data[i+1]
		}
	} else if direction == "right" {
		for i := a.Length - 1; i > index; i-- {
			a.Data[i] = a.Data[i-1]
		}
	}

}
