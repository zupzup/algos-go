package mergesort

import "math"

func Sort(in []int) []int {
	if in == nil {
		return nil
	}
	MergeSort(in, 0, len(in))
	return in
}

func MergeSort(in []int, p, r int) {
	if p < r {
		q := int(math.Floor(float64((p + r) / 2)))
		MergeSort(in, p, q)
		MergeSort(in, q+1, r)
		Merge(in, p, q, r)
	}
}

func Merge(in []int, p, q, r int) {
	lenLeft := q - p
	lenRight := r - q
	left := make([]int, lenLeft+1)
	right := make([]int, lenRight+1)
	for i := 0; i < lenLeft; i++ {
		left[i] = in[p+i]
	}
	for i := 0; i < lenRight; i++ {
		right[i] = in[q+i]
	}
	left[lenLeft] = int(^uint(0) >> 1)
	right[lenRight] = int(^uint(0) >> 1)
	i := 0
	j := 0
	for k := p; k < r; k++ {
		if left[i] <= right[j] {
			in[k] = left[i]
			i = i + 1
		} else {
			in[k] = right[j]
			j = j + 1
		}
	}

}
