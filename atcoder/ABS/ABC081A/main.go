package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func Inputln() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	s := Inputln()
	var sum int
	if is1(string(s[0])) {
		sum += 1
	}
	if is1(string(s[1])) {
		sum += 1
	}
	if is1(string(s[2])) {
		sum += 1
	}
	fmt.Println(sum)
}

func is1(s string) bool {
	return s == "1"
}
