package heapsort

import "math"

func BuildMaxHeap(in []int) []int {
	if in == nil {
		return nil
	}
	heapSize := len(in)
	for i := int(math.Floor(float64(len(in) / 2))); i >= 0; i-- {
		maxHeapify(in, i, heapSize)
	}
	return in
}

func HeapSort(in []int) []int {
	heap := BuildMaxHeap(in)
	heapSize := len(heap)
	for i := len(heap) - 1; i >= 0; i-- {
		heap[0], heap[i] = heap[i], heap[0]
		heapSize -= 1
		maxHeapify(heap, 0, heapSize)
	}
	return heap
}

func maxHeapify(in []int, i, heapSize int) {
	l := left(i)
	r := right(i)
	largest := 0
	if l < heapSize && in[l] > in[i] {
		largest = l
	} else {
		largest = i
	}
	if r < heapSize && in[r] > in[largest] {
		largest = r
	}
	if largest != i {
		in[i], in[largest] = in[largest], in[i]
		maxHeapify(in, largest, heapSize)
	}
}

func parent(i int) int {
	return int(math.Floor(float64((i - 1) / 2)))
}

func left(i int) int {
	return 2 * i
}

func right(i int) int {
	return 2*i + 1
}
