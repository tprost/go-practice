package sorting

import "sort"

func Quicksort(sortable sort.Interface) {

	partition := func(sortable sort.Interface, lo, hi int) int {
		pivot := hi
		i := lo
		for j := lo; j < hi; j++ {
			if sortable.Less(j, pivot) {
				sortable.Swap(i, j)
				i++
			}
		}
		sortable.Swap(i, hi)
		return i
	}

	var quicksortRecursive func(sortable sort.Interface, lo, hi int)

	quicksortRecursive = func(sortable sort.Interface, lo, hi int) {
		if lo < hi {
			p := partition(sortable, lo, hi)
			quicksortRecursive(sortable, lo, p - 1)
			quicksortRecursive(sortable, p+1, hi)
		}
	}

	quicksortRecursive(sortable, 0, sortable.Len() - 1)

}
