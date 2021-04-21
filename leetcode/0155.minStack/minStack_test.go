package minStack

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-1)
	minStack.Push(0)
	minStack.Push(-2)
	if val := minStack.GetMin(); val != -2 {
		t.Errorf("GetMin() = %v want =%v", val, -2)
	}
	minStack.Pop()
	if val := minStack.Top(); val != 0 {
		t.Errorf("Top() = %v want =%v", val, 0)
	}
	if val := minStack.GetMin(); val != -1 {
		t.Errorf("GetMin() = %v want =%v", val, -1)
	}
}
