package redBlackTree

import (
	"cmp"
	"fmt"
	"github.com/ArtroxGabriel/GoliasReforge/src/utils"
)

type color bool

const (
	BLACK, RED color = true, false
)

type Tree[K comparable, V any] struct {
	Root       *Node[K, V]
	size       int
	Comparator utils.Comparator[K]
}

type Node[K comparable, V any] struct {
	Key         K
	Value       V
	Color       color
	Left, Right *Node[K, V]
	Parent      *Node[K, V]
}

func New[K cmp.Ordered, V any]() *Tree[K, V] {
	return &Tree[K, V]{Comparator: cmp.Compare[K]}
}

func (tree *Tree[K, V]) Put(key K, value V) {
	var insertedNode *Node[K, V]
	if tree.Root == nil {
		tree.Root = &Node[K, V]{Key: key, Value: value, Color: RED}
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
					node.Left = &Node[K, V]{Key: key, Value: value, Color: RED}
					insertedNode = node.Left
					loop = false
				} else {
					node = node.Left
				}
			case compare > 0:
				if node.Right == nil {
					node.Right = &Node[K, V]{Key: key, Value: value, Color: RED}
					insertedNode = node.Right
					loop = false
				} else {
					node = node.Right
				}
			}
		}
		insertedNode.Parent = node
	}
	tree.insertCase1(insertedNode)
	tree.size++
}

func (tree *Tree[K, V]) Get(Key K) (value V, found bool) {
	node := tree.lookup(Key)
	if node != nil {
		return node.Value, true
	}
	return value, false
}

func (tree *Tree[K, V]) GetNode(key K) *Node[K, V] {
	return tree.lookup(key)
}

func (tree *Tree[K, V]) Remove(key K) {
	var child *Node[K, V]
	node := tree.lookup(key)
	if node == nil {
		return
	}
	if node.Left != nil && node.Right != nil {
		pred := node.Left.maximumNode()
		node.Key = pred.Key
		node.Value = pred.Value
		node = pred
	}
	if node.Left == nil || node.Right == nil {
		if node.Right == nil {
			child = node.Left
		} else {
			child = node.Right
		}

		if node.Color == BLACK {
			node.Color = nodeColor(child)
			tree.deleteCase1(node)
		}
		tree.replaceNode(node, child)
		if node.Parent == nil && child != nil {
			child.Color = BLACK
		}
	}
	tree.size--
}

func (tree *Tree[K, V]) Empty() bool {
	return tree.size == 0
}

func (tree *Tree[K, V]) Size() int {
	return tree.size
}

func (node *Node[K, V]) Size() int {
	if node == nil {
		return 0
	}
	size := 1
	if node.Left != nil {
		size += node.Left.Size()
	}
	if node.Right != nil {
		size += node.Right.Size()
	}
	return size
}

func (tree *Tree[K, V]) Keys() []K {
	keys := make([]K, tree.size)
	it := tree.Iterator()
	for i := 0; it.Next(); i++ {
		keys[i] = it.Key()
	}
	return keys
}

