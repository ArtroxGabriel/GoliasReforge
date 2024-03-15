package main

import (
	"fmt"
	"github.com/ArtroxGabriel/GoliasReforge/src/queue"
)

func main() {
	fila := queue.New[int]()

	fmt.Println(fila.IsEmpty())

	fila.Enqueue(1)
	fila.Enqueue(2)
	fila.Enqueue(3)
	fila.Enqueue(4)

	fmt.Println(fila.Size())

	fmt.Println(fila.Dequeue())
	fmt.Println(fila.Dequeue())
	fmt.Println(fila.Dequeue())
	fmt.Println(fila.Dequeue())

}
