package tree

import (
	"cmp"
	"fmt"
	"strings"
)

const (
	PREFIX_LEFT  = "├──"
	PREFIX_RIGHT = "└──"

	EXTRA_LEFT  = "│   "
	EXTRA_RIGHT = "    "
)

const (
	/*
	   FormatLinuxTree formats the tree nicely as a string.

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
	FormatLinuxTree = "FormatLinuxTree"
)

func FormatTree[T cmp.Ordered](t Tree[T]) string {
	if t == nil {
		return ""
	}

	if t.Root() == nil {
		return "*"
	}

	if t.Root().Left() == nil && t.Root().Right() == nil {
		return fmt.Sprint(t.Root().Value())
	}

	out := fmt.Sprintf("%v\n", t.Root().Value())
	prefix := []string{}

	type stkobj struct {
		n   Node[T]
		cnt int
	}

	stack := []*stkobj{}
	stack = append(stack, &stkobj{
		n:   t.Root(),
		cnt: 0,
	})

	for len(stack) != 0 {
		cobj := stack[len(stack)-1]
		n := cobj.n

		// processed both left and right
		if cobj.cnt >= 2 {
			stack = stack[:len(stack)-1]
			if len(prefix) != 0 {
				prefix = prefix[:len(prefix)-1]
			}
		}

		if n == nil || (n.Left() == nil && n.Right() == nil) {
			stack = stack[:len(stack)-1]
			if len(prefix) != 0 {
				prefix = prefix[:len(prefix)-1]
			}
			continue
		}

		if cobj.cnt == 0 {
			l := n.Left()
			var toprint string
			if l == nil {
				toprint = "*"
			} else {
				toprint = fmt.Sprint(l.Value())
			}
			out += fmt.Sprintf("%s%s %v\n", strings.Join(prefix, ""), PREFIX_LEFT, toprint)
			prefix = append(prefix, EXTRA_LEFT)
			stack = append(stack, &stkobj{
				n:   l,
				cnt: 0,
			})
			cobj.cnt += 1
			continue
		} else if cobj.cnt == 1 {
			r := n.Right()
			var toprint string
			if r == nil {
				toprint = "*"
			} else {
				toprint = fmt.Sprint(r.Value())
			}
			out += fmt.Sprintf("%s%s %v\n", strings.Join(prefix, ""), PREFIX_RIGHT, toprint)
			prefix = append(prefix, EXTRA_RIGHT)
			stack = append(stack, &stkobj{
				n:   n.Right(),
				cnt: 0,
			})
			cobj.cnt += 1
			continue
		}
	}

	return strings.TrimSuffix(out, "\n")
}
