package tree

// BstNode represents a regular binary search tree node.
type BstNode[T comparable] struct {
	Left  *BstNode[T]
	Right *BstNode[T]
	P     *BstNode[T]
	Key   T
	Root  *BstNode[T] // may be removed later if not useful in algorithms
}

func NewBST[T comparable](Key T, opts ...BstreeOpt[T]) *BstNode[T] {
	return nil
}

type BstreeOpt[T comparable] func(n *BstNode[T])

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
