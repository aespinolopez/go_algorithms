package sort

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	var input = []int{8, 15, 4, 65, 41}
	var expectedOutput = []int{4, 8, 15, 41, 65}

	InsertionSort(input[:])

	for i, el := range expectedOutput {
		if el != input[i] {
			t.Error("Array isn't sorted")
		}
	}
}
