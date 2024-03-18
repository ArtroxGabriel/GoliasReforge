package main

import (
	"fmt"
	"github.com/ArtroxGabriel/GoliasReforge/cmd/tree/redBlackTree"
	"github.com/ArtroxGabriel/GoliasReforge/cmd/tree/redBlackTreeGab"
)

func main() {
	tree := redBlackTreeGab.New()
	rbt := redBlackTree.New[int, int]()

	baseInput := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for _, input := range baseInput {
		redBlackTreeGab.Insert(tree, input)
		rbt.Insert(input, input)

	}

	redBlackTreeGab.InOrder(tree, tree.Root)

	redBlackTreeGab.Remove(tree, 7)
	aqui := redBlackTreeGab.Search(tree, 6)

	fmt.Println()
	fmt.Printf("%d%v\n", aqui.Key, aqui.Color)
	fmt.Printf("%d%v", aqui.Right.Key, aqui.Right.Color)
	println()
	redBlackTreeGab.InOrder(tree, tree.Root)

}
