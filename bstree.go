package tree

import (
	"fmt"
	"strings"
)

// BstNode represents a regular binary search tree node.
type BstNode[T comparable] struct {
	Left  *BstNode[T]
	Right *BstNode[T]
	P     *BstNode[T]
	Key   T
	Root  *BstNode[T] // may be removed later if not useful in algorithms
}

func NewBST[T comparable](key T, opts ...BstreeOpt[T]) *BstNode[T] {
	root := &BstNode[T]{
		Left:  nil,
		Right: nil,
		P:     nil,
		Key:   key,
		Root:  nil,
	}
	root.Root = root

	for _, o := range opts {
		o(root)
	}

	return root
}

// Format returns a string representation of the tree, based on its layout.
func (n *BstNode[T]) Format(layout string) string {

	if n == nil {
		return "*"
	}

	out := n.formatLinuxTree("", nil, true)
	return strings.TrimRight(out, "\n")
}

// TraverseInorder traverses the tree using an inorder traversal, applying the
// supplied function to each node.
func (n *BstNode[T]) TraverseInorder(f func(*BstNode[T])) {
	if n == nil {
		return
	}
	n.Left.TraverseInorder(f)
	f(n)
	n.Right.TraverseInorder(f)
}

// TraversePreorder traverses the tree using an inorder traversal, applying the
// supplied function to each node.
func (n *BstNode[T]) TraversePreorder(f func(*BstNode[T])) {
	if n == nil {
		return
	}
	f(n)
	n.Left.TraversePreorder(f)
	n.Right.TraversePreorder(f)
}

// TraversePostorder traverses the tree using an inorder traversal, applying the
// supplied function to each node.
func (n *BstNode[T]) TraversePostorder(f func(*BstNode[T])) {
	if n == nil {
		return
	}
	n.Left.TraversePostorder(f)
	n.Right.TraversePostorder(f)
	f(n)
}

func (n *BstNode[T]) formatLinuxTree(prefix string, parent *BstNode[T], isRoot bool) string {

	if n == nil && parent == nil {
		panic("nil tree")
	}

	// shadow nil node of leaf, no asterisk
	if n == nil && parent.Left == nil && parent.Right == nil {
		return ""
	}
	// left node should have asterisk
	if n == nil && parent.Left == nil {
		return fmt.Sprintf("%s%s %v\n", prefix, PREFIX_LEFT, "*")
	}
	// left node should have asterisk
	if n == nil && parent.Right == nil {
		return fmt.Sprintf("%s%s %v\n", prefix, PREFIX_RIGHT, "*")
	}

	out := ""
	newprefix := prefix

	if isRoot {
		out += fmt.Sprintf("%v\n", n.Key)
	} else if n == n.P.Left {
		out += fmt.Sprintf("%s%s %v\n", prefix, PREFIX_LEFT, n.Key)
		newprefix = prefix + EXTRA_LEFT
	} else if n == n.P.Right {
		out += fmt.Sprintf("%s%s %v\n", prefix, PREFIX_RIGHT, n.Key)
		newprefix = prefix + EXTRA_RIGHT
	}

	out += n.Left.formatLinuxTree(newprefix, n, false)
	out += n.Right.formatLinuxTree(newprefix, n, false)
	return out
}

type BstreeOpt[T comparable] func(n *BstNode[T])

// for now behaviour is undefined for options below

func WithLeftChild[T comparable](c *BstNode[T]) BstreeOpt[T] {
	return func(n *BstNode[T]) {
		n.Left = c
		c.P = n
		c.Root = n.Root
	}
}

func WithRightChild[T comparable](c *BstNode[T]) BstreeOpt[T] {
	return func(n *BstNode[T]) {
		n.Right = c
		c.P = n
		c.Root = n.Root
	}
}
