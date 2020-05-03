package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var next string
	if s == "ABC" {
		next = "ARC"
	} else {
		next = "ABC"
	}
	fmt.Println(next)
}
