package dataStructure

import (
	"errors"
)

type queue struct {
	Body []int
	Head int
	Tail int
	Max  int
}

func NewQueue(max int) *queue {
	q := new(queue)
	q.Body = make([]int, max)
	q.Head, q.Tail = 0, 0
	return q
}

func (q *queue) isEmpty() bool {
	return q.Head == q.Tail
}

func (q *queue) isFull() bool {
	return q.Head == (q.Tail-1)%q.Max
}

func (q *queue) Enqueue(x int) error {
	if q.isFull() {
		err := errors.New("overflow: body is full")
		return err
	}
	q.Body[q.Tail] = x
	if q.Tail+1 == q.Max {
		q.Tail = 0
	} else {
		q.Tail++
	}
	return nil
}

func (q *queue) Dequeue(x int) (int, error) {
	if q.isEmpty() {
		err := errors.New("underflow: body is empty")
		return 0, err
	}
	x = q.Body[q.Tail]
	if q.Head+1 == q.Max {
		q.Head = 0
	} else {
		q.Head++
	}
	return x, nil
}
