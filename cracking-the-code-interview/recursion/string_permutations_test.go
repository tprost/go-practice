package recursion

import "testing"
import "errors"

func TestStringPermutations(t *testing.T) {

	expected := []string{"cat", "cta", "act", "atc", "tca", "tac"}
	output := StringPermutations("cat")

	var error error
	index := 0
	if len(expected) != len(output) {
		error = errors.New("different amount of permutations")
	}
	for error == nil && index < len(expected) {
		if expected[index] != output[index] {
			error = errors.New("did not match")
		}
		index++
	}
	if error != nil {
		t.Error(error)
		t.Error("output was", output)
		t.Error("expected", expected)
	}
}
