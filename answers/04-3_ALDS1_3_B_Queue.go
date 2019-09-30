package answers

import (
	"bufio"
	"errors"
	"fmt"
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

type Queues struct {
	Objs []Obj
	Head int
	Tail int
	Max  int
}

func NewQueue(max int) *Queues {
	queue := new(Queues)
	queue.Objs = make([]Obj, max)
	queue.Head, queue.Tail = 0, 0
	return queue
}

func (q *Queues) isEmpty() bool {
	return q.Head == q.Tail
}

func (q *Queues) isFull() bool {
	return q.Head == (q.Tail-1)%q.Max
}

func (q *Queues) Enqueue(obj Obj) error {
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

func (q *Queues) Dequeue() (Obj, error) {
	if q.isEmpty() {
		err := errors.New("underflow: body is empty")
		return Obj{}, err
	}
	obj := q.Objs[q.Tail]
	if q.Head+1 == q.Max {
		q.Head = 0
	} else {
		q.Head++
	}
	return obj, nil
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
		fmt.Println(q.Objs)
	}
	q.Head = 1
	q.Tail = n + 1
	var elaps int

	for q.Head != q.Tail {
		obj, _ := q.Dequeue()
		lowerTime := min(qt, obj.Body)
		obj.Body -= lowerTime
		elaps += lowerTime
		fmt.Println("kko", q.Objs, lowerTime, elaps)
		if obj.Body > 0{
			q.Enqueue(obj)
		}else {
			fmt.Println(obj.Name, elaps)
		}
	}
}
