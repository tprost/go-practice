package noplusadd

import "testing"

func TestAdd(t *testing.T) {
	answer := Add(1, 1)
	if answer != 2 {
		t.Error("1 + 1 should equal 2 but came out to", answer)
	}
}
