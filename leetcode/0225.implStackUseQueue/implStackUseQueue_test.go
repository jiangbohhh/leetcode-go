package implStackUseQueue

import (
	"testing"
)

func Test_StackUseQueue(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if v := stack.Top(); v != 3 {
		t.Errorf("the stack top value is error: got = %v want = %v", v, 2)
	}
	if v := stack.Pop(); v != 3 {
		t.Errorf("the stack top value is error: got = %v want = %v", v, 2)
	}
	if stack.Empty() != false {
		t.Errorf("the stack is Empty error: got = %v want = %v", stack.Empty(), false)
	}

}
