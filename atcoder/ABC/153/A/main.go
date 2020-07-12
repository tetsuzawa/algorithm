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
	a, _ := strconv.Atoi(ss[1])

	cnt := h / a

	if h%a != 0 {
		cnt += 1
	}
	fmt.Println(cnt)

}
