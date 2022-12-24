/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
	if nil == root {
		return nil
	}

	q := []*TreeNode{root}

	var rtn []int

	// BFS
	for {
		length := len(q)
		tail := q[length-1]
		rtn = append(rtn, tail.Val)

		for i := 0; i < length; i++ {
			node := q[i]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if len(q) == length {
			break
		}
		q = q[length:]
	}

	return rtn
}

// @lc code=end
