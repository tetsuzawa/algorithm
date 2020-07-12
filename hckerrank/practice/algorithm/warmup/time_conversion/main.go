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
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
	hh, err := strconv.Atoi(s[0:2])
	checkError(err)
	mm := s[3:5]
	checkError(err)
	ss := s[6:8]
	checkError(err)
	XM := s[8:]
	if hh == 12{
		hh = 0
	}
	if XM == "AM" {
		return fmt.Sprintf("%02d:%s:%s", hh, mm, ss)
	} else {
		return fmt.Sprintf("%02d:%s:%s", hh+12, mm, ss)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

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
