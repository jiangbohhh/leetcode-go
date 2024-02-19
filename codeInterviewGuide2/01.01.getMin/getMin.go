package getMin

import "jiangbohhh/leetcode-go/structure/stack"

type MinStack struct {
	dataStack *stack.Stack
	minStack  *stack.Stack
}

func NewMinStack() *MinStack {
	dataStack := stack.NewStack()
	minStack := stack.NewStack()
	return &MinStack{
		dataStack: dataStack,
		minStack:  minStack,
	}
}

func (m *MinStack) Push(n int) {
	if m.minStack.IsEmpty() {
		m.minStack.Push(n)
	} else if m.GetMin() > n {
		m.minStack.Push(n)
	} else {
		m.minStack.Push(m.GetMin())
	}
	m.dataStack.Push(n)
	return
}

func (m *MinStack) Pop() int {
	value := m.dataStack.Pop()
	m.minStack.Pop()
	return value
}

func (m *MinStack) GetMin() int {
	if m.minStack.IsEmpty() {
		return 0
	}
	return m.minStack.Peek()
}
