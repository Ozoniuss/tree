package main

import (
	"fmt"

	"github.com/Ozoniuss/tree"
)

func main() {
	t := tree.NewBST[int]()
	t.Insert(4)
	t.Insert(2)
	t.Insert(12)
	t.Insert(1)
	t.Insert(8)

	t.Insert(8)
	t.Insert(6)
	t.Insert(9)
	t.Insert(5)

	t.Insert(11)
	t.Insert(13)

	fmt.Println(tree.FormatTree(t))
}
