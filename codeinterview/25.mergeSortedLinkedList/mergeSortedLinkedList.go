package mergeSortedLinkedList

// 剑指 Offer 25. merge Sorted LinkedList
// 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
//
// 示例1：
//
// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4
// 限制：
//
// 0 <= 链表长度 <= 1000
// 注意：本题与主站 21 题相同：https://leetcode-cn.com/problems/merge-two-sorted-lists/
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var newHead = new(ListNode)
	if l1.Val < l2.Val {
		newHead = l1
		newHead.Next = mergeTwoLists(l1.Next, l2)
	} else {
		newHead = l2
		newHead.Next = mergeTwoLists(l1, l2.Next)
	}
	return newHead
}
