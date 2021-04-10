package queueWithTwoStacks

import (
	"testing"
)

func TestQueuesWithTwoStack(t *testing.T) {
	t.Run("['CQueue','appendTail','deleteHead','deleteHead']", func(t *testing.T) {
		cQueue := Constructor()
		cQueue.AppendTail(3)
		deleteHead := cQueue.DeleteHead()
		t.Logf("deleteHead = %v", deleteHead)
		deleteHead = cQueue.DeleteHead()
		t.Logf("deleteHead = %v", deleteHead)
	})
}
