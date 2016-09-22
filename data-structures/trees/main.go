package main

import "errors"

type Node struct {
	Previous *Node
	Data interface{}
}

func NewNode(data interface{}) *Node {
	node := Node{}
	node.Data = data
	return &node
}

type Stack struct {
	top *Node
}

func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(data interface{}) {
	if stack.top == nil {
		stack.top = NewNode(data)
	} else {
		previous := stack.top
		stack.top = NewNode(data)
		stack.top.Previous = previous
	}
}

func (stack *Stack) Pop() (interface{}, error) {
	popped := stack.top
	if popped == nil {
		return nil, errors.New("empty stack")
	}
	stack.top = stack.top.Previous
	return popped.Data, nil
}

func main() {

}
