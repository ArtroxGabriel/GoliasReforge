package redBlackTree

type Iterator[K comparable, V any] struct {
	tree     *Tree[K, V]
	node     *Node[K, V]
	position position
}

type position byte

const (
	begin, between, end position = 0, 1, 2
)

func (tree *Tree[K, V]) Iterator() *Iterator[K, V] {
	return &Iterator[K, V]{tree: tree, node: nil, position: begin}
}

func (tree *Tree[K, V]) IteratorAt(node *Node[K, V]) *Iterator[K, V] {
	return &Iterator[K, V]{tree: tree, node: node, position: between}
}

func (iterator *Iterator[K, V]) Next() bool {
	if iterator.position == end {
		goto end
	}
	if iterator.position == begin {
		left := iterator.tree.Left()
		if left == nil {
			goto end
		}
		iterator.node = left
		goto between
	}
	if iterator.node.Right != nil {
		iterator.node = iterator.node.Right
		for iterator.node.Left != nil {
			iterator.node = iterator.node.Left
		}
		goto between
	}
	for iterator.node.Parent != nil {
		node := iterator.node
		iterator.node = iterator.node.Parent
		if node == iterator.node.Left {
			goto between
		}
	}

end:
	iterator.node = nil
	iterator.position = end
	return false

between:
	iterator.position = between
	return true
}

func (iterator *Iterator[K, V]) Prev() bool {
	if iterator.position == begin {
		goto begin
	}
	if iterator.position == end {
		right := iterator.tree.Right()
		if right == nil {
			goto begin
		}
		iterator.node = right
		goto between
	}
	if iterator.node.Left != nil {
		iterator.node = iterator.node.Left
		for iterator.node.Right != nil {
			iterator.node = iterator.node.Right
		}
		goto between
	}
	for iterator.node.Parent != nil {
		node := iterator.node
		iterator.node = iterator.node.Parent
		if node == iterator.node.Right {
			goto between
		}
	}

begin:
	iterator.node = nil
	iterator.position = begin
	return false

between:
	iterator.position = between
	return true
}

func (iterator *Iterator[K, V]) Value() V {
	return iterator.node.Value
}

func (iterator *Iterator[K, V]) Key() K {
	return iterator.node.Key
}

func (iterator *Iterator[K, V]) Node() *Node[K, V] {
	return iterator.node
}

func (iterator *Iterator[K, V]) Begin() {
	iterator.node = nil
	iterator.position = begin
}

func (iterator *Iterator[K, V]) End() {
	iterator.node = nil
	iterator.position = end
}

func (iterator *Iterator[K, V]) First() bool {
	iterator.Begin()
	return iterator.Next()
}

func (iterator *Iterator[K, V]) Last() bool {
	iterator.End()
	return iterator.Prev()
}

func (iterator *Iterator[K, V]) NextTo(f func(key K, value V) bool) bool {
	for iterator.Next() {
		key, value := iterator.Key(), iterator.Value()
		if f(key, value) {
			return true
		}
	}
	return false
}

func (iterator *Iterator[K, V]) PrevTo(f func(key K, value V) bool) bool {
	for iterator.Prev() {
		key, value := iterator.Key(), iterator.Value()
		if f(key, value) {
			return true
		}
	}
	return false
}
