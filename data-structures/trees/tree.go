package main

type Tree struct {
	Integer int
	Left *Tree
	Right *Tree
}

func NewTree(integer int) *Tree {
	tree := &Tree{}
	tree.Integer = integer
	return tree
}

func (tree *Tree) Insert(integer int) *Tree {
	if integer < tree.Integer {
		if tree.Left != nil {
			tree.Left.Insert(integer)
		} else {
			tree.Left = NewTree(integer)
		}
	} else {
		if tree.Right != nil {
			tree.Right.Insert(integer)
		} else {
			tree.Right = NewTree(integer)
		}
	}
	return tree
}

func (tree *Tree) InOrder() []int {
	order := []int{}
	if tree.Left != nil {
		order = append(order, tree.Left.InOrder()...)
	}
	order = append(order, tree.Integer)
	if tree.Right != nil {
		order = append(order, tree.Right.InOrder()...)
	}
	return order
}
