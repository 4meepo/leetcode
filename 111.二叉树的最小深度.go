/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 */
package leetcode

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := minDepth(root.Left)
	r := minDepth(root.Right)
	if l == 0 {
		return r + 1
	}
	if r == 0 {
		return l + 1
	}
	return minDepth_min(l, r) + 1
}

func minDepth_min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
