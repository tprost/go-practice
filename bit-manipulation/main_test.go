package bitmip

import "testing"
import "strconv"

type Test struct {
	N string
	M string
	I int
	J int
  Output string
}

var tests = []Test{
	{"10000000000", "10101", 2, 6, "10001010100"},
	{"10000000000", "10101", 3, 7, "10010101000"},
	{"10000000000", "10101", 3, 4, "10000001000"},
}

func TestBitManipulation(t *testing.T) {
	for _, test := range tests {
		nParsed, error := strconv.ParseUint(test.N, 2, 32)
		if error != nil {
			t.Error("could not parse input bit pattern")
		}
		mParsed, error := strconv.ParseUint(test.M, 2, 32)
		if error != nil {
			t.Error("could not parse input bit pattern")
		}
		n := uint32(nParsed)
		m := uint32(mParsed)
		i := test.I
		j := test.J
		error = BitManipulation(&n, &m, i, j)
		if error != nil {
			t.Error("BitManipulation produced error")
		}
		outputString := strconv.FormatInt(int64(n), 2)
		if outputString != test.Output {
			t.Error(outputString + " did not match expected " + test.Output)
		}
	}
}

func TestRequiredBits(t *testing.T) {
	output := RequiredBits(31, 14)
	if output != 2 {
		t.Error("RequiredBits should have output 2 but it gave", output) 
	}
}
