package qsort

import (
	"testing"
)

func TestQuickSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	QuickSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 {
		t.Error("QuickSort() failed")
	}
}
