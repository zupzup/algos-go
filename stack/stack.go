package stack

import "errors"

type Stack struct {
	data []int
	top  int
}

func New() *Stack {
	return &Stack{[]int{}, 0}
}

func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return 0, errors.New("Stack underflow")
	}
	s.top = s.top - 1
	return s.data[s.top], nil
}

func (s *Stack) Push(element int) {
	s.top = s.top + 1
	s.data = append(s.data, element)
}

func (s *Stack) Empty() bool {
	return s.top == 0
}

func (s *Stack) Size() int {
	return len(s.data)
}
