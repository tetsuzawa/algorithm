package lib

//O(N^2) (N^2 - N) / 2
// it does not depends on how data is arranged
//It operates at high speed for data that is arranged to some extent.
//while the computational complexity becomes O(N^2) if the data arranged descending order.
func InsertionSort(a []int, n int) {
	var v int

	for i := 1; i < n; i++ {
		v = a[i]
		j := i - 1
		for j >= 0 && a[j] > v {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = v
	}
}

//O(N^2) (N^2 - N) / 2
//if the array is arranged, it terminates in O(N) time
//note that " if a[j] < a[j-1] " is stable, but it will not stable if "<" is "<=".
//it is known that number of swaps describes disorder of row
func BubbleSort(a []int, n int) {
	var isArranged = true

	for isArranged {
		isArranged = false
		for j := n - 1; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
				isArranged = true
			}
		}
	}
}

//O(N^2) (N^2 - N) / 2
//NOT STABLE Algorithm
// it does not depends on how data is arranged
func SelectionSort(a []int, n int) {
	var minj int //index of min value

	for i := 0; i < n; i++ {
		minj = i
		for j := i; j < n; j++ {
			if a[j] < a[minj] {
				minj = j
			}
		}
		a[i], a[minj] = a[minj], a[i]
	}
}
