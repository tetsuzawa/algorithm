package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func Inputln() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	ss := strings.Split(Inputln(), " ")
	a, _ := strconv.Atoi(ss[0])
	b, _ := strconv.Atoi(ss[1])
	if isEven(a * b) {
		fmt.Println("Even")
		return
	}
	fmt.Println("Odd")
}

func isEven(x int) bool {
	return x%2 == 0
}

func isOdd(x int) bool {
	return x%2 != 0
}
