package answers

const (
	M    int = 1046527
	NULL int = -1
	L    int = 14
)

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
	sum, p := 0, 1

	for i := 0; i < len(ss); i++ {
		sum += p * getStr(ss[i])
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

func find(ss []string) int {

}
