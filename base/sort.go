package base

func getTopK(arr []int, start, end, k int) int {
	if k >= len(arr) {
		return -1
	}
	i := partition(arr, start, end)
	if i == k {
		return i
	}
	if i > k {
		return getTopK(arr, start, i, k)
	}
	return getTopK(arr, i+1, end, k)
}

func partition(arr []int, start, end int) int {
	if start == end {
		return start
	}
	baseValue := arr[start]
	i, j := start, end-1
	for i < j {
		for i < j && arr[j] > baseValue {
			j--
		}
		for i < j && arr[i] <= baseValue {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[start], arr[j] = arr[j], arr[start]
	return j
}
