package containers

type IteratorWithIndex[T any] interface {
	Next() bool

	Value() T

	Index() int

	Begin()

	First() bool

	NextTo(func(index int, value T) bool) bool
}

type IteratorWithKey[K, V any] interface {
	Next() bool
	Value() V
	Key() K
	Begin()
	First() bool
	NextTo(func(key K, value V) bool) bool
}

type ReverseIteratorWithIndex[T any] interface {
	Prev() bool
	End()
	Last() bool
	PrevTo(func(index int, value T) bool) bool
	IteratorWithIndex[T]
}

type ReverseIteratorWithKey[K, V any] interface {
	Prev() bool
	End()
	Last() bool
	PrevTo(func(key K, value V) bool) bool
	IteratorWithKey[K, V]
}
