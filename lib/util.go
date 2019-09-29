package lib

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

/////////////////// inputln ////////////////
var sc = bufio.NewScanner(os.Stdin)

func Inputln() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	s, t := Inputln(), Inputln()
	fmt.Println(s)
	fmt.Println(t)
}

/////////////////// inputln ////////////////

////////////////// min max //////////////////
func MultiMin(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton min() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Min(float64(res), float64(nums[i])))
	}
	return res
}

func MultiMax(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton max() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Max(float64(res), float64(nums[i])))
	}
	return res
}

////////////////// min max //////////////////

////////////////// print array without parenthesis //////////////////
func PrintNoParenthesis(a []int) {
	for i, n := range a {
		fmt.Print(n)
		if i == len(a)-1 {
			fmt.Printf("\n")
		} else {
			fmt.Print(" ")
		}
	}
}
////////////////// print array without parenthesis //////////////////
