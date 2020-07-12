package dataStructure

import "errors"

type stack struct {
	Body []int
	Top  int
	Max  int
}

func NewStack(max int) *stack {
	s := new(stack)
	s.Top = 0
	s.Max = max
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

func (s *stack) Pop(x int) (int, error) {
	if s.isEmpty() {
		err := errors.New("body is empty")
		return 0, err
	}
	s.Top--
	return s.Body[s.Top+1], nil
}
