package containers

import (
	"TrabalhosEDAvancada/src/utils"
	"cmp"
	"slices"
)

type Container[T any] interface {
	Empty() bool
	Size() int
	Clear()
	Values() []T
}

func GetSortedValues[T cmp.Ordered](container Container[T]) []T {
	values := container.Values()
	if len(values) < 2 {
		return values
	}
	slices.Sort(values)
	return values
}

func GetSortedValuesFunc[T cmp.Ordered](container Container[T], comparator utils.Comparator[T]) []T {
	values := container.Values()
	if len(values) < 2 {
		return values
	}
	slices.SortFunc(values, comparator)
	return values
}
