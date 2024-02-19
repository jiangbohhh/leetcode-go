package getMin

import "jiangbohhh/leetcode-go/structure/stack"

// 设计一个有getMin功能的栈
// 实现一个特殊的栈，在实现栈的基础功能上，再实现返回栈中最小元素的操作
// 要求：
// 1.pop、push、getMin的操作的时间复杂度都是O(1)
// 2.设计栈可以使用现成栈结构

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
