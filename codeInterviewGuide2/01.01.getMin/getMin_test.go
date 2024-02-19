package getMin

import (
	"testing"
)

func TestNewMinStack(t *testing.T) {
	stack := NewMinStack()
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(1)
	stack.Push(2)
	stack.Push(1)
	for {
		if stack.dataStack.IsEmpty() {
			break
		}
		minVal := stack.GetMin()
		t.Logf("stack min = %d", minVal)
		value := stack.Pop()
		t.Logf("stack data value=%d", value)
	}
}
