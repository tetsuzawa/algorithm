package main

import (
	"bufio"
	"fmt"
	"os"
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

func nextString() string { return nextReader() }

func main() {
	nextReader = newScanner()
	s := nextString()
	for i := 0; i < len(s); i++ {
		fmt.Print("x")
	}
	fmt.Println("")
}
