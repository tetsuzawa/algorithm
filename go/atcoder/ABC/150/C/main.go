package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var nextReader func() string

func newScanner() func() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024), int(1e11))
	sc.Split(bufio.ScanWords)
	return func() string {
		sc.Scan()
		return sc.Text()
	}
}

func nextInt() int { n, _ := strconv.Atoi(nextReader()); return n }

func nextInts(size int) []int {
	ns := make([]int, size)
	for i := 0; i < size; i++ {
		ns[i] = nextInt()
	}
	return ns
}

func init() {
	nextReader = newScanner()
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

func SortedPermute(nums []int) [][]int {
	n := factorial(len(nums))
	ret := make([][]int, 0, n)
	permute(makeCopy(nums), &ret)
	ret = SortPermutation(ret)
	return ret
}

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

func Equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	} else {
		return a
	}
}

func main() {
	N := nextInt()
	P := nextInts(N)
	Q := nextInts(N)
	permutation := SortedPermute(P)
	var a, b int
	for i, v := range permutation {
		if Equals(P, v) {
			a = i + 1
		}
		if Equals(Q, v) {
			b = i + 1
		}
	}
	fmt.Println(abs(a - b))
}
