package main

import "fmt"

// calculate `n` Fibonacci numbers
func Fibonacci(n int64) []int64 {
	if n == 0 {
		return []int64{}
	}
	if n == 1 {
		return []int64{1}
	}
	if n == 2 {
		return []int64{1, 1}
	}
	previous := Fibonacci(n - 1)
	return append(previous, previous[n - 2] + previous[n - 3])
}

func main() {
	fmt.Println(Fibonacci(1000))
}
