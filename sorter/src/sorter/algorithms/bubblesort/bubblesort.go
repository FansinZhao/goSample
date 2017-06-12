package bubblesort

func BubbleSort(values []int) {
	for i := len(values) - 1; i >= 0; i-- {
		for j := i + 1; j < len(values); j++ {
			if values[j-1] > values[j] {
				values[j-1], values[j] = values[j], values[j-1]
			}
		}

	}
}
