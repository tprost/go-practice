package stack

type Stack struct {
	Elements []interface{}
}

func NewStack() *Stack {
	elements := []interface{}{}
	return &Stack{elements}
}

func (stack *Stack) Push(element interface{}) {
	stack.Elements = append(stack.Elements, element)
}

func (stack *Stack) Pop() interface{} {
	element, remaining := stack.Elements[stack.Size()-1], stack.Elements[:stack.Size()-1]
	stack.Elements = remaining
	return element
}

func (stack *Stack) Size() int {
	return len(stack.Elements)
}
