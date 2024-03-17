package redBlackTreeGab

import "fmt"

type color bool

const (
	RED, BLACK color = false, true
)

type Node struct {
	Key                 int
	Color               color
	Left, Right, Parent *Node
}

type Rbt struct {
	Root *Node
	Null *Node
}

func New() *Rbt {
	nulo := &Node{Color: BLACK}
	return &Rbt{Root: nulo, Null: nulo}
}

// Insercao na Red Black Tree
// INICIO Insert
func rotateLeft(rbt *Rbt, x *Node) {
	y := x.Right
	x.Right = y.Left

	if y.Left != rbt.Null {
		y.Left.Parent = x
	}

	y.Parent = x.Parent

	if x.Parent == rbt.Null {
		rbt.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

func rotateRight(rbt *Rbt, y *Node) {
	x := y.Left
	y.Left = x.Right

	if x.Right != rbt.Null {
		x.Right.Parent = y
	}

	x.Parent = y.Parent

	if y.Parent == rbt.Null {
		rbt.Root = x
	} else if y == y.Parent.Left {
		y.Parent.Left = x
	} else {
		y.Parent.Right = x
	}

	x.Right = y
	y.Parent = x
}

func fixInsert(rbt *Rbt, z *Node) {
	var uncle *Node

	for rbt.Root != z && z.Color == RED && z.Parent.Color == RED {
		if z.Parent == z.Parent.Parent.Left {
			uncle = z.Parent.Parent.Right

			//case 1
			if uncle.Color == RED {
				z.Parent.Color = BLACK
				uncle.Color = BLACK
				z = z.Parent.Parent
				z.Color = RED
			} else {
				if z == z.Parent.Right {
					z = z.Parent
					rotateLeft(rbt, z)
				}
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				rotateRight(rbt, z.Parent.Parent)
			}
		} else {
			uncle = z.Parent.Parent.Left

			//case 1
			if uncle.Color == RED {
				z.Parent.Color = BLACK
				uncle.Color = BLACK
				z = z.Parent.Parent
				z.Color = RED
			} else {
				if z == z.Parent.Left {
					z = z.Parent
					rotateRight(rbt, z)
				}
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				rotateLeft(rbt, z.Parent.Parent)
			}
		}
	}
	rbt.Root.Color = BLACK
}

func Insert(T *Rbt, key int) {
	insertedNode, raiz, pai := &Node{Key: key, Color: RED}, T.Root, T.Null

	for raiz != T.Null {
		pai = raiz
		if key <= raiz.Key {
			raiz = raiz.Left
		} else {
			raiz = raiz.Right
		}
	}

	insertedNode.Parent = pai

	if pai == T.Null {
		T.Root = insertedNode
	} else if insertedNode.Key <= pai.Key {
		pai.Left = insertedNode
	} else {
		pai.Right = insertedNode
	}

	insertedNode.Left = T.Null
	insertedNode.Right = T.Null
	fixInsert(T, insertedNode)
}

// FIM Insert

// Remoção na Red Black Tree
// INICIO Remove
func transplat(T *Rbt, u *Node, v *Node) {
	if u.Parent == T.Null {
		T.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	v.Parent = u.Parent
}

func fixRemove(T *Rbt, z *Node) {
	var w *Node

	for z != T.Root && z.Color == BLACK {
		if z == z.Parent.Left {
			w = z.Parent.Right
			if w.Color == RED {
				w.Color = BLACK
				z.Parent.Color = RED
				rotateLeft(T, z.Parent)
				w = z.Parent.Right
			}
			if w.Left.Color == BLACK && w.Right.Color == BLACK {
				w.Color = RED
				z = z.Parent
			} else if w.Right.Color == BLACK {
				w.Left.Color = BLACK
				w.Color = RED
				rotateRight(T, z.Parent)
				w = z.Parent.Right
			}
			w.Color = z.Parent.Color
			z.Parent.Color = BLACK
			w.Right.Color = BLACK
			rotateLeft(T, z.Parent)
			z = T.Root
		} else {
			w = z.Parent.Left
			if w.Color == RED {
				w.Color = BLACK
				z.Parent.Color = RED
				rotateRight(T, z.Parent)
				w = z.Parent.Left
			}
			if w.Right.Color == BLACK && w.Left.Color == BLACK {
				w.Color = RED
				z = z.Parent
			} else if w.Left.Color == BLACK {
				w.Right.Color = BLACK
				w.Color = RED
				rotateLeft(T, z.Parent)
				w = z.Parent.Left
			}
			w.Color = z.Parent.Color
			z.Parent.Color = BLACK
			w.Left.Color = BLACK
			rotateRight(T, z.Parent)
			z = T.Root
		}
		z.Color = BLACK
	}

}

func sucessor(T *Rbt, node *Node) *Node {
	caba := node.Right
	salva := caba
	for caba.Left != T.Null {
		salva = caba.Left
		caba = caba.Left
	}
	return salva
}

func Remove(T *Rbt, z *Node) {
	var x *Node
	y := z
	originalColor := y.Color

	if z.Left == T.Null {
		x = z.Right
		transplat(T, z, x)
	} else if z.Right == T.Null {
		x = z.Left
		transplat(T, z, x)
	} else {
		y = sucessor(T, z)
		x = y.Right
		originalColor = y.Color
		transplat(T, y, x)
		y.Left = z.Left
		z.Left.Parent = y
		y.Right = z.Right
		z.Right.Parent = y
		transplat(T, z, y)
		y.Color = z.Color
	}
	if originalColor == BLACK {
		fixRemove(T, x)
	}
}

// FIM Remove

func InOrder(T *Rbt, node *Node) {
	if node != T.Null {
		InOrder(T, node.Left)
		fmt.Println(node.Key)
		InOrder(T, node.Right)
	}
}

func Search(T *Rbt, key int) *Node {
	if T.Root == T.Null {
		return nil
	}
	encontrei := T.Root
	for encontrei != T.Null && encontrei.Key != key {
		if encontrei.Key <= key {
			encontrei = encontrei.Right
		} else {
			encontrei = encontrei.Left
		}
	}
	return encontrei
}
