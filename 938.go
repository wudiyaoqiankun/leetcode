package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	} else if root.Val > R {
		return rangeSumBST(root.Left, L, R)
	} else {
		return root.Val + rangeSumBST(root.Left, L, root.Val) + rangeSumBST(root.Right, root.Val, R)
	}
}
