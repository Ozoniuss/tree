package tree

import (
	"fmt"
	"iter"
	"strings"
)

// BstNode represents a regular binary search tree node.
type BstNode[T comparable] struct {
	Left  *BstNode[T]
	Right *BstNode[T]
	P     *BstNode[T]
	Key   T
}

func NewBST[T comparable](key T, opts ...BstreeOpt[T]) *BstNode[T] {
	root := &BstNode[T]{
		Left:  nil,
		Right: nil,
		P:     nil,
		Key:   key,
	}

	for _, o := range opts {
		o(root)
	}

	return root
}

// Values returns a DFS (preorder) iterator over values in the tree.
func (n *BstNode[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		var traverse func(nd *BstNode[T]) bool
		traverse = func(nd *BstNode[T]) bool {
			if nd == nil {
				return true
			}
			if !yield(nd.Key) {
				return false
			}
			if !traverse(nd.Left) {
				return false
			}
			return traverse(nd.Right)
		}
		traverse(n)
	}
}

// All returns a DFS (preorder) iterator over nodes in the tree.
func (n *BstNode[T]) All() iter.Seq[*BstNode[T]] {
	return func(yield func(*BstNode[T]) bool) {
		var traverse func(nd *BstNode[T]) bool
		traverse = func(nd *BstNode[T]) bool {
			if nd == nil {
				return true
			}
			if !yield(nd) {
				return false
			}
			if !traverse(nd.Left) {
				return false
			}
			return traverse(nd.Right)
		}
		traverse(n)
	}
}

// Equal reports whether two trees are equal, checking that their structure and
// elements are identical.
func Equal[T comparable](t1, t2 *BstNode[T]) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil && t2 != nil {
		return false
	}
	if t1 != nil && t2 == nil {
		return false
	}

	if t1.Key != t2.Key {
		return false
	}
	if !Equal(t1.Left, t2.Left) {
		return false
	}
	if !Equal(t1.Right, t2.Right) {
		return false
	}
	return true
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

// TraversePreorder traverses the tree using a preorder traversal, applying the
// supplied function to each node.
func (n *BstNode[T]) TraversePreorder(f func(*BstNode[T])) {
	if n == nil {
		return
	}
	f(n)
	n.Left.TraversePreorder(f)
	n.Right.TraversePreorder(f)
}

// TraversePostorder traverses the tree using a postorder traversal, applying
// the supplied function to each node.
func (n *BstNode[T]) TraversePostorder(f func(*BstNode[T])) {
	if n == nil {
		return
	}
	n.Left.TraversePostorder(f)
	n.Right.TraversePostorder(f)
	f(n)
}

// TraverseLevelorder traverses the tree using a leve order, applying the
// supplied function to each node.
func (n *BstNode[T]) TraverseLevelorder(f func(*BstNode[T])) {
	if n == nil {
		return
	}

	q := make([]*BstNode[T], 0)
	q = append(q, n)
	for len(q) != 0 {
		top := q[0]
		q = q[1:]

		f(top)
		if top.Left != nil {
			q = append(q, top.Left)
		}
		if top.Right != nil {
			q = append(q, top.Right)
		}
	}
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
	}
}

func WithRightChild[T comparable](c *BstNode[T]) BstreeOpt[T] {
	return func(n *BstNode[T]) {
		n.Right = c
		c.P = n
	}
}
