package _4_2_ALDS1_3_A_Stack

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack struct {
	Body []int
	Top  int
	Max  int
}

func NewStack(max int) *stack {
	s := new(stack)
	s.Top = 0
	s.Max = max
	s.Body = make([]int, max)
	return s
}

func (s *stack) isEmpty() bool {
	return s.Top == 0
}

func (s *stack) isFull() bool {
	return s.Top >= s.Max-1
}

func (s *stack) Push(x int) error {
	if s.isFull() {
		err := errors.New("body is full")
		return err
	}
	s.Top++
	s.Body[s.Top] = x
	return nil
}

func (s *stack) Pop() (int, error) {
	if s.isEmpty() {
		err := errors.New("body is empty")
		return 0, err
	}
	s.Top--
	return s.Body[s.Top+1], nil
}

func MainStack() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	sl := strings.Split(sc.Text(), " ")
	st := NewStack(len(sl))
	for _, s := range sl {
		switch s {
		case "+":
			a, _ := st.Pop()
			b, _ := st.Pop()
			_ = st.Push(a + b)
		case "-":
			a, _ := st.Pop()
			b, _ := st.Pop()
			_ = st.Push(b - a)
		case "*":
			a, _ := st.Pop()
			b, _ := st.Pop()
			_ = st.Push(a * b)
		default:
			n, _ := strconv.Atoi(s)
			_ = st.Push(n)
		}
	}
	n, _ := st.Pop()
	fmt.Println(n)

}

//func MainStack() {
//	sc = bufio.NewScanner(os.Stdin)
//	sc.Scan()
//	str	:= sc.Text()
//	//var keys = make([]string, len(str))
//	keys := strings.Split(sc.Text(), " ")
//	numr := regexp.MustCompile(`[1-9]+`)
//	oper := regexp.MustCompile(`[\+\-\*\/]+`)
//	nums := numr.FindAllStringSubmatch(str, -1)
//	opes := oper.FindAllStringSubmatch(str, -1)
//	nst := NewStack(len(nums))
//	ost := NewStack(len(opes))
//
//	for _, p := range nums{
//		nst.Push()
//
//	}
//
//	for _, key := range keys{
//		k, err := strconv.Atoi(key)
//		if err != nil{
//
//		}else {
//			_ = st.Push(k)
//		}
//	}
//
//}

//func MainStack() {
//	sc := bufio.NewScanner(os.Stdin)
//	sc.Scan()
//	str := sc.Text()
//	st := NewStack(len(str))
//	numr := regexp.MustCompile(`[1-9]+`)
//	isMulti := false
//	tmpIndex := 0
//	for i := 0; i < len(str); i++ {
//		if string(str[i]) == "+" {
//			a, _ := st.Pop()
//			b, _ := st.Pop()
//			_ = st.Push(a + b)
//		} else if string(str[i]) == "-" {
//			b, _ := st.Pop()
//			a, _ := st.Pop()
//			_ = st.Push(a - b)
//		} else if string(str[i]) == "*" {
//			a, _ := st.Pop()
//			b, _ := st.Pop()
//			_ = st.Push(a * b)
//		} else if numr.MatchString(string(str[i])) {
//			if i < tmpIndex && isMulti{
//				continue
//			} else if i == tmpIndex && tmpIndex != 0{
//				isMulti = false
//				continue
//			}
//			j := i
//			cnt := 0
//			for string(str[j]) != " " && j<len(str){
//				j++
//				cnt++
//				if j != 1{isMulti = true}
//			}
//			tmpIndex = i + cnt
//			n, _ := strconv.Atoi(str[i:i+cnt])
//			_ = st.Push(n)
//			if j == len(str)-1{
//				break
//			}
//			//if is2digits {
//			//	continue
//			//}else if i+1 < len(str) && !is2digits && numr.MatchString(string(str[i+1])) {
//			//	is2digits = true
//			//	n, _ := strconv.Atoi(str[i:i+1])
//			//	_ = st.Push(n)
//			//}else{
//			//	n, _ := strconv.Atoi(string(str[i]))
//			//	_ = st.Push(n)
//			//}
//		} else {
//			continue
//		}
//		fmt.Println(st.Body)
//	}
//	ans, _ := st.Pop()
//	fmt.Println(ans)
//}
