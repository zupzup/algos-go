package binarysearchtree

type Tree struct {
	left  *Tree
	right *Tree
	value int
}

func New(value int) *Tree {
	return &Tree{nil, nil, value}
}

func (t *Tree) Insert(value int) {
	if value < t.value {
		if t.left != nil {
			t.left.Insert(value)
		} else {
			t.left = New(value)
		}
	}
	if value >= t.value {
		if t.right != nil {
			t.right.Insert(value)
		} else {
			t.right = New(value)
		}
	}
}

func (t *Tree) Search(value int) *Tree {
	if value == t.value {
		return t
	}
	if value < t.value && t.left != nil {
		return t.left.Search(value)
	}
	if value >= t.value && t.right != nil {
		return t.right.Search(value)
	}
	return nil
}

func (t *Tree) Size() int {
	res := 1
	if t.left != nil {
		res += t.left.Size()
	}
	if t.right != nil {
		res += t.right.Size()
	}
	return res
}
