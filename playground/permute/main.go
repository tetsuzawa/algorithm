package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/* --------------- 順列 permutation  --------------- */

// 速い
func Permute(nums []int) [][]int {
	n := factorial(len(nums))
	ret := make([][]int, 0, n)
	permute(nums, &ret)
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

func SortPermutation(nums [][]int) [][]int {
	is := Ints2dToInts(nums)
	sort.Ints(is)
	nums = IntsToInts2d(is)
	return nums
}

func main() {
	sl := []int{1, 2, 3, 4}
	permutation := Permute(sl)
	for i, v := range permutation {
		fmt.Printf("%v: \t%v\n", i, v)
	}
	fmt.Printf("\n\n")
	permutation = SortPermutation(permutation)
	for i, v := range permutation {
		fmt.Printf("%v: \t%v\n", i, v)
	}
}
