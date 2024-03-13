package main

import (
	"fmt"
	rbt "github.com/ArtroxGabriel/GoliasReforge/src/tree/redBlackTree"
)

func main() {
	fmt.Println("Red-Black Tree")

	tree := rbt.New[int, string]()

	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a")

	algo := tree.Values()
	for _, a := range algo {
		fmt.Printf("%v\t", a)
	}
}
