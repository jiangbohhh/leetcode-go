package naryTreePostorderTraversal

import "jiangbohhh/leetcode-go/structure/tree"

// 0590. N 叉树的后序遍历
// 给定一个 N 叉树，返回其节点值的 后序遍历 。
// N 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。
// 进阶：
// 递归法很简单，你可以使用迭代法完成此题吗?
// 示例 1：
// 输入：root = [1,null,3,2,4,null,5,6]
// 输出：[5,6,3,2,4,1]
// 示例 2：
// 输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
// 输出：[2,6,14,11,7,3,12,8,4,13,9,10,5,1]
// 提示：
// N 叉树的高度小于或等于 1000
// 节点总数在范围 [0,10^4] 内
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
type Node = tree.NNode

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var list = make([]int, 0)
	for _, v := range root.Children {
		if child := postorder(v); child != nil {
			list = append(list, child...)
		}
	}
	list = append(list, root.Val)
	return list
}
