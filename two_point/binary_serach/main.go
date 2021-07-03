package binary_serach

func getKth(a []int, b []int, k int) int {
	return getKthRec(a, 0, len(a), b, 0, len(b), k)
}

func getKthRec(a []int, start1, end1 int, b []int, start2, end2 int, k int) int {
	if k == 1 {
		return min(a[start1], b[start2])
	}
	step := k / 2 - 1
	pos1 := a[start1+step]
	pos2 := b[start2+step]
	if pos1 > pos2 {
		return getKthRec(a, start1+step, end1, b, start2, end2, k-(start1-step+1))
	} else {
		return getKthRec(a, start1, end1, b, start2+step, end2, k-(start2-step+1))
	}
}

func min(a, b int) int{
	if a > b {
		return b
	}
	return a
}