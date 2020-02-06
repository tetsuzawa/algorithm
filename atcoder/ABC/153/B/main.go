package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	ss := strings.Split(sc.Text(), " ")
	h, _ := strconv.Atoi(ss[0])
	n, _ := strconv.Atoi(ss[1])
	sc.Scan()
	ss = strings.Split(sc.Text(), " ")
	var ai = make([]int, n)
	for i, v := range ss {
		ai[i], _ = strconv.Atoi(v)
		h -= ai[i]
	}
	if h<=0 {
		fmt.Println("Yes")
	}else {
		fmt.Println("No")
	}
}
