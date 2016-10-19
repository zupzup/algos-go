package linkedlist

type LinkedList struct {
	head *LinkedElement
	tail *LinkedElement
	size int
}

type LinkedElement struct {
	data int
	prev *LinkedElement
	next *LinkedElement
}

func New() *LinkedList {
	return &LinkedList{nil, nil, 0}
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) Search(element int) *LinkedElement {
	x := l.head
	for x != nil && x.data != element {
		x = x.next
	}
	return x
}

func (l *LinkedList) Insert(element int) {
	x := &LinkedElement{element, nil, l.head}
	if l.head != nil {
		l.head.prev = x
	}
	l.head = x
	l.size += 1
}

func (l *LinkedList) Delete(element *LinkedElement) {
	if element.prev != nil {
		element.prev.next = element.next
	} else {
		l.head = element.next
	}
	if element.next != nil {
		element.next.prev = element.prev
	}
	l.size -= 1
}
