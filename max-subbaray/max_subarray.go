package maxsubarray

import "math"

func FindMaxSubarray(in []int) (low, high, sum int) {
	if in == nil || len(in) == 0 {
		return -1, -1, -1
	}
	return findMaximumSubarray(in, 0, len(in)-1)
}

func findMaximumSubarray(in []int, low, high int) (lowRes, highRes, sumRes int) {
	if high == low {
		return low, high, in[low]
	} else {
		mid := int(math.Floor(float64((low + high) / 2)))
		leftLow, leftHigh, leftSum := findMaximumSubarray(in, low, mid)
		rightLow, rightHigh, rightSum := findMaximumSubarray(in, mid+1, high)
		crossLow, crossHigh, crossSum := findMaxCrossing(in, low, mid, high)
		if leftSum >= rightSum && leftSum >= crossSum {
			return leftLow, leftHigh, leftSum
		} else if rightSum >= leftSum && rightSum >= crossSum {
			return rightLow, rightHigh, rightSum
		} else {
			return crossLow, crossHigh, crossSum
		}
	}
}

func findMaxCrossing(in []int, low, mid, high int) (maxLeft, maxRight, fullSum int) {
	leftSum := int(math.MinInt64)
	sum := 0
	for i := mid; i > low; i-- {
		sum = sum + in[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	rightSum := int(math.MinInt64)
	sum = 0
	for i := mid + 1; i < high; i++ {
		sum = sum + in[i]
		if sum > rightSum {
			rightSum = sum
			maxRight = i
		}
	}
	return maxLeft, maxRight, leftSum + rightSum
}
