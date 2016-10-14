package mergesort

import (
	"reflect"
	"testing"
)

func TestNil(t *testing.T) {
	actual := Sort(nil)
	if actual != nil {
		t.Errorf("Results don't match: actual %v expected %v", actual, nil)
	}
}

func TestEmptyArray(t *testing.T) {
	actual := Sort([]int{})
	expected := []int{}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Results don't match: actual %v expected %v", actual, expected)
	}
}

func TestIsSorted(t *testing.T) {
	actual := Sort([]int{8, 3, 1, 9, 2})
	expected := []int{1, 2, 3, 8, 9}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Results don't match: actual %v expected %v", actual, expected)
	}
}

func TestIsSorted2(t *testing.T) {
	actual := Sort([]int{8, 3, 1, 9, 2, 5, 6, 7})
	expected := []int{1, 2, 3, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Results don't match: actual %v expected %v", actual, expected)
	}
}
