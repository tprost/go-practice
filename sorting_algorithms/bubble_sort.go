package sorting_algorithms

import "sort"

func BubbleSort(data sort.Interface) sort.Interface {
	length := data.Len()
	sorted := false
	for sorted == false {
		sorted = true
		for i := 0; i < length - 1; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
				sorted = false
			}
		}
	}
	return data
}

func BubbleSortInts(integers []int) []int {
	length := len(integers)
	sorted := false
	for sorted == false {
		sorted = true
		for i := 0; i < length - 1; i++ {
			if integers[i] > integers[i+1] {
				temp := integers[i]
				integers[i] = integers[i+1]
				integers[i+1] = temp
				sorted = false
			}
		}
	}
	return integers
}

func BubbleSortStrings(strings []string) []string {
	length := len(strings)
	sorted := false
	for sorted == false {
		sorted = true
		for i := 0; i < length - 1; i++ {
			if strings[i] > strings[i+1] {
				temp := strings[i]
				strings[i] = strings[i+1]
				strings[i+1] = temp
				sorted = false
			}
		}
	}
	return strings
}
