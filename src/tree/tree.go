package tree

import "github.com/ArtroxGabriel/GoliasReforge/src/containers"

type Tree[V any] interface {
	containers.Container[V]
}
