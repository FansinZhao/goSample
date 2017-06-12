package bubblesort

import (
	"testing"
)

func TestBubblSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[4] != 5 {
		t.Error("Bubblesort failed ,Got", values, "1 2 3 4 5")
	}
}
