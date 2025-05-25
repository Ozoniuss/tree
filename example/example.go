package main

import (
	"fmt"

	"github.com/Ozoniuss/tree"
)

func main() {

	t1 := tree.NewBST[int]()
	t1.Insert(4)

	t1.Insert(1)
	t1.Insert(8)

	t1.Insert(8)
	t1.Insert(6)
	t1.Insert(9)
	t1.Insert(5)

	t1.Insert(11)
	t1.Insert(13)

	fmt.Println(tree.FormatTree(t1, tree.FormatLinuxTree))
	fmt.Println(tree.FormatTree(t1, tree.FormatHorizontal))
	fmt.Println(tree.FormatTree(t1, tree.FormatHorizontalSquared))

	t2 := tree.NewBST[string]()
	t2.Insert("5sfhskfuceskjvsdnkvjkdsn")
	t2.Insert("1dbfalkfbdslkjfbadslkfbl")
	t2.Insert("3dbfalkfbdslkjfbadslkfbl")
	t2.Insert("8dsbflkjsdbfjzhklsdbfljkds")
	t2.Insert("7dsbflkjsdbfjzhklsdbfljkds")
	t2.Insert("9dsbflkjsdbfjzhklsdbfljkds")

	fmt.Println(tree.FormatTree(t2, tree.FormatLinuxTree))
	fmt.Println(tree.FormatTree(t2, tree.FormatHorizontal))
	fmt.Println(tree.FormatTree(t2, tree.FormatHorizontalSquared))
}
