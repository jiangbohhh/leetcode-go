package reverseNodesInKGroup

import (
	"jiangbohhh/leetcode-go/structure/linkedlist"
)

// 0025.Reverse Nodes In K-Group
// 给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。
// k是一个正整数，它的值小于或等于链表的长度。
// 如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。
// 进阶：
//
// 你可以设计一个只使用常数额外空间的算法来解决此问题吗？
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
// 示例 1：
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[2,1,4,3,5]
// 示例 2：
// 输入：head = [1,2,3,4,5], k = 3
// 输出：[3,2,1,4,5]
// 示例 3：
// 输入：head = [1,2,3,4,5], k = 1
// 输出：[1,2,3,4,5]
// 示例 4：
// 输入：head = [1], k = 1
// 输出：[1]
//
// 提示：
// 列表中节点的数量在范围 sz 内
// 1 <= sz <= 5000
// 0 <= Node.val <= 1000
// 1 <= k <= sz
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
type ListNode = linkedlist.ListNode

// reverseKGroup
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	l := 0
	for n := head; n != nil; n = n.Next {
		l++
	}
	group := l / k
	var dummy = &ListNode{Next: head}
	var prev, curr = dummy, dummy
	for i := 0; i < group; i++ {
		for j := 0; j < k; j++ {
			curr = curr.Next
		}
		pre, cur := reverse(prev.Next, curr)
		prev.Next = pre
		prev = cur
		curr = cur
	}
	// [1,2,3,4,5],2 => [2,1,4,3,5]
	return dummy.Next
}

// reverse
func reverse(head, tail *ListNode) (newHead *ListNode, newTail *ListNode) {
	next := tail.Next
	cur := head
	// cur 代表当前翻转的 node，next 代表尾部指向的 node
	for next != tail {
		tmp := cur.Next
		// 将 cur的 Next 指向 next，表示将前节点指向最后节点
		cur.Next = next
		// 每一次将前一个节点挂在最后的节点上，并将最后节点更新为当前的节点
		next = cur
		cur = tmp
	}
	return tail, head
}

func reverseKGroup1(head *ListNode, k int) *ListNode {
	dummyHead := new(ListNode)
	tail := dummyHead
	p := head
	for p != nil {
		count := 0
		q := p
		for q != nil {
			count++
			if count == k {
				break
			}
			q = q.Next
		}
		if q == nil {
			tail.Next = p
			return dummyHead.Next
		} else {
			tmp := q.Next
			newHead, newTail := reverse1(p, q)
			tail.Next = newHead
			tail = newTail
			p = tmp
		}
	}
	return dummyHead.Next
}

func reverse1(head, tail *ListNode) (newHead, newTail *ListNode) {
	newTail = head
	p := head
	for p != tail {
		tmp := p.Next
		p.Next = newHead
		newHead = p
		p = tmp
	}
	tail.Next = newHead
	newHead = tail
	return
}
