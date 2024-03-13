package tree

import "TrabalhosEDAvancada/src/containers"

type Tree[V any] interface {
	containers.Container[V]
}
