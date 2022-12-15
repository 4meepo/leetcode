/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
 */
package leetcode

import "testing"

func Test_recoverTree(t *testing.T) {
	_1 := &TreeNode{Val: 1}
	_2 := &TreeNode{Val: 2}
	_3 := &TreeNode{Val: 3}
	_4 := &TreeNode{Val: 4}

	_3.Left = _1
	_3.Right = _4
	_4.Left = _2

	recoverTree(_3)
}
