package trees

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func insert(root *Node, val int) {
	if root.Val >= val {
		if root.Left == nil {
			root.Left = &Node{Val: val}
			return
		}
		insert(root.Left, val)
	} else {
		if root.Right == nil {
			root.Right = &Node{Val: val}
			return
		}
		insert(root.Right, val)
	}
}

func print(n *Node) {
	if n == nil {
		return
	}
	fmt.Println("N")
	fmt.Println(n.Val)
	fmt.Println("L")
	print(n.Left)
	fmt.Println("R")
	print(n.Right)
}

func PrintTree(t *Tree) {
	print(t.Root)
}

func (t *Tree) Insert(val int) {
	if t.Root == nil {
		t.Root = &Node{
			Val: val,
		}
		return
	}
	insert(t.Root, val)
}

func find(n *Node, val int) bool {
	if n == nil {
		return false
	}
	if n.Val == val {
		return true
	}

	if n.Val <= val {
		return find(n.Right, val)
	}
	return find(n.Left, val)
}

func (t *Tree) Find(val int) bool {
	return find(t.Root, val)
}
