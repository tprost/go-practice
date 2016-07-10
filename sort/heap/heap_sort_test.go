package heap

import (
	"testing"
	"reflect"
	"fmt"
)

type StringSlice []string

func (slice StringSlice) Len() int {
	return len(slice)
}

func (slice StringSlice) Less(i, j int) bool {
	return slice[i] < slice[j];
}

func (slice StringSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func TestLeft(t *testing.T) {
	if left(0) != 1 {
		t.Error("left(0) produced " + fmt.Sprintf("%v", left(0)) + " instead of 1") 
	}
}

func TestRight(t *testing.T) {
	if right(2) != 6 {
		t.Error("right(2) produced " + fmt.Sprintf("%v", right(2)) + " instead of 6") 
	}
}

func TestMaxHeapify(t *testing.T) {
	input := StringSlice{"a", "d", "e", "b", "c"}
	expected := StringSlice{"e", "d", "a", "b", "c"} 
	output := maxHeapify(input, 0, input.Len())
	if !reflect.DeepEqual(output, expected) {
		t.Error("did not successfully max-heapify: " + fmt.Sprintf("%v", output))
	}
}

func TestBuildMaxHeap(t *testing.T) {
	input := StringSlice{"a", "b", "c", "d", "e"}
	expected := StringSlice{"e", "d", "c", "a", "b"} 
	output := buildMaxHeap(input)
	if !reflect.DeepEqual(output, expected) {
		t.Error("did not successfully build max heap: " + fmt.Sprintf("%v", output))
	}
}

func TestHeapsort(t *testing.T) {
	input := StringSlice{"e", "b", "d", "a", "c"}
	expected := StringSlice{"a", "b", "c", "d", "e"} 
	output := Heapsort(input)
	if !reflect.DeepEqual(output, expected) {
		t.Error("did not successfully sort: " + fmt.Sprintf("%v", output))
	}
}


