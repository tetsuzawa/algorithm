package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var usage = "usage: make-sample-file 2"

func touch(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		defer file.Close()
		return err
	} else {
		currentTime := time.Now().Local()
		err = os.Chtimes(path, currentTime, currentTime)
		return err
	}
}

func main() {
	args := os.Args
	if len(args) < 2 || 3 < len(args) {
		fmt.Println(usage)
		os.Exit(1)
	}
	var start, end int
	var err error
	if len(args) == 2 {
		start = 1
		end, err = strconv.Atoi(os.Args[1])
		check(err)
	} else {
		start, err = strconv.Atoi(os.Args[1])
		check(err)
		end, err = strconv.Atoi(os.Args[2])
		check(err)

	}
	check(err)
	for i := start; i <= end; i++ {
		err = touch(fmt.Sprintf("sample-%d.in", i))
		check(err)
		err = touch(fmt.Sprintf("sample-%d.out", i))
		check(err)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
