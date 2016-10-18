package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	s := New()
	if s.Size() != 0 {
		t.Errorf("Error, actual: %v expected: %v", s.Size(), 0)
		return
	}
}

func TestPushPop(t *testing.T) {
	s := New()
	s.Push(1)
	i, err := s.Pop()
	if i != 1 || err != nil {
		t.Errorf("Error, actual: %v expected: %v", i, 1)
		return
	}
}

func TestPopEmpty(t *testing.T) {
	s := New()
	_, err := s.Pop()
	if err == nil {
		t.Errorf("Error expected on Popping from empty stack")
		return
	}
}
