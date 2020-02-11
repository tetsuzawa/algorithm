package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// O(n)

const (
	MAX = 10000
	NIL = -1
)

type Node struct {
	p int
	l int
	r int
}

var n int
var T = make([]Node, MAX)

// 先行順巡回
func preParse(u int) {
	if u == NIL {
		return
	}
	fmt.Printf(" %d", u)
	preParse(T[u].l)
	preParse(T[u].r)
}

// 中間順巡回
func inParse(u int) {
	if u == NIL {
		return
	}
	inParse(T[u].l)
	fmt.Printf(" %d", u)
	inParse(T[u].r)
}

// 後行順巡回
func postParse(u int) {
	if u == NIL {
		return
	}
	postParse(T[u].l)
	postParse(T[u].r)
	fmt.Printf(" %d", u)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var v, l, r, root int
	var err error

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		T[i].p = NIL
	}

	var ss []string
	for i := 0; i < n; i++ {
		sc.Scan()
		ss = strings.Split(sc.Text(), " ")
		v, err = strconv.Atoi(ss[0])
		checkLog(err)
		l, err = strconv.Atoi(ss[1])
		checkLog(err)
		r, err = strconv.Atoi(ss[2])
		checkLog(err)
		T[v].l = l
		T[v].r = r
		if l != NIL {
			T[l].p = v
		}
		if r != NIL {
			T[r].p = v
		}
	}

	for i := 0; i < n; i++ {
		if T[i].p == NIL {
			root = i
		}
	}

	fmt.Println("Preorder")
	preParse(root)
	fmt.Println("")
	fmt.Println("Inorder")
	inParse(root)
	fmt.Println("")
	fmt.Println("Postorder")
	postParse(root)
	fmt.Println("")
}

func checkLog(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
