package reverseLinkedList

import "jiangbohhh/leetcode-go/structure/linkedlist"

// 0206. Reverse LinkedList
// 反转一个单链表。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
// 进阶:
// 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/reverse-linked-list
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
type ListNode = linkedlist.ListNode

// reverseList
func reverseList(head *ListNode) *ListNode {
	var newHead = head
	for next := head; next != nil; {
		tmp := next.Next
		if next == newHead {
			newHead.Next = nil
		} else {
			next.Next = newHead
		}
		newHead = next
		next = tmp
	}
	return newHead
}
