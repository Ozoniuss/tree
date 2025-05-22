package tree

import (
	"cmp"
	"fmt"
)

type BST[T cmp.Ordered] struct {
	root *BSTNode[T]
	size int
}

func NewBST[T cmp.Ordered]() *BST[T] {
	return &BST[T]{
		size: 0,
		root: nil,
	}
}

func (t *BST[T]) Root() Node[T] {
	return t.root
}
func (t *BST[T]) Size() int {
	return t.size
}

func (t *BST[T]) Insert(value T) error {
	if t.root == nil {
		t.root = &BSTNode[T]{
			parent: nil,
			left:   nil,
			right:  nil,
			value:  value,
		}
		t.size++
		return nil
	}

	var y *BSTNode[T]
	c := t.root
	for c != nil {
		y = c
		if value < c.value {
			c = c.left
		} else if value > c.value {
			c = c.right
		} else {
			return fmt.Errorf("value already exists")
		}
	}
	if value < y.value {
		y.left = &BSTNode[T]{
			parent: y,
			left:   nil,
			right:  nil,
			value:  value,
		}
	} else {
		y.right = &BSTNode[T]{
			parent: y,
			left:   nil,
			right:  nil,
			value:  value,
		}
	}
	t.size++
	return nil
}

func (t *BST[T]) MustInsert(value T) {
	t.Insert(value)
}

func (t *BST[T]) Count(value T) int {
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

type BSTNode[T cmp.Ordered] struct {
	parent *BSTNode[T]
	left   *BSTNode[T]
	right  *BSTNode[T]
	value  T
}

func (n *BSTNode[T]) Parent() Node[T] {
	return n.parent
}
func (n *BSTNode[T]) Left() Node[T] {
	return n.left
}
func (n *BSTNode[T]) Right() Node[T] {
	return n.right
}
func (n *BSTNode[T]) Value() T {
	return n.value
}
func (n *BSTNode[T]) Count() int {
	return 1
}
