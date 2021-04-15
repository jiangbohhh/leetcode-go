package addTwoNumber

import "jiangbohhh/leetcode-go/structure/linkedlist"

// 0002. Add Two Number
// 给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0开头。
//
// 示例 1：
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
// 示例 2：
// 输入：l1 = [0], l2 = [0]
// 输出：[0]
// 示例 3：
// 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// 输出：[8,9,9,9,0,0,0,1]
//
// 提示：
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/add-two-numbers
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode = linkedlist.ListNode

// addTwoNumbers
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var long, short = l1, l2
	var len1, len2 int
	for i := l1; i != nil; i = i.Next {
		len1++
	}
	for j := l2; j != nil; j = j.Next {
		len2++
	}
	if len1 > len2 {
		long = l1
		short = l2
	} else {
		long = l2
		short = l1
	}
	carry := false
	newHead := new(ListNode)
	ret := newHead
	for long != nil {
		node := new(ListNode)
		if carry {
			node.Val++
			carry = false
		}
		node.Val += long.Val + short.Val
		if node.Val > 9 {
			carry = true
			node.Val -= 10
		}
		newHead.Next = node
		newHead = node
		if short.Next == nil {
			short.Val = 0
		} else {
			short = short.Next
		}
		long = long.Next
	}
	if carry {
		newHead.Next = &ListNode{
			Val: 1,
		}
	}
	return ret.Next
}
