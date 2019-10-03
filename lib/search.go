package lib

//O(n)
func LinearSearch(arr []int, n, key int) (int, bool) {
	for i := 0; i < n; i++ {
		if arr[i] == key {
			return i, true
		}
	}
	return 0, false
}

//O(n)
//faster than LinearSearch because the number of comparison is less
func LinearSearchGardner(arr []int, n, key int) (int, bool) {
	i := 0
	arr = append(arr, key)
	for arr[i] != key {
		i++
	}
	if i == n {
		return 0, false
	}
	return i, true
}
