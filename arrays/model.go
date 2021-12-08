package main

import (
	"encoding/json"
	"fmt"
)

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

// get does get and return the element at the specified index of the `Array` caller.
func (a *Array) get(index int) (interface{}, error) {
	// Validate the `Array` caller
	if a == nil || a.Data == nil {
		return a, nil
	}

	// Check if the specified index is valid/exists
	isValid := a.checkValidIndex(index)
	if isValid {
		// Case valid/exists, return and print the element at the specified index of the `Array` caller
		bt, err := json.Marshal(a.Data[index])
		if err != nil {
			return nil, fmt.Errorf("failed to marshal data at index of %v", index)
		}
		fmt.Printf("Array at index of %v has the value of %v\n", index, string(bt))
		return a.Data[index], nil
	}

	// Case invalid/not-exists, print and return the error
	fmt.Printf("This array doesn't have item at index %v\n", index)
	return nil, fmt.Errorf("invalid index")
}

// push does push the specified element into the `Array` caller at the new last index.
func (a *Array) push(item interface{}) (*Array, error) {
	// Validate the `Array` caller
	if a == nil || a.Data == nil {
		return a, nil
	}

	// If the `Array` caller is empty/nil then initialize one.
	if a.Data == nil {
		a.Data = map[int]interface{}{}
	}

	// Handle pushing
	a.Data[a.Length] = item
	a.Length++

	bt, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Array after push %v is %v\n", item, string(bt))

	return a, nil
}

func (a *Array) delete(index int) (*Array, error) {
	// Validate the `Array` caller
	if a == nil || a.Data == nil {
		return a, nil
	}

	// Check if the specified index is valid/exists
	isTrue := a.checkValidIndex(index)
	if isTrue {
		// Case valid/exists, handle delete the specified element
		a.shiftItems(index)
		delete(a.Data, a.Length-1)

		bt, err := json.Marshal(a.Data)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal data at index of %v", index)
		}
		fmt.Printf("Array after deleting item at index of %v is %v\n", index, string(bt))

		return a, nil
	}

	// Case invalid/not-exists, print and return the error
	fmt.Printf("This array doesn't have item at index %v\n", index)
	return nil, fmt.Errorf("invalid index")
}

// shiftItems does shift elements to the previous index, from the specified index to the last index (BUT not including) of the `Array` caller.
func (a *Array) shiftItems(index int) int {
	if a == nil || a.Data == nil {
		return a.Length
	}

	for i := index; i < a.Length-1; i++ {
		a.Data[i] = a.Data[i+1]
	}
	return a.Length
}

func main() {
	fmt.Println("Hello, this is ARRAYS playground!")

	myArray := Array{}

	myArray.push(1)
	myArray.push("5")
	myArray.push([]string{"john", "mike", "shawn"})
	myArray.push(map[string]interface{}{"key1": 22})
	myArray.push(10)

	myArray.delete(2)

	myArray.get(3)
}
