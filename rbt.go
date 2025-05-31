package tree

import (
	"cmp"
	"errors"
)

const (
	COLOR_RED   = "red"
	COLOR_BLACK = "black"
)

// sentinel value
func sentinel[T cmp.Ordered]() *RBTNode[T] {
	return &RBTNode[T]{
		color: COLOR_BLACK,
	}
}

func getColor[T cmp.Ordered](n *RBTNode[T]) string {
	if n == nil {
		return COLOR_BLACK
	}
	return n.color
}

type RBT[T cmp.Ordered] struct {
	root *RBTNode[T]
	size int
	tnil *RBTNode[T]
}

func NewRBT[T cmp.Ordered]() *RBT[T] {
	tnil := sentinel[T]()
	return &RBT[T]{
		size: 0,
		root: tnil,
		tnil: tnil,
	}
}

func (t *RBT[T]) Root() Node[T] {
	panicIfNilTree(t)

	if t.root.isSentinel() {
		return nil
	}
	return t.root
}
func (t *RBT[T]) Size() int {
	panicIfNilTree(t)

	return t.size
}

func (t *RBT[T]) Insert(value T) error {
	panicIfNilTree(t)

	if t.root == t.tnil {
		t.root = &RBTNode[T]{
			parent: t.tnil,
			left:   t.tnil,
			right:  t.tnil,
			value:  value,
			color:  COLOR_BLACK,
		}
		t.size = 1
		return nil
	}

	y := t.tnil
	x := t.root
	z := &RBTNode[T]{
		value: value,
	}

	for x != t.tnil {
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

	if y == t.tnil {
		t.root = z
	} else if z.value < y.value {
		y.left = z
	} else {
		y.right = z
	}
	z.left = t.tnil
	z.right = t.tnil
	z.color = COLOR_RED

	insertFixup(t, z)
	t.size++

	return nil
}

func (t *RBT[T]) Delete(value T) error {
	panicIfNilTree(t)

	// uninitialized tree
	if t.root == nil {
		return errors.New("value not found")
	}

	// find z
	z := t.root
	for z != t.tnil {
		if value < z.value {
			z = z.left
		} else if value > z.value {
			z = z.right
		} else {
			break
		}
	}
	if z == t.tnil {
		return errors.New("value not found")
	}

	y := z
	yorigcolor := getColor(y)
	var x *RBTNode[T]

	if z.left == t.tnil {
		x = z.right
		rbtransplant(t, z, z.right)
	} else if z.right == t.tnil {
		x = z.left
		rbtransplant(t, z, z.left)
	} else {
		y = treeMinimumRbt(t, z.right)
		yorigcolor = getColor(y)
		x = y.right
		if y.parent == z {
			x.parent = y
		} else {
			rbtransplant(t, y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		rbtransplant(t, z, y)
		y.left = z.left
		y.left.parent = y
		y.color = getColor(z)
	}
	if yorigcolor == COLOR_BLACK {
		rbDeleteFixup(t, x)
	}
	t.size--
	return nil
}

func rbDeleteFixup[T cmp.Ordered](t *RBT[T], x *RBTNode[T]) {
	for x != t.root && getColor(x) == COLOR_BLACK {
		if x == x.parent.left {
			w := x.parent.right
			if getColor(w) == COLOR_RED {
				w.color = COLOR_BLACK
				x.parent.color = COLOR_RED
				leftRotate(t, x.parent)
				w = x.parent.right
			}
			if getColor(w.left) == COLOR_BLACK && getColor(w.right) == COLOR_BLACK {
				w.color = COLOR_RED
				x = x.parent
			} else if getColor(w.right) == COLOR_BLACK {
				w.left.color = COLOR_BLACK
				w.color = COLOR_RED
				rightRotate(t, w)
				w = x.parent.right
			} else {
				w.color = getColor(x.parent)
				x.parent.color = COLOR_BLACK
				w.right.color = COLOR_BLACK
				leftRotate(t, x.parent)
				x = t.root
			}
		} else {
			w := x.parent.left
			if getColor(w) == COLOR_RED {
				w.color = COLOR_BLACK
				x.parent.color = COLOR_RED
				rightRotate(t, x.parent)
				w = x.parent.left
			}
			if getColor(w.right) == COLOR_BLACK && getColor(w.left) == COLOR_BLACK {
				w.color = COLOR_RED
				x = x.parent
			} else if getColor(w.left) == COLOR_BLACK {
				w.right.color = COLOR_BLACK
				w.color = COLOR_RED
				leftRotate(t, w)
				w = x.parent.left
			} else {
				w.color = getColor(x.parent)
				x.parent.color = COLOR_BLACK
				w.left.color = COLOR_BLACK
				rightRotate(t, x.parent)
				x = t.root
			}
		}
	}
	x.color = COLOR_BLACK
}

func (t *RBT[T]) String() string {
	panicIfNilTree(t)

	return FormatTree(t, string(FormatHorizontal))
}

func (t *RBT[T]) Count(value T) int {
	panicIfNilTree(t)

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
	panicIfNilOrSentinelNode(n)

	if n.parent.isSentinel() {
		return nil
	}
	return n.parent
}
func (n *RBTNode[T]) Left() Node[T] {
	panicIfNilOrSentinelNode(n)

	if n.left.isSentinel() {
		return nil
	}
	return n.left
}
func (n *RBTNode[T]) Right() Node[T] {
	panicIfNilOrSentinelNode(n)

	if n.left.isSentinel() {
		return nil
	}
	return n.right
}
func (n *RBTNode[T]) Value() T {
	panicIfNilOrSentinelNode(n)

	return n.value
}
func (n *RBTNode[T]) Count() int {
	panicIfNilOrSentinelNode(n)

	return 1
}

func (n *RBTNode[T]) ttycolor() string {
	panicIfNilNode(n)

	return n.color
}

func (n *RBTNode[T]) isSentinel() bool {
	// all children of non-sentinel rbt nodes are represented as tnil.
	return n.left == nil && n.right == nil
}

// panicIfNilOrSentinelNode will panic if the current node is nil or a sentinel.
// Sentinels are included because they are an implementation detail to simplify
// edge case handling in algorithms, but in essence they still represent nil
// nodes and should be exposed as nil from public methods.
//
// Note that this method is typically called for public methods, as algorithms
// use the internal representation directly.
func panicIfNilOrSentinelNode[T cmp.Ordered](n *RBTNode[T]) {
	if n == nil {
		panic("nil node")
	} else if n.isSentinel() {
		panic("nil node")
	}
}

func leftRotate[T cmp.Ordered](t *RBT[T], x *RBTNode[T]) {
	y := x.right
	x.right = y.left
	if y.left != t.tnil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == t.tnil {
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
	if x.right != t.tnil {
		x.right.parent = y
	}
	x.parent = y.parent
	if y.parent == t.tnil {
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

// transplant replaces one subtree with another subtree
func rbtransplant[T cmp.Ordered](t *RBT[T], u *RBTNode[T], v *RBTNode[T]) {
	// u is root
	if u.parent == t.tnil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

func treeMinimumRbt[T cmp.Ordered](t *RBT[T], x *RBTNode[T]) *RBTNode[T] {
	for x.left != t.tnil {
		x = x.left
	}
	return x
}
