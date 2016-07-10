package bubble

import (
	"testing"
	"reflect"
	"fmt"
)

func TestFiveIntegers(t *testing.T) {
	integers := []int{5, 4, 9, 2, 3}
	sorted := []int{2, 3, 4, 5, 9}
	BubbleSortInts(integers)
	if !reflect.DeepEqual(integers, sorted) {
		t.Error("did not sort the integers. got "	+ fmt.Sprintf("%v", integers) + " instead of " + fmt.Sprintf("%v", sorted))
	}
}

func TestFiveStrings(t *testing.T) {
	// original := []string{"banana", "apple", "aardvark", "zebra", "dog"}
	words := []string{"banana", "apple", "aardvark", "zebra", "dog"}
	sorted := []string{"aardvark", "apple", "banana", "dog", "zebra"}
	BubbleSortStrings(words)
	if !reflect.DeepEqual(words, sorted) {
		t.Error("did not sort the stringss. got "	+ fmt.Sprintf("%v", words) + " instead of " + fmt.Sprintf("%v", sorted))
	}
}
