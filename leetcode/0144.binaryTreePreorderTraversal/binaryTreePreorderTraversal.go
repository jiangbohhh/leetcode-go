package binaryTreePreorderTraversal

import "jiangbohhh/leetcode-go/structure/tree"

// 0144. 二叉搜索树的前序遍历
// 给你二叉树的根节点 root ，返回它节点值的前序遍历。
// 示例 1：
// 输入：root = [1,null,2,3]
// 输出：[1,2,3]
// 示例 2：
// 输入：root = []
// 输出：[]
// 示例 3：
// 输入：root = [1]
// 输出：[1]
// 示例 4：
// 输入：root = [1,2]
// 输出：[1,2]
// 示例 5：
// 输入：root = [1,null,2]
// 输出：[1,2]
// 提示：
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/binary-tree-preorder-traversal
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
type TreeNode = tree.TreeNode

//
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var list = make([]int, 0)
	list = append(list, root.Val)
	if left := preorderTraversal(root.Left); left != nil {
		list = append(list, left...)
	}
	if right := preorderTraversal(root.Right); right != nil {
		list = append(list, right...)
	}
	return list
}
