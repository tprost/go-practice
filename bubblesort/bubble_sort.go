package bubble_sort

func BubbleSort(integers []int) []int {
	length := len(integers)
	sorted := false
	for sorted == false {
		sorted = true
		for i := 0; i < length - 1; i++ {
			if integers[i] > integers[i+1] {
				swap := integers[i]
				integers[i] = integers[i+1]
				integers[i+1] = swap
				sorted = false
			}
		}
	}
	return integers
}
