package main

import (
	"fmt"
	"github.com/ArtroxGabriel/GoliasReforge/src/tree/redBlackTree"
	"github.com/ArtroxGabriel/GoliasReforge/src/tree/redBlackTreeGab"
)

func main() {
	tree := redBlackTreeGab.New()
	rbt := redBlackTree.New[int, int]()

	redBlackTreeGab.Insert(tree, 1)
	rbt.Insert(1, 1)
	redBlackTreeGab.Insert(tree, 2)
	rbt.Insert(2, 3)
	redBlackTreeGab.Insert(tree, 3)
	rbt.Insert(3, 3)
	redBlackTreeGab.Insert(tree, 4)
	rbt.Insert(4, 4)
	redBlackTreeGab.Insert(tree, 5)
	rbt.Insert(5, 5)

	fmt.Println(rbt)
	fmt.Println("\n", tree.Root.Right.Right.Key)
}
