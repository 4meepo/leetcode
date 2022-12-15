/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
 type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
*/

func isValidBST(root *TreeNode) bool {
	var data []int
	return isValidBST_middleTraverse(root, &data)
}

// 利用中序递归遍历判断
func isValidBST_middleTraverse(root *TreeNode, data *[]int) bool {
	if root == nil {
		return true
	}

	if b := isValidBST_middleTraverse(root.Left, data); b == false {
		return false
	}

	d := *data
	if len(d) > 0 {
		if root.Val <= d[len(d)-1] {
			return false
		}
	}
	*data = append(*data, root.Val)

	if b := isValidBST_middleTraverse(root.Right, data); b == false {
		return false
	}
	return true
}

// @lc code=end
