package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'breakPalindrome' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING palindromeStr as parameter.
 */

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func BreakPalindrome(palindromeStr string) string {
	for i, v := range palindromeStr {
		if v > 'a' {
			for r := 'a'; r <= 'z'; r++ {
				res := palindromeStr[:i] + string(r) + palindromeStr[i+1:]
				if !isPalindrome(res) {
					return res
				}
			}
		}
	}
	return "IMPOSSIBLE"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	palindromeStr := readLine(reader)

	result := BreakPalindrome(palindromeStr)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
