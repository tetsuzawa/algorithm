package lib

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

/////////////////// inputln ////////////////
var sc = bufio.NewScanner(os.Stdin)

func inputln() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	s, t := inputln(), inputln()
	fmt.Println(s)
	fmt.Println(t)
}


func multiMin(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton min() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Min(float64(res), float64(nums[i])))
	}
	return res
}

func multiMax(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton max() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Max(float64(res), float64(nums[i])))
	}
	return res
}
