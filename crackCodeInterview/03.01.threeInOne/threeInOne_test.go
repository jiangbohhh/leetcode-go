package threeInOne

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	obj := Constructor(3)
	obj.Push(0, 1)
	if v := obj.Pop(0); v != 1 {
		t.Errorf("Pop() = %v want =%v", v, 1)
	}
	if v := obj.Peek(0); v != -1 {
		t.Errorf("Peek() = %v want =%v", v, -1)
	}
	if v := obj.IsEmpty(0); v != true {
		t.Errorf("IsEmpty() = %v want = %v", v, true)
	}
}
