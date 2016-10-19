package linkedlist

import (
	"testing"
)

func TestNewList(t *testing.T) {
	l := New()
	if l.Size() != 0 {
		t.Errorf("Error, actual: %v expected: %v", l.Size(), 0)
		return
	}
}

func TestAddRemove(t *testing.T) {
	l := New()
	l.Insert(5)
	elem := l.Search(5)
	l.Delete(elem)
	if l.Size() != 0 {
		t.Errorf("Error, actual: %v expected: %v", l.Size(), 0)
		return
	}
}

func TestMultiple(t *testing.T) {
	l := New()
	l.Insert(5)
	l.Insert(2)
	l.Insert(3)
	elem1 := l.Search(5)
	elem2 := l.Search(2)
	elem3 := l.Search(3)
	l.Delete(elem2)
	if l.Size() != 2 || elem1.data != 5 || elem3.data != 3 {
		t.Errorf("Error, actual: %v %v %v expected: %v %v %v", l.Size(), elem1.data, elem3.data, 5, 2, 3)
		return
	}
}
