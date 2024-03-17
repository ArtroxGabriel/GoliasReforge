package binaryTree

import (
	"cmp"
	"fmt"
	"github.com/ArtroxGabriel/GoliasReforge/src/utils"
)

type Tree[K comparable, V any] struct {
	Root       *Node[K, V]
	size       int
	Comparator utils.Comparator[K]
}

type Node[K comparable, V any] struct {
	Key         K
	Value       V
	Left, Right *Node[K, V]
	Parent      *Node[K, V]
}

func New[K cmp.Ordered, V any]() *Tree[K, V] {
	return &Tree[K, V]{Comparator: cmp.Compare[K]}
}

func (tree *Tree[K, V]) Insert(key K, value V) {
	var insertedNode *Node[K, V]
	if tree.Root == nil {
		tree.Root = &Node[K, V]{Key: key, Value: value}
		insertedNode = tree.Root
	} else {
		node := tree.Root
		loop := true
		for loop {
			compare := tree.Comparator(key, node.Key)
			switch {
			case compare == 0:
				node.Key = key
				node.Value = value
				return
			case compare < 0:
				if node.Left == nil {
					node.Left = &Node[K, V]{Key: key, Value: value}
					insertedNode = node.Left
					loop = false
				} else {
					node = node.Left
				}
			case compare > 0:
				if node.Right == nil {
					node.Right = &Node[K, V]{Key: key, Value: value}
					insertedNode = node.Right
					loop = false
				} else {
					node = node.Right
				}
			}
		}
		insertedNode.Parent = node
	}
	tree.size++
}

func (tree *Tree[K, V]) Get(key K) (value V, found bool) {
	node := tree.Root
	for node != nil {
		compare := tree.Comparator(key, node.Key)
		switch {
		case compare == 0:
			return node.Value, true
		case compare > 0:
			node = node.Right
		case compare < 0:
			node = node.Left
		}
	}
	return value, false
}

func (tree *Tree[K, V]) GetNode(key K) *Node[K, V] {
	node := tree.Root
	for node != nil {
		compare := tree.Comparator(key, node.Key)
		switch {
		case compare == 0:
			return node
		case compare > 0:
			node = node.Right
		case compare < 0:
			node = node.Left
		}
	}
	return nil
}

func (tree *Tree[K, V]) Remove(key K) {
	node := tree.Root
	for node != nil {
		compare := tree.Comparator(key, node.Key)
		switch {
		case compare == 0:
			node = nil
			return
		case compare > 0:
			node = node.Right
		case compare < 0:
			node = node.Left
		}
	}
	tree.size--
	return
}

func (tree *Tree[K, V]) Empty() bool {
	return tree.size == 0
}

func (tree *Tree[K, V]) Size() int {
	return tree.size
}

func (tree *Tree[K, V]) Clear() {
	tree.Root = nil
	tree.size = 0
}

func (tree *Tree[K, V]) Values() []V {
	values := make([]V, 0, tree.size)
	it := tree.Iterator()
	for i := 0; it.Next(); i++ {
		values[i] = it.Value()
	}
	return values
}

func (tree *Tree[K, V]) Left() *Node[K, V] {
	var parent *Node[K, V]
	current := tree.Root
	for current != nil {
		parent = current
		current = current.Left
	}
	return parent
}

func (tree *Tree[K, V]) Right() *Node[K, V] {
	var parent *Node[K, V]
	current := tree.Root
	for current != nil {
		parent = current
		current = current.Right
	}
	return parent
}

func (tree *Tree[K, V]) String() string {
	str := "Binary Tree\n"
	if !tree.Empty() {
		output(tree.Root, "", true, &str)
	}
	return str
}

func (node *Node[K, V]) String() string {
	return fmt.Sprintf("%v", node.Key)
}

func output[K comparable, V any](node *Node[K, V], prefix string, isTail bool, str *string) {
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.Right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.Left, newPrefix, true, str)
	}
}
