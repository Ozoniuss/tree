package main

import (
	"fmt"

	"github.com/Ozoniuss/tree"
)

func debug() {
	t := tree.Instantiate1()
	fmt.Println(t.Key, t.Left.Key, t.Right.Key, t.Right.Left.Key, t.Right.Right.Key, t.Right.Left.Left.Key, t.Right.Left.Right.Key)
	fmt.Println(t.Key, t.Left.P, t.Right.P, t.Right.Left.P, t.Right.Right.P, t.Right.Left.Left.P, t.Right.Left.Right.P)
	fmt.Println(t.ToStringPreorder())
	fmt.Println(t.Left.ToStringPreorder())
	fmt.Println(t.Right.ToStringPreorder())

	t = tree.Instantiate2()
	fmt.Println(t.Key, t.Left.Key, t.Right.Key, t.Left.Left.Key, t.Right.Left.Key, t.Right.Right.Key, t.Right.Left.Left.Key, t.Right.Left.Right.Key, t.Right.Left.Left.Left.Key, t.Right.Left.Right.Right.Key)

	fmt.Println(t.ToStringPreorder())

}

func main() {
	debug()
}
