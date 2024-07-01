package stack

import "fmt"

// 定义一个通用的栈结构体
type TStack[T any] struct {
	items []T
}

func NewTStack[T any]() *TStack[T] {
	return &TStack[T]{}
}

// 入栈操作
func (s *TStack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// 出栈操作
func (s *TStack[T]) Pop() (T, error) {
	var val T
	if len(s.items) == 0 {
		return val, fmt.Errorf("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// 获取栈顶元素但不出栈
func (s *TStack[T]) Peek() (T, error) {
	var val T
	if len(s.items) == 0 {
		return val, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// 检查栈是否为空
func (s *TStack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
