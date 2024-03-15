package queue

type Queue[K any] struct {
	First *Node[K]
}

type Node[K any] struct {
	Value K
	Next  *Node[K]
}

func New[K any]() *Queue[K] {
	return &Queue[K]{}
}

func (q *Queue[K]) IsEmpty() bool {
	return q.First == nil
}

func (q *Queue[K]) Enqueue(value K) {
	if q.First == nil {
		q.First = &Node[K]{Value: value}
		return
	}
	currentNode := q.First
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = &Node[K]{Value: value}
}

func (q *Queue[K]) Dequeue() *Node[K] {
	firstNode := q.First
	q.First = firstNode.Next
	return firstNode
}

func (q *Queue[K]) Size() int {
	size := 0
	for current := q.First; current != nil; current = current.Next {
		size++
	}
	return size
}
