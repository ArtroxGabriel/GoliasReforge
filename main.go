package main

import (
	"fmt"
	bt "github.com/ArtroxGabriel/GoliasReforge/src/tree/binaryTree"
	rbt "github.com/ArtroxGabriel/GoliasReforge/src/tree/redBlackTree"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Red-Black Tree")

	rbtTree := rbt.New[int, string]()
	binaryTree := bt.New[int, string]()

	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	for i := 1; i <= 20; i++ {
		randomInt := rand.Intn(100) + 1 // Generate a random integer between 1 and 100
		rbtTree.Put(randomInt, fmt.Sprintf("value%d", randomInt))
		binaryTree.Insert(randomInt, fmt.Sprintf("value%d", randomInt))
	}

	fmt.Println(rbtTree)
	fmt.Println(binaryTree)
}
