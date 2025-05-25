package main

import (
	"fmt"

	"github.com/Ozoniuss/tree"
)

func main() {
	// t := tree.NewBST[string]()
	// t.Insert("4sfhskfuceskjvsdnkvjkdsnfvnjkdfsvnjkdfsnvdksfjvsdfvdfs")
	// t.Insert("jsdbfalkfbdslkjfbadslkfbladsjkfnbadskljfnbadslkjfbnadsjkfblasd")
	// t.Insert("1dsbflkjsdbfjzhklsdbfljkdsbfjkalbfjkdasbfalsdbfkjasdbfdkkfald")

	t := tree.NewBST[int]()
	t.Insert(4)

	t.Insert(1)
	t.Insert(8)

	t.Insert(8)
	t.Insert(6)
	t.Insert(9)
	t.Insert(5)

	t.Insert(11)
	t.Insert(13)

	fmt.Println(tree.FormatTree(t, tree.FormatLinuxTree))
	fmt.Println(tree.FormatTree(t, tree.FormatHorizontal))
	fmt.Println(tree.FormatTree(t, tree.FormatHorizontalSquared))

	t2 := tree.NewBST[int]()
	t2.Insert(5)

}
