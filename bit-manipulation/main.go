package bit_manipulation

import "errors"

// sets the bits between i and j in n to equal m
// i and j index from the least significant bit
func BitSubbiness(n, m uint32, i, j int) (uint32, error) {
	if i < 0 || j < 0 || i > 31 || j > 31 {
		return 0, errors.New("indices are not between 0 and 31")
	}

	// `a` is a 32-bit number such that all bits between i and j (inclusive)
	// are equal to 1, and others are 0
	// `b` is the opposite of `a`
	var a, b uint32

	for k := 31; k >= 0; k-- {
		a = a << 1
		b = b << 1
		if k >= i && k <= j {
			a++
		} else {
			b++
		}
	}

	// c is n, with bits i through j replaced with zeroes
	c := n & b

  return c | (a & (m << uint32(i))), nil
}
