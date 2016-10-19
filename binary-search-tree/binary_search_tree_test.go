package binarysearchtree

import (
	"testing"
)

//TESTING: create, insert, search, size

func TestNewTree(t *testing.T) {
	tree := New(10)
	if tree.Size() != 1 {
		t.Errorf("Error, actual: %v expected: %v", tree.Size(), 1)
		return
	}
}

func TestInsertSearch(t *testing.T) {
	tree := New(10)
	tree.Insert(5)
	tree.Insert(11)
	t2 := tree.Search(11)
	if t2.value != 11 || tree.Size() != 3 {
		t.Errorf("Error, actual: %v %v  expected: %v %v ", t2.value, tree.Size(), 11, 3)
		return
	}
}

func TestEmptySearch(t *testing.T) {
	tree := New(10)
	tree.Insert(11)
	t2 := tree.Search(100)
	if t2 != nil {
		t.Errorf("Error, actual: %v expected: %v ", t2, nil)
		return
	}
}

func TestInsertMultiple(t *testing.T) {
	tree := New(10)
	tree.Insert(5)
	tree.Insert(11)
	t2 := tree.Search(5)
	t2.Insert(3)
	t2.Insert(7)
	if t2.value != 5 || tree.Size() != 5 {
		t.Errorf("Error, actual: %v %v  expected: %v %v ", t2.value, tree.Size(), 5, 5)
		return
	}
}

func TestInsertMultiple2(t *testing.T) {
	tree := New(10)
	tree.Insert(5)
	tree.Insert(11)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(15)
	if tree.Size() != 6 {
		t.Errorf("Error, actual: %v expected: %v", tree.Size(), 5)
		return
	}
}
