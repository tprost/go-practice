package main

import "testing"


func TestPop(t *testing.T) {
	stack := NewStack()
	stack.Push("hello")
	stack.Push("test")
	popped, _ := stack.Pop()
	test := popped.(string)
	if test != "test" {
		t.Error("did not pop test string")
	}
}

func TestPush(t *testing.T) {
	stack := NewStack()
	stack.Push("hello")
	stack.Push("test")
	popped, _ := stack.Pop()
	test := popped.(string)
	if test != "test" {
		t.Error("did not pop test string")
	}
}

func TestPoppedTooManyTimes(t *testing.T) {
	stack := NewStack()
	stack.Push("hello")
	stack.Push("test")
	stack.Pop()
	stack.Pop()
	_, error := stack.Pop()
	if error == nil {
		t.Error("should have thrown an error")
	}
}
