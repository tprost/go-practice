package quicksort

import "sort"

func partition(data sort.Interface, p int, r int) int {
	i := p - 1
	for j := p; j < r; j++ {
		if !data.Less(r, j) {
			i = i + 1
			data.Swap(i, j)
		}
	}
	data.Swap(i + 1, r)
	return i + 1
}

func quicksort(data sort.Interface, p int, r int) sort.Interface {
	if p < r {
		q := partition(data, p, r)
		quicksort(data, p, q - 1)
		quicksort(data, q, r)
	}
	return data
}

func Quicksort(data sort.Interface) sort.Interface {
	return quicksort(data, 0, data.Len() - 1)
}
