package quicksort

func quickSort(values []int, left, right int) {
	temp := values[left]
	t := left
	i, j := left, right
	for i < j {
		for j >= t && values[j] >= temp {
			j--
		}
		if j >= t {
			values[t] = values[j]
			t = j
		}
		for i <= t && values[i] <= temp {
			i++
		}
		if i <= t {
			values[t] = values[i]
			t = i
		}
	}
	values[t] = temp
	if t-left > 1 {
		quickSort(values, left, t-1)
	}
	if right-t > 1 {
		quickSort(values, t+1, right)
	}
}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}
