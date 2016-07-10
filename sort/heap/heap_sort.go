package heap

import "sort"

func parent(index int) int {
	return (index + 1) / 2 - 1
}

func left(index int) int {
	return 2 * (index + 1) - 1
}

func right(index int) int {
	return 2 * (index + 1)
}

func maxHeapify(data sort.Interface, index int, heapsize int) sort.Interface {
	l := left(index)
	r := right(index)
	largest := index
	if l < heapsize && data.Less(index, l) {
		largest = l
	}
	if r < heapsize && data.Less(largest, r) {
		largest = r
	}
	if largest != index {
		data.Swap(index, largest)
		return maxHeapify(data, largest, heapsize)
	}
	return data
}

func buildMaxHeap(data sort.Interface) sort.Interface {
	for i := data.Len()/2 - 1; i >= 0; i-- {
		maxHeapify(data, i, data.Len())
	}
	return data
}

func Heapsort(data sort.Interface) sort.Interface {
	buildMaxHeap(data)
	heapsize := data.Len()
	for i := data.Len() - 1; i > 0; i-- {
		data.Swap(0, i)
		heapsize--
		maxHeapify(data, 0, heapsize)
	}
	return data
}
