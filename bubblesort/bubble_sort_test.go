package bubble_sort

import (
	"testing"
	"reflect"
	"fmt"
)

func TestFiveIntegers(t *testing.T) {
	original := []int{5, 4, 9, 2, 3}
	integers := []int{5, 4, 9, 2, 3}
	integers = BubbleSort(integers)
	if !reflect.DeepEqual(integers, []int{2, 3, 4, 5, 9}) {
		t.Error("did not sort the integers. got "	+ fmt.Sprintf("%v", integers) + " instead of " + fmt.Sprintf("%v", original))
	}
}
