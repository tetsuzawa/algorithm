package lib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/* ------------------------------- reader ------------------------------- */
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

func nextString() string { return nextReader() }

func nextInt() int { n, _ := strconv.Atoi(nextReader()); return n }

func nextInts(size int) []int {
	ns := make([]int, size)
	for i := 0; i < size; i++ {
		ns[i] = nextInt()
	}
	return ns
}
/* ------------------------------- reader ------------------------------- */

/* ------------------------------- inputln ------------------------------- */
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

/* ------------------------------- inputln ------------------------------- */


/* ------------------------------- print array without parenthesis ------------------------------- */
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
/* ------------------------------- print array without parenthesis ------------------------------- */
