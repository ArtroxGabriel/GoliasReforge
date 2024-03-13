package utils

import "time"

type Comparator[T any] func(x, y T) int

func TimeComparator(a, b time.Time) int {
	switch {
	case a.After(b):
		return 1
	case b.After(a):
		return -1
	default:
		return 0
	}
}
