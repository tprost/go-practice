package bitmip

import "math"

func RequiredBits(a, b uint32) int {
	var count int
	for i := 31; i >= 0; i-- {
		t := uint32(math.Pow(2, float64(i)))
		var different bool
		if a >= t {
			different = !different
			a = a - t
		}
		if b >= t {
			different = !different
			b = b - t
		}
		if different {
			count++
		}
	}
	return count
}
