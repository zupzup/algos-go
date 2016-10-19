package binarysearchtree

type Tree struct {
	left   *Tree
	right  *Tree
	value  int
	parent *Tree
}

func New(value int) *Tree {
	return &Tree{nil, nil, value, nil}
}

func (t *Tree) Insert(value int) {
	if value < t.value {
		if t.left != nil {
			t.left.Insert(value)
		} else {
			left := New(value)
			left.parent = t
			t.left = left
		}
	}
	if value >= t.value {
		if t.right != nil {
			t.right.Insert(value)
		} else {
			right := New(value)
			right.parent = t
			t.right = right
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

func (t *Tree) Minimum() *Tree {
	for t.left != nil {
		t = t.left
	}
	return t
}

func (t *Tree) Delete(z *Tree) {
	if z.left == nil {
		transplant(t, z, z.right)
	} else if z.right == nil {
		transplant(t, z, z.left)
	} else {
		y := z.right.Minimum()
		if y.parent != z {
			transplant(t, y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		transplant(t, z, y)
		y.left = z.left
		y.left.parent = y
	}
}

func transplant(root, u, v *Tree) {
	if u.parent == nil {
		root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}
