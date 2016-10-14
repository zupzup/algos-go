package maxsubarray

import (
	// "reflect"
	"testing"
)

func TestNil(t *testing.T) {
	low, high, sum := FindMaxSubarray(nil)
	if low != -1 && high != -1 && sum != -1 {
		t.Errorf("Results don't match: actual %d expected %d", low, -1)
	}
}

func TestEmptyArray(t *testing.T) {
	low, high, sum := FindMaxSubarray([]int{})
	resLow, resHigh, resSum := -1, -1, -1
	if resLow != low || resHigh != high || resSum != sum {
		t.Errorf("Results don't match: actual %d %d %d expected %d %d %d", low, high, sum, resLow, resHigh, resSum)
	}
}

func TestRealCase(t *testing.T) {
	low, high, sum := FindMaxSubarray([]int{10, 15, 85, 19, 8, 2, 109})
	resLow, resHigh, resSum := 1, 5, 129
	if resLow != low || resHigh != high || resSum != sum {
		t.Errorf("Results don't match: actual %d %d %d expected %d %d %d", low, high, sum, resLow, resHigh, resSum)
	}
}

func TestRealCase2(t *testing.T) {
	low, high, sum := FindMaxSubarray([]int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7})
	resLow, resHigh, resSum := 7, 10, 43
	if resLow != low || resHigh != high || resSum != sum {
		t.Errorf("Results don't match: actual %d %d %d expected %d %d %d", low, high, sum, resLow, resHigh, resSum)
	}
}

func TestRealCase3(t *testing.T) {
	low, high, sum := FindMaxSubarray([]int{1, -4, 3, -4})
	resLow, resHigh, resSum := 2, 2, 3
	if resLow != low || resHigh != high || resSum != sum {
		t.Errorf("Results don't match: actual %d %d %d expected %d %d %d", low, high, sum, resLow, resHigh, resSum)
	}
}
