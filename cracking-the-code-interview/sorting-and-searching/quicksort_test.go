package sorting

import "testing"
import "fmt"

func TestQuicksort(t *testing.T) {

	people := ByAge{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	sorted := ByAge{
		{"Michael", 17},
		{"Jenny", 26},
		{"Bob", 31},
		{"John", 42},
	}

	Quicksort(people)
	for index, person := range people {
		if person != sorted[index] {
			fmt.Println(people)
			fmt.Println(sorted)
			t.Error("sorted does not match")
			return
		}
	}
}
