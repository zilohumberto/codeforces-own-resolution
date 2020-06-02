package stack

import (
	"testing"
)

var s ItemStack

func initStack() *ItemStack {
	if s.items == nil {
		s = ItemStack{}
		s.New()
	}
	return &s
}

func TestEmpty(t *testing.T) {
	s := initStack()
	if !s.isEmpty() {
		t.Error("IsEmpty should be true")
	}
	s.Push(3)
	if s.isEmpty() {
		t.Error("IsEmpty should be false")
	}
	s.Pop()
	if !s.isEmpty() {
		t.Error("IsEmpty should be true")
	}
}

func TestPush(t *testing.T) {
	s := initStack()
	s.Push(5)
	item := s.Pop()
	if item != 5 {
		t.Error("Push error. excepted 5. got: ", item)
	}
	s.Push(0)
	item = s.Pop()
	if item != 0 {
		t.Error("Push error. excepted 0. got: ", item)
	}
	s.Push(100)
	item = s.Pop()
	if item != 100 {
		t.Error("Push error. excepted 100. got: ", item)
	}
}

func TestPop(t *testing.T) {
	s := initStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	item := s.Pop()
	if item != 3 {
		t.Error("Pop error. Excepted 3. got: ", item)
	}
	item = s.Pop()
	if item != 2 {
		t.Error("Pop error. Excepted 2. got: ", item)
	}
	s.Push(5)
	item = s.Pop()
	if item != 5 {
		t.Error("Pop error. Excepted 5. got: ", item)
	}
	item = s.Pop()
	if item != 1 {
		t.Error("Pop error. Excepted 1. got: ", item)
	}
}
