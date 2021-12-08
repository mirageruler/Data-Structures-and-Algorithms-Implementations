package main

import (
	"encoding/json"
	"fmt"
)

type Array struct {
	Length int
	Data   map[int]interface{}
}

func (a *Array) checkValidIndex(index int) bool {
	if _, ok := a.Data[index]; ok {
		return true
	}
	return false
}

func (a *Array) get(index int) (interface{}, error) {
	isValid := a.checkValidIndex(index)
	if isValid {
		bt, err := json.Marshal(a.Data[index])
		if err != nil {
			return nil, fmt.Errorf("failed to marshal data at index of %v", index)
		}
		fmt.Printf("Array at index of %v has the value of %v\n", index, string(bt))
		return a.Data[index], nil
	}

	fmt.Printf("This array doesn't have item at index %v\n", index)
	return nil, fmt.Errorf("invalid index")
}

func (a *Array) push(item interface{}) (*Array, error) {
	if a.Data == nil {
		a.Data = map[int]interface{}{}
	}

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
	if a == nil || a.Data == nil {
		return a, nil
	}

	isTrue := a.checkValidIndex(index)
	if isTrue {
		a.shiftItems(index)
		delete(a.Data, a.Length-1)

		bt, err := json.Marshal(a.Data)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal data at index of %v", index)
		}
		fmt.Printf("Array after deleting item at index of %v is %v\n", index, string(bt))

		return a, nil
	}

	fmt.Printf("This array doesn't have item at index %v\n", index)
	return nil, fmt.Errorf("invalid index")
}

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
