package binary_search_tree

type Comparable interface {
	Compare(comparable Comparable) bool
	Equals(comparable Comparable) bool
}

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Data Comparable
}

func NewTreeNode(data Comparable) *TreeNode {
	return &TreeNode{nil, nil, data}
}

func (node *TreeNode) Insert(data Comparable) *TreeNode {
	if node.Data.Compare(data) {
		if node.Left != nil {
			node.Left.Insert(data)
		}
		node.Left = NewTreeNode(data)
	} else {
		if node.Right != nil {
			node.Right.Insert(data)
		}
		node.Right = NewTreeNode(data)
	}
	return node
}

func (node *TreeNode) Search(data Comparable) bool {
	if !node.Data.Equals(data) && node.Left == nil && node.Right == nil {
		return false
	}
	if node.Data.Equals(data) {
		return true
	}
	if node.Data.Compare(data) {
		if node.Left == nil {
			return false
		}
		return node.Left.Search(data)
	} else {
		if node.Right == nil {
			return false
		}
		return node.Right.Search(data)
	}
}

func (node *TreeNode) Length() int {
	return 0
}
