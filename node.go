package gobinarytree

import (
	"fmt"
	"strings"
)

type Node[A ordered] struct {
	Parent *Node[A]
	Left   *Node[A]
	Right  *Node[A]
	Value  A
}

func Print[A ordered](node *Node[A], space int) {
	_print(node, space)
	fmt.Println()
}

func _print[A ordered](node *Node[A], space int) {
	if node == nil {
		return
	}

	space += 4

	_print(node.Right, space)
	fmt.Println()
	fmt.Printf("%s%+v", strings.Repeat(" ", space), node.Value)

	_print(node.Left, space)
}

func (this *Node[A]) Min() *Node[A] {
	current := this

	for {
		if current.Left == nil {
			return current
		}

		current = current.Left
	}
}

func (this *Node[A]) Max() *Node[A] {
	current := this

	for {
		if current.Right == nil {
			return current
		}

		current = current.Right
	}
}
