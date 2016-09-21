package main

import (
	"bytes"
	"fmt"
)

type Node struct {
	Next *Node
	String string
}

func NewNode(string string) *Node {
	node := &Node{}
	node.String = string
	return node
}

func Equals(node1, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return node1.String == node2.String
}

func HasNode(head, node *Node) bool {
	if Equals(head, node) {
		return true
	}
	if head.Next == nil {
		return false
	}
	return HasNode(head.Next, node)
}

func Length(head *Node) int {
	if head == nil {
		return 0
	} else {
		return Length(head.Next) + 1
	}
}

// assume first is the first node in a doubly-linked list with no loops
func RemoveDuplicates(head *Node) *Node {
	if head.Next == nil {
		return head
	}
	node := head.Next
	head.Next = nil
	tail := head
	for node != nil {
		if !HasNode(head, node) {
			tail.Next = node
			tail = node
			node = node.Next
			tail.Next = nil
		} else {
			node = node.Next
		}
	}
	return head
}

func (head *Node) GoString() string {
	var buffer bytes.Buffer
	node := head
	for node != nil {
		buffer.WriteString(node.String)
		buffer.WriteString(" ")
		node = node.Next
	}
	return buffer.String()
}

func main() {
	strings := []string{"apple", "banana", "cat", "dog", "fish", "fish", "fish"}
	head := NewNode(strings[0])
	node := head
	for _, str := range strings[1:] {
		node.Next = NewNode(str)
		node = node.Next
	}
	fmt.Printf("%#v\n", head)
	RemoveDuplicates(head)
	fmt.Printf("%#v\n", head)
}
