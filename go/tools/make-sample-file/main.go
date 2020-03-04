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
	if len(os.Args) != 2 {
		fmt.Println(usage)
		os.Exit(1)
	}
	num, err := strconv.Atoi(os.Args[1])
	check(err)
	for i := 1; i <= num; i++ {
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

