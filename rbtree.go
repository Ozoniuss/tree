package rbtree

/*
source: Thomas H. Cormen, Charles E. Leiserson, Ronald L. Rivest, Clifford Stein - "Introduction to Algorithms 3rd Edition" (2009)
A Red-Black Tree have the following properties:

1. Every node is either red or black.
2. The root is black.
3. Every leaf (NIL) is black.
4. If a node is red, then both its children are black.
5. For each node, all simple paths from the node to descendant leaves contain the same number of black nodes.
*/

/*
Left rotation example: lrot(10)
  5         |      10
 / \        |     / \
2  10       |    5   12
   / \      |   / \
  8   12    |  2   8
 / \        |     / \
6   9       |    6   9
*/

/*
Right rotation example (reverse of left rotation): rrot(10)
    10      |    5
   / \      |   / \
  5   12    |  2  10
 / \        |     / \
2   8       |    8   12
   / \      |   / \
  6   9     |  6   9
*/

type RbtNode[T comparable] struct {
	Root  *RbtNode[T]
	Color int
	Key   T
	Left  *RbtNode[T]
	Right *RbtNode[T]
	P     *RbtNode[T]
}

func (n *RbtNode[T]) GetColor() int {
	if n == nil {
		return 0
	}
	return n.Color
}
