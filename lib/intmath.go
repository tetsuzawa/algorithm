package lib

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func maxMapKeys(mp map[int]int) []int {
	maxKeys := make([]int, 0)
	max := 0
	for k, v := range mp {
		if v > max {
			maxKeys = append([]int{}, k)
			max = v
		} else if v == max {
			maxKeys = append(maxKeys, k)
		}
	}
	return maxKeys
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func sum(xs []int) int {
	var s int
	for _, x := range xs {
		s += x
	}
	return s
}

/* --------------- find the digits of the number  --------------- */
// eg. x=18532, returns 1
func maxDigit(x int) int {
	a := 0
	for x != 0 {
		a = x
		x /= 10
	}
	return a
}

// eg. x=18532, returns 2
func leastDigit(x int) int {
	return x % 10
}

// eg. x=18532, returns 5
func nDigits(x int) int {
	if x == 0 {
		return 1
	}
	var cnt int
	for x != 0 {
		cnt++
		x /= 10
	}
	return cnt
}

/* --------------- find the digits of the number  --------------- */

/* --------------- 順列 permutation  --------------- */
func Permute(nums []int) [][]int {
	n := factorial(len(nums))
	ret := make([][]int, 0, n)
	permute(makeCopy(nums), &ret)
	return ret
}

func SortedPermute(nums []int) [][]int {
	n := factorial(len(nums))
	ret := make([][]int, 0, n)
	permute(makeCopy(nums), &ret)
	ret = SortPermutation(ret)
	return ret
}

func permute(nums []int, ret *[][]int) {
	*ret = append(*ret, makeCopy(nums))

	n := len(nums)
	p := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		p[i] = i
	}
	for i := 1; i < n; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}

		nums[i], nums[j] = nums[j], nums[i]
		*ret = append(*ret, makeCopy(nums))
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}

func factorial(n int) int {
	ret := 1
	for i := 2; i <= n; i++ {
		ret *= i
	}
	return ret
}

func makeCopy(nums []int) []int {
	return append([]int{}, nums...)
}

/* --------------- 2次元順列スライスの並び替え [[1,2,3],[3,2,1],[1,3,2]], return  [[1,2,3],[1,3,2],[3,2,1]] --------------- */

func SortPermutation(nums [][]int) [][]int {
	is := Ints2dToInts(nums)
	sort.Ints(is)
	nums = IntsToInts2d(is)
	return nums
}

func Ints2dToInts(is2d [][]int) []int {
	is := make([]int, len(is2d))
	for i, v := range is2d {
		is[i] = IntsToInt(v)
	}
	return is
}

func IntsToInt(is []int) int {
	s := IntsToString(is)
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func IntsToString(is []int) string {
	var ss = make([]string, len(is), len(is))
	for i, v := range is {
		ss[i] = strconv.Itoa(v)
	}
	return strings.Join(ss, "")
}

func IntsToInts2d(is []int) [][]int {
	is2d := make([][]int, len(is), len(is))
	for i, v := range is {
		is2d[i] = IntToInts(v)
	}
	return is2d
}

func IntToInts(a int) []int {
	s := strconv.Itoa(a)
	is := StringToInts(s)
	return is
}

func StringToInts(s string) []int {
	is := make([]int, len(s))
	var err error
	for i := 0; i < len(s); i++ {
		is[i], err = strconv.Atoi(string(s[i]))
		if err != nil {
			panic(err)
		}
	}
	return is
}

/* --------------- 2次元順列スライスの並び替え [[1,2,3],[3,2,1],[1,3,2]], return  [[1,2,3],[1,3,2],[3,2,1]] --------------- */
/* --------------- 順列 permutation  --------------- */

