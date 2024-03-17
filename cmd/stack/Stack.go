package stack

type Stack[K any] struct {
	top *Node[K]
}

type Node[K any] struct {
	Value K
	Next  *Node[K]
}

func New[K any]() *Stack[K] {
	return &Stack[K]{}
}

func (s *Stack[K]) Push(value K) {
	s.top = &Node[K]{Value: value, Next: s.top}
}

func (s *Stack[K]) Pop() *Node[K] {
	removedNode := s.top
	s.top = s.top.Next
	return removedNode
}

func (s *Stack[K]) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack[K]) Size() int {
	size := 0
	for node := s.top; node != nil; node = node.Next {
		size++
	}
	return size
}
