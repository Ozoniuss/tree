package tree

import (
	"fmt"
	"strings"
)

const (
	PREFIX_LEFT  = "├──"
	PREFIX_RIGHT = "└──"

	EXTRA_LEFT  = "│   "
	EXTRA_RIGHT = "    "
)

/*
ToStringPreorder formats the tree nicely as a string.

	    4
	   / \
	  2   10
	 /   / \
	1   8   12
	   / \
	  6   9
	 /     \
	5       11

would get converted to

	4
	├── 2
	│   ├── 1
	│   └── *
	└── 10
		├── 8
		│   ├── 6
		│   │   ├── 5
		│   │   └── *
		│   └── 9
		│       ├── *
		│       └── 11
		└── 12
*/
func (n *RbtNode[T]) ToStringPreorder() string {

	if n == nil {
		return "*"
	}

	// copy tree because algorithm checks if tree is root
	c := &RbtNode[T]{
		Key:   n.Key,
		Right: n.Right,
		Left:  n.Left,
		P:     nil,
	}
	c.Root = c

	out := c._printPreorder("", nil)
	return strings.TrimRight(out, "\n")
}

func (n *RbtNode[T]) _printPreorder(prefix string, parent *RbtNode[T]) string {

	if n == nil && parent == nil {
		panic("nil node with nil parent")
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

	if n == n.Root {
		out += fmt.Sprintf("%v\n", n.Key)
	} else if n == n.P.Left {
		out += fmt.Sprintf("%s%s %v\n", prefix, PREFIX_LEFT, n.Key)
		newprefix = prefix + EXTRA_LEFT
	} else if n == n.P.Right {
		out += fmt.Sprintf("%s%s %v\n", prefix, PREFIX_RIGHT, n.Key)
		newprefix = prefix + EXTRA_RIGHT
	}

	out += n.Left._printPreorder(newprefix, n)
	out += n.Right._printPreorder(newprefix, n)
	return out
}
