package tree

import (
	"cmp"
	"errors"
)

const (
	COLOR_RED   = "red"
	COLOR_BLACK = "black"
)

func getColor[T cmp.Ordered](n *RBTNode[T]) string {
	if n == nil {
		return COLOR_BLACK
	}
	return n.color
}

type RBT[T cmp.Ordered] struct {
	root *RBTNode[T]
	size int
}

func panicIfNilRBT[T cmp.Ordered](t *RBT[T]) {
	if t == nil {
		panic("nil tree")
	}
}

func panicIfNilRBTNode[T cmp.Ordered](n *RBTNode[T]) {
	if n == nil {
		panic("nil node")
	}
}

func NewRBT[T cmp.Ordered]() *RBT[T] {
	return &RBT[T]{
		size: 0,
		root: nil,
	}
}

func (t *RBT[T]) Root() Node[T] {
	panicIfNilRBT(t)

	if t.root == nil {
		return nil
	}
	return t.root
}
func (t *RBT[T]) Size() int {
	panicIfNilRBT(t)

	return t.size
}

func (t *RBT[T]) Insert(value T) error {
	panicIfNilRBT(t)

	if t.root == nil {
		t.root = &RBTNode[T]{
			parent: nil,
			left:   nil,
			right:  nil,
			value:  value,
			color:  COLOR_BLACK,
		}
		t.size++
		return nil
	}

	var y *RBTNode[T] = nil
	x := t.root
	z := &RBTNode[T]{
		value: value,
	}

	for x != nil {
		y = x
		if z.value < x.value {
			x = x.left
		} else if z.value > x.value {
			x = x.right
		} else {
			return errors.New("value already exists")
		}
	}
	z.parent = y

	if y == nil {
		t.root = z
	} else if z.value < y.value {
		y.left = z
	} else {
		y.right = z
	}
	z.left = nil
	z.right = nil
	z.color = COLOR_RED

	// so far this was regular insertion. now rebalance

	insertFixup(t, z)

	return nil
}

func (t *RBT[T]) Delete(value T) error {
	panic("unimplemetned")
}

func (t *RBT[T]) String() string {
	panicIfNilRBT(t)

	return FormatTree(t, string(FormatHorizontal))
}

func (t *RBT[T]) Count(value T) int {
	panicIfNilRBT(t)

	if t.root == nil {
		return 0
	}
	c := t.root
	for c != nil {
		if value < c.value {
			c = c.left
		} else if value > c.value {
			c = c.right
		} else {
			return c.Count()
		}
	}
	return 0
}

type RBTNode[T cmp.Ordered] struct {
	parent *RBTNode[T]
	left   *RBTNode[T]
	right  *RBTNode[T]
	value  T
	color  string
}

func (n *RBTNode[T]) Parent() Node[T] {
	panicIfNilNode(n)

	if n.parent == nil {
		return nil
	}
	return n.parent
}
func (n *RBTNode[T]) Left() Node[T] {
	panicIfNilNode(n)

	if n.left == nil {
		return nil
	}
	return n.left
}
func (n *RBTNode[T]) Right() Node[T] {
	panicIfNilNode(n)

	if n.right == nil {
		return nil
	}
	return n.right
}
func (n *RBTNode[T]) Value() T {
	panicIfNilNode(n)

	return n.value
}
func (n *RBTNode[T]) Count() int {
	panicIfNilNode(n)

	return 1
}

func (n *RBTNode[T]) Color() string {
	panicIfNilNode(n)

	return n.color
}

func (n *RBTNode[T]) ttycolor() string {
	panicIfNilNode(n)

	return n.color
}

func leftRotate[T cmp.Ordered](t *RBT[T], x *RBTNode[T]) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func rightRotate[T cmp.Ordered](t *RBT[T], y *RBTNode[T]) {
	x := y.left
	y.left = x.right
	if x.right != nil {
		x.right.parent = y
	}
	x.parent = y.parent
	if y.parent == nil {
		t.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}
	x.right = y
	y.parent = x
}

func insertFixup[T cmp.Ordered](t *RBT[T], z *RBTNode[T]) {
	for getColor(z.parent) == COLOR_RED {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			if getColor(y) == COLOR_RED {
				z.parent.color = COLOR_BLACK
				y.color = COLOR_BLACK
				z.parent.parent.color = COLOR_RED
				z = z.parent.parent
			} else if z == z.parent.right {
				z = z.parent
				leftRotate(t, z)
			} else {
				z.parent.color = COLOR_BLACK
				z.parent.parent.color = COLOR_RED
				rightRotate(t, z.parent.parent)
			}
		} else {
			y := z.parent.parent.left
			if getColor(y) == COLOR_RED {
				z.parent.color = COLOR_BLACK
				y.color = COLOR_BLACK
				z.parent.parent.color = COLOR_RED
				z = z.parent.parent
			} else if z == z.parent.left {
				z = z.parent
				rightRotate(t, z)
			} else {
				z.parent.color = COLOR_BLACK
				z.parent.parent.color = COLOR_RED
				leftRotate(t, z.parent.parent)
			}
		}
	}
	t.root.color = COLOR_BLACK
}
