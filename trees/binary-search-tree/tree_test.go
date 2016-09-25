package binary_search_tree

import "testing"

type Integer struct {
	Integer int
}

func (integer *Integer) Equals(otherComparable Comparable) bool {
	otherInteger := otherComparable.(*Integer)
	return integer.Integer == otherInteger.Integer
}

func (integer *Integer) Compare(otherComparable Comparable) bool {
	otherInteger := otherComparable.(*Integer)
	return integer.Integer > otherInteger.Integer
}

func NewInteger(integer int) *Integer {
	return &Integer{integer}
}

func TestInsert(t *testing.T) {
	tree := NewTreeNode(NewInteger(6))
	tree.Insert(NewInteger(1))
	tree.Insert(NewInteger(2))
	tree.Insert(NewInteger(3))
}

func TestSearch(t *testing.T) {
	tree := NewTreeNode(NewInteger(6))
	tree.Insert(NewInteger(1))
	tree.Insert(NewInteger(2))
	tree.Insert(NewInteger(3))
	if !tree.Search(NewInteger(3)) {
		t.Error("should have found 3 but didn't")
	}
}
