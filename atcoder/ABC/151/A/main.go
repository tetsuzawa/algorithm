package main

import (
	"fmt"
	"strings"
)

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	var C string
	fmt.Scan(&C)
	byteIdx := strings.Index(alphabet, C)
	fmt.Println(string(alphabet[byteIdx+1]))
}
