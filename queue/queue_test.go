package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := New()
	if q == nil {
		t.Errorf("Error, actual: %v expected: %v", q, nil)
		return
	}
}

func TestEnqueueDequeue(t *testing.T) {
	q := New()
	q.Enqueue(1)
	i, err := q.Dequeue()
	if i != 1 || err != nil {
		t.Errorf("Error, actual: %v expected: %v", i, 1)
		return
	}
}

func TestMultiple(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	i, err := q.Dequeue()
	i2, err := q.Dequeue()
	i3, err := q.Dequeue()
	i4, err := q.Dequeue()
	if i != 4 || i2 != 2 || i3 != 3 || i4 != 4 || err != nil {
		t.Errorf("Error, actual: %v %v %v %v expected: %v %v %v %v", i, i2, i3, i4, 4, 2, 3, 4)
		return
	}
}

func TestDequeueEmpty(t *testing.T) {
	q := New()
	_, err := q.Dequeue()
	if err == nil {
		t.Errorf("Error expected on Popping from empty Queue")
		return
	}
}
