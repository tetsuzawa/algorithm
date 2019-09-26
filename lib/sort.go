package lib

//O(N^2)
//It operates at high speed for data that is arranged to some extent.
//while the computational complexity becomes O(N^2) if the data arranged descending order.
func InsertionSort(a []int, n int) {
	var i, j int
	var v int

	for i = 1; i < n; i++ {
		v = a[i]
		j = i - 1
		for j >= 0 && a[j] > v {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = v
	}
}

func BubbleSort(a []int, n int) {
	var isArranged = true
	var j int

	for isArranged {
		isArranged = false
		for j = n - 1; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
			isArranged = true
		}
	}
}
