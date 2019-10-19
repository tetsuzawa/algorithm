package lib

//O(NlogN), guaranteed
//it need certain size of memory to keep input data

var cnt int

func Merge(A []int, n, left, mid, right int) {
	var sentinel = 2000000000
	n1 := mid - left
	n2 := right - mid
	var L = make([]int, n1)
	var R = make([]int, n2)
	for i := 0; i < n1; i++ {
		L[i] = A[left+1]
	}
	for i := 0; i < n2; i++ {
		L[i] = A[mid+1]
	}
	L = append(L, sentinel)
	R = append(R, sentinel)
	var i, j = 0, 0
	for k := left; k < right; k++ {
		cnt++
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}
}

func MergeSort(A []int, n, left, right int) {
	if left+1 < right {
		mid := (left + right) / 2
		MergeSort(A, n, left, mid)
		MergeSort(A, n, mid, right)
		Merge(A, n, left, mid, right)
	}

}

//////////// Partition ////////////

func Partion(A []int, p, r int) int {
	x := A[r]
	i := p - 1

	for j := p; j < r; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[i] = a[i], a[i+1]
	return i+1
}

//////////// Partition ////////////

//////////// Quick Sort////////////
func QuickSort(A []int, p, r int)  {
	if p<r{
		p = Partion(A, p, r)
	}
}
