package main

import (
	"fmt"

	"encoding/json"
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
		return a, nil
	}

	// Check if the specified index is valid/exists
	if a.checkValidIndex(index) {
		// Case valid/exists, print and return the element at the specified index of the `Array` caller
		bs, err := json.Marshal(a.Data[index])
		if err != nil {
			return nil, err
		}
		fmt.Printf("Array at index %v has the value of %v\n", index, string(bs))
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

	// Printing
	{
		bs, err := json.Marshal(a)
		if err != nil {
			return err
		}
		fmt.Printf("Array after being popped is %v\n", string(bs))
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

	// Printing
	{
		bs, err := json.Marshal(a)
		if err != nil {
			return err
		}
		fmt.Printf("Array after being pushed `%v` is %v\n", item, string(bs))
	}

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
		{
			err := a.shiftItems(index, "left")
			if err != nil {
				return fmt.Errorf("Delete failed shiftItems")
			}

			delete(a.Data, a.Length-1)
			a.Length--
		}

		// Printing
		{
			bs, err := json.Marshal(a)
			if err != nil {
				return fmt.Errorf("failed to marshal data at index %v", index)
			}
			fmt.Printf("Array after deleting item at index %v is %v\n", index, string(bs))
		}

		return nil
	}

	// Case invalid/not-exists, print and return the error
	fmt.Printf("This array doesn't have item at index %v\n", index)
	return fmt.Errorf("invalid index")
}

func (a *Array) Insert(item interface{}, index int) error {
	// Validate the `Array` caller
	if a == nil {
		return fmt.Errorf("array is nil! please check again")
	}

	// If the `Array` caller is empty/nil then initialize one.
	if a.Data == nil {
		a.Data = map[int]interface{}{}
	}

	// Handle inserting
	{
		a.Length++

		err := a.shiftItems(index, "right")
		if err != nil {
			return fmt.Errorf("Insert failed shiftItems")
		}
		a.Data[index] = item
	}

	// Printing
	{
		bs, err := json.Marshal(a)
		if err != nil {
			return err
		}
		fmt.Printf("Array after inserting item `%v` at index %v is: %v\n", item, index, string(bs))
	}

	return nil

}

// shiftItems does shift elements to the previous index, from the specified index to the last index (BUT not including) of the `Array` caller.
func (a *Array) shiftItems(index int, direction string) error {
	if a == nil {
		return fmt.Errorf("invalid index")
	}

	if direction == "left" {
		for i := index; i < a.Length-1; i++ {
			a.Data[i] = a.Data[i+1]
		}
	} else if direction == "right" {
		for i := a.Length - 1; i > index; i-- {
			a.Data[i] = a.Data[i-1]
		}
	}

	return nil
}

func main() {
	fmt.Println("Hello, this is ARRAYS playground!")
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------------------")

	myArray := Array{}

	myArray.Push(1)
	myArray.Push("5")
	myArray.Push([]string{"john", "mike", "shawn"})
	myArray.Push(map[string]interface{}{"key1": 22})
	myArray.Push(10)

	myArray.Insert("INSERT", 2)

	myArray.Delete(2)
	myArray.Pop()

	myArray.Get(3)
}
