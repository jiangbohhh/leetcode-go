package sortStack

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Pop()
	if v := obj.Peek(); v != 2 {
		t.Errorf("Peek() = %v want = %v", v, 2)
	}
	if v := obj.IsEmpty(); v == true {
		t.Errorf("IsEmpty() = %v want = %v", v, true)
	}
}
