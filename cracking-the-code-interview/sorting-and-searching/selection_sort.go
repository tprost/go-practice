package sorting

import "sort"

func SelectionSort(sortable sort.Interface) {
	for i := 0; i < sortable.Len(); i++ {
		smallest := i
		for j := i + 1; j < sortable.Len(); j++ {
			if sortable.Less(j, smallest) {
				smallest = j
			}
		}
		sortable.Swap(i, smallest)
	}
}
