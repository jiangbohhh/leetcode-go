package linkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a Double Node.
type DoubleNode struct {
	Val  int
	Prev *DoubleNode
	Next *DoubleNode
}

// ListToNode
func ListToNode(list []int) (head *ListNode) {
	if len(list) >= 1 {
		head = &ListNode{
			Val:  list[len(list)-1],
			Next: nil,
		}

	}
	for i := len(list) - 2; i >= 0; i-- {
		node := &ListNode{
			Val:  list[i],
			Next: nil,
		}
		node.Next = head
		head = node
	}
	return
}

// NodeToList
func NodeToList(head *ListNode) (list []int) {
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	return
}
