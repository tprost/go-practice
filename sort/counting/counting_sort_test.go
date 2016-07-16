package counting_sort

import (
	"testing"
	"reflect"
	"fmt"
)

func TestCountingSort(t *testing.T) {
	input := []int{2, 5, 3, 0, 2, 3, 0, 3}
	expected := []int{0, 0, 2, 2, 3, 3, 3, 5}
	output := []int{0, 0, 0, 0, 0, 0, 0, 0}
	CountingSort(input, output, 5)
	if !reflect.DeepEqual(output, expected) {
		t.Error("did not successfully sort: " + fmt.Sprintf("%v", output))
	}
}
