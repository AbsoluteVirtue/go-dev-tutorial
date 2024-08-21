// Equivalent Binary Trees, https://go.dev/tour/concurrency/7

// A function to check whether two binary trees store the same sequence.
// This example uses the tree package, which defines the type:
// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }
package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for i := 0; i < 10; i++ {
		v1, e1 := <- c1
		v2, e2 := <- c2
		if e1 == false || e2 == false {
			break // bug if channels have diff len
		}
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))

	// t := tree.New(1)
	// ch := make(chan int, 10)
	// go Walk(t, ch)

	// for i := 0; i < 10; i++ {
	// 	v, ok := <- ch
	// 	if ok == false {
	// 		break
	// 	}
	// 	fmt.Println(v)
	// }
}
