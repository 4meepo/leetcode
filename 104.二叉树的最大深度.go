/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {
	return maxDepth_dfs(root)
}

func maxDepth_dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth_dfs(root.Left) + 1
	r := maxDepth_dfs(root.Right) + 1
	if l > r {
		return l
	}
	return r
}

// @lc code=end
