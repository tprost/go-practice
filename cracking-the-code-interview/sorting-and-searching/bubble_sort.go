package sorting

import "sort"

func BubbleSort(arr sort.Interface) {
	if arr == nil {
		return
	}
	var sorted bool
	for !sorted {
		sorted = true
		for i := 0; i < arr.Len() - 1; i++ {
			if arr.Less(i+1, i) {
				arr.Swap(i, i+1)
				sorted = false
			}
		}
	}
}
