package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stdout, err)
	}
}

func run() error {
	args := os.Args
	exePath := args[0]
	//targetDir := args[1]
	//pattern := filepath.Dir(targetDir) + "*.txt"
	pattern := "*.txt"
	files, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}

	for i, file := range files {
		fmt.Println("Test txt:", i+1)
		//cmd := exec.Command("cat", file, "|", exePath)
		content, err := ioutil.ReadFile(file)
		fmt.Println("content:", content)
		if err != nil {
			return err
		}
		fmt.Println("cmd")
		fmt.Println(exePath)
		cmd := exec.Command(exePath)
		fmt.Println("stdin")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			return err
		}
		fmt.Println("write")
		if err = binary.Write(stdin, binary.LittleEndian, content); err != nil {
			return err
		}
		fmt.Println("close")
		if err = stdin.Close(); err != nil {
			return err
		}
		fmt.Println("out")
		out, err := cmd.Output()
		if err != nil {
			return err
		}
		fmt.Printf("Result: %v\n\n", string(out))
	}
	fmt.Println("End!")
	return nil
}

func ReadFile(path string) (string, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
