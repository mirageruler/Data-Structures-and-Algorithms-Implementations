package sorting

import (
	"reflect"
	"testing"
)

func TestBubble(t *testing.T) {
	input := []int{6, 5, 3, 1, 8, 7, 2, 4}
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8}
	got := Bubble(input)

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("got: %#v, expect: %#v", got, expect)
	}

}

func TestSelection(t *testing.T) {
	input := []int{6, 5, 3, 1, 8, 7, 2, 4}
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8}
	got := Selection(input)

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("got: %#v, expect: %#v", got, expect)
	}
}

func TestInsertion(t *testing.T) {
	input := []int{6, 5, 3, 1, 8, 7, 2, 4}
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8}
	got := Insertion(input)

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("got: %#v, expect: %#v", got, expect)
	}
}

func TestMerge(t *testing.T) {
	input := []int{6, 5, 3, 1, 8, 7, 2, 4}
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8}
	got := Merge(input)

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("got: %#v, expect: %#v", got, expect)
	}
}

func TestQuick(t *testing.T) {

}

func TestHeap(t *testing.T) {

}
