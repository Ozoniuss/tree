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
	if t == nil {
		panic("root: nil tree")
	}
	if t.root == nil {
		return nil
	}
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

	var y *BSTNode[T] = nil
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

func (t *BST[T]) Delete(value T) error {
	panic("unimplemented")
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
	if n == nil {
		panic("parent: nil node")
	}
	if n.parent == nil {
		return nil
	}
	return n.parent
}
func (n *BSTNode[T]) Left() Node[T] {
	if n == nil {
		panic("left: nil node")
	}
	if n.left == nil {
		return nil
	}
	return n.left
}
func (n *BSTNode[T]) Right() Node[T] {
	if n == nil {
		panic("right: nil node")
	}
	if n.right == nil {
		return nil
	}
	return n.right
}
func (n *BSTNode[T]) Value() T {
	if n == nil {
		panic("value: nil node")
	}
	return n.value
}
func (n *BSTNode[T]) Count() int {
	if n == nil {
		panic("count: nil node")
	}
	return 1
}
