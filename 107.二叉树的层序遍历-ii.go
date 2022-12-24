/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层序遍历 II
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
func levelOrderBottom(root *TreeNode) [][]int {
	var rtn [][]int
	if root == nil {
		return rtn
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)
	n := 1

	for {
		var arr []int
		delta := 0
		for i := 0; i < n; i++ {
			arr = append(arr, q[i].Val)

			if q[i].Left != nil {
				q = append(q, q[i].Left)
				delta++
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
				delta++
			}
		}
		rtn = append(rtn, arr)

		if delta <= 0 {
			// 无增量,结束
			break
		}
		q = q[n:]
		n = delta
	}

	rtn2 := make([][]int, len(rtn))

	l := len(rtn)
	for i := 0; i < l; i++ {
		rtn2[l-i-1] = rtn[i]
	}

	return rtn2
}

// @lc code=end