// Values returns all values in-order based on the key.
func (tree *Tree[K, V]) Values() []V {
	values := make([]V, tree.size)
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

func (tree *Tree[K, V]) Clear() {
	tree.Root = nil
	tree.size = 0
}

func (tree *Tree[K, V]) lookup(key K) *Node[K, V] {
	node := tree.Root
	for node != nil {
		compare := tree.Comparator(key, node.Key)
		switch {
		case compare == 0:
			return node
		case compare < 0:
			node = node.Left
		case compare > 0:
			node = node.Right
		}
	}
	return nil
}

func (node *Node[K, V]) grandparent() *Node[K, V] {
	if node != nil && node.Parent != nil {
		return node.Parent.Parent
	}
	return nil
}

func (node *Node[K, V]) uncle() *Node[K, V] {
	if node == nil || node.Parent == nil || node.Parent.Parent == nil {
		return nil
	}
	return node.Parent.sibling()
}

func (node *Node[K, V]) sibling() *Node[K, V] {
	if node == nil || node.Parent == nil {
		return nil
	}
	if node == node.Parent.Left {
		return node.Parent.Right
	}
	return node.Parent.Left
}

func (tree *Tree[K, V]) rotateLeft(node *Node[K, V]) {
	right := node.Right
	tree.replaceNode(node, right)
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Left = node
	node.Parent = right
}

func (tree *Tree[K, V]) rotateRight(node *Node[K, V]) {
	left := node.Left
	tree.replaceNode(node, left)
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Right = node
	node.Parent = left
}

func (tree *Tree[K, V]) replaceNode(old *Node[K, V], new *Node[K, V]) {
	if old.Parent == nil {
		tree.Root = new
	} else {
		if old == old.Parent.Left {
			old.Parent.Left = new
		} else {
			old.Parent.Right = new
		}
	}
	if new != nil {
		new.Parent = old.Parent
	}
}

func (tree *Tree[K, V]) insertCase1(node *Node[K, V]) {
	if node.Parent == nil {
		node.Color = BLACK
	} else {
		tree.insertCase2(node)
	}
}
func (tree *Tree[K, V]) insertCase2(node *Node[K, V]) {
	if nodeColor(node.Parent) == BLACK {
		return
	}
	tree.insertCase3(node)
}

func (tree *Tree[K, V]) insertCase3(node *Node[K, V]) {
	uncle := node.uncle()
	if nodeColor(uncle) == RED {
		node.Parent.Color = BLACK
		uncle.Color = BLACK
		node.grandparent().Color = RED
		tree.insertCase1(node.grandparent())
	} else {
		tree.insertCase4(node)
	}
}

func (tree *Tree[K, V]) insertCase4(node *Node[K, V]) {
	grandparent := node.grandparent()
	if node == node.Parent.Right && node.Parent == grandparent.Left {
		tree.rotateLeft(node.Parent)
		node = node.Left
	} else if node == node.Parent.Left && node.Parent == grandparent.Right {
		tree.rotateRight(node.Parent)
		node = node.Right
	}
	tree.insertCase5(node)
}

func (tree *Tree[K, V]) insertCase5(node *Node[K, V]) {
	node.Parent.Color = BLACK
	grandparent := node.grandparent()
	grandparent.Color = RED
	if node == node.Parent.Left && node.Parent == grandparent.Left {
		tree.rotateRight(grandparent)
	} else if node == node.Parent.Right && node.Parent == grandparent.Right {
		tree.rotateLeft(grandparent)
	}
}

func (node *Node[K, V]) maximumNode() *Node[K, V] {
	if node == nil {
		return nil
	}
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func (tree *Tree[K, V]) deleteCase1(node *Node[K, V]) {
	if node.Parent == nil {
		return
	}
	tree.deleteCase2(node)
}

func (tree *Tree[K, V]) deleteCase2(node *Node[K, V]) {
	sibling := node.sibling()
	if nodeColor(sibling) == RED {
		node.Parent.Color = RED
		sibling.Color = BLACK
		if node == node.Parent.Left {
			tree.rotateLeft(node.Parent)
		} else {
			tree.rotateRight(node.Parent)
		}
	}
	tree.deleteCase3(node)
}

func (tree *Tree[K, V]) deleteCase3(node *Node[K, V]) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == BLACK &&
		nodeColor(sibling) == BLACK &&
		nodeColor(sibling.Left) == BLACK &&
		nodeColor(sibling.Right) == BLACK {
		sibling.Color = RED
		tree.deleteCase1(node.Parent)
	} else {
		tree.deleteCase4(node)
	}
}

func (tree *Tree[K, V]) deleteCase4(node *Node[K, V]) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == RED &&
		nodeColor(sibling) == BLACK &&
		nodeColor(sibling.Left) == BLACK &&
		nodeColor(sibling.Right) == BLACK {
		sibling.Color = RED
		node.Parent.Color = BLACK
	} else {
		tree.deleteCase5(node)
	}
}

func (tree *Tree[K, V]) deleteCase5(node *Node[K, V]) {
	sibling := node.sibling()
	if node == node.Parent.Left &&
		nodeColor(sibling) == BLACK &&
		nodeColor(sibling.Left) == RED &&
		nodeColor(sibling.Right) == BLACK {
		sibling.Color = RED
		sibling.Left.Color = BLACK
		tree.rotateRight(sibling)
	} else if node == node.Parent.Right &&
		nodeColor(sibling) == BLACK &&
		nodeColor(sibling.Right) == RED &&
		nodeColor(sibling.Left) == BLACK {
		sibling.Color = RED
		sibling.Right.Color = BLACK
		tree.rotateLeft(sibling)
	}
	tree.deleteCase6(node)
}

func (tree *Tree[K, V]) deleteCase6(node *Node[K, V]) {
	sibling := node.sibling()
	sibling.Color = nodeColor(node.Parent)
	node.Parent.Color = BLACK
	if node == node.Parent.Left && nodeColor(sibling.Right) == RED {
		sibling.Right.Color = BLACK
		tree.rotateLeft(node.Parent)
	} else if nodeColor(sibling.Left) == RED {
		sibling.Left.Color = BLACK
		tree.rotateRight(node.Parent)
	}
}

func nodeColor[K comparable, V any](node *Node[K, V]) color {
	if node == nil {
		return BLACK
	}
	return node.Color
}

func (tree *Tree[K, V]) String() string {
	str := "RedBlackTree\n"
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
