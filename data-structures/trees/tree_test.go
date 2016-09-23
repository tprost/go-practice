package main

import "testing"

func TestInOrder(t *testing.T) {
	tree := NewTree(5)
	tree.Insert(3)
	tree.Insert(8)
	tree.Insert(1)
	tree.Insert(2)
	rightOrder := []int{1, 2, 3, 5, 8}
	order := tree.InOrder()
	for index, number := range rightOrder {
		if order[index] != number {
			t.Error("order did not match", order[index], number)
		}
	}
}
