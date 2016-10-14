package insertion

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

func TestIsSortedDoubles(t *testing.T) {
	actual := Sort([]int{8, 8, 3, 1, 9, 2})
	expected := []int{1, 2, 3, 8, 8, 9}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Results don't match: actual %v expected %v", actual, expected)
	}
}

func TestIsSortedNegatives(t *testing.T) {
	actual := Sort([]int{-8, 8, 3, 1, 9, 2})
	expected := []int{-8, 1, 2, 3, 8, 9}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Results don't match: actual %v expected %v", actual, expected)
	}
}
