package answers

import "fmt"

const (
	M    int = 1046527
	NULL int = -1
	L    int = 14
)

var H []string

func getStr(s string) int {
	if s == "A" {
		return 1
	} else if s == "C" {
		return 2
	} else if s == "G" {
		return 3
	} else if s == "T" {
		return 4
	} else {
		return 0
	}
}

func getKey(ss []string) int {
	//TODO is ss should be slice or row?
	sum, p := 0, 1

	for _, s := range ss{
		sum += p * getStr(s)
		p *= 5
	}
	return sum
}

func h1(key int) int {
	return key % M
}

func h2(key int) int {
	return 1 + (key % (M - 1))
}

func find(ss []string) bool {
	key := getKey(ss)
	i:=0
	for {
		h := (h1(key) + i * h2(key)) % M
		fmt.Println(h)
		//TODO
		//if H[h] == ss{
		//	return true
		//}
	}
}
