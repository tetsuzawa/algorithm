package d04_RemoveCoveredIntervals

func DropChildSlice(sl [][]int, i int) [][]int {
	if sl == nil {
		return nil
	}
	return append(sl[:i], sl[i+1:]...)
}

func removeCoveredIntervals(intervals [][]int) int {
	count := len(intervals)
	covered := make(map[int]bool)
	for i, interval := range intervals {
		l, r := interval[0], interval[1]
		for j, cinterval := range intervals {
			if i == j {
				continue
			}
			if covered[j]{
				continue
			}
			if l <= cinterval[0] && cinterval[1] <= r {
				covered[j] = true
				count--
			}
		}
	}
	return count
}
