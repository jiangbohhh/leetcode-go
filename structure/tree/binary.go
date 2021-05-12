package tree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// N-ary Tree
type NNode struct {
	Val      int
	Children []*NNode
}
