package palindromeLinkedList

import (
	"container/list"
	"jiangbohhh/leetcode-go/structure/linkedlist"
)

// 0234. palindrome Linked List
// 请判断一个链表是否为回文链表。
//
// 示例 1:
// 输入: 1->2
// 输出: false
// 示例 2:
// 输入: 1->2->2->1
// 输出: true
// 进阶：
// 你能否用O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/palindrome-linked-list
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
type ListNode = linkedlist.ListNode

// isPalindrome
func isPalindrome(head *ListNode) bool {
	stack := list.New()
	for node := head; node != nil; node = node.Next {
		stack.PushBack(node)
	}
	for node := head; node != nil; node = node.Next {
		val := stack.Back()
		if stack.Remove(val).(*ListNode).Val != node.Val {
			return false
		}
	}
	return true
}

// isPalindrome2
func isPalindrome2(head *ListNode) bool {
	var (
		mid        = new(ListNode)
		slow, fast = head, head
	)
	for slow != nil && fast != nil && fast.Next != nil {
		tmp := slow.Next
		fast = fast.Next.Next
		if slow == head {
			mid = slow
			mid.Next = nil
		} else {
			slow.Next = mid
			mid = slow
		}
		slow = tmp

	}
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		if slow.Val != mid.Val {
			return false
		}
		slow = slow.Next
		mid = mid.Next
	}
	return true
}
