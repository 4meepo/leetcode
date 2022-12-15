/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
 */
package leetcode

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}

	data := []int{}
	recoverTree_traverse(root, &data)

	var m, n *int
	length := len(data)
	for i := 0; i < length; i++ {
		if i+1 < length && data[i] > data[i+1] {
			m = &data[i]
			break
		}
	}
	for i := length - 1; i > 0; i-- {
		if data[i] < data[i-1] {
			n = &data[i]
			break
		}
	}

	// swap
	recoverTree_swapValue(root, *m, *n)

}

func recoverTree_traverse(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}

	recoverTree_traverse(root.Left, data)

	*data = append(*data, root.Val)

	recoverTree_traverse(root.Right, data)
}

func recoverTree_swapValue(root *TreeNode, m, n int) {
	if root == nil {
		return
	}
	recoverTree_swapValue(root.Left, m, n)

	if root.Val == m {
		root.Val = n
	} else if root.Val == n {
		root.Val = m

	}
	recoverTree_swapValue(root.Right, m, n)
	return
}

// @lc code=end
