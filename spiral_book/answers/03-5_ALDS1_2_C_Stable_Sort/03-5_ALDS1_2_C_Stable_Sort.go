package _3_5_ALDS1_2_C_Stable_Sort

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	Sym string
	Num int
}

func BubbleSort03_5(a []Card, n int) {
	var isArranged = true

	for isArranged {
		isArranged = false
		for j := n - 1; j > 0; j-- {
			if a[j].Num < a[j-1].Num {
				a[j], a[j-1] = a[j-1], a[j]
				isArranged = true
			}
		}
	}
}

func SelectionSort03_5(a []Card, n int) {
	var minj int //index of min value

	for i := 0; i < n; i++ {
		minj = i
		for j := i; j < n; j++ {
			if a[j].Num < a[minj].Num {
				minj = j
			}
		}
		a[i], a[minj] = a[minj], a[i]
	}
}

func PrintNoParenthesis03_5(a []Card) {
	for i, n := range a {
		fmt.Printf("%s%d", n.Sym, n.Num)
		if i == len(a)-1 {
			fmt.Printf("\n")
		} else {
			fmt.Print(" ")
		}
	}
}

func MainStableSort() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	cards1 := make([]Card, n)
	cards2 := make([]Card, n)
	sc.Scan()
	for i, s := range strings.Split(sc.Text(), " ") {
		num, _ := strconv.Atoi(string(s[1]))
		card := Card{Sym: string(s[0]), Num: num}
		cards1[i] = card
	}

	copy(cards2, cards1)

	BubbleSort03_5(cards1, n)
	PrintNoParenthesis03_5(cards1)
	fmt.Println("Stable")
	SelectionSort03_5(cards2, n)
	PrintNoParenthesis03_5(cards2)
	isStable := true
	for i, _ := range cards1 {
		if cards1[i] != cards2[i] {
			isStable = false
			break
		}
	}
	if isStable {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
}
