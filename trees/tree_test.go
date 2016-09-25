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

func TestFind(t *testing.T) {
	tree := NewTree(5)
	tree.Insert(3)
	tree.Insert(8)
	tree.Insert(1)
	tree.Insert(2)
	if tree.Find(4) {
		t.Error("found 4 but it is not there")
	}
	if !tree.Find(8) {
		t.Error("should have found 8")
	}
}
