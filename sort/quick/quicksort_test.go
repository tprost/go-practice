package quicksort

import (
	"testing"
	"reflect"
	"fmt"
)

type IntSlice []int

func (slice IntSlice) Len() int {
	return len(slice)
}

func (slice IntSlice) Less(i, j int) bool {
	return slice[i] < slice[j]
}

func (slice IntSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}


func TestPartition(t *testing.T) {
	input := IntSlice{2, 8, 7, 1, 3, 5, 6, 4}
	expected := IntSlice{2, 1, 3, 4, 7, 5, 6, 8}
	output := partition(input, 0, 7)
	if output != 3 {
		t.Error("expected partition to output 2 but it produced " + fmt.Sprintf("%v", output))
	}
	if !reflect.DeepEqual(input, expected) {
		t.Error("did not match expected")
	}
}

func TestQuicksort(t *testing.T) {
	input := IntSlice{0, 4, 5, 2, 9, 7, 8, 1, 3, 6}
	expected := IntSlice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	output := Quicksort(input)
	if !reflect.DeepEqual(output, expected) {
		t.Error("did not successfully sort: " + fmt.Sprintf("%v", output))
	}
}
