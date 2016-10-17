package quicksort

import (
	"fmt"
	"reflect"
	"testing"
)

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
		"dupes",
		[]int{1, 2, 2, 5, 3, 3, 6},
		[]int{1, 2, 2, 3, 3, 5, 6},
	},
	{
		"negs",
		[]int{-1, 2, -5, 3, 6},
		[]int{-5, -1, 2, 3, 6},
	},
	{
		"long",
		[]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7},
		[]int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16},
	},
	{
		"longer",
		[]int{23, 1, 6, 19, 14, 18, 8, 24, 15},
		[]int{1, 6, 8, 14, 15, 18, 19, 23, 24},
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

func TestQuickSort(t *testing.T) {
	for _, tc := range testsSort {
		t.Run(fmt.Sprintf("Test: %s", tc.description), func(t *testing.T) {
			res := Sort(tc.in)
			if !reflect.DeepEqual(res, tc.result) {
				t.Errorf("Error, actual: %v expected: %v", res, tc.result)
				return
			}
		})
	}
}
