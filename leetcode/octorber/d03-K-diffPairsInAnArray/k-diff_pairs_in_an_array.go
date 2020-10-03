package d03_K_diffPairsInAnArray

func findPairs(nums []int, k int) int {
	var seenPair = make(map[int]int)
	var count int
	for i, a := range nums {
		for j, b := range nums {
			if i == j {
				continue
			}
			if v, ok := seenPair[a]; ok && v == b {
				continue
			}
			if b-a == k {
				count++
				seenPair[a] = b
			}
		}
	}
	return count
}
