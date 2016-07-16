package counting_sort

func CountingSort(A []int, B []int, k int) []int {
	C := make([]int, k + 1)
	for i := 0; i < len(A); i++ {
		C[A[i]] = C[A[i]] + 1
	}
	// `C[i]` contains the number of occurences of `i`
	for i := 1; i < k + 1; i++ {
		C[i] = C[i] + C[i - 1]
	}
	// `C[i]` contains the number of items less
  // than or equal to`i`
	for i := len(A) - 1; i >= 0; i-- {
		B[C[A[i]] - 1] = A[i]
		C[A[i]] = C[A[i]] - 1
	}
	return B
}
