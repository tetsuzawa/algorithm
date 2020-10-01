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
 * Complete the 'countStudents' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY height as parameter.
 */
//func countStudents(height []int32) int32 {
//	var res int32
//	for i := 0; i < len(height)-1; i++ {
//		swapped := false
//		for j := len(height) - 1; j > i; j-- {
//			if height[j] < height[j-1] {
//				t := height[j]
//				height[j] = height[j-1]
//				height[j-1] = t
//				swapped = true
//			}
//		}
//		if swapped {
//			res += 1
//		}
//	}
//	return res
//}

func countStudents(height []int32) int32 {
	var res int32
	for i := len(height); i > 0; i-- {
		k := height[i]
		j := len(height) - 1
		for j >= 0 && height[j] > k {
			//height[j+1] = height[j]
			j--
			res++
		}
		//height[j+1] = k
	}

	//var res int32
	//for i := 0; i < len(height)-1; i++ {
	//	for j := len(height) - 1; j > i; j-- {
	//		if height[j] < height[j-1] {
	//			t := height[j]
	//			height[j] = height[j-1]
	//			height[j-1] = t
	//			res++
	//		}
	//	}
	//}
	fmt.Println(height)
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	heightCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var height []int32

	for i := 0; i < int(heightCount); i++ {
		heightItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		heightItem := int32(heightItemTemp)
		height = append(height, heightItem)
	}

	result := countStudents(height)

	fmt.Fprintf(writer, "%d\n", result)

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
