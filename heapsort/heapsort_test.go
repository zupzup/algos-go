package heapsort

import (
	"fmt"
	"reflect"
	"testing"
)

var testsNew = []struct {
	description string
	in          []int
	result      []int
}{
	{
		"basic",
		[]int{1, 2, 5, 3, 6},
		[]int{6, 5, 2, 3, 1},
	},
	{
		"basiclong",
		[]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7},
		[]int{16, 14, 9, 10, 8, 1, 4, 2, 3, 7},
	},
	{
		"basiclonger",
		[]int{23, 1, 6, 19, 14, 18, 8, 24, 15},
		[]int{24, 23, 18, 19, 15, 6, 8, 1, 14},
	},
	{
		"empty",
		[]int{},
		[]int{},
	},
	{
		"nil",
		nil,
		nil,
	},
}

var testsSort = []struct {
	description string
	in          []int
	result      []int
}{
	{
		"basic",
		[]int{1, 2, 5, 3, 6},
		[]int{1, 2, 3, 5, 6},
	},
	{
		"basiclong",
		[]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7},
		[]int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16},
	},
	{
		"basiclonger",
		[]int{23, 1, 6, 19, 14, 18, 8, 24, 15},
		[]int{1, 6, 8, 14, 15, 18, 19, 23, 24},
	},
}

func TestBuildMaxHeap(t *testing.T) {
	for _, tc := range testsNew {
		t.Run(fmt.Sprintf("Test: %s", tc.description), func(t *testing.T) {
			res := BuildMaxHeap(tc.in)
			if !reflect.DeepEqual(res, tc.result) {
				t.Errorf("Error, actual: %v expected: %v", res, tc.result)
				return
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	for _, tc := range testsSort {
		t.Run(fmt.Sprintf("Test: %s", tc.description), func(t *testing.T) {
			res := HeapSort(tc.in)
			if !reflect.DeepEqual(res, tc.result) {
				t.Errorf("Error, actual: %v expected: %v", res, tc.result)
				return
			}
		})
	}
}
