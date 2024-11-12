package tree

type TreeNode[T comparable] interface {
	Root()
	Parent()
	Left()
	Right()
	Key() T
}
