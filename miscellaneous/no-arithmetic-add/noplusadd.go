package noplusadd

import "math"

// add two numbers without using arithmetic operators
func Add(x, y uint32) uint32 {
	var answer uint32
	var carry bool
	// function to return either 0 or 1, the bit of `num` at `position`

	for i := 0; i<32; i++ {
		var b uint32 = uint32(math.Exp2(float64(i))) // need to do some casting and this syntax is likely wrong
		if (b & x) & (b & y) > 0 {
			if carry == true {
				answer = answer | b
			}
			carry = true
		} else if (b & x) | (b & y) > 0 {
			if carry == false {
				answer = answer | b
			}
		} else {
			if carry == true {
				answer = answer | b
			}
			carry = false
		}
	}
	return answer
}
