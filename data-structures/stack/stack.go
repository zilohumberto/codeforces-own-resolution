package stack

// Item type of the Stack
type Item interface {
}

// ItemStack the stack of items
type ItemStack struct {
	items []Item
}

// New instance a new stack
func (s *ItemStack) New() *ItemStack {
	s.items = []Item{}
	return s
}

// isEmpty if stack is empty or not
func (s *ItemStack) isEmpty() bool {
	return len(s.items) == 0
}

// Push add a element at end of the stack
func (s *ItemStack) Push(v Item) {
	s.items = append(s.items, v)
}

// Pop delete a element of the front
func (s *ItemStack) Pop() Item {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}
