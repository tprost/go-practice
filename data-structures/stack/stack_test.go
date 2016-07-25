package stack

import (
	"testing"
	"fmt"
)

type Thing struct {
	String string
}

func TestPop(t *testing.T) {
	stack := NewStack()
	stack.Push(&Thing{"hello"})
	stack.Push(&Thing{"hello again"})
	if stack.Size() != 2 {
		t.Error("size was " + fmt.Sprintf("%v", stack.Size()))
	}
	stack.Pop()
	if stack.Size() != 1 {
		t.Error("size was " + fmt.Sprintf("%v", stack.Size()))
	}
}
