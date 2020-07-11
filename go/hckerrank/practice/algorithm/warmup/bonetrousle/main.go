package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the bonetrousle function below.
 */
func bonetrousle(n int64, k int64, b int32) []int64 {
	var i int64
	for bit := 0; bit < (1 << k); bit++ {
		var num int
		for bits := bit; bits != 0; bits &= bits - 1 {
			num++
		}
		fmt.Println("num", num)
		if num != int(b) {
			continue
		}
		fmt.Printf("bit: %010b\n", bit)
		var sum int64
		basket := make([]int64, 0)
		for i = 0; i < k; i++ {
			if bit&(1<<i) != 0 {
				sum += i
				basket = append(basket, i)
			}
			fmt.Println("basket:", basket)
			if sum == n && len(basket) == int(b) {
				return basket
			} else if sum > n || len(basket) > int(b) {
				break
			}
		}
	}
	return []int64{-1}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nkb := strings.Split(readLine(reader), " ")

		n, err := strconv.ParseInt(nkb[0], 10, 64)
		checkError(err)

		k, err := strconv.ParseInt(nkb[1], 10, 64)
		checkError(err)

		bTemp, err := strconv.ParseInt(nkb[2], 10, 64)
		checkError(err)
		b := int32(bTemp)

		result := bonetrousle(n, k, b)

		for resultItr, resultItem := range result {
			fmt.Fprintf(writer, "%d", resultItem)

			if resultItr != len(result)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		fmt.Fprintf(writer, "\n")
	}

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
