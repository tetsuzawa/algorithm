package _4_3_ALDS1_3_B_Queue

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Obj struct {
	Name string
	Body int
}

func NewObj(name string, body int) *Obj {
	obj := new(Obj)
	obj.Name = name
	obj.Body = body
	return obj
}

type Queue struct {
	Objs []Obj
	Head int
	Tail int
	Max  int
}

func NewQueue(max int) *Queue {
	q := new(Queue)
	q.Objs = make([]Obj, max*3)
	q.Head, q.Tail = 0, 0
	q.Max = max
	return q
}

func (q *Queue) isEmpty() bool {
	return q.Head == q.Tail
}

func (q *Queue) isFull() bool {
	return q.Head == (q.Tail-1)%q.Max
	//return q.Head == (q.Tail-1)%q.Max
}

func (q *Queue) Enqueue(obj Obj) error {
	if q.isFull() {
		err := errors.New("overflow: body is full")
	return err
	}
	//obj := Obj{name, x}
	q.Objs[q.Tail] = obj
	if q.Tail+1 == q.Max {
		q.Tail = 0
	} else {
		q.Tail++
	}
	return nil
}

func (q *Queue) Dequeue() (Obj, error) {
	if q.isEmpty() {
		err := errors.New("underflow: body is empty")
		return Obj{}, err
	}
	obj := q.Objs[q.Head]
	if q.Head+1 == q.Max {
		q.Head = 0
	} else {
		q.Head++
	}
	return obj, nil
}

func min_04_3(x, y int) int {
	if x < y{
		return x
	} else {
		return y
	}
}

func MainQueue() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	sl := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(sl[0])
	qt, _ := strconv.Atoi(sl[1])
	q := NewQueue(1000)

	for i := 0; i < n; i++ {
		sc.Scan()
		sl = strings.Split(sc.Text(), " ")
		qt, _ := strconv.Atoi(sl[1])
		q.Objs[i] = Obj{sl[0], qt}
	}
	q.Head = 0
	q.Tail = n
	var elaps int

	for q.Head != q.Tail {
		obj, err := q.Dequeue()
		if err != nil{
			fmt.Println(err)
		}
		lowerTime := min_04_3(qt, obj.Body)
		obj.Body -= lowerTime
		elaps += lowerTime
		if obj.Body > 0{
			if err = q.Enqueue(obj); err != nil{
				log.Fatalln(err)
			}
		}else {
			fmt.Println(obj.Name, elaps)
		}
	}
}
