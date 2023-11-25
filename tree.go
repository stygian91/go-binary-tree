package gobinarytree

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type ordered = constraints.Ordered

type Tree[A ordered] struct {
	Root Node[A]
}

type Node[A ordered] struct {
	Parent *Node[A]
	Left   *Node[A]
	Right  *Node[A]
	Value  A
}

func New[A ordered](value A) Tree[A] {
	return Tree[A]{
		Root: Node[A]{Value: value},
	}
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

func (this *Tree[A]) Add(value A) bool {
	current := &this.Root

	for {
		if value < current.Value {
			if current.Left != nil {
				current = current.Left
				continue
			}

			current.Left = &Node[A]{
				Parent: current,
				Value:  value,
			}
			break
		}

		if value > current.Value {
			if current.Right != nil {
				current = current.Right
				continue
			}

			current.Right = &Node[A]{
				Parent: current,
				Value:  value,
			}
			break
		}

		return true
	}

	return false
}

func (this *Tree[A]) Remove(value A) bool {
	toDelete := this.Search(value)
	if toDelete == nil {
		return false
	}

	parent := toDelete.Parent

	if toDelete.Left == nil && toDelete.Right == nil {
		if parent == nil {
			this.Root = Node[A]{}
			return true
		}

		if toDelete.Value < parent.Value {
			parent.Left = nil
		} else {
			parent.Right = nil
		}

		return true
	}

	if toDelete.Left == nil {
		if parent == nil {
			this.Root = *toDelete.Right
			toDelete.Right.Parent = nil
			return true
		}

		if toDelete.Value < parent.Value {
			parent.Left = toDelete.Right
		} else {
			parent.Right = toDelete.Right
		}

		toDelete.Right.Parent = parent
		return true
	}

	if toDelete.Right == nil {
		if parent == nil {
			this.Root = *toDelete.Left
			toDelete.Left.Parent = nil
			return true
		}

		if toDelete.Value < parent.Value {
			parent.Left = toDelete.Left
		} else {
			parent.Right = toDelete.Left
		}

		toDelete.Left.Parent = parent
		return true
	}

	rightMin := toDelete.Right.Min()
	rightMin.Left = toDelete.Left
	toDelete.Left.Parent = rightMin

	if parent == nil {
		this.Root = *toDelete.Right
		toDelete.Right.Parent = nil
		return true
	}

	if toDelete.Value < parent.Value {
		parent.Left = toDelete.Right
	} else {
		parent.Right = toDelete.Right
	}

	toDelete.Right.Parent = parent
	return true
}

func (this Tree[A]) Search(needle A) *Node[A] {
	current := &this.Root

	for {
		if needle < current.Value {
			if current.Left == nil {
				return nil
			}

			current = current.Left
			continue
		}

		if needle > current.Value {
			if current.Right == nil {
				return nil
			}

			current = current.Right
			continue
		}

		return current
	}
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

func (this *Tree[A]) BreadthFirstVisit(cb func(*Node[A])) {
	visited := map[*Node[A]]bool{}
	queue := []*Node[A]{&this.Root}

	addToQueue := func(node *Node[A]) {
		if node == nil || visited[node] == true {
			return
		}

		queue = append(queue, node)
	}

	for {
		if len(queue) == 0 {
			return
		}

		next := queue[0]
		queue = queue[1:]

		if visited[next] == true {
			continue
		}

		visited[next] = true
		cb(next)

		addToQueue(next.Left)
		addToQueue(next.Parent)
		addToQueue(next.Right)
	}
}

func (this *Tree[A]) DepthFirstVisit(cb func(*Node[A])) {
	visited := map[*Node[A]]bool{}
	stack := []*Node[A]{&this.Root}

	addToQueue := func(node *Node[A]) {
		if node == nil || visited[node] == true {
			return
		}

		stack = append(stack, node)
	}

	for {
		if len(stack) == 0 {
			return
		}

		next := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[next] == true {
			continue
		}

		visited[next] = true
		cb(next)

		addToQueue(next.Right)
		addToQueue(next.Parent)
		addToQueue(next.Left)
	}
}
