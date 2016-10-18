package queue

import "errors"
import "fmt"

type Queue struct {
	data       []int
	head, tail int
}

func New() *Queue {
	q := make([]int, 3)
	return &Queue{q, 0, 0}
}

func (q *Queue) Enqueue(element int) {
	q.data[q.tail] = element
	if q.tail == len(q.data)-1 {
		q.tail = 0
	} else {
		q.tail = q.tail + 1
	}
}

func (q *Queue) Dequeue() (int, error) {
	if q.head == q.tail && q.head == 0 {
		return 0, errors.New("Queue is empty")
	}
	x := q.data[q.head]
	if q.head == len(q.data)-1 {
		q.head = 0
	} else {
		q.head = q.head + 1
	}
	return x, nil
}
