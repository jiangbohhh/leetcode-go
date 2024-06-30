package implementQueueUseStack

import "testing"

func Test_MyQueue(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	if v := queue.Pop(); v != 1 {
		t.Errorf("the queue top value is error: got = %v want = %v", v, 1)
	}
	if v := queue.Peek(); v != 2 {
		t.Errorf("the stack top value is error: got = %v want = %v", v, 2)
	}
	if queue.Empty() != false {
		t.Errorf("the stack is Empty error: got = %v want = %v", queue.Empty(), false)
	}
}
