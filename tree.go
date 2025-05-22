package tree

import "cmp"

type Tree[T cmp.Ordered] interface {
	Root() Node[T]
	Size() int
	Insert(value T) error
	MustInsert(value T)
	Count(value T) int
}

type Node[T cmp.Ordered] interface {
	Parent() Node[T]
	Value() T
	Left() Node[T]
	Right() Node[T]
	Count() int
}
